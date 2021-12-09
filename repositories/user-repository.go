package repositories

import (
	"project/models"
)

type UserRepository interface {
	CreateUser(u *models.User) *models.User
	GetUser(Id int) (*models.User, error)
	GetAllUsers() []*models.User
	UpdateUser(Id int, user *models.User) (*models.User, error)
	DeleteUser(Id int) error
}
