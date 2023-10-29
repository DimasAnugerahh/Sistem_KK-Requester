package controller

import "github.com/labstack/echo/v4"

type KkController interface {
	GetKK() echo.HandlerFunc
	CreateKK() echo.HandlerFunc
}
