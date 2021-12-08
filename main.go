package main

import (
	"fmt"
	"log"
	"net/http"
	"project/controllers"
	"project/repositories"
	"project/services"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var (
	userRepo       repositories.UserRepository = repositories.NewGormRepository()
	userService    services.UserService        = services.NewUserService(userRepo)
	userController controllers.UserController  = controllers.NewUserController(userService)
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	sm := mux.NewRouter()

	getRouter := sm.Methods("GET").Subrouter()
	getRouter.HandleFunc("/{id:[0-9]+}", userController.GetUser)

	postRouter := sm.Methods("POST").Subrouter()
	postRouter.HandleFunc("/", userController.CreateUser)

	port := ":9090"
	fmt.Printf("Mux HTTP server running on port %v", port)
	http.ListenAndServe(port, sm)
}
