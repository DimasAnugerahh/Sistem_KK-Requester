package service

import (
	"kk-requester/model/domain"

	"github.com/labstack/echo/v4"
)

type SuratKematianService interface {
	CreateSuratKematian(ctx echo.Context, NewSuratKematian *domain.SuratKematian) (*domain.SuratKematian, error)
	SuratKematianUpdate(ctx echo.Context, UpdatedSuratKematian *domain.SuratKematian, SuratKematianId float64, accountId uint) (*domain.SuratKematian, error)
	GetSuratKematian(ctx echo.Context, accountId uint) ([]domain.SuratKematian, error)
}
