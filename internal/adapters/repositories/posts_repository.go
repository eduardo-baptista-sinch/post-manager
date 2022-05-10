package repositories

import (
	"main/internal/core/domain"

	"gorm.io/gorm"
)

type PostsRepository struct {
	db *gorm.DB
}

func NewPostsRepository(db *gorm.DB) *PostsRepository {
	return &PostsRepository{db: db}
}

func (repository *PostsRepository) GetAll() *[]domain.Post {
	posts := new([]domain.Post)
	repository.db.Find(&posts)
	return posts
}
