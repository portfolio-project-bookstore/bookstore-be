package middleware

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const secret = "RahasiaBanget"

func CreateToken(userId int, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userId"] = userId
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour * 5).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func ExtractTokenId(e echo.Context) (uint, string) {
	users := e.Get("user").(*jwt.Token)
	if users.Valid {
		claims := users.Claims.(jwt.MapClaims)
		user_id := claims["user_id"].(float64)
		role := claims["role"].(string)
		return uint(user_id), role
	}
	return 0, ""
}

func JWTMiddleware() echo.MiddlewareFunc {
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningMethod: "HS256",
		SigningKey:    []byte(secret),
	})
}

func ValidateToken(e echo.Context) bool {
	login := e.Get("user").(*jwt.Token)

	return login.Valid
}
