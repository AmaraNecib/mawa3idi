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
		workdays, err := db.GetWorkdaysByServiceID(ctx.Context(), int32(serviceID))
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err,
			})
		}
		if len(workdays) != 0 {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": "You already have days of work",
			})
		}
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
		startTime := time.Now().Second()
		idStr := ctx.Params("id")
		id, err := strconv.ParseInt(idStr, 10, 32)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err.Error(),
			})
		}

		Workdays, err := db.GetWorkdaysInRange(ctx.Context(), int32(id))
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"ok":    false,
				"error": err.Error(),
			})
		}
		numOfWorkDay := 0
		daysOfWork := map[string]bool{}
		for _, workday := range Workdays {
			daysOfWork[workday.Name] = workday.OpenToWork
			fmt.Println(workday.Name, workday.OpenToWork)
			if workday.OpenToWork {
				numOfWorkDay++
			}
		}
		if numOfWorkDay == 0 {
			return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
				"ok":       true,
				"workdays": []GetWorkDay{},
			})
		}
		var allWorkDays []GetWorkDay
		currentTime := time.Now().In(time.FixedZone("Algeria", 1*60*60))
		currentDay := int(currentTime.Weekday())

		workdaysInArabic := map[string]string{
			"Sunday":    "الأحد",
			"Monday":    "الأثنين",
			"Tuesday":   "الثلاثاء",
			"Wednesday": "الأربعاء",
			"Thursday":  "الخميس",
			"Friday":    "الجمعه",
			"Saturday":  "السبت",
		}
		// fmt.Println(Workdays)
		// fmt.Println(daysOfWork)
		daysProcessed := 0
		i := -(currentDay)
		for daysProcessed < 7 || i < 60 {
			if daysProcessed >= 7 || i > 60 {
				break
			}
			// workday := Workdays[i]
			// if !workdaysInArabic[time.Weekday((currentDay+daysProcessed)%7).String()] {
			// 	fmt.Println("Day is not open to work")
			// 	i++
			// 	continue
			// }

			dayIndex := (currentDay + i) % 7
			dayName := time.Weekday(dayIndex).String()
			workdayStartTime := Workdays[time.Now().Weekday()].StartTime.Format("15:04")
			workdayEndTime := Workdays[time.Now().Weekday()].EndTime.Format("15:04")
			workdayDate := currentTime.AddDate(0, 0, i)
			// fmt.Println(workdayDate)
			// if daysOfWork[workdaysInArabic[dayName]] {
			// 	fmt.Println("Day is open to work", dayName, workdaysInArabic[dayName], daysOfWork[workdaysInArabic[dayName]])
			// }
			if time.Now().Format("15:30") >= workdayEndTime && currentTime.Format("2006-01-02") >= workdayDate.Format("2006-01-02") {
				// Skip this day if the current time is beyond end time
				fmt.Println(time.Now().Format("15:30"), workdayEndTime, currentTime.Format("2006-01-02"), workdayDate.Format("2006-01-02"))
				i++
				continue
			}
			if !daysOfWork[workdaysInArabic[dayName]] {
				// fmt.Println("Day is not open to work", dayName, daysOfWork[dayName])
				i++
				continue
			}
			// fmt.Println("Day is open to work", dayName)
			// Convert StartTime and EndTime to time strings for comparison

			reservations, err := db.GetReservationsByWeekdayID(ctx.Context(), workdayDate)
			if err != nil {
				return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"ok":    false,
					"error": err.Error(),
				})
			}

			if len(reservations) == 0 {
				reservations = []DB.Reservation{}
			}
			// fmt.Println("Day is not open to work")
			allWorkDays = append(allWorkDays, GetWorkDay{
				Name:                  workdaysInArabic[dayName],
				NumberOfReservistions: len(reservations),
				Date:                  workdayDate.Format("2006-01-02"),
				MaxClients:            int(Workdays[time.Now().Weekday()].MaxClients),
				StartTime:             workdayStartTime,
				EndTime:               workdayEndTime,
				OpenToWork:            true,
			})
			daysProcessed++
			i++
		}
		fmt.Println(i)
		fmt.Println("end time: ", time.Now().Second()-startTime)
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"ok":       true,
			"workdays": allWorkDays,
		})

	}
}

