package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rogerok/wflow-backend/configs"
	_ "github.com/rogerok/wflow-backend/docs"
	"github.com/rogerok/wflow-backend/forms"
	"github.com/rogerok/wflow-backend/router"
	"os"
)

// @title Word-Flow app API
// @version 1.0
// @description Word-Flow API docs
// @host 127.0.0.1:5000
// @BasePath /api
func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Could not load environment %s", err.Error())
		return
	}

	forms.InitTranslation()

	app := fiber.New()

	db, err := router.SetupRouter(app)

	app.Get("/swagger/*", swagger.HandlerDefault)

	err = app.Listen(":" + os.Getenv("PORT"))

	if err != nil {
		panic(err.Error())
	}

	defer configs.CloseConnectionToDb(db)

}
