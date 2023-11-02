package routes

import (
	"kk-requester/controller"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewAktaKematianRoutes(e *echo.Echo, controller controller.AktaKematianController) {
	AktaKematianGroup := e.Group("/akta-kematians")

	AktaKematianGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	AktaKematianGroup.POST("/requests/:id", (controller.CreateAktaKematian()))
	AktaKematianGroup.GET("", (controller.GetAktaKematian()))
	AktaKematianGroup.PUT("/:id", (controller.AktaKematianUpdate()))

}
