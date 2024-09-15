package api

import (
	"mawa3id/DB"

	"github.com/gofiber/fiber/v2"
)

func GetAllComplaintTypes(db *DB.Queries) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		complaintTypes, err := db.GetAllComplaintTypes(c.Context())
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"ok":              true,
			"complaint_types": complaintTypes,
		})
	}
}

func CreateComplaintType(db *DB.Queries) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var complaintType ComplaintType
		if err := c.BodyParser(&complaintType); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err.Error(),
			})
		}
		err := db.CreateComplaintType(c.Context(), complaintType.Name)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"ok":  true,
			"msg": "Complaint type created successfully",
		})
	}
}
