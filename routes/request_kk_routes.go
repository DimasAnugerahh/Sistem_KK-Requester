package routes

import (
	"kk-requester/controller"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewKKRequestRoutes(e *echo.Echo, controller controller.RequestKKController) {
	KKRequestGroup := e.Group("/requests")

	KKRequestGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	KKRequestGroup.POST("", (controller.CreateRequestKK()))
	KKRequestGroup.GET("", (controller.GetRequestKK()))
	KKRequestGroup.GET("/me", (controller.GetUserRequestKK()))

	KKRequestGroup.PUT("/:id", (controller.RequestKKUpdate()))

}
