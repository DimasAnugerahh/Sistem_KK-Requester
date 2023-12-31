package service

import (
	"kk-requester/model/domain"

	"github.com/labstack/echo/v4"
)

type KTPService interface {
	CreateKTP(ctx echo.Context, NewKTP *domain.KTP, request *domain.ReqDetailKtp) (*domain.KTP, *domain.ReqDetailKtp, error)
	KTPUpdate(ctx echo.Context, UpdatedKTP *domain.KTP, ktpId float64, accountId uint) (*domain.KTP, error)
	GetKTP(ctx echo.Context, idAccount uint) ([]domain.KTP, error)
}
