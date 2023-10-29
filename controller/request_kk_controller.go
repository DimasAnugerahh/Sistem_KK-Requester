package controller

import "github.com/labstack/echo/v4"

type RequestKKController interface {
	CreateRequestKK() echo.HandlerFunc
	GetRequestKK() echo.HandlerFunc
	RequestKKUpdate() echo.HandlerFunc
}
