package routes

import (
	"kk-requester/controller"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewKTPRoutes(e *echo.Echo, controller controller.KTPController) {
	KTPGroup := e.Group("/ktps")

	KTPGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	KTPGroup.POST("/requests/:id", (controller.CreateKTP()))
	KTPGroup.GET("", (controller.GetKTP()))
	KTPGroup.PUT("/:id", (controller.KTPUpdate()))

}
