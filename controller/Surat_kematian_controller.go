package controller

import "github.com/labstack/echo/v4"

type SuratKematianController interface {
	CreateSuratKematian() echo.HandlerFunc
	SuratKematianUpdate() echo.HandlerFunc
	GetSuratKematian() echo.HandlerFunc
}
