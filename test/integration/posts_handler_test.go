package integration

import (
	"io/ioutil"
	"main/cmd/post_manager/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHealth(t *testing.T) {
	server := http.NewServer()

	req := httptest.NewRequest("GET", "/api/health", nil)

	resp, _ := server.Test(req, -1)

	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equalf(t, resp.StatusCode, 204, "Expected status code to be 204, got %d", resp.StatusCode)
	assert.Equalf(t, string(body), "", "Expected body to be nil, got %s", string(body))
}
