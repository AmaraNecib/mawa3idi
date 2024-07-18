package api

import (
	"fmt"
	"mawa3id/DB"
	auth "mawa3id/jwt"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(db *DB.Queries) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		var user DB.User
		if err := ctx.BodyParser(&user); err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err,
			})
		}
		hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": "something went wrong",
			})
		}
		res, err := db.CreateUser(ctx.Context(), DB.CreateUserParams{
			FirstName:   user.FirstName,
			LastName:    user.LastName,
			Email:       user.Email,
			Password:    string(hash),
			PhoneNumber: user.PhoneNumber,
			RoleID:      user.RoleID,
		})
		if err != nil {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": fmt.Sprintf("something went wrong: %v", err),
			})
		}
		user.ID = res.ID
		response := fiber.Map{
			"ok":   true,
			"user": user,
		}
		return ctx.Status(fiber.StatusCreated).JSON(response)
	}
}

func GetAllUsers(db *DB.Queries) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		users, err := db.GetUsers(ctx.Context())
		if err != nil {
			ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err,
			})
		}
		if len(users) == 0 {
			users = []DB.User{}
		}
		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"ok":    true,
			"users": users,
		})
	}
}

func login(db *DB.Queries) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var Login Login
		if err := c.BodyParser(&Login); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": err,
			})
		}
		// Throws Unauthorized error
		// if email != "john" || pass != "doe" {
		// 	return c.SendStatus(fiber.StatusUnauthorized)
		// }
		// get user if by email
		// get user by email

		res, err := db.GetUserByEmail(c.Context(), Login.Email)

		// Create the Claims
		if err != nil {
			fmt.Println("heress ", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"ok":    false,
				"error": "Invalid Credentials"},
			)
		}
		err = bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(Login.Password))
		// claims := auth.Claims{
		// 	"id":   res.ID,
		// 	"role": res.RoleName,
		// 	"exp":  time.Now().Add(time.Hour * 24 * 30).Unix(),
		// }
		if err != nil {
			fmt.Println("hello ", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"ok":    false,
				"error": "Invalid Credentials"},
			)
		}
		// Create token
		token, err := auth.CreateToken(int(res.ID), res.RoleName)
		if err != nil {
			fmt.Println("here ", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"ok":    false,
				"error": "Invalid Credentials",
			})
		}
		fmt.Println(res.RoleName)

		return c.Status(fiber.StatusAccepted).JSON(fiber.Map{
			"ok":    true,
			"token": token})
	}
}
