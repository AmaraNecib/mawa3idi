package api

import (
	"fmt"
	"mawa3id/DB"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func SearchServices(db *DB.Queries) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		categoryStr := ctx.Query("category")

		if categoryStr != "" {
			category, err := strconv.ParseInt(categoryStr, 10, 32)
			if err != nil || category == 0 {
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"ok":    false,
					"error": err,
					"msg":   "category not found",
				})
			}
			services, err := db.SearchServicesByCategory(ctx.Context(), DB.SearchServicesByCategoryParams{
				CategoryID: int32(category),
				Limit:      25,
				Offset:     0,
			})
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
			subcategoryID, err := strconv.ParseInt(subcategory, 10, 32)
			if err != nil || subcategoryID == 0 {
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"ok":    false,
					"error": err,
					"msg":   "subcategory not found",
				})
			}
			// s, err := database.Query("SELECT * FROM subcategory WHERE name = $1", subcategory)
			// if err != nil {
			// 	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			// 		"ok":    false,
			// 		"error": err,
			// 	})
			// }
			// fmt.Println(s)
			services, err := db.SearchServicesBySubCategory(ctx.Context(), DB.SearchServicesBySubCategoryParams{
				SubcategoryID: int32(subcategoryID),
				Limit:         25,
				Offset:        0})
			// , DB.SearchServicesBySubCategoryParams{
			// 	Name:   subcategory,
			// 	Limit:  25,
			// 	Offset: 0}
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
