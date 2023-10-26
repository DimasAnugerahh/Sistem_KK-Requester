package service

import (
	"fmt"
	"kk-requester/model/domain"
	"kk-requester/repository"

	"github.com/labstack/echo/v4"
)

type SuratKematianServiceImpl struct {
	SuratKematianRepository repository.SuratKematianRepository
}

func NewSuratKematianService(kr repository.SuratKematianRepository) *SuratKematianServiceImpl {
	return &SuratKematianServiceImpl{
		SuratKematianRepository: kr,
	}
}

func (service *SuratKematianServiceImpl) CreateSuratKematian(ctx echo.Context, NewSuratKematian *domain.SuratKematian) (*domain.SuratKematian, error) {
	result, err := service.SuratKematianRepository.CreateSuratKematian(NewSuratKematian)
	if err != nil {
		return nil, fmt.Errorf("ERROR CREATING SuratKematian: %s", err.Error())
	}

	return result, nil
}

func (service *SuratKematianServiceImpl) SuratKematianUpdate(ctx echo.Context, UpdatedSuratKematian *domain.SuratKematian, SuratKematianId float64, accountId uint) (*domain.SuratKematian, error) {
	result, err := service.SuratKematianRepository.SuratKematianUpdate(UpdatedSuratKematian, SuratKematianId, accountId)
	if err != nil {
		return nil, fmt.Errorf("ERROR LOGIN: %s", err.Error())
	}

	return result, nil
}

func (service *SuratKematianServiceImpl) GetSuratKematian(ctx echo.Context, accountId uint) ([]domain.SuratKematian, error) {
	result, err := service.SuratKematianRepository.GetSuratKematian(accountId)
	if err != nil {
		return nil, fmt.Errorf("ERROR LOGIN: %s", err.Error())
	}

	return result, nil
}
