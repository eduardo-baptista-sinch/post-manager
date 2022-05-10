package services

import (
	"main/internal/core/domain"
	"main/internal/core/ports"
)

type PostsService struct {
	repository ports.PostsRepository
}

func NewPostsService(repository ports.PostsRepository) *PostsService {
	return &PostsService{repository: repository}
}

func (service *PostsService) GetAll() *[]domain.Post {
	return service.repository.GetAll()
}
