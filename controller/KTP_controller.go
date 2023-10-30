package controller

import "github.com/labstack/echo/v4"

type KTPController interface {
	CreateKTP() echo.HandlerFunc
	KTPUpdate() echo.HandlerFunc
	GetKTP() echo.HandlerFunc
}
