package api

import (
	"mawa3id/DB"
	auth "mawa3id/jwt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func CreateComplaint(db *DB.Queries) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var complaint Complaint
		if err := c.BodyParser(&complaint); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err.Error(),
			})
		}
		// get the user id from the token

		userID := auth.GetUserID(strings.Split(c.Get("Authorization"), " ")[1])
		if userID == 0 || complaint.Complaint == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": "UserID and Complaint are required",
			})
		}
		err := db.CreateComplaint(c.Context(), DB.CreateComplaintParams{
			UserID:    int32(userID),
			Complaint: complaint.Complaint,
			TypeID:    int32(complaint.TypeId),
		})
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"ok":  true,
			"msg": "Complaint created successfully",
		})
	}
}

func GetAllComplaints(db *DB.Queries) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		pageStr := c.Query("page")
		page, err := strconv.Atoi(pageStr)
		if err != nil || page < 0 {
			page = 1
		}
		limitStr := c.Query("limit")
		limit, err := strconv.Atoi(limitStr)
		if err != nil || limit < 0 {
			limit = 10
		}
		complaints, err := db.GetAllComplaints(c.Context(), DB.GetAllComplaintsParams{
			Limit:  int32(limit),
			Offset: int32(page-1) * int32(limit),
		})
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"ok":         true,
			"complaints": complaints,
		})
	}
}

func GetComplaintByID(db *DB.Queries) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		idStr := c.Params("id")
		id, err := strconv.Atoi(idStr)
		if err != nil || id <= 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": "Invalid complaint id",
			})
		}
		complaint, err := db.GetComplaintByID(c.Context(), int64(id))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"ok":        true,
			"complaint": complaint,
		})
	}
}
