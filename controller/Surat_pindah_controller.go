package controller

import "github.com/labstack/echo/v4"

type SuratPindahController interface {
	CreateSuratPindah() echo.HandlerFunc
	SuratPindahUpdate() echo.HandlerFunc
	GetSuratPindah() echo.HandlerFunc
}
