package services

import (
	"project/models"
	"project/repositories"
)

type UserService interface {
	Create(user *models.User) *models.User
	GetBy(id int) (*models.User, error)
	GetAllUsers() []*models.User
	UpdateUser(Id int, user *models.User) (*models.User, error)
	DeleteUser(Id int) error
}

type service struct{}

var (
	repo repositories.UserRepository
)

func NewUserService(repository repositories.UserRepository) UserService {
	repo = repository
	return &service{}
}

func (*service) Create(user *models.User) *models.User {
	return repo.CreateUser(user)
}

func (*service) GetBy(id int) (*models.User, error) {
	return repo.GetUser(id)
}

func (*service) GetAllUsers() []*models.User {
	return repo.GetAllUsers()
}

func (*service) UpdateUser(Id int, user *models.User) (*models.User, error) {
	return repo.UpdateUser(Id, user)
}

func (*service) DeleteUser(Id int) error {
	return repo.DeleteUser(Id)
}
