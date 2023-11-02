package routes

import (
	"kk-requester/controller"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewSuratPindahRoutes(e *echo.Echo, controller controller.SuratPindahController) {
	SuratPindahGroup := e.Group("")

	SuratPindahGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	SuratPindahGroup.POST("/requests/:id/surat-pindahs", (controller.CreateSuratPindah()))
	SuratPindahGroup.GET("/surat-pindahs", (controller.GetSuratPindah()))
	SuratPindahGroup.PUT("/surat-pindahs/:id", (controller.SuratPindahUpdate()))

}
