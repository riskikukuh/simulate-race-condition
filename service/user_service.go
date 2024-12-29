package service

import (
	"simulation-race-condition/models"
)

type UserService interface {
	Create(data models.UserRequest) (*models.User, error)
	Update(data *models.User) (*models.User, error)
	FindAll() ([]models.User, error)
	FindById(userId int64) (*models.User, error)
	DeleteByUserId(userId int64) error
}
