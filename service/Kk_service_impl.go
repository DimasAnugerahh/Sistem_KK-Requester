package service

import (
	"fmt"
	"kk-requester/model/domain"
	"kk-requester/repository"

	"github.com/labstack/echo/v4"
)

type KkServiceImpl struct {
	KkRepository repository.KkRepository
}

func NewKkService(ks repository.KkRepository) *KkServiceImpl {
	return &KkServiceImpl{
		KkRepository: ks,
	}
}

func (service *KkServiceImpl) GetKK(ctx echo.Context) ([]domain.KK, error) {
	kk, err := service.KkRepository.GetKK()
	if err != nil {
		return nil, fmt.Errorf("ERROR GET KK: %s", err.Error())
	}

	return kk, nil
}

func (service *KkServiceImpl) CreateKK(ctx echo.Context, NewKK *domain.KK) (*domain.KK, error) {
	kk, err := service.KkRepository.CreateKK(NewKK)
	if err != nil {
		return nil, fmt.Errorf("ERROR CREATE KK: %s", err.Error())
	}

	return kk, nil
}
