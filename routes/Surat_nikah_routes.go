package routes

import (
	"kk-requester/controller"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewSuratNikahRoutes(e *echo.Echo, controller controller.SuratNikahController) {
	SuratNikahGroup := e.Group("")

	SuratNikahGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	SuratNikahGroup.POST("/requests/:id/suratnikah", (controller.CreateSuratNikah()))
	SuratNikahGroup.GET("/suratnikah", (controller.GetSuratNikah()))
	SuratNikahGroup.PUT("/suratnikah/:id", (controller.SuratNikahUpdate()))

}
