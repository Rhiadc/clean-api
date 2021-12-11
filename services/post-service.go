package services

import (
	"project/models"
	"project/repositories"
)

type PostService interface {
	CreatePost(post *models.Post) *models.Post
}

type postService struct{}

var (
	postRepo repositories.PostRepository
)

func NewPostService(repo repositories.PostRepository) PostService {
	postRepo = repo
	return &postService{}
}

func (*postService) CreatePost(post *models.Post) *models.Post {
	return postRepo.CreatePost(post)
}
