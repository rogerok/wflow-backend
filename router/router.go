package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rogerok/wflow-backend/api"

	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRouter(app *fiber.App) {
	app.Use(logger.New())

	apiv1 := app.Group("/api")
	apiv1.Get("/user", api.HandleGetUsers)
	apiv1.Get("/user/:id", api.HandleGetUser)
}
