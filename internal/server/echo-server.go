package server

import (
	"context"
	"go-echo/internal/config"
	"go-echo/internal/database"
	"go-echo/internal/shared/customvalidator"
	"go-echo/internal/user/handlers"
	"go-echo/internal/user/repositories"
	"go-echo/internal/user/usecases"
	"html/template"
	"io"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

type echoServer struct {
	app *echo.Echo
	db  database.Database
	cv  *customvalidator.CustomValidator
	config config.Server
}

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func NewEchoServer(db database.Database, serverConfig config.Server) Server {
	echoApp := echo.New()
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("web/*.html")),
	}
	echoApp.Renderer = renderer
	jwtConfig := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(jwtCustomClaims)
		},
		SigningKey: []byte(serverConfig.JwtSecret),
	}
	echoApp.Use(echojwt.WithConfig(jwtConfig))
	echoApp.Logger.SetLevel(log.DEBUG)
	cvalidator := customvalidator.NewCustomValidator()

	return &echoServer{
		app: echoApp,
		db:  db,
		cv:  cvalidator,
		config: serverConfig,
	}
}

type jwtCustomClaims struct {
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
	jwt.RegisteredClaims
}


func (s *echoServer) Start() {

	s.app.Use(middleware.Recover())
	s.app.Use(middleware.Logger())

	// Create a group with the prefix "/api"

	apiGroup := s.app.Group("/api")
	guiGroup := s.app.Group("/gui")

	// Health check adding
	s.app.GET("v1/health", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	s.app.GET("/", func(c echo.Context) error {
		// content := map[string]string{
		// 	"message": "Go Echo",
		// }
		routes := s.app.Routes()
		output := make(map[string][]echo.Route)
		output["Routes"] = make([]echo.Route, 0)
		for i := 0; i < len(routes); i++ {
			route := *routes[i]
			output["Routes"] = append(output["Routes"], route)
		}
		err := c.Render(http.StatusOK, "home", output)

		return err
	})

	InitUserHttpHandler(s, apiGroup, guiGroup)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()
	// Start server
	go func() {
		if err := s.app.Start(":" + s.config.Port); err != nil && err != http.ErrServerClosed {
			s.app.Logger.Fatal(err.Error(), " shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server with a timeout of 10 seconds.
	<-ctx.Done()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := s.app.Shutdown(ctx); err != nil {
		s.app.Logger.Fatal(err)
	}
}

func (s *echoServer) GetValidator() *customvalidator.CustomValidator {
	return s.cv
}

func (s *echoServer) GetDatabase() database.Database {
	return s.db
}

func InitUserHttpHandler(s Server, api *echo.Group, gui *echo.Group) {
	userRepository := repositories.InitUserPostgresRepository(s.GetDatabase().GetDb())
	userService := usecases.Init(userRepository)
	userHandler := handlers.InitUserHandler(userService, s.GetValidator())

	api.POST("/login", userHandler.Login)

	userRoute := api.Group("/users")
	userRoute.GET("", userHandler.Find)
	userRoute.DELETE("", userHandler.Delete)
	userRoute.POST("/seed", userHandler.Seed)

	guiUserRoute := gui.Group("/users")
	guiUserRoute.GET("", userHandler.FindGUI)
}
