package api

import (
	auth "mawa3id/jwt"
	"mawa3id/static"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func AdminMiddleware(c *fiber.Ctx) error {
	// Assuming the user role is stored in locals after JWT verification or session
	userRole, err := auth.GetUserRole((strings.Split(c.Get("Authorization"), " ")[1]))

	// Check if the user role is "admin"
	if userRole != static.Admin || err != nil {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Access denied. Admins only.",
		})
	}

	// Continue to the next handler if the user is an admin
	return c.Next()
}
