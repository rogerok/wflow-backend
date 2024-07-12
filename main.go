package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rogerok/wflow-backend/router"
	"os"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Could not load environment %s", err.Error())
		return
	}

	app := fiber.New()
	router.SetupRouter(app)
	err = app.Listen(":" + os.Getenv("PORT"))

	if err != nil {
		panic(err.Error())
	}

}
