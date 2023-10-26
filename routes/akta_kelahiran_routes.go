package routes

import (
	"kk-requester/controller"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewAktaKelahiranRoutes(e *echo.Echo, controller controller.AktaKelahiranController) {
	AktaKelahiranGroup := e.Group("")

	AktaKelahiranGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	AktaKelahiranGroup.POST("/aktekelahiran", (controller.CreateAktaKelahiran()))
	AktaKelahiranGroup.GET("/aktekelahiran", (controller.GetAktaKelahiran()))
	AktaKelahiranGroup.PUT("/aktekelahiran/:id", (controller.AktaKelahiranUpdate()))

}
