package routes

import (
	"kk-requester/controller"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewSuratPindahRoutes(e *echo.Echo, controller controller.SuratPindahController) {
	SuratPindahGroup := e.Group("/surat-pindahs")

	SuratPindahGroup.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	SuratPindahGroup.POST("/requests/:id", (controller.CreateSuratPindah()))
	SuratPindahGroup.GET("", (controller.GetSuratPindah()))
	SuratPindahGroup.PUT("/:id", (controller.SuratPindahUpdate()))

}
