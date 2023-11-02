package routes

import (
	"kk-requester/controller"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewSuratNikahRoutes(e *echo.Echo, controller controller.SuratNikahController) {
	SuratNikahGroup := e.Group("/surat-nikahs")

	SuratNikahGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	SuratNikahGroup.POST("/requests/:id", (controller.CreateSuratNikah()))
	SuratNikahGroup.GET("", (controller.GetSuratNikah()))
	SuratNikahGroup.PUT("/:id", (controller.SuratNikahUpdate()))

}
