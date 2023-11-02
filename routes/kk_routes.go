package routes

import (
	"kk-requester/controller"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewkkRoutes(e *echo.Echo, controller controller.KkController) {
	kkGroup := e.Group("/kks")

	kkGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	kkGroup.POST("", (controller.CreateKK()))
	kkGroup.GET("", (controller.GetKK()))
}
