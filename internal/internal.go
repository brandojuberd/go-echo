package internal

import (
	"fmt"
	"go-echo/internal/database"
	"go-echo/internal/user"
	"html/template"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

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

func InitInternal() {
	err := godotenv.Load("development.env")

	if err != nil {
		fmt.Println(err)
	}

	port := os.Getenv("PORT")

	fmt.Println("Start go-echo Application")

	db := database.InitDatabaseConnection()

	e := echo.New()

	// Create a group with the prefix "/api"
	apiGroup := e.Group("/api")
	guiGroup := e.Group("/gui")

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("web/*.html")),
	}
	e.Renderer = renderer

	e.GET("/", func(c echo.Context) error {
		// content := map[string]string{
		// 	"message": "Go Echo",
		// }
		routes := e.Routes()
		output := make(map[string][]echo.Route)
		output["Routes"] = make([]echo.Route, 0)
		for i := 0; i < len(routes); i++ {
			route := *routes[i]
			output["Routes"] = append(output["Routes"], route)
		}
		err := c.Render(http.StatusOK, "home", output)

		return err
	})

	user.UserInit(db, apiGroup, guiGroup)

	e.Logger.Fatal(e.Start(":" + port))

}
