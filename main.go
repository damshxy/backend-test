package main

import (
	"github.com/damshxy/manajemen-backend-test/database"
	"github.com/damshxy/manajemen-backend-test/database/migrations"
	"github.com/damshxy/manajemen-backend-test/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Run Connect To Database
	database.ConnectDatabase()
	
	// Run Migraition
	migrations.RunMigration()

	// Routes
	api := app.Group("/api/v1")
	routes.MainRoutes(api)

	// Start Server
	app.Listen(":3000")
}