package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rogerok/wflow-backend/configs"
	_ "github.com/rogerok/wflow-backend/docs"
	"github.com/rogerok/wflow-backend/forms/validators"
	"github.com/rogerok/wflow-backend/router"
	"os"
)

// @title Word-Flow app API
// @version 1.0
// @description Word-Flow API docs
// @host  http://127.0.0.1:5000
// @BasePath /api
func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Could not load environment %s", err.Error())
		return
	}

	validators.InitTranslation()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowMethods:     "GET,POST,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders:     "Content-Type,Authorization,Accept,Origin,Access-Control-Request-Method,Access-Control-Request-Headers,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Access-Control-Allow-Methods,Access-Control-Expose-Headers,Access-Control-Max-Age,Access-Control-Allow-Credentials",
		AllowCredentials: true,
		AllowOriginsFunc: func(origin string) bool {
			return origin == "http://193.46.217.8:4200" || origin == "http://localhost:4200"
		},
	}))
	db, dbError := configs.ConnectToDb()

	if dbError != nil {
		fmt.Printf("Could connect to database %s=", dbError.Error())
	}

	router.SetupRouter(app, db)

	app.Get("/swagger/*", swagger.HandlerDefault)

	err = app.Listen(":" + os.Getenv("PORT"))

	if err != nil {
		panic(err.Error())
	}

	defer configs.CloseConnectionToDb(db)

}
