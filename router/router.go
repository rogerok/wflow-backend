package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jmoiron/sqlx"
	"github.com/rogerok/wflow-backend/handlers"
	"github.com/rogerok/wflow-backend/repositories"
	"github.com/rogerok/wflow-backend/services"
)

func SetupRouter(app *fiber.App, db *sqlx.DB) {

	app.Use(logger.New())

	api := app.Group("/api")

	user := api.Group("/user")
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)

	user.Get("/", handlers.UsersList(userService))
	user.Get("/:id", handlers.UserById(userService))

}
