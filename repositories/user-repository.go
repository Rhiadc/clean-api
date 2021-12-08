package repositories

import (
	"project/models"
)

type UserRepository interface {
	CreateUser(u *models.User) *models.User
	GetUser(Id int) *models.User
	// GetUsers() ([]models.User, error)
	// UpdateUser(id int) (models.User, error)
	// DeleteUser(id int) error
}
