package handlers

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func setup() *fiber.App {
	app := fiber.New()
	app.Get("/api/health", GetHealth)
	return app
}

func TestGetHealth(t *testing.T) {
	app := setup()

	req := httptest.NewRequest("GET", "/api/health", nil)

	resp, _ := app.Test(req, -1)

	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equalf(t, resp.StatusCode, 204, "Expected status code to be 204, got %d", resp.StatusCode)
	assert.Equalf(t, string(body), "", "Expected body to be nil, got %s", string(body))
}
