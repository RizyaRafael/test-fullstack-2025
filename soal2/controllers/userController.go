package controllers

import (
	"app/handler"
	"app/model"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	var user model.User
	var foundUser model.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid body request",
		})
	}

	if user.Email == "" || user.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "email and password required",
		})
	}

	// find the email inside the database
	foundUser.Password = "samplePassword"

	// check if the unhashed password is the same as the inputed password
	checkPass := handler.ComparePass(user.Password, foundUser.Password)
	if checkPass != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "invalid email or password",
		})
	} else {
		token, err := handler.SignToken(foundUser.Email, c)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"message": "error in jwt sign",
			})
		}

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"access_token": "Bearer " + token,
		})
	}
}
