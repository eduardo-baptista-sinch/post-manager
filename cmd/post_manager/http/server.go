package http

import (
	"main/internal/adapters/repositories"
	"main/internal/core/services"
	"main/internal/http/handlers"
	"main/internal/infra/database"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Server struct {
	App *fiber.App
	DB  *gorm.DB
}

func NewServer() *Server {
	app := fiber.New()

	// TODO: read config from env variables
	// start infrastructure
	db := database.ConnectDB()

	// Start repositories
	postsRepository := repositories.NewPostsRepository(db)

	// Start services
	postsService := services.NewPostsService(postsRepository)

	// Start handlers
	postsHandler := handlers.NewPostsHandler(postsService)

	// Register routes
	app.Get("/api/health", handlers.GetHealth)
	app.Get("/api/posts", postsHandler.GetPosts)

	return &Server{
		App: app,
		DB:  db,
	}
}

func (server *Server) Listen() {
	server.App.Listen(":3000")
}
