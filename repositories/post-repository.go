package repositories

import "project/models"

type PostRepository interface {
	CreatePost(p *models.Post) *models.Post
	// GetPost(Id int) ([]*models.Post, error)
	// GetPosts() []*models.Post
	// UpdatePost(Id int, post *models.Post) (*models.Post, error)
	// DeletePost(Id int) error
}
