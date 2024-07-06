package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/rogerok/wflow-backend/configs"
	"os"
)

func main() {

	envErr := godotenv.Load()

	if envErr != nil {
		fmt.Printf("Could not load environment")
	}

	db, err := configs.ConnectToDb()

	defer db.Close(context.Background())

	app := fiber.New()

	err = app.Listen(":" + os.Getenv("PORT"))

	if err != nil {
		panic(err.Error())
	}

}
