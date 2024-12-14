package services

import (
	"fmt"
	"go-echo/internal/user/entities"
	"go-echo/internal/user/models"
	"go-echo/internal/user/repositories"
)

type UserService struct {
	repo repositories.UserRepository
}

func Init(repo repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *entities.User) error {
	return s.repo.Create(user)
}

// func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
// 	return s.repo.FindByEmail(email)
// }

func (s *UserService) Find(filter *models.GetUserFilter) (*[]entities.User, error) {
	return s.repo.Find(filter)
}

func (s *UserService) Delete(filter *models.GetUserFilter) error {
	return s.repo.Delete(filter)
}

func (s *UserService) Seed() (*[]entities.User, error) {
	var createdUsers []entities.User
	for i := 0; i < 100; i++ {
		user := entities.User{
			ID:       uint64(i),
			Email:    fmt.Sprint(i) + "@example.com",
			Username: "User No. " + fmt.Sprint(i),
			Age:      18,
			Password: "Not Hashed Password",
		}
		s.repo.Create(&user)
		createdUsers = append(createdUsers, user)
	}
	return &createdUsers, nil
}
