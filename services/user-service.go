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

type userService struct{}

var (
	userRepo repositories.UserRepository
)

func NewUserService(repository repositories.UserRepository) UserService {
	userRepo = repository
	return &userService{}
}

func (*userService) Create(user *models.User) *models.User {
	return userRepo.CreateUser(user)
}

func (*userService) GetBy(id int) (*models.User, error) {
	return userRepo.GetUser(id)
}

func (*userService) GetAllUsers() []*models.User {
	return userRepo.GetAllUsers()
}

func (*userService) UpdateUser(Id int, user *models.User) (*models.User, error) {
	return userRepo.UpdateUser(Id, user)
}

func (*userService) DeleteUser(Id int) error {
	return userRepo.DeleteUser(Id)
}
