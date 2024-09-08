package api

import (
	"mawa3id/DB"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllReservitionStatus(db *DB.Queries) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		reservitionStatus, err := db.GetAllReserveStatus(ctx.Context())
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err,
			})
		}
		// if len(reservitionStatus) == 0 {
		// 	reservitionStatus := []DB.ReservationStatus{}
		// }
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"ok":                true,
			"reservitionStatus": reservitionStatus,
		})
	}
}

func CreateReservitionStatus(db *DB.Queries) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		reservitionStatus := new(ReservationStatus)
		if err := ctx.BodyParser(reservitionStatus); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err,
			})
		}
		err := db.CreateReserveStatus(ctx.Context(), reservitionStatus.Name)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": "oops, can't create reservition status",
			})
		}
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"ok":                true,
			"reservitionStatus": reservitionStatus.Name,
		})
	}
}

func UpdateReservitionStatusByID(db *DB.Queries) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		reservitionStatus := new(ReservationStatus)
		id, err := strconv.ParseInt(ctx.Params("id"), 10, 32)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err,
			})
		}
		reservitionStatus.Id = int32(id)
		if err := ctx.BodyParser(reservitionStatus); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err,
			})
		}
		err = db.UpdateReserveStatusName(ctx.Context(), DB.UpdateReserveStatusNameParams{
			ID:   int64(reservitionStatus.Id),
			Name: reservitionStatus.Name})
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": "oops, can't update reservition status",
			})
		}
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"ok":                true,
			"reservitionStatus": reservitionStatus.Name,
		})
	}
}
