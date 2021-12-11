package main

import (
	"fmt"
	"log"
	"net/http"
	"project/controllers"
	"project/repositories"
	"project/routes"
	"project/services"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var (
	userRepo       repositories.UserRepository = repositories.NewGormRepository()
	userService    services.UserService        = services.NewUserService(userRepo)
	userController controllers.UserController  = controllers.NewUserController(userService)

	postRepo       repositories.PostRepository = repositories.NewGormRepository()
	postService    services.PostService        = services.NewPostService(postRepo)
	postController controllers.PostController  = controllers.NewPostController(postService)
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	sm := mux.NewRouter()
	routes.InitRoutes(sm, userController, postController)

	port := ":9090"
	fmt.Printf("Mux HTTP server running on port %v", port)
	log.Fatal(http.ListenAndServe(port, sm))
}
