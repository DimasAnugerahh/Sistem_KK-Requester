package service

import (
	"kk-requester/model/domain"

	"github.com/labstack/echo/v4"
)

type AktaKelahiranService interface {
	CreateAktaKelahiran(ctx echo.Context, NewAktaKelahiran *domain.AktaKelahiran, request *domain.ReqDetailAktaKelahiran) (*domain.AktaKelahiran, *domain.ReqDetailAktaKelahiran, error)
	AktaKelahiranUpdate(ctx echo.Context, UpdatedAktaKelahiran *domain.AktaKelahiran, AktaKelahiranId float64, accountId uint) (*domain.AktaKelahiran, error)
	GetAktaKelahiran(ctx echo.Context, accountId uint) ([]domain.AktaKelahiran, error)
}
