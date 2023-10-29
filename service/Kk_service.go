package service

import (
	"kk-requester/model/domain"

	"github.com/labstack/echo/v4"
)

type KkService interface {
	GetKK(ctx echo.Context) ([]domain.KK, error)
	CreateKK(ctx echo.Context, NewKK *domain.KK) (*domain.KK, error)
}
