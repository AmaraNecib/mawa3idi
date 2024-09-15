package api

import (
	"mawa3id/DB"
	auth "mawa3id/jwt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func RequestDeleteAccount(db *DB.Queries) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		// get the user id from the token
		user_id := auth.GetUserID(strings.Split(c.Get("Authorization"), " ")[1])

		if user_id == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": "Invalid token",
			})
		}
		err := db.CreateDeleteAccountRequest(c.Context(), int32(user_id))
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"ok":    false,
				"error": "Internal server error",
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"ok":  true,
			"msg": "Request sent",
		})
	}
}
