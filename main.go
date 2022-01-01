package main

import (
	"github.com/arthur-rl/applyx/database"
	"github.com/arthur-rl/applyx/routes/api"
	"github.com/arthur-rl/applyx/routes/auth"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//setup routes
	api.SetupRoutes(app)
	auth.SetupRoutes(app)
	// connect to database. might not be needed anymroe
	database.Connect()
	app.Listen(":3002")
}
