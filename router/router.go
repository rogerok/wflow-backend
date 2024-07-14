package router

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rogerok/wflow-backend/configs"
	"github.com/rogerok/wflow-backend/handlers"
	"github.com/rogerok/wflow-backend/repositories"
	"github.com/rogerok/wflow-backend/services"
)

func SetupRouter(app *fiber.App) {
	db, dbError := configs.ConnectToDb()

	if dbError != nil {
		fmt.Printf("Could connect to database %s=", dbError.Error())
		return
	}

	defer configs.CloseConnectionToDb(db)

	app.Use(logger.New())

	api := app.Group("/api")

	user := api.Group("/user")
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	user.Get("/user", handlers.UsersList(userService))

	//api.Get("/user/:id", handlers.UserHandler.)
}
