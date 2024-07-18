package auth

import (
	"fmt"
	"mawa3id/static"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func encrypt(payload Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return tokenString, nil
}
func decrypt(input string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(input, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		fmt.Println("hi", err)
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}

func GetUserRole(token string) (string, error) {
	payload, err := decrypt(token)
	if err != nil {
		return "", err
	}
	return payload.Role, nil
}

func ValidToken(token string) bool {
	claims, err := decrypt(token)
	fmt.Println(claims)
	if err != nil || claims.Exp <= time.Now().Unix() {
		return false
	}
	return true
}

func CreateToken(id int, role string) (string, error) {
	claims := Claims{
		ID:   id,
		Role: role,
		Exp:  time.Now().Add(time.Hour * 24 * time.Duration(static.ValidTokenDays)).Unix(),
	}
	return encrypt(claims)
}

// Protected protect routes
func GetUserID(token string) int {
	payload, err := decrypt(token)
	if err != nil {
		return 0
	}
	return payload.ID
}
