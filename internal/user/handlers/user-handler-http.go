package handlers

import (
	"fmt"
	"go-echo/internal/user/entities"
	"go-echo/internal/user/models"
	"go-echo/internal/user/usecases"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userService *usecases.UserUsecase
}

func InitUserHandler(userService *usecases.UserUsecase) UserHandler {

	return &userHandler{userService: userService}

}

func (controller *userHandler) CreateUser(e echo.Context) error {
	var user entities.User
	// if err := e.ShouldBindJSON(&user); err != nil {
	// 	// c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }

	// if err := h.userService.CreateUser(&user); err != nil {
	// 	// c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	// 	return
	// }

	response := entities.User{
		ID:    user.ID,
		Email: user.Email,
	}

	return e.JSON(http.StatusCreated, response)
}

func (controller *userHandler) Find(c echo.Context) error {
	filter := new(models.GetUserFilter)
	err := c.Bind(filter)

	if err != nil {
		return err
	}

	users, err := controller.userService.Find(filter)
	if err != nil {
		return err
	}
	output := map[string]*[]entities.User{"users": users}
	return c.JSON(http.StatusOK, output)
}

func (controller *userHandler) FindGUI(c echo.Context) error {
	filter := new(models.GetUserFilter)
	err := c.Bind(filter)

	if err != nil {
		return err
	}

	users, err := controller.userService.Find(filter)
	if err != nil {
		return err
	}
	fmt.Println(map[string]*[]entities.User{"Users": users})
	output := map[string]*[]entities.User{"Users": users}

	err = c.Render(http.StatusOK, "users", output)
	fmt.Println(err)
	return err
}

func (controller *userHandler) Delete(c echo.Context) error {
	filter := new(models.GetUserFilter)
	err := c.Bind(filter)

	if err != nil {
		return err
	}

	err = controller.userService.Delete(filter)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, "successfully ")
}

func (controller *userHandler) Seed(c echo.Context) error {
	users, err := controller.userService.Seed()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}
