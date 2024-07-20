package api

import (
	"fmt"
	"mawa3id/DB"
	auth "mawa3id/jwt"
	"strconv"
	"strings"
	"time"

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
		numOfService, err := db.GetServiceByUserID(ctx.Context(), int32(service.UserID))
		if err != nil || len(numOfService) != 0 {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": "You already have a service",
			})
		}
		err = db.CreateService(ctx.Context(), DB.CreateServiceParams{
			UserID:           service.UserID,
			Description:      service.Description,
			GoogleMapAddress: service.GoogleMapAddress,
			Willaya:          service.Willaya,
			Baladia:          service.Baladia,
			SubcategoryID:    service.SubcategoryID,
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

func GetServiceByID(db *DB.Queries) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		idStr := ctx.Params("id")
		id, err := strconv.ParseInt(idStr, 10, 32)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err,
			})
		}
		service, err := db.GetServiceByID(ctx.Context(), id)

		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err,
			})
		}
		// get the days of work for the service
		days, err := db.GetDaysOfWorkByServiceID(ctx.Context(), id)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err,
			})
		}
		// get the days of work for the service
		numberOfReservistions := int32(0)
		daysOfWork := []DaysOfWork{}
		for _, day := range days {
			daysOfWork = append(daysOfWork, DaysOfWork{
				Name:          day.Name,
				From:          day.StartTime.String(),
				To:            day.EndTime.String(),
				Limit:         day.MaxClients,
				CurrentNumber: numberOfReservistions,
				Date:          time.Now().Format("2006-01-02"),
			})
		}
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"ok":      true,
			"service": service,
			"days":    days,
		})
	}
}
