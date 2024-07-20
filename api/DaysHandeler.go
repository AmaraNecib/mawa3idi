package api

import (
	"fmt"
	"mawa3id/DB"
	auth "mawa3id/jwt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateWorkDays(db *DB.Queries) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var day workdays
		var err error
		if err = ctx.BodyParser(&day); err != nil {
			fmt.Println(err)
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err,
			})
		}
		token := strings.Split(string(ctx.Get("Authorization")), " ")[1]
		UserID := int32(auth.GetUserID(string(token)))
		if UserID == 0 {
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

		service, err := db.GetServiceByUserID(ctx.Context(), int32(UserID))
		fmt.Println(service, UserID)
		if err != nil || len(service) == 0 {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": "You don't have a service",
			})
		}
		// add the days of work
		serviceID := service[0].ID
		// get all days of the week
		// for each day check if it is in the days of work
		days, err := db.GetAllDays(ctx.Context())
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": fmt.Sprintf("something went wrong: %v", err),
			})
		}
		if day.Saturday {

			startTime, err := time.Parse("15:04", day.StartTime)
			if err != nil {
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"ok":    false,
					"error": fmt.Sprintf("something went wrong: %v", err),
				})
			}
			endTime, err := time.Parse("15:04", day.EndTime)
			if err != nil {
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"ok":    false,
					"error": fmt.Sprintf("something went wrong: %v", err),
				})
			}

			err = db.CreateWorkday(ctx.Context(), DB.CreateWorkdayParams{
				ServiceID:  int32(serviceID),
				Name:       days[0].Name,
				StartTime:  startTime,
				EndTime:    endTime,
				MaxClients: int32(day.MaxClients),
				DayID:      int32(days[0].ID),
			})
			if err != nil {
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"ok":    false,
					"error": fmt.Sprintf("something went wrong: %v", err),
				})
			}
		}

		if day.Sunday {
			startTime, err := time.Parse("15:04", day.StartTime)
			if err != nil {
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"ok":    false,
					"error": fmt.Sprintf("something went wrong: %v", err),
				})
			}
			endTime, err := time.Parse("15:04", day.EndTime)
			if err != nil {
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"ok":    false,
					"error": fmt.Sprintf("something went wrong: %v", err),
				})
			}

			err = db.CreateWorkday(ctx.Context(), DB.CreateWorkdayParams{
				ServiceID:  int32(serviceID),
				Name:       days[1].Name,
				StartTime:  startTime,
				EndTime:    endTime,
				MaxClients: int32(day.MaxClients),
				DayID:      int32(days[1].ID),
			})
			if err != nil {
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"ok":    false,
					"error": fmt.Sprintf("something went wrong: %v", err),
				})
			}
		}
		if day.Monday {
			startTime, err := time.Parse("15:04", day.StartTime)
			if err != nil {
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"ok":    false,
					"error": fmt.Sprintf("something went wrong: %v", err),
				})
			}
			endTime, err := time.Parse("15:04", day.EndTime)
			if err != nil {
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"ok":    false,
					"error": fmt.Sprintf("something went wrong: %v", err),
				})
			}
			err = db.CreateWorkday(ctx.Context(), DB.CreateWorkdayParams{
				ServiceID:  int32(serviceID),
				Name:       days[2].Name,
				StartTime:  startTime,
				EndTime:    endTime,
				MaxClients: int32(day.MaxClients),
				DayID:      int32(days[2].ID),
			})
			if err != nil {
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"ok":    false,
					"error": fmt.Sprintf("something went wrong: %v", err),
				})
			}
		}
		if day.Tuesday {
			startTime, err := time.Parse("15:04", day.StartTime)
			if err != nil {
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"ok":    false,
					"error": fmt.Sprintf("something went wrong: %v", err),
				})
			}
			endTime, err := time.Parse("15:04", day.EndTime)
			if err != nil {
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"ok":    false,
					"error": fmt.Sprintf("something went wrong: %v", err),
				})
			}
			err = db.CreateWorkday(ctx.Context(), DB.CreateWorkdayParams{
				ServiceID:  int32(serviceID),
				Name:       days[3].Name,
				StartTime:  startTime,
				EndTime:    endTime,
				MaxClients: int32(day.MaxClients),
				DayID:      int32(days[3].ID),
			})
			if err != nil {
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"ok":    false,
					"error": fmt.Sprintf("something went wrong: %v", err),
				})
			}
		}
		if day.Wednesday {
			startTime, err := time.Parse("15:04", day.StartTime)
			if err != nil {
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"ok":    false,
					"error": fmt.Sprintf("something went wrong: %v", err),
				})
			}
			endTime, err := time.Parse("15:04", day.EndTime)
			if err != nil {
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"ok":    false,
					"error": fmt.Sprintf("something went wrong: %v", err),
				})
			}
			err = db.CreateWorkday(ctx.Context(), DB.CreateWorkdayParams{
				ServiceID:  int32(serviceID),
				Name:       days[4].Name,
				StartTime:  startTime,
				EndTime:    endTime,
				MaxClients: int32(day.MaxClients),
				DayID:      int32(days[4].ID),
			})
			if err != nil {
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"ok":    false,
					"error": fmt.Sprintf("something went wrong: %v", err),
				})
			}
		}
		if day.Thursday {
			startTime, err := time.Parse("15:04", day.StartTime)
			if err != nil {
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"ok":    false,
					"error": fmt.Sprintf("something went wrong: %v", err),
				})
			}
			endTime, err := time.Parse("15:04", day.EndTime)
			if err != nil {
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"ok":    false,
					"error": fmt.Sprintf("something went wrong: %v", err),
				})
			}
			err = db.CreateWorkday(ctx.Context(), DB.CreateWorkdayParams{
				ServiceID:  int32(serviceID),
				Name:       days[5].Name,
				StartTime:  startTime,
				EndTime:    endTime,
				MaxClients: int32(day.MaxClients),
				DayID:      int32(days[5].ID),
			})
			if err != nil {
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"ok":    false,
					"error": fmt.Sprintf("something went wrong: %v", err),
				})
			}
		}
		if day.Friday {
			startTime, err := time.Parse("15:04", day.StartTime)
			if err != nil {
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"ok":    false,
					"error": fmt.Sprintf("something went wrong: %v", err),
				})
			}
			endTime, err := time.Parse("15:04", day.EndTime)
			if err != nil {
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"ok":    false,
					"error": fmt.Sprintf("something went wrong: %v", err),
				})
			}
			err = db.CreateWorkday(ctx.Context(), DB.CreateWorkdayParams{
				ServiceID:  int32(serviceID),
				Name:       days[6].Name,
				StartTime:  startTime,
				EndTime:    endTime,
				MaxClients: int32(day.MaxClients),
				DayID:      int32(days[6].ID),
			})
			if err != nil {
				return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"ok":    false,
					"error": fmt.Sprintf("something went wrong: %v", err),
				})
			}
		}

		response := fiber.Map{
			"ok":      true,
			"message": "Days of work created successfully",
		}
		return ctx.JSON(response)
	}
}

func GetAllWorkDays(db *DB.Queries) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		userID := int32(auth.GetUserID(strings.Split(string(ctx.Get("Authorization")), " ")[1]))
		if userID == 0 {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": "Unauthorized",
			})
		}
		serviceId, err := db.GetServiceByUserID(ctx.Context(), userID)
		if err != nil || len(serviceId) == 0 {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": "You don't have a service",
			})
		}
		serviceID := int32(serviceId[0].ID)
		if err != nil || len(serviceId) == 0 || serviceID == 0 {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": "You don't have a service",
			})
		}
		workdays, err := db.GetWorkdaysByServiceID(ctx.Context(), int32(serviceID))
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err,
			})
		}
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"ok":       true,
			"workdays": workdays,
		})
	}
}
