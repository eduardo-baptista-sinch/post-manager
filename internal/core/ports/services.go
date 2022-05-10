package ports

import "main/internal/core/domain"

type CratePostCommand struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type PostsService interface {
	GetAll() *[]domain.Post
	Create(command CratePostCommand) (*domain.Post, error)
}
