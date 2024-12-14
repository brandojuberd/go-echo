package handlers

import "github.com/labstack/echo/v4"

type UserHandler interface {
	CreateUser(e echo.Context) error
	Find(c echo.Context) error
	FindGUI(c echo.Context) error
	Delete(c echo.Context) error
	Seed(c echo.Context) error
}
