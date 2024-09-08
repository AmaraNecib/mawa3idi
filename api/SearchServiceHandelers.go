package api

import (
	"database/sql"
	"fmt"
	"mawa3id/DB"

	"github.com/gofiber/fiber/v2"
)

func SearchServices(db *DB.Queries) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		category := ctx.Query("category")
		// fmt.Println(category == "")
		if category != "" {
			services, err := db.SearchServicesByCategory(ctx.Context(), DB.SearchServicesByCategoryParams{
				Name:   category,
				Limit:  25,
				Offset: 0})
			if err != nil {
				ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"ok":    false,
					"error": err,
				})
			}
			if len(services) == 0 {
				services = []DB.Service{}
			}
			return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
				"ok":       true,
				"services": services,
			})
		}
		subcategory := ctx.Query("subcategory")
		if subcategory != "" {
			services, err := db.SearchServicesBySubCategory(ctx.Context(), DB.SearchServicesBySubCategoryParams{
				Column1: sql.NullString{String: subcategory, Valid: true},
				Limit:   25,
				Offset:  0,
			})
			fmt.Println(services)
			if err != nil {
				ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"ok":    false,
					"error": err,
				})
			}
			if len(services) == 0 {
				services = []DB.Service{}
			}
			return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
				"ok":       true,
				"services": services,
			})
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"ok":    false,
			"error": "category or subcategory not found",
		})
	}
}
