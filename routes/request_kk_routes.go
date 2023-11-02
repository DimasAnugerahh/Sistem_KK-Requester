package routes

import (
	"kk-requester/controller"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewKKRequestRoutes(e *echo.Echo, controller controller.RequestKKController) {
	KKRequestGroup := e.Group("")

	KKRequestGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	KKRequestGroup.POST("/requests", (controller.CreateRequestKK()))
	KKRequestGroup.GET("/requests", (controller.GetRequestKK()))
	KKRequestGroup.GET("/request/me", (controller.GetUserRequestKK()))

	KKRequestGroup.PUT("/requests/:id", (controller.RequestKKUpdate()))

}
