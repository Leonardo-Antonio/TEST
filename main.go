package main

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(map[string]interface{}{
			"name":    "Leonardo",
			"message": "Welcome a docker and golang in rasberry pi",
		})
	})
	app.Listen(":8080")
}
