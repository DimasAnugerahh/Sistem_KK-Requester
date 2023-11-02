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

	SuratNikahGroup.POST("/requests/:id/surat-nikahs", (controller.CreateSuratNikah()))
	SuratNikahGroup.GET("/surat-nikahs", (controller.GetSuratNikah()))
	SuratNikahGroup.PUT("/surat-nikahs/:id", (controller.SuratNikahUpdate()))

}
