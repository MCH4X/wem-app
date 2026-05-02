package service

import (
	model "github.com/MCH4X/wem-app/backend/internal/models"
	repository "github.com/MCH4X/wem-app/backend/internal/repositories"
)

type UserService interface {
	GetUsers() ([]model.User, error)
	CreateUser(user *model.User) error
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) GetUsers() ([]model.User, error) {
	return s.repo.GetAll()
}

func (s *userService) CreateUser(user *model.User) error {
	return s.repo.Create(user)
}
