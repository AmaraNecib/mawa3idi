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
		timeObj, err := time.Parse("2006-01-02", reservation.Time)
		// check if the day is the day of request or after
		if err != nil || timeObj.Before(time.Now().Truncate(24*time.Hour)) {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": "Invalid date. Reservations can only be made for today or future dates.",
			})
		}
		// get the number of the day
		dayId := timeObj.Weekday()
		fmt.Println("the day is ", dayId, int32(dayId))
		res, err := db.GetWorkdaysByServiceID(c.Context(), int32(reservation.ServiceID))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err.Error(),
			})
		}

		fmt.Println("week day", dayId, res[(dayId+1)%7].Name, res[(dayId+1)%7].ID, "open to work", res[(dayId+1)%7].OpenToWork, "day id", res[(dayId+1)%7].ID)
		if !res[(dayId+1)%7].OpenToWork {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": "The day is not open",
				"day":   res[(dayId+1)%7].Name,
				"res":   res,
			})
		}
		token := strings.Split(string(c.Get("Authorization")), " ")[1]
		UserID := int32(auth.GetUserID(string(token)))
		// check if the date is matched with the day and the day is open or not
		if reservation.ServiceID == 0 || UserID == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": "DayID, ServiceID and UserID are required",
			})
		}
		myReservations, err := db.GetReservationsCountByUserIdAndServiceId(c.Context(), DB.GetReservationsCountByUserIdAndServiceIdParams{
			UserID:    UserID,
			ServiceID: int32(reservation.ServiceID),
			Time:      timeObj,
			WeekdayID: int32(res[(dayId+1)%7].ID),
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
			WeekdayID: int32(res[(dayId+1)%7].ID),
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
			WeekdayID: int32(res[(dayId+1)%7].ID),
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
		if err != nil || (role != "customer" && role != "admin") {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": "Unauthorized",
				"role":  role,
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
				"msg":   "Invalid body",
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
				"msg":   "Error updating reservation status",
			})
		}
		// send notification to the next 4 users by fcm
		// get the next 4 users
		resInfo, err := db.GetReservationInfoByID(c.Context(), int64(id))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err.Error(),
				"msg":   "Error getting reservation info",
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
				"msg":   "Error getting next users",
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

func GetAllReservationsByUserID(db *DB.Queries) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		token := strings.Split(string(c.Get("Authorization")), " ")[1]
		UserID := int32(auth.GetUserID(string(token)))
		if UserID == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": "Invalid token",
			})
		}
		limit := c.QueryInt("limit")
		if limit == 0 {
			limit = 10
		}
		page := c.QueryInt("page")
		if page == 0 {
			page = 1
		}
		reservations, err := db.GetReservationsByUserID(c.Context(), DB.GetReservationsByUserIDParams{
			UserID: int32(UserID),
			Limit:  int32(limit),
			Offset: int32((page - 1) * limit),
		})
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"ok":           true,
			"reservations": reservations,
		})
	}
}

func GetReservationByID(db *DB.Queries) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": "Invalid ID",
			})
		}
		reservation, err := db.GetReservationByID(c.Context(), int64(id))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"ok":          true,
			"reservation": reservation,
		})
	}
}

func GetReservationsByServiceID(db *DB.Queries) func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": "Invalid ID",
			})
		}
		limit := c.QueryInt("limit")
		if limit == 0 {
			limit = 10
		}
		page := c.QueryInt("page")
		if page == 0 {
			page = 1
		}
		reservations, err := db.GetReservationsByServiceID(c.Context(), DB.GetReservationsByServiceIDParams{
			ServiceID: int32(id),
			Limit:     int32(limit),
			Offset:    int32((page - 1) * limit),
		})
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err.Error(),
			})
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"ok":           true,
			"reservations": reservations,
		})
	}
}
