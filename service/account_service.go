package service

import (
	"kk-requester/model/domain"

	"github.com/labstack/echo/v4"
)

type AccountService interface {
	GetAllAccounts(ctx echo.Context) ([]domain.Account, error)
	GetAccounts(ctx echo.Context, id int) (*domain.Account, error)

	CreateAccount(ctx echo.Context, NewAccount *domain.Account) (*domain.Account, error)
	AccountLogin(ctx echo.Context, NewAccount *domain.Account) (*domain.Account, error)
	AccountUpdate(ctx echo.Context, UpdatedAccount *domain.Account, id float64) (*domain.Account, error)
	AccountDelete(ctx echo.Context, DeletedAccount *domain.Account, id int) (*domain.Account, error)
}
