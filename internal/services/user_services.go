package services

import (
	"CashMini/internal/models"
	"CashMini/internal/repository"
	"errors"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo}
}
func (s *UserService) CreateUser(username, password string) (*models.UserModel, error) {
	exist, err := s.userRepo.ExistsByUsername(username)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, errors.New("user already exists")
	}
	user := models.UserModel{
		Username: username,
		Password: password,
	}
	err = s.userRepo.CreateUser(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
