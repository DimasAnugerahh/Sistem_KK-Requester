package routes

import (
	"kk-requester/controller"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewAccountRoutes(e *echo.Echo, controller controller.AccountController) {
	AccountGroup := e.Group("")

	AccountGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	AccountGroup.GET("/accounts", (controller.GetAllAccounts()))
	AccountGroup.GET("/accounts/me", (controller.GetAccounts()))

	AccountGroup.PUT("/accounts", (controller.AccountUpdate()))

	AccountGroup.DELETE("/accounts/:id", (controller.AccountDelete()))
	AccountGroup.DELETE("/accounts", (controller.AccountDelete()))

	e.POST("/accounts", controller.CreateAccount())
	e.POST("/accounts/login", controller.AccountLogin())
}
