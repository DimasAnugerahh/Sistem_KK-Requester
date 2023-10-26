package routes

import (
	"kk-requester/controller"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewSuratKematianRoutes(e *echo.Echo, controller controller.SuratKematianController) {
	SuratKematianGroup := e.Group("")

	SuratKematianGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	SuratKematianGroup.POST("/suratkematian", (controller.CreateSuratKematian()))
	SuratKematianGroup.GET("/suratkematian", (controller.GetSuratKematian()))
	SuratKematianGroup.PUT("/suratkematian/:id", (controller.SuratKematianUpdate()))

}
