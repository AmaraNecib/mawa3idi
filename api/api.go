package api

import (
	"log"
	"os"
	"strings"

	"mawa3id/DB" // Adjust the import path as necessary
	auth "mawa3id/jwt"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Protected() func(*fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
	})
}
func Init(db *DB.Queries) (*fiber.App, error) {
	app := fiber.New(
		fiber.Config{
			Prefork: true,
		},
	)

	app.Use(logger.New())

	api := app.Group("/api")
	v1 := api.Group("/v1")

	// User
	v1.Post("/register", CreateUser(db))
	v1.Post("/login", login(db))
	// authorized routes
	v1.Get(("/token"), restricted)
	v1.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"ok":    false,
				"error": "Unauthorized",
			})
		},
	}))
	v1.Post("/service", CreateServices(db))
	// get all services
	v1.Get("/service", GetAllServices(db))
	v1.Get("/service/:id", GetServiceByID(db))
	// category
	v1.Post("/category", CreateCategory(db))

	// days of the week
	v1.Post("/day", CreateDay(db))
	// get all categories
	v1.Get("/category", GetAllCategories(db))
	v1.Put("/category/:id", UpdateCategoryByID(db))
	v1.Delete("/category/:id", DeleteCategory(db))
	// subcategory
	v1.Post("/subcategory", CreateSubCategory(db))
	// get all subcategories
	v1.Get("/subcategory", GetAllSubCategories(db))
	// v1.Delete("/subcategory/:id", DeleteSubCategory(db))

	// get all users
	// v1.Get("/user", Protected(), GetAllUsers(db))
	v1.Post("/workdays", CreateWorkDays(db))
	// get all workdays
	v1.Get("/workdays", GetAllWorkDays(db))
	v1.Put("/workdays/:id", UpdateWorkDaysByID(db))
	v1.Get("/workdays/:id", GetWorkDaysByID(db))

	// reservation
	v1.Post("/reservation", CreateReservation(db))
	// get all reservations
	// v1.Get("/reservation", GetAllReservations(db))
	// Pass db instance to handler functions
	v1.Post("/role", CreateRole(db))
	// get all roles
	v1.Get("/role", GetAllRoles(db))
	v1.Put("/role/:id", UpdateRoleByID(db))
	v1.Delete("/role/:id", DeleteRole(db))

	// reservation
	v1.Get("/reservation-type", GetAllReservitionTypes(db))
	v1.Post("/reservation-type", CreateReservitionType(db))
	v1.Delete("/reservation-type/:id", DeleteReservitionType(db))

	log.Fatal(app.Listen(":3000"))
	return app, nil
}

func CreateDay(db *DB.Queries) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		day := new(DB.Day)
		if err := ctx.BodyParser(day); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err,
			})
		}
		err := db.CreateDay(ctx.Context(), day.Name)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err,
			})
		}
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"ok":  true,
			"day": day.Name,
		})
	}
}

func restricted(c *fiber.Ctx) error {
	if auth.ValidToken(strings.Split(c.Get("Authorization"), " ")[1]) {
		return c.SendString("Welcome to the restricted area")
	} else {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"ok":    false,
			"error": "Unauthorized",
		})
	}
}
