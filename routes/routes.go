package routes

import (
	"project/controllers"

	"github.com/gorilla/mux"
)

func InitRoutes(sm *mux.Router, userController controllers.UserController, postController controllers.PostController) {
	getRouter := sm.Methods("GET").Subrouter()
	getRouter.HandleFunc("/users/{id:[0-9]+}", userController.GetUser)
	getRouter.HandleFunc("/users/", userController.GetAllUser)

	putRouter := sm.Methods("PUT").Subrouter()
	putRouter.HandleFunc("/users/{id:[0-9]+}", userController.UpdateUser)

	postRouter := sm.Methods("POST").Subrouter()
	postRouter.HandleFunc("/users", userController.CreateUser)
	postRouter.HandleFunc("/posts", postController.CreatePost)

	deleteRouter := sm.Methods("DELETE").Subrouter()
	deleteRouter.HandleFunc("/users/{id:[0-9]+}", userController.DeleteUser)
}
