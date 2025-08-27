package controllers

import (
	"app/model"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	var user model.User
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
	

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": "You've succesfully registered",
	})
}
