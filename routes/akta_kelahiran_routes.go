package routes

import (
	"kk-requester/controller"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewAktaKelahiranRoutes(e *echo.Echo, controller controller.AktaKelahiranController) {
	AktaKelahiranGroup := e.Group("/akte-kelahirans")

	AktaKelahiranGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	AktaKelahiranGroup.POST("/requests/:id", (controller.CreateAktaKelahiran()))
	AktaKelahiranGroup.GET("", (controller.GetAktaKelahiran()))
	AktaKelahiranGroup.PUT("/:id", (controller.AktaKelahiranUpdate()))
}
