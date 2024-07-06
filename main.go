package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/rogerok/wflow-backend/configs"
	"github.com/rogerok/wflow-backend/router"
	"os"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Could not load environment %s", err.Error())
		return
	}

	db, dbError := configs.ConnectToDb()

	if dbError != nil {
		fmt.Printf("Could connect to database %s=", dbError.Error())
		return
	}

	defer configs.CloseConnectionToDb(db, context.Background())

	app := fiber.New()
	router.SetupRouter(app)
	err = app.Listen(":" + os.Getenv("PORT"))

	if err != nil {
		panic(err.Error())
	}

}
