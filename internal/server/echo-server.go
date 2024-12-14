package server

import (
	"go-echo/internal/database"
	"go-echo/internal/user/handlers"
	"go-echo/internal/user/repositories"
	"go-echo/internal/user/services"
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type echoServer struct {
	app *echo.Echo
	db  database.Database
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

func NewEchoServer(db database.Database) Server {
	echoApp := echo.New()
	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("web/*.html")),
	}
	echoApp.Renderer = renderer
	echoApp.Logger.SetLevel(log.DEBUG)

	return &echoServer{
		app: echoApp,
		db:  db,
	}
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

	InitUserHttpHandler(s.db.GetDb(), apiGroup, guiGroup)

	port := os.Getenv("PORT")
	s.app.Logger.Fatal(s.app.Start(":" + port))
}

func InitUserHttpHandler(db *gorm.DB, api *echo.Group, gui *echo.Group) {
	userRepository := repositories.InitUserPostgresRepository(db)
	userService := services.Init(userRepository)
	userHandler := handlers.InitUserHandler(userService)

	userRoute := api.Group("/users")
	userRoute.GET("", userHandler.Find)
	userRoute.DELETE("", userHandler.Delete)
	userRoute.POST("/seed", userHandler.Seed)

	guiUserRoute := gui.Group("/users")
	guiUserRoute.GET("", userHandler.FindGUI)
}
