package ports

import "main/internal/core/domain"

type PostsRepository interface {
	GetAll() *[]domain.Post
}
