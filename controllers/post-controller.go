package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"project/models"
	"project/services"
)

var (
	postService services.PostService
)

type PostController interface {
	CreatePost(w http.ResponseWriter, r *http.Request)
}

type postController struct{}

func NewPostController(service services.PostService) PostController {
	postService = service
	return &postController{}
}

func (*postController) CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	post := &models.Post{}
	err := json.NewDecoder(r.Body).Decode(post)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error unmarshelling the posts data", http.StatusInternalServerError)
		return
	}

	result := postService.CreatePost(post)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
