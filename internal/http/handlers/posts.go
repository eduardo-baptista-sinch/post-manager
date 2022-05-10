package handlers

import (
	"main/internal/core/services"

	"github.com/gofiber/fiber/v2"
)

type PostsHandler struct {
	postsService *services.PostsService
}

func NewPostsHandler(postsService *services.PostsService) *PostsHandler {
	return &PostsHandler{postsService: postsService}
}

func (handler *PostsHandler) GetPosts(c *fiber.Ctx) error {
	posts := handler.postsService.GetAll()
	return c.JSON(posts)
}
