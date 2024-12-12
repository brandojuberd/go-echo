package user

import (
	"go-echo/internal/user/handler"
	"go-echo/internal/user/repository"
	"go-echo/internal/user/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func UserInit(db *gorm.DB, api *echo.Group, gui *echo.Group) {
	userRepository := repository.Init(db)
	userService := service.Init(userRepository)
	userHandler := handler.Init(userService)

	userRoute := api.Group("/users")
	userRoute.GET("", userHandler.FindAll)
	userRoute.POST("/seed", userHandler.Seed)

	guiUserRoute := gui.Group("/users")
	guiUserRoute.GET("", userHandler.FindAllGUI)

}
