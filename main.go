package main

import (
	"os"

	"github.com/M-Faheem-Khan/goTemplateEngine-Test/engine"
	"github.com/gofiber/fiber/v2"
)

func main() {
	templatePath := "/workspaces/goTemplateEngine-Test/templates/index.html"
	htmlFile, err := os.ReadFile(templatePath)
	if err != nil {
		panic(err)
	}

	vars := map[string]string{
		"title":           "POC - Template Engine written in Go",
		"sample_variable": "Hello World - I'm rendered via Golang template engine",
		"boolean":         "true",
	}

	s := engine.Parse(htmlFile, vars)

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Send([]byte(s))
	})

	app.Listen(":3000")

}

// EOF
