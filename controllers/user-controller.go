package controllers

import (
	"encoding/json"
	"net/http"
	"project/models"
	"project/services"
	"strconv"

	"github.com/gorilla/mux"
)

var (
	userService services.UserService
)

type controller struct{}

type UserController interface {
	GetUser(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	GetAllUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)
}

func NewUserController(service services.UserService) UserController {
	userService = service
	return &controller{}
}

func (*controller) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	id, err := getId(r)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadGateway)
	}

	user, err := userService.GetBy(id)

	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)

}

func (*controller) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user := &models.User{}
	err := json.NewDecoder(r.Body).Decode(user)

	if err != nil {
		http.Error(w, "Error unmarshelling the posts data", http.StatusInternalServerError)
		return
	}

	result := userService.Create(user)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)

}

func (*controller) GetAllUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	users := userService.GetAllUsers()
	err := json.NewEncoder(w).Encode(users)

	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func (*controller) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := getId(r)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadGateway)
		return
	}

	user := &models.User{}
	err = json.NewDecoder(r.Body).Decode(user)

	updatedUser, err := userService.UpdateUser(id, user)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadGateway)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(updatedUser)
}

func (*controller) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id, err := getId(r)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadGateway)
		return
	}

	err = userService.DeleteUser(id)

	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func getId(r *http.Request) (int, error) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		return -1, err
	}
	return id, nil
}
