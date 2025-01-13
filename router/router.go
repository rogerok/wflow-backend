package router

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jmoiron/sqlx"
	"github.com/rogerok/wflow-backend/configs"
	"github.com/rogerok/wflow-backend/handlers"
	"github.com/rogerok/wflow-backend/middleware"
	"github.com/rogerok/wflow-backend/repositories"
	"github.com/rogerok/wflow-backend/services"
)

func SetupRouter(app *fiber.App) (*sqlx.DB, error) {
	db, dbError := configs.ConnectToDb()

	if dbError != nil {
		fmt.Printf("Could connect to database %s=", dbError.Error())
		return nil, dbError
	}

	app.Use(logger.New())

	api := app.Group("/api")
	usersRepo := repositories.NewUserRepository(db)
	userService := services.NewUsersService(usersRepo)

	auth := api.Group("/auth")
	authRepo := repositories.NewAuthRepository(db)
	authService := services.NewAuthService(usersRepo, authRepo)

	// public
	api.Post("/users", handlers.CreateUser(userService))

	auth.Post("/", handlers.AuthUser(authService))
	auth.Post("/refresh", handlers.Refresh(authService))

	// private
	apiPrivate := api.Group("/private", middleware.AuthMiddleware())

	users := apiPrivate.Group("/users")

	users.Get("/", handlers.UsersList(userService))
	users.Get("/:id", handlers.UserById(userService))

	books := apiPrivate.Group("/books")
	booksRepo := repositories.NewBooksRepository(db)
	booksService := services.NewBooksService(booksRepo)
	books.Post("/", handlers.CreateBook(booksService))

	return db, nil

}
