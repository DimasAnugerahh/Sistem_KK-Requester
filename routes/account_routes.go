package routes

import (
	"kk-requester/controller"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewAccountRoutes(e *echo.Echo, controller controller.AccountController) {
	AccountGroup := e.Group("/accounts")

	AccountGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	AccountGroup.GET("", (controller.GetAllAccounts()))
	AccountGroup.GET("/me", (controller.GetAccounts()))

	AccountGroup.PUT("", (controller.AccountUpdate()))

	AccountGroup.DELETE("/:id", (controller.AccountDelete()))
	AccountGroup.DELETE("", (controller.AccountDelete()))

	e.POST("/accounts", controller.CreateAccount())
	e.POST("/accounts/login", controller.AccountLogin())
}
