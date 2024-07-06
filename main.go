package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/rogerok/wflow-backend/api"
	"os"
)

func main() {

	envErr := godotenv.Load()

	if envErr != nil {
		fmt.Printf("Could not load environment")
	}

	db, err := connectToDb(dbConfig())

	db.Ping(context.Background())

	defer db.Close(context.Background())

	app := fiber.New()

	apiv1 := app.Group("/api")

	apiv1.Get("/user", api.HandleGetUsers)
	apiv1.Get("/user/:id", api.HandleGetUser)

	err = app.Listen(":" + os.Getenv("PORT"))

	if err != nil {
		panic(err.Error())
	}

}
