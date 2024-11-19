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

func CreateRole(db *DB.Queries) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var role DB.Role
		if role.Name == "" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": "Role name is required",
			})
		}
		if err := ctx.BodyParser(&role); err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err,
			})
		}
		err := db.CreateRole(ctx.Context(), role.Name)
		if err != nil {
			return err
		}
		response := fiber.Map{
			"ok":   true,
			"role": role.Name,
		}
		return ctx.Status(fiber.StatusCreated).JSON(response)
	}
}

func UpdateRoleByID(db *DB.Queries) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var role DB.Role
		if err := ctx.BodyParser(&role); err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err,
			})
		}
		token := strings.Split(string(ctx.Get("Authorization")), " ")[1]
		UserRole, err := auth.GetUserRole((string(token)))
		fmt.Print(UserRole, err)
		if UserRole != static.Admin || err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": "Unauthorized",
			})
		}
		//parse to int
		id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
		if err != nil {
			return err
		}
		err = db.UpdateRoleByID(ctx.Context(), DB.UpdateRoleByIDParams{
			ID:   id,
			Name: role.Name,
		})
		if err != nil {
			return err
		}
		response := fiber.Map{
			"ok":   true,
			"role": role.Name,
		}
		return ctx.Status(fiber.StatusCreated).JSON(response)
	}
}
