package controller

import "github.com/labstack/echo/v4"

type AktaKematianController interface {
	CreateAktaKematian() echo.HandlerFunc
	AktaKematianUpdate() echo.HandlerFunc
	GetAktaKematian() echo.HandlerFunc
}
