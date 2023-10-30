package service

import (
	"kk-requester/model/domain"

	"github.com/labstack/echo/v4"
)

type AktaKematianService interface {
	CreateAktaKematian(ctx echo.Context, NewAktaKematian *domain.AktaKematian, request *domain.ReqDetailAktaKematian) (*domain.AktaKematian, *domain.ReqDetailAktaKematian, error)
	AktaKematianUpdate(ctx echo.Context, UpdatedAktaKematian *domain.AktaKematian, AktaKematianId float64, accountId uint) (*domain.AktaKematian, error)
	GetAktaKematian(ctx echo.Context, accountId uint) ([]domain.AktaKematian, error)
}
