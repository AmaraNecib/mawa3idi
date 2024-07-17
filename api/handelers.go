package api

import (
	"fmt"
	"log"
	"mawa3id/DB"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreateUser(db *DB.Queries) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// var user DB.CreateUserParams
		// has
		var role Role
		if err := ctx.BodyParser(&role); err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err,
			},
			)
		}
		fmt.Println("Error parsing body")

		if err := db.CreateRole(ctx.Context(), role.Name); err != nil {
			return err
		}

		log.Printf("Received user: %+v", role)

		response := fiber.Map{
			"Ok":   true,
			"role": role.Name,
		}

		return ctx.Status(fiber.StatusCreated).JSON(response)
	}
}

func GetAllRoles(db *DB.Queries) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		roles, err := db.GetRoles(ctx.Context())
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err,
			})
		}
		if len(roles) == 0 {
			roles = []DB.Role{}
		}
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"ok":    true,
			"roles": roles,
		})
	}
}

func DeleteRole(db *DB.Queries) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		//parse to int
		id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
		if err != nil {
			return err
		}
		if err := db.DeleteRoleByID(ctx.Context(), id); err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err,
			})
		}
		return ctx.Status(fiber.StatusNoContent).JSON(fiber.Map{
			"ok":      true,
			"message": "Role deleted successfully",
		})
	}
}
