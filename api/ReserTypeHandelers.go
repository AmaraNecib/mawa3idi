package api

import (
	"log"
	"mawa3id/DB"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreateReservitionType(db *DB.Queries) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var reservitionType ReservationType
		if err := ctx.BodyParser(&reservitionType); err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err,
			})
		}
		res, err := db.CreateReserveType(ctx.Context(), reservitionType.Name)
		if err != nil {
			return err
		}
		log.Printf("Received reservition type: %+v", res)
		response := fiber.Map{
			"ok":          true,
			"reservition": reservitionType.Name,
		}
		return ctx.Status(fiber.StatusCreated).JSON(response)
	}
}

func GetAllReservitionTypes(db *DB.Queries) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		reservitionTypes, err := db.GetReserveTypes(ctx.Context())
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err,
			})
		}
		if len(reservitionTypes) == 0 {
			reservitionTypes = []DB.ReserveType{}
		}
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"ok":              true,
			"reservitionType": reservitionTypes,
		})
	}
}

func DeleteReservitionType(db *DB.Queries) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		id, err := strconv.ParseInt(ctx.Params("id"), 10, 64)
		if err != nil {
			return err
		}
		if err := db.DeleteReserveTypeByID(ctx.Context(), id); err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err,
			})
		}
		return ctx.Status(fiber.StatusAccepted).JSON(fiber.Map{
			"ok":      true,
			"message": "Reservition type deleted successfully",
		})
	}
}
