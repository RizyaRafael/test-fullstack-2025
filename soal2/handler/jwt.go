package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func SignToken(email string, c *fiber.Ctx) (string, error) {
	jwtSecret := "SECRET_ENV"
	claims := jwt.MapClaims{
		"email": email,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(jwtSecret))

	if err != nil {
		return "", c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	"message": "error in jwt sign",
})
	}
	return signedToken, nil
}
