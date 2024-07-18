package api

import (
	"fmt"
	"mawa3id/DB"
	auth "mawa3id/jwt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func CreateCategory(db *DB.Queries) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var category DB.Category
		var err error
		if err = ctx.BodyParser(&category); err != nil {
			fmt.Println(err)
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err,
			})
		}
		token := strings.Split(string(ctx.Get("Authorization")), " ")[1]
		role, err := auth.GetUserRole((string(token)))
		if role != "admin" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": "Unauthorized",
			})
		}
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err,
			})
		}
		err = db.CreateCategory(ctx.Context(), DB.CreateCategoryParams{
			Name:        category.Name,
			Description: category.Description,
		})
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": fmt.Sprintf("something went wrong: %v", err),
			})
		}
		response := fiber.Map{
			"ok":   true,
			"user": category,
		}
		return ctx.Status(fiber.StatusCreated).JSON(response)
	}
}
