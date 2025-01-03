package handlers

import (
	"fmt"
	"go-echo/internal/shared/customvalidator"
	"go-echo/internal/user/entities"
	"go-echo/internal/user/models"
	"go-echo/internal/user/usecases"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userHandler struct {
	userService *usecases.UserUsecase
	cv          *customvalidator.CustomValidator
}

func InitUserHandler(userService *usecases.UserUsecase, cv *customvalidator.CustomValidator) UserHandler {
	return &userHandler{userService: userService, cv: cv}
}

func (controller *userHandler) CreateUser(c echo.Context) error {
	var user *entities.User
	if err := c.Bind(user); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	if err := controller.userService.CreateUser(user); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	return c.JSON(http.StatusCreated, "register success")
}

func (controller *userHandler) Login(c echo.Context) error {

	data := &models.UserLogin{}
	err := c.Bind(data)
	if err != nil {
		return EchoResponse(c, http.StatusBadRequest, err.Error())
	}

	validate := controller.cv.Validate(data)

	if len(validate) > 0 {
		return EchoResponse(c, http.StatusBadRequest, controller.cv.HumanizeMessage(validate))
	}

	_, err = controller.userService.Login(data)
	if err != nil {
		return EchoResponse(c, http.StatusBadRequest, err.Error())
	}

	return EchoResponse(c, http.StatusOK, "ok")
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
	if err := c.Bind(filter); err != nil {
		return EchoResponse(c, http.StatusBadRequest, err.Error())
	}

	users, err := controller.userService.Find(filter)
	if err != nil {
		return EchoResponse(c, http.StatusBadRequest, err.Error())
	}

	output := map[string]*[]entities.User{"Users": users}

	err = c.Render(http.StatusOK, "users", output)
	fmt.Println(err)
	return err
}

func (controller *userHandler) Delete(c echo.Context) error {
	filter := new(models.GetUserFilter)

	if err := c.Bind(filter); err != nil {
		return EchoResponse(c, http.StatusBadRequest, err.Error())
	}

	if err := controller.userService.Delete(filter); err != nil {
		return EchoResponse(c, http.StatusBadRequest, err.Error())
	}

	return c.String(http.StatusOK, "delete success")
}

func (controller *userHandler) Seed(c echo.Context) error {
	users, err := controller.userService.Seed()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, users)
}
