package handler

import (
	"fmt"
	"go-echo/internal/user/dto"
	"go-echo/internal/user/model"
	"go-echo/internal/user/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	userService *service.UserService
}

func Init(userService *service.UserService) *UserController {

	return &UserController{userService: userService}

}

func (controller *UserController) CreateUser(e echo.Context) {
	var user model.User
	// if err := e.ShouldBindJSON(&user); err != nil {
	// 	// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// if err := h.userService.CreateUser(&user); err != nil {
	// 	// c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	response := model.User{
		ID:    user.ID,
		Email: user.Email,
	}

	e.JSON(http.StatusCreated, response)
}

func (controller *UserController) FindAll(c echo.Context) error {
	filter := new(dto.GetUserFilter)
	err := c.Bind(filter)

	if err != nil {
		return err
	}

	users, err := controller.userService.Find(filter)
	if err != nil {
		return err
	}
	output := map[string]*[]model.User{"users": users}
	return c.JSON(http.StatusOK, output)
}

func (controller *UserController) FindAllGUI(c echo.Context) error {
	filter := new(dto.GetUserFilter)
	err := c.Bind(filter)

	if err != nil {
		return err
	}

	users, err := controller.userService.Find(filter)
	if err != nil {
		return err
	}
	fmt.Println(map[string]*[]model.User{"Users": users})
	output := map[string]*[]model.User{"Users": users}

	err = c.Render(http.StatusOK, "users", output)
	fmt.Println(err)
	return err
}

func (controller *UserController) Seed(c echo.Context) error {
	users, err := controller.userService.Seed()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}
