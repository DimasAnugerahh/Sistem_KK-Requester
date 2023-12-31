package service

import (
	"kk-requester/model/domain"

	"github.com/labstack/echo/v4"
)

type RequestKKService interface {
	CreateRequestKK(ctx echo.Context, NewkkRequest *domain.RequestKK) (*domain.RequestKK, error)
	GetRequestKK(ctx echo.Context, page int, limit int, sortby string, order string) ([]domain.RequestKK, int, error)
	GetUserRequestKK(ctx echo.Context, page int, limit int, sortby string, order string, accountId uint) ([]domain.RequestKK, int, error)
	RequestKKUpdate(ctx echo.Context, UpdatedRequestKK *domain.RequestKK, RequestKKId float64) (*domain.RequestKK, error)
}
