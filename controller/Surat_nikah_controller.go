package controller

import "github.com/labstack/echo/v4"

type SuratNikahController interface {
	CreateSuratNikah() echo.HandlerFunc
	SuratNikahUpdate() echo.HandlerFunc
	GetSuratNikah() echo.HandlerFunc
}
