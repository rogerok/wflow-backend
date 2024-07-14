package router

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/rogerok/wflow-backend/configs"
)

func SetupRouter(app *fiber.App) {
	db, dbError := configs.ConnectToDb()

	if dbError != nil {
		fmt.Printf("Could connect to database %s=", dbError.Error())
		return
	}

	defer configs.CloseConnectionToDb(db)

	app.Use(logger.New())

	//apiv1 := app.Group("/api")
	//apiv1.Get("/user", handlers.HandleGetUsers)
	//apiv1.Get("/user/:id", handlers.HandleGetUser)
}
