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
		fmt.Println(service, int32(UserID))
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
		fmt.Println(days)
		if err != nil || len(days) == 0 {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": fmt.Sprintf("something went wrong: %v", err),
			})
		}
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
			OpenToWork: day.Saturday,
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

		err = db.CreateWorkday(ctx.Context(), DB.CreateWorkdayParams{
			ServiceID:  int32(serviceID),
			Name:       days[1].Name,
			OpenToWork: day.Sunday,
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

		err = db.CreateWorkday(ctx.Context(), DB.CreateWorkdayParams{
			ServiceID:  int32(serviceID),
			Name:       days[2].Name,
			StartTime:  startTime,
			OpenToWork: day.Monday,
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

		err = db.CreateWorkday(ctx.Context(), DB.CreateWorkdayParams{
			ServiceID:  int32(serviceID),
			Name:       days[3].Name,
			StartTime:  startTime,
			EndTime:    endTime,
			OpenToWork: day.Tuesday,
			MaxClients: int32(day.MaxClients),
			DayID:      int32(days[3].ID),
		})
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
			OpenToWork: day.Wednesday,
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

		err = db.CreateWorkday(ctx.Context(), DB.CreateWorkdayParams{
			ServiceID:  int32(serviceID),
			Name:       days[5].Name,
			StartTime:  startTime,
			OpenToWork: day.Thursday,
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
		err = db.CreateWorkday(ctx.Context(), DB.CreateWorkdayParams{
			ServiceID:  int32(serviceID),
			Name:       days[6].Name,
			StartTime:  startTime,
			OpenToWork: day.Friday,
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

// func GetWorkDaysByID(db *DB.Queries) fiber.Handler {
// 	return func(ctx *fiber.Ctx) error {
// 		idStr := ctx.Params("id")
// 		id, err := strconv.ParseInt(idStr, 10, 32)
// 		if err != nil {
// 			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 				"ok":    false,
// 				"error": err,
// 			})
// 		}
// 		workday, err := db.GetWorksdayByID(ctx.Context(), id)
// 		if err != nil {
// 			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 				"ok":    false,
// 				"error": err,
// 			})
// 		}
// 		workdays := []DB.Weekday{workday}
// 		if err != nil {
// 			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 				"ok":    false,
// 				"error": err,
// 			})
// 		}
// 		// the arabic name of the day

// 		today := time.Now().Weekday().String()

// 		allWorkDays := make([]GetWorkDay, 0)
// 		for _, workday := range workdays {
// 			allWorkDays = append(allWorkDays, GetWorkDay{
// 				Name:                  workday.Name,
// 				NumberOfReservistions: workday.MaxClients,
// 				Date:                  today,
// 				StartTime:             workday.StartTime.String(),
// 				EndTime:               workday.EndTime.String(),
// 			})
// 		}

// 		// make the workdays in this structure to be able to send it to the client
// 		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
// 			"ok":       true,
// 			"workdays": allWorkDays,
// 		})
// 	}
// }

func GetWorkDaysByID(db *DB.Queries) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		idStr := ctx.Params("id")
		id, err := strconv.ParseInt(idStr, 10, 32)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err.Error(),
			})
		}

		// Calculate the start and end dates for the next 7 workdays
		// startDate := time.Now()

		weekdays, err := db.GetWeekdaysInRange(ctx.Context(), int32(id))
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"ok":    false,
				"error": err.Error(),
			})
		}
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": fmt.Sprintf("something went wrong: %v", err),
			})
		}

		weekdays, err = db.GetWeekdaysInRange(ctx.Context(), int32(id))
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"ok":    false,
				"error": err.Error(),
			})
		}

		allWorkDays := make([]GetWorkDay, 0)
		today := time.Now()
		daysProcessed := 0
		for daysProcessed < 7 {
			for _, workday := range weekdays {
				if daysProcessed >= 7 {
					break
				}
				if !workday.OpenToWork {
					continue
				}

				workdayDate := today.AddDate(0, 0, int(workday.DayID))
				reservations, err := db.GetReservationsByWeekdayID(ctx.Context(), workdayDate)
				if err != nil {
					return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
						"ok":    false,
						"error": err.Error(),
					})
				}
				fmt.Println("reservations", reservations, len(reservations))
				if len(reservations) == 0 {
					reservations = []DB.Reservation{}
				}
				allWorkDays = append(allWorkDays, GetWorkDay{
					Name: workday.Name,
					// NumberOfReservations: len(reservations),
					NumberOfReservistions: int(len(reservations)),
					Date:                  workdayDate.Format("2006-01-02"),
					MaxClients:            int(workday.MaxClients),
					StartTime:             workday.StartTime.Format("15:04"),
					EndTime:               workday.EndTime.Format("15:04"),
					OpenToWork:            workday.OpenToWork,
				})
				daysProcessed++
			}

		}
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"ok":       true,
			"workdays": allWorkDays,
			"day":      daysProcessed,
		})
	}
}

func UpdateWorkDaysByID(db *DB.Queries) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		idStr := ctx.Params("id")
		id, err := strconv.ParseInt(idStr, 10, 32)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err,
			})
		}
		var day workdays
		if err = ctx.BodyParser(&day); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err,
			})
		}
		startTime, err := time.Parse("15:04", day.StartTime)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err,
			})
		}
		endTime, err := time.Parse("15:04", day.EndTime)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err,
			})
		}
		err = db.UpdateWorkdayByID(ctx.Context(), DB.UpdateWorkdayByIDParams{
			ID:         int64(id),
			OpenToWork: day.Saturday,
			StartTime:  startTime,
			EndTime:    endTime,
			MaxClients: int32(day.MaxClients),
		})
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err,
			})
		}
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"ok": true,
		})
	}
}
