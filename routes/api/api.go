package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	fmt.Println("Loading API Routes")
	api := app.Group("/api")
	api.Use("/", GetIndex)
}

func GetIndex(c *fiber.Ctx) error {
	c.SendString("ApplyX API Version 1")
	return nil
}
