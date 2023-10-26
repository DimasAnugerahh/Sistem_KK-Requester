package service

import (
	"kk-requester/model/domain"

	"github.com/labstack/echo/v4"
)

type SuratNikahService interface {
	CreateSuratNikah(ctx echo.Context, NewSuratNikah *domain.SuratNikah) (*domain.SuratNikah, error)
	SuratNikahUpdate(ctx echo.Context, UpdatedSuratNikah *domain.SuratNikah, SuratNikahId float64, accountId uint) (*domain.SuratNikah, error)
	GetSuratNikah(ctx echo.Context, accountId uint) ([]domain.SuratNikah, error)
}
