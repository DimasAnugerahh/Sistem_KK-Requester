package controller

import "github.com/labstack/echo/v4"

type AccountController interface {
	GetAllAccounts() echo.HandlerFunc
	GetAccounts() echo.HandlerFunc
	CreateAccount() echo.HandlerFunc
	AccountLogin() echo.HandlerFunc
	AccountUpdate() echo.HandlerFunc
	AccountDelete() echo.HandlerFunc
}
