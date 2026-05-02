package repository

import (
	model "github.com/MCH4X/wem-app/backend/internal/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAll() ([]model.User, error)
	Create(user *model.User) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) GetAll() ([]model.User, error) {
	var users []model.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepo) Create(user *model.User) error {
	return r.db.Create(user).Error
}
