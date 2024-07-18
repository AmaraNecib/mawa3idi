package api

import (
	"fmt"
	"mawa3id/DB"
	auth "mawa3id/jwt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func CreateServices(db *DB.Queries) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var service DB.Service
		var err error
		if err = ctx.BodyParser(&service); err != nil {
			fmt.Println(err)
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err,
			})
		}
		token := strings.Split(string(ctx.Get("Authorization")), " ")[1]
		service.UserID = int32(auth.GetUserID(string(token)))
		if service.UserID == 0 {
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
		err = db.CreateService(ctx.Context(), DB.CreateServiceParams{
			UserID:           service.UserID,
			Description:      service.Description,
			GoogleMapAddress: service.GoogleMapAddress,
			Willaya:          service.Willaya,
			Baladia:          service.Baladia,
		})
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": fmt.Sprintf("something went wrong: %v", err),
			})
		}
		response := fiber.Map{
			"ok":   true,
			"user": service,
		}
		return ctx.Status(fiber.StatusCreated).JSON(response)
	}
}

type Service struct {
	UserID           int32  `json:"user_id"`
	Description      string `json:"description"`
	GoogleMapAddress string `json:"google_map_address"`
	Willaya          string `json:"willaya"`
	Baladia          string `json:"baladia"`
}

func GetAllServices(db *DB.Queries) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		services, err := db.GetServices(ctx.Context())
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err,
			})
		}
		if len(services) == 0 {
			services = []DB.Service{}
		}
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"ok":       true,
			"services": services,
		})
	}
}
