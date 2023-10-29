package service

import (
	"kk-requester/model/domain"

	"github.com/labstack/echo/v4"
)

type SuratPindahService interface {
	CreateSuratPindah(ctx echo.Context, NewSuratPindah *domain.SuratPindah, request *domain.ReqDetailSuratPindah) (*domain.SuratPindah, *domain.ReqDetailSuratPindah, error)
	SuratPindahUpdate(ctx echo.Context, UpdatedSuratPindah *domain.SuratPindah, SuratPindahId float64, accountId uint) (*domain.SuratPindah, error)
	GetSuratPindah(ctx echo.Context, accountId uint) ([]domain.SuratPindah, error)
}
