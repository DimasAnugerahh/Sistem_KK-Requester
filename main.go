package main

import (
	"kk-requester/config"
	"kk-requester/controller"
	"kk-requester/repository"
	"kk-requester/routes"
	"kk-requester/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	app := echo.New()
	db, err := config.InitDB()
	if err != nil {
		panic(err)
	}
	config.InitialMigration(db)

	AccountRepository := repository.NewAccountRepository(db)
	AccountService := service.NewAccountService(AccountRepository)
	AccountController := controller.NewAccountController(AccountService)

	app.Use(middleware.LoggerWithConfig(
		middleware.LoggerConfig{
			Format: "method=${method}, uri=${uri}, status=${status}, time=${time_rfc3339}\n",
		},
	))

	routes.NewAccountRoutes(app, AccountController)

	app.Start(":8080")
}
