package swagger

import (
	"io"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

var fileContent = `{
"swagger": "2.0",
"info": {
	"title": "Testing API",
	"version": "v1.0.0"
}
}`

func TestSwaggerHandler(t *testing.T) {
	app := fiber.New()
	SwaggerHandler(app, []byte(fileContent))

	req, err := http.NewRequest("GET", "/swagger/doc.json", nil)
	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)

	body, err := io.ReadAll(resp.Body)
	assert.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, fileContent, string(body))
}

func TestSwaggerPrefix(t *testing.T) {
	app := fiber.New()
	SwaggerHandler(app, []byte(fileContent))

	req, err := http.NewRequest("GET", "/swagger", nil)
	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, 301, resp.StatusCode)
	loc, err := resp.Location()
	assert.NoError(t, err)
	assert.Equal(t, "/swagger/", loc.Path)
}
