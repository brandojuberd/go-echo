package repository

import (
	"go-echo/internal/user/dto"
	"go-echo/internal/user/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func Init(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *model.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) FindByEmail(email string) (*model.User, error) {
	var user model.User
	err := r.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (r *UserRepository) FindById(id uint) (*model.User, error) {
	var user model.User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *UserRepository) Find(filter *dto.GetUserFilter) (*[]model.User, error) {
	var users []model.User
	err := r.db.Find(&users, filter).Error
	return &users, err
}

func (r *UserRepository) Delete() {
	// r.db.Raw("DELETE from users")
	// return &users, err
}
