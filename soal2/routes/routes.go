package routes

import (
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {

	user := app.Group("/user")
	UserRouter(user)
}
 