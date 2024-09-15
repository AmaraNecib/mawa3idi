package api

import (
	"fmt"
	"mawa3id/DB"
	auth "mawa3id/jwt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func CreateReservation(db *DB.Queries) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var reservation MyReservation
		if err := c.BodyParser(&reservation); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err.Error(),
			})
		}
		fmt.Println(reservation.ServiceID)
		timeObj, err := time.Parse("2006-01-02", reservation.Time)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err.Error(),
			})
		}
		// get the number of the day
		dayId := timeObj.Weekday()
		res, err := db.GetWorkdaysByServiceID(c.Context(), int32(reservation.ServiceID))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err.Error(),
			})
		}
		weekdayId := 0 // get the weekday id
		for _, day := range res {
			if int32(dayId) == day.DayID {
				weekdayId = int(day.ID)
			}
		}
		if weekdayId == 0 || !res[dayId].OpenToWork {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": "The day is not open",
			})
		}
		token := strings.Split(string(c.Get("Authorization")), " ")[1]
		UserID := int32(auth.GetUserID(string(token)))
		// check if the date is matched with the day and the day is open or not
		if dayId == 0 || reservation.ServiceID == 0 || UserID == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": "DayID, ServiceID and UserID are required",
			})
		}
		myReservations, err := db.GetReservationsCountByUserIdAndServiceId(c.Context(), DB.GetReservationsCountByUserIdAndServiceIdParams{
			UserID:    UserID,
			ServiceID: int32(reservation.ServiceID),
			Time:      timeObj,
			WeekdayID: int32(weekdayId),
		})
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err.Error(),
			})
		}
		fmt.Println(myReservations)
		if myReservations >= 1 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": "You already have a reservation",
			})
		}

		// get the reservation count
		count, err := db.GetReservationsCount(c.Context(), DB.GetReservationsCountParams{
			WeekdayID: int32(weekdayId),
			ServiceID: int32(reservation.ServiceID),
			Time:      timeObj,
		})
		if err != nil {
			fmt.Println(err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err.Error(),
			})

		}
		if count >= int64(res[dayId].MaxClients) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": "The day is full",
			})
		}
		fmt.Println(count)
		myReservation, err := db.CreateReservation(c.Context(), DB.CreateReservationParams{
			UserID:    UserID,
			WeekdayID: int32(weekdayId),
			ServiceID: int32(reservation.ServiceID),
			Time:      timeObj,
			Ranking:   int32(count + 1),
		})
		// res, err = db.CreateReservation(c.Context(), DB.CreateReservationParams{
		// 	UserID:      UserID,
		// 	WeekdayID:   int32(weekdayId),
		// 	ServiceID:   int32(reservation.ServiceID),
		// 	Reservation: timeObj,
		// })
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"ok":          true,
			"reservation": myReservation,
		})
	}
}

func UpdateReservationByID(db *DB.Queries) func(*fiber.Ctx) error {
	// get only reserv_status and update it
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": "Invalid ID",
			})
		}
		var reservation UpdateReservationStatus
		if err := c.BodyParser(&reservation); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err.Error(),
			})
		}
		_, err = db.UpdateReservationStatusByID(c.Context(), DB.UpdateReservationStatusByIDParams{
			ID:           int64(id),
			ReservStatus: int32(reservation.Status),
		})
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"ok":  true,
			"msg": "Reservation status updated",
		})
	}
}

func ReservationCompleted(db *DB.Queries) func(*fiber.Ctx) error {
	// get only reserv_status and update it
	return func(c *fiber.Ctx) error {
		role, err := auth.GetUserRole(strings.Split(c.Get("Authorization"), " ")[1])
		if err != nil || role != "admin" || role != "customer" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": "Unauthorized",
			})
		}
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": "Invalid ID",
			})
		}
		var reservation UpdateReservationStatus
		if err := c.BodyParser(&reservation); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err.Error(),
			})
		}
		_, err = db.UpdateReservationStatusByID(c.Context(), DB.UpdateReservationStatusByIDParams{
			ID:           int64(id),
			ReservStatus: int32(reservation.Status),
		})
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err.Error(),
			})
		}
		// send notification to the next 4 users by fcm
		// get the next 4 users
		resInfo, err := db.GetReservationInfoByID(c.Context(), int64(id))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err.Error(),
			})
		}
		// get the next 4 users
		nextUsers, err := db.GetNextUserReservations(c.Context(), DB.GetNextUserReservationsParams{
			ServiceID: resInfo.ServiceID,
			WeekdayID: resInfo.WeekdayID,
			ID:        int64(id),
			Limit:     4,
		})
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err.Error(),
			})
		}
		// send notification to the next 4 users
		for _, user := range nextUsers {
			fmt.Println(user.UserID)
			// send notification to the user
			// get the user token
			// userToken, err := db.GetUserToken(c.Context(), user.UserID)
			// if err != nil {
			// 	return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			// 		"ok":    false,
			// 		"error": err.Error(),
			// 	})
			// }
			// send notification to the user
			// send notification to the user
			// send notification to the user
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"ok": true,
		})
	}
}
