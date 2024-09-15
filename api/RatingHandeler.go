package api

import (
	"mawa3id/DB"
	auth "mawa3id/jwt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func CreateRating(db *DB.Queries) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var rating DB.Rating
		if err := c.BodyParser(&rating); err != nil || rating.UserID == 0 || rating.ServiceID == 0 || rating.Rating < 1 || rating.Rating > 5 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": "Invalid request",
			})
		}
		err := db.CreateRating(c.Context(), DB.CreateRatingParams{
			UserID:    rating.UserID,
			ServiceID: rating.ServiceID,
			Rating:    rating.Rating,
			Comment:   rating.Comment,
		})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"ok":    false,
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"ok": true,
		})
	}
}

func GetAllRatingByServiceID(db *DB.Queries) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		serviceID, err := c.ParamsInt("id")
		limit := c.QueryInt("limit")
		page := c.QueryInt("page")
		if err != nil || serviceID == 0 || limit == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": "Invalid request",
			})
		}
		if page == 0 {
			page = 1
		}
		ratings, err := db.GetAllRatingByServiceID(c.Context(), DB.GetAllRatingByServiceIDParams{
			ServiceID: int32(serviceID),
			Limit:     int32(limit),
			Offset:    int32((page - 1) * limit),
		})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"ok":    false,
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"ok":      true,
			"ratings": ratings,
		})
	}
}

func DeleteRating(db *DB.Queries) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		ratingID, err := c.ParamsInt("id")
		if err != nil || ratingID == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": "Invalid request",
			})
		}
		// get user id from token
		user_id := auth.GetUserID(strings.Split(c.Get("Authorization"), " ")[1])

		if user_id == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": "Invalid token",
			})
		}
		// check if the user is the owner of the rating
		err = db.DeleteRating(c.Context(), DB.DeleteRatingParams{
			ID:     int64(ratingID),
			UserID: int32(user_id),
		})
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"ok":    false,
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"ok": true,
		})
	}
}
