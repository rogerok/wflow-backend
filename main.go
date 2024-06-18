package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/rogerok/wflow-backend/api"
	"os"
)

func main() {

	envErr := godotenv.Load(".env.dev")

	if envErr != nil {
		fmt.Printf("Could not load environment")
	}

	port := os.Getenv("PORT")

	app := fiber.New()

	apiv1 := app.Group("/api")

	app.Get("/foo", func(c *fiber.Ctx) error {
		return c.JSON(map[string]string{"go": "mod"})
	})
	apiv1.Get("/user", api.HandleGetUsers)
	apiv1.Get("/user/:id", api.HandleGetUser)

	err := app.Listen(":" + port)

	if err != nil {
		fmt.Printf(err.Error())
	}

}
