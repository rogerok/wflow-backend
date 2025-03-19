package router

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
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
	quotes := api.Group("/quotes")
	quotesRepo := repositories.NewQuotesRepository(db)
	quotesService := services.NewQuotesService(quotesRepo)
	quotes.Get("/", handlers.GetRandomQuote(quotesService))
	api.Post("/users", handlers.CreateUser(userService))

	auth.Post("/", handlers.AuthUser(authService))
	auth.Post("/refresh", handlers.Refresh(authService))
	auth.Post("/logout", handlers.Logout(authService))

	// private
	apiPrivate := api.Group("/private", middleware.AuthMiddleware())

	users := apiPrivate.Group("/users")

	users.Get("/", handlers.UsersList(userService))
	users.Get("/:id", handlers.UserById(userService))

	books := apiPrivate.Group("/books")
	booksRepo := repositories.NewBooksRepository(db)
	booksService := services.NewBooksService(booksRepo)
	books.Post("/", handlers.CreateBook(booksService))
	books.Put("/update/:id", handlers.UpdateBook(booksService))
	books.Delete("/delete/:id", handlers.DeleteBook(booksService))
	books.Get("/", handlers.GetBooksList(booksService))
	books.Get("/:id", handlers.GetBookById(booksService))

	goals := apiPrivate.Group("/goals")
	goalsRepo := repositories.NewGoalsRepository(db)
	goalsService := services.NewGoalsService(goalsRepo)
	goals.Post("/", handlers.CreateGoal(goalsService))
	goals.Put("/edit/:id", handlers.UpdateGoal(goalsService))
	goals.Delete("/delete/:id", handlers.DeleteGoal(goalsService))
	goals.Get("/", handlers.GetList(goalsService))
	goals.Get("/:id", handlers.GetGoalById(goalsService))

	reports := apiPrivate.Group("/reports")
	reportsRepo := repositories.NewReportsRepository(db)
	reportsService := services.NewReportsService(reportsRepo, goalsRepo)
	reports.Post("/", handlers.CreateReport(reportsService))

	// inner commands
	command := app.Group("/command")
	command.Get("/goals", handlers.RecalculateGoals(goalsService))

	statistics := apiPrivate.Group("/statistics")
	statisticsRepo := repositories.NewStatisticsRepository(db)
	statisticsService := services.NewStatisticService(statisticsRepo)
	statistics.Get("/user", handlers.GetUserStatistics(statisticsService))
	statistics.Get("/user/full", handlers.GetFullProfileChartData(statisticsService))
	statistics.Get("/goal/:id", handlers.GetGoalStatistics(statisticsService))

	return db, nil

}
