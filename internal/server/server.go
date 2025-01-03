package server

import (
	"go-echo/internal/database"
	"go-echo/internal/shared/customvalidator"
)

type Server interface {
	Start()
	GetValidator() *customvalidator.CustomValidator
	GetDatabase() database.Database
}