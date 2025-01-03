package usecases

import (
	"fmt"
	"go-echo/internal/user/entities"
	"go-echo/internal/user/models"
	"go-echo/internal/user/repositories"

	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	repo repositories.UserRepository
}

func Init(repo repositories.UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (s *UserUsecase) CreateUser(user *entities.User) error {
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return s.repo.Create(user)
}

func (s *UserUsecase) Login(loginInformation *models.UserLogin) (*entities.User, error) {
	user, err := s.repo.FindOne(&models.GetUserFilter{Email: loginInformation.Email})
	
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserUsecase) Find(filter *models.GetUserFilter) (*[]entities.User, error) {
	return s.repo.Find(filter)
}

func (s *UserUsecase) Delete(filter *models.GetUserFilter) error {
	err := s.repo.Delete(filter)
	return err
}

func (s *UserUsecase) Seed() (*[]entities.User, error) {
	var createdUsers []entities.User
	for i := 0; i < 100; i++ {
		user := entities.User{
			ID:       uint64(100 + i),
			Email:    fmt.Sprint(i) + "@example.com",
			Username: "User No. " + fmt.Sprint(i),
			Age:      18,
			Password: "Test4",
		}
		err := s.CreateUser(&user)
		if err != nil {
			return nil, err
		}
		createdUsers = append(createdUsers, user)
	}
	return &createdUsers, nil
}

func hashPassword(password string) (string, error) {
	// The cost parameter can be bcrypt.DefaultCost or a value between 4 and 31
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
