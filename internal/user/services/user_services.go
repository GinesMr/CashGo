package services

import (
	"CashMini/internal/user/models"
	"CashMini/internal/user/repository"
	passwordHashing "CashMini/internal/utils/hasher"
	"errors"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo}
}

func (s *UserService) CreateUser(username, password, email string) (*models.UserModel, error) {
	exist, err := s.userRepo.ExistsByUsername(username)
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, errors.New("user already exists")
	}
	passwordHash, err := passwordHashing.HashPassword(password)
	if err != nil {
		return nil, err
	}
	user := models.UserModel{
		Username: username,
		Password: passwordHash,
		Email:    email,
	}
	err = s.userRepo.CreateUser(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
func (s *UserService) DeleteUserByUsernamePassword(username string, password string) bool {
	err := s.userRepo.DeleteUserByUsernamePassword(username, password)
	if err != nil {
		return false
	}
	return true

}
