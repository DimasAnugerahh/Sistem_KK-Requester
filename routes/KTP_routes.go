package routes

import (
	"kk-requester/controller"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewKTPRoutes(e *echo.Echo, controller controller.KTPController) {
	KTPGroup := e.Group("")

	KTPGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	KTPGroup.POST("/ktp", (controller.CreateKTP()))
	KTPGroup.GET("/ktp", (controller.GetKTP()))
	KTPGroup.PUT("/ktp/:id", (controller.KTPUpdate()))

}
