package user

import (
	"go-echo/internal/user/handlers"
	"go-echo/internal/user/repositories"
	"go-echo/internal/user/services"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func UserInit(db *gorm.DB, api *echo.Group, gui *echo.Group) {
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
