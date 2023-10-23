package routes

import (
	"kk-requester/controller"
	"kk-requester/helper"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewAccountRoutes(e *echo.Echo, controller controller.AccountController) {
	AccountGroup := e.Group("")

	AccountGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	AccountGroup.GET("/accounts", helper.RoleAuth(controller.GetAllAccounts()))
	AccountGroup.PUT("/accounts/:id", helper.IdAuth(controller.AccountUpdate()))
	AccountGroup.DELETE("/accounts/:id", helper.RoleAuth(controller.AccountDelete()))
	e.POST("/accounts", controller.CreateAccount())
	e.POST("/accounts/login", controller.AccountLogin())
}
