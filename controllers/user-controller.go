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
}

func NewUserController(service services.UserService) UserController {
	userService = service
	return &controller{}
}

func (*controller) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Unable to convert id", http.StatusBadRequest)
		return
	}

	user := userService.GetBy(id)

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
