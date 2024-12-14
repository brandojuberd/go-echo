package repositories

import (
	"go-echo/internal/user/entities"
	"go-echo/internal/user/models"

	"gorm.io/gorm"
)

type UserPostgresRepository struct {
	db *gorm.DB
}

func InitUserPostgresRepository(db *gorm.DB) UserRepository {
	return &UserPostgresRepository{db: db}
}	

func (r *UserPostgresRepository) Create(user *entities.User) error {
	return r.db.Create(user).Error
}

func (r *UserPostgresRepository) FindById(id uint) (*entities.User, error) {
	var user entities.User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *UserPostgresRepository) Find(filter *models.GetUserFilter) (*[]entities.User, error) {
	var users []entities.User
	err := r.db.Find(&users, filter).Error
	return &users, err
}

func (r *UserPostgresRepository) Delete(filter *models.GetUserFilter) error {
	return r.db.Delete(&entities.User{}, filter).Error
}