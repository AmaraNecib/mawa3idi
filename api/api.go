package api

import (
	"log"

	"mawa3id/DB" // Adjust the import path as necessary

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Init(db *DB.Queries) (*fiber.App, error) {
	app := fiber.New(
		fiber.Config{
			Prefork: true,
		},
	)
	app.Use(logger.New())

	api := app.Group("/api")

	// Pass db instance to handler functions
	api.Post("/role", CreateUser(db))
	// get all roles
	api.Get("/role", GetAllRoles(db))
	api.Delete("/role/:id", DeleteRole(db))
	log.Fatal(app.Listen(":3000"))
	return app, nil
}

type Role struct {
	Name string `json:"name"`
}
