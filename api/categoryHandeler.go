package api

import (
	"fmt"
	"mawa3id/DB"
	auth "mawa3id/jwt"
	"mawa3id/static"
	"strconv"
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
		if category.Name == "" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": "Name is required",
			})
		}
		token := strings.Split(string(ctx.Get("Authorization")), " ")[1]
		role, err := auth.GetUserRole((string(token)))
		if role != static.Admin {
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

func GetAllCategories(db *DB.Queries) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		categories, err := db.GetCategories(ctx.Context())
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": fmt.Sprintf("something went wrong: %v", err),
			})
		}
		if len(categories) == 0 {
			categories = []DB.Category{}
		}
		response := fiber.Map{
			"ok":         true,
			"categories": categories,
		}
		return ctx.Status(fiber.StatusOK).JSON(response)
	}
}

func UpdateCategoryByID(db *DB.Queries) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var category DB.Category
		var err error
		if err = ctx.BodyParser(&category); err != nil {
			fmt.Println(err)
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": fmt.Sprintf("something went wrong2: %v", err),
			})
		}
		if category.Name == "" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": "Name is required",
			})
		}
		token := strings.Split(string(ctx.Get("Authorization")), " ")[1]
		role, err := auth.GetUserRole((string(token)))
		if role != static.Admin {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": "Unauthorized",
			})
		}
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": fmt.Sprintf("something went wrong1: %v", err),
			})
		}
		id := ctx.Params("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": "Invalid ID",
			})
		}
		err = db.UpdateCategoryByID(ctx.Context(), DB.UpdateCategoryByIDParams{
			ID:          int64(idInt),
			Name:        category.Name,
			Description: category.Description,
		})
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": fmt.Sprintf("something went wrong: %v", err),
			})
		}
		category.ID = int64(idInt)
		response := fiber.Map{
			"ok":         true,
			"categories": category,
		}
		return ctx.Status(fiber.StatusCreated).JSON(response)
	}
}

func DeleteCategory(db *DB.Queries) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		token := strings.Split(string(ctx.Get("Authorization")), " ")[1]
		role, err := auth.GetUserRole((string(token)))
		if role != static.Admin {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": "Unauthorized",
			})
		}
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": fmt.Sprintf("something went wrong1: %v", err),
			})
		}
		id := ctx.Params("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": "Invalid ID",
			})
		}
		err = db.DeleteCategoryByID(ctx.Context(), int64(idInt))
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": fmt.Sprintf("something went wrong: %v", err),
			})
		}
		response := fiber.Map{
			"ok":      true,
			"message": fmt.Sprintf("Category with ID %v deleted successfully", id),
		}
		return ctx.Status(fiber.StatusOK).JSON(response)
	}
}

func CreateSubCategory(db *DB.Queries) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var subCategory DB.Subcategory
		var err error
		if err = ctx.BodyParser(&subCategory); err != nil {
			fmt.Println(err)
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err,
			})
		}
		if subCategory.Name == "" || subCategory.CategoryID == 0 || subCategory.Description == "" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": "All fields are required",
			})
		}
		token := strings.Split(string(ctx.Get("Authorization")), " ")[1]
		role, err := auth.GetUserRole((string(token)))
		if role != static.Admin {
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
		err = db.CreateSubcategory(ctx.Context(), DB.CreateSubcategoryParams{
			Name:        subCategory.Name,
			Description: subCategory.Description,
			CategoryID:  subCategory.CategoryID,
		})
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": fmt.Sprintf("something went wrong: %v", err),
			})
		}
		response := fiber.Map{
			"ok":   true,
			"user": subCategory,
		}
		return ctx.Status(fiber.StatusCreated).JSON(response)
	}
}

func GetAllSubCategories(db *DB.Queries) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		subcategories, err := db.GetSubcategories(ctx.Context())
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": fmt.Sprintf("something went wrong: %v", err),
			})
		}
		if len(subcategories) == 0 {
			subcategories = []DB.GetSubcategoriesRow{}
		}
		response := fiber.Map{
			"ok":            true,
			"subcategories": subcategories,
		}
		return ctx.Status(fiber.StatusOK).JSON(response)
	}
}

func GetSubCategoriesByID(db *DB.Queries) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": "Invalid ID",
			})
		}
		subcategories, err := db.GetSubCatgoriesByCatgoryId(ctx.Context(), int32(idInt))
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": fmt.Sprintf("something went wrong: %v", err),
			})
		}
		if len(subcategories) == 0 {
			subcategories = []DB.Subcategory{}
		}
		response := fiber.Map{
			"ok":            true,
			"subcategories": subcategories,
		}
		return ctx.Status(fiber.StatusOK).JSON(response)
	}
}
