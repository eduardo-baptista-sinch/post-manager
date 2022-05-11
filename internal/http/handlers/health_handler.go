package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func GetHealth(c *fiber.Ctx) error {
	return c.SendStatus(http.StatusNoContent)
}
