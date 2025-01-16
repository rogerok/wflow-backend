package router

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/encryptcookie"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"os"

	"time"

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

	app.Use(logger.New(),
		limiter.New(limiter.Config{
			Max:        1000,
			Expiration: 60 * time.Second,
		}),
		compress.New(),
		encryptcookie.New(encryptcookie.Config{
			Key: os.Getenv("COOKIES_SECRET_KEY"),
		}),
		func(c *fiber.Ctx) error {
			c.Set("Accept-Encoding", "gzip, deflate, br")
			return c.Next()
		},
	)

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
	books.Get("/", handlers.GetBooksList(booksService))

	goals := apiPrivate.Group("/goals")
	goalsRepo := repositories.NewGoalsRepository(db)
	goalsService := services.NewGoalsService(goalsRepo)
	goals.Post("/", handlers.CreateGoal(goalsService))
	goals.Get("/", handlers.GetListByBookId(goalsService))
	goals.Get("/:id", handlers.GetById(goalsService))

	return db, nil

}
