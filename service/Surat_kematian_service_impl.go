package service

import (
	"fmt"
	"kk-requester/model/domain"
	"kk-requester/repository"

	"github.com/labstack/echo/v4"
)

type AktaKematianServiceImpl struct {
	AktaKematianRepository repository.AktaKematianRepository
}

func NewAktaKematianService(kr repository.AktaKematianRepository) *AktaKematianServiceImpl {
	return &AktaKematianServiceImpl{
		AktaKematianRepository: kr,
	}
}

func (service *AktaKematianServiceImpl) CreateAktaKematian(ctx echo.Context, NewAktaKematian *domain.AktaKematian, request *domain.ReqDetailAktaKematian) (*domain.AktaKematian, *domain.ReqDetailAktaKematian, error) {
	aktaKematian, req, err := service.AktaKematianRepository.CreateAktaKematian(NewAktaKematian, request)
	if err != nil {
		return nil, nil, fmt.Errorf("ERROR CREATING AktaKematian: %s", err.Error())
	}

	return aktaKematian, req, nil
}

func (service *AktaKematianServiceImpl) AktaKematianUpdate(ctx echo.Context, UpdatedAktaKematian *domain.AktaKematian, AktaKematianId float64, accountId uint) (*domain.AktaKematian, error) {
	result, err := service.AktaKematianRepository.AktaKematianUpdate(UpdatedAktaKematian, AktaKematianId, accountId)
	if err != nil {
		return nil, fmt.Errorf("ERROR LOGIN: %s", err.Error())
	}

	return result, nil
}

func (service *AktaKematianServiceImpl) GetAktaKematian(ctx echo.Context, accountId uint) ([]domain.AktaKematian, error) {
	result, err := service.AktaKematianRepository.GetAktaKematian(accountId)
	if err != nil {
		return nil, fmt.Errorf("ERROR LOGIN: %s", err.Error())
	}

	return result, nil
}
