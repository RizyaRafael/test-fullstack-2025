package routes

import (
	"app/controllers"

	"github.com/gofiber/fiber/v2"
)

func UserRouter(app fiber.Router) {
	app.Post("/login", controllers.Login)
}