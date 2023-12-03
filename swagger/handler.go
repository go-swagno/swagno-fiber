package swagger

import (
	swaggerUi "github.com/go-swagno/swagno-files"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

type Config struct {
	Prefix string
}

var defaultConfig = Config{
	Prefix: "/swagger",
}

type option func(*Config)

func WithPrefix(prefix string) option {
	return func(c *Config) {
		c.Prefix = prefix
	}
}

func SwaggerHandler(a *fiber.App, doc []byte, opts ...option) {
	config := Config{}

	for _, opt := range opts {
		opt(&config)
	}

	if config.Prefix == "" {
		config.Prefix = defaultConfig.Prefix
	}

	a.Use(defaultConfig.Prefix+"/doc.json", func(c *fiber.Ctx) error {
		c.Set("Content-type", "application/json; charset=utf-8")
		return c.SendString(string(doc))
	})

	a.Use(func(c *fiber.Ctx) error {
		if c.Path() == config.Prefix {
			return c.Redirect(config.Prefix+"/", 301)
		}
		return c.Next()
	})

	a.Use(defaultConfig.Prefix, filesystem.New(filesystem.Config{
		Root: swaggerUi.HTTP,
	}))
}