func UpdateWorkDaysByID(db *DB.Queries) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var day workdays
		if err := ctx.BodyParser(&day); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err,
			})
		}
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
		if len(serviceId) == 0 || serviceID == 0 {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": "You don't have a service",
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
		// err = db.UpdateWorkdayByID(ctx.Context(), DB.UpdateWorkdayByIDParams{
		//  ID:         int64(id),
		// 	OpenToWork: day.Saturday,
		// 	StartTime:  startTime,
		// 	EndTime:    endTime,
		// 	MaxClients: int32(day.MaxClients),
		// })
		// if err != nil {
		// 	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		// 		"ok":    false,
		// 		"error": err,
		// 	})
		// }
		days, err := db.GetWorkdaysByServiceID(ctx.Context(), serviceID)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err,
			})
		}
		err = db.UpdateWorkdayByID(ctx.Context(), DB.UpdateWorkdayByIDParams{
			StartTime:  startTime,
			OpenToWork: day.Sunday,
			EndTime:    endTime,
			MaxClients: int32(day.MaxClients),
			ID:         int64(days[0].ID),
		})
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": fmt.Sprintf("something went wrong: %v", err),
			})
		}
		err = db.UpdateWorkdayByID(ctx.Context(), DB.UpdateWorkdayByIDParams{
			StartTime:  startTime,
			OpenToWork: day.Monday,
			EndTime:    endTime,
			MaxClients: int32(day.MaxClients),
			ID:         int64(days[1].ID),
		})
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": fmt.Sprintf("something went wrong: %v", err),
			})
		}
		err = db.UpdateWorkdayByID(ctx.Context(), DB.UpdateWorkdayByIDParams{
			StartTime:  startTime,
			OpenToWork: day.Tuesday,
			EndTime:    endTime,
			MaxClients: int32(day.MaxClients),
			ID:         int64(days[2].ID),
		})
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": fmt.Sprintf("something went wrong: %v", err),
			})
		}
		err = db.UpdateWorkdayByID(ctx.Context(), DB.UpdateWorkdayByIDParams{
			StartTime:  startTime,
			OpenToWork: day.Wednesday,
			EndTime:    endTime,
			MaxClients: int32(day.MaxClients),
			ID:         int64(days[3].ID),
		})
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": fmt.Sprintf("something went wrong: %v", err),
			})
		}
		err = db.UpdateWorkdayByID(ctx.Context(), DB.UpdateWorkdayByIDParams{
			StartTime:  startTime,
			OpenToWork: day.Thursday,
			EndTime:    endTime,
			MaxClients: int32(day.MaxClients),
			ID:         int64(days[4].ID),
		})
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": fmt.Sprintf("something went wrong: %v", err),
			})
		}
		err = db.UpdateWorkdayByID(ctx.Context(), DB.UpdateWorkdayByIDParams{
			StartTime:  startTime,
			OpenToWork: day.Friday,
			EndTime:    endTime,
			MaxClients: int32(day.MaxClients),
			ID:         int64(days[5].ID),
		})
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": fmt.Sprintf("something went wrong: %v", err),
			})
		}
		err = db.UpdateWorkdayByID(ctx.Context(), DB.UpdateWorkdayByIDParams{
			StartTime:  startTime,
			OpenToWork: day.Saturday,
			EndTime:    endTime,
			MaxClients: int32(day.MaxClients),
			ID:         int64(days[6].ID),
		})
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": fmt.Sprintf("something went wrong: %v", err),
			})
		}
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"ok":      true,
			"message": "Workdays updated successfully",
		})
	}
}

func UpdateAllWorkDays(db *DB.Queries) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var day workdays
		if err := ctx.BodyParser(&day); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err,
			})
		}
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
		if len(serviceId) == 0 || serviceID == 0 {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": "You don't have a service",
			})
		}

		Workdays, err := db.GetWorkdaysInRange(ctx.Context(), int32(serviceID))
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"ok":    false,
				"error": err.Error(),
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
			StartTime:  startTime,
			OpenToWork: day.Saturday,
			EndTime:    endTime,
			MaxClients: int32(day.MaxClients),
			ID:         int64(Workdays[0].ID),
		})
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": "Something went wrong when updating the workday",
			})
		}
		err = db.UpdateWorkdayByID(ctx.Context(), DB.UpdateWorkdayByIDParams{
			StartTime:  startTime,
			OpenToWork: day.Sunday,
			EndTime:    endTime,
			MaxClients: int32(day.MaxClients),
			ID:         int64(Workdays[1].ID),
		})
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": "Something went wrong when updating the workday",
			})
		}
		err = db.UpdateWorkdayByID(ctx.Context(), DB.UpdateWorkdayByIDParams{
			StartTime:  startTime,
			OpenToWork: day.Monday,
			EndTime:    endTime,
			MaxClients: int32(day.MaxClients),
			ID:         int64(Workdays[2].ID),
		})
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": "Something went wrong when updating the workday",
			})
		}
		err = db.UpdateWorkdayByID(ctx.Context(), DB.UpdateWorkdayByIDParams{
			StartTime:  startTime,
			OpenToWork: day.Tuesday,
			EndTime:    endTime,
			MaxClients: int32(day.MaxClients),
			ID:         int64(Workdays[3].ID),
		})
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": "Something went wrong when updating the workday",
			})
		}
		err = db.UpdateWorkdayByID(ctx.Context(), DB.UpdateWorkdayByIDParams{
			StartTime:  startTime,
			OpenToWork: day.Wednesday,
			EndTime:    endTime,
			MaxClients: int32(day.MaxClients),
			ID:         int64(Workdays[4].ID),
		})
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": "Something went wrong when updating the workday",
			})
		}
		err = db.UpdateWorkdayByID(ctx.Context(), DB.UpdateWorkdayByIDParams{
			StartTime:  startTime,
			OpenToWork: day.Thursday,
			EndTime:    endTime,
			MaxClients: int32(day.MaxClients),
			ID:         int64(Workdays[5].ID),
		})
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": "Something went wrong when updating the workday",
			})
		}
		err = db.UpdateWorkdayByID(ctx.Context(), DB.UpdateWorkdayByIDParams{
			StartTime:  startTime,
			OpenToWork: day.Friday,
			EndTime:    endTime,
			MaxClients: int32(day.MaxClients),
			ID:         int64(Workdays[6].ID),
		})
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": "Something went wrong when updating the workday",
			})
		}
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"ok":      true,
			"message": "Workdays updated successfully",
		})
	}
}
