package middlewares

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const secretjwt = "S3CR3T"

func JWTMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte(secretjwt),
	})
}

//fungsi untuk men generate token
func CreateToken(id int, name string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["id"] = id
	claims["name"] = name
	claims["exp"] = time.Now().Add(time.Hour * 200).Unix() //Token expires after 200 hour
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretjwt))
}

func ExtractToken(c echo.Context) (int, error) {
	loginToken := c.Get("user").(*jwt.Token)
	if loginToken.Valid {
		claims := loginToken.Claims.(jwt.MapClaims)
		id := int(claims["id"].(float64))
		return id, nil
	}
	return -1, fmt.Errorf("unauthorized")
}
