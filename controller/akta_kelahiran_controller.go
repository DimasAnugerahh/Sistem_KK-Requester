package controller

import "github.com/labstack/echo/v4"

type AktaKelahiranController interface {
	CreateAktaKelahiran() echo.HandlerFunc
	AktaKelahiranUpdate() echo.HandlerFunc
	GetAktaKelahiran() echo.HandlerFunc
}
