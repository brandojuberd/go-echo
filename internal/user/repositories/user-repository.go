package repositories

import (
	"go-echo/internal/user/entities"
	"go-echo/internal/user/models"
)

type UserRepository interface {
	Create(user *entities.User) error
	FindById(id uint) (*entities.User, error)
	Find(filter *models.GetUserFilter) (*[]entities.User, error)
	Delete(filter *models.GetUserFilter) error
}
