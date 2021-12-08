package services

import (
	"project/models"
	"project/repositories"
)

type UserService interface {
	Create(user *models.User) *models.User
	GetBy(id int) *models.User
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

func (*service) GetBy(id int) *models.User {
	return repo.GetUser(id)
}
