package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/rogerok/wflow-backend/handlers"

	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRouter(app *fiber.App) {
	app.Use(logger.New())

	apiv1 := app.Group("/api")
	apiv1.Get("/user", handlers.HandleGetUsers)
	apiv1.Get("/user/:id", handlers.HandleGetUser)
}
