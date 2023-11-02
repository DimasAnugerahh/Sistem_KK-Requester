package routes

import (
	"kk-requester/controller"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewAktaKematianRoutes(e *echo.Echo, controller controller.AktaKematianController) {
	AktaKematianGroup := e.Group("")

	AktaKematianGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	AktaKematianGroup.POST("/requests/:id/akta-kematian", (controller.CreateAktaKematian()))
	AktaKematianGroup.GET("/akta-kematians", (controller.GetAktaKematian()))
	AktaKematianGroup.PUT("/akta-kematians/:id", (controller.AktaKematianUpdate()))

}
