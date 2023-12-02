package main

import (
	"github.com/go-swagno/swagno-fiber/swagger"
	"github.com/gofiber/fiber/v2"
)

func main() {
	var fileContent = `{
    "swagger": "2.0",
    "info": {
      "title": "Testing API",
      "version": "v1.0.0"
    }
  }`

	app := fiber.New()
	swagger.SwaggerHandler(app, []byte(fileContent))

	app.Listen(":8080")
}
