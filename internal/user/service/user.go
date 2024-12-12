package service

import (
	"fmt"
	"go-echo/internal/user/dto"
	"go-echo/internal/user/model"
	"go-echo/internal/user/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func Init(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user *model.User) error {
	return s.repo.Create(user)
}

func (s *UserService) GetUserByEmail(email string) (*model.User, error) {
	return s.repo.FindByEmail(email)
}

func (s *UserService) Find(filter *dto.GetUserFilter) (*[]model.User, error) {
	return s.repo.Find(filter)
}

func (s *UserService) DeleteAll(filter *dto.GetUserFilter) {
}

func (s *UserService) Seed() (*[]model.User, error) {
	var createdUsers []model.User
	for i := 0; i < 100; i++ {
		user := model.User{
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
