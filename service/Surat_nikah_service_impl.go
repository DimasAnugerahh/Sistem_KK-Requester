package service

import (
	"fmt"
	"kk-requester/model/domain"
	"kk-requester/repository"

	"github.com/labstack/echo/v4"
)

type SuratNikahServiceImpl struct {
	SuratNikahRepository repository.SuratNikahRepository
}

func NewSuratNikahService(kr repository.SuratNikahRepository) *SuratNikahServiceImpl {
	return &SuratNikahServiceImpl{
		SuratNikahRepository: kr,
	}
}

func (service *SuratNikahServiceImpl) CreateSuratNikah(ctx echo.Context, NewSuratNikah *domain.SuratNikah) (*domain.SuratNikah, error) {
	result, err := service.SuratNikahRepository.CreateSuratNikah(NewSuratNikah)
	if err != nil {
		return nil, fmt.Errorf("ERROR CREATING SuratNikah: %s", err.Error())
	}

	return result, nil
}

func (service *SuratNikahServiceImpl) SuratNikahUpdate(ctx echo.Context, UpdatedSuratNikah *domain.SuratNikah, SuratNikahId float64, accountId uint) (*domain.SuratNikah, error) {
	result, err := service.SuratNikahRepository.SuratNikahUpdate(UpdatedSuratNikah, SuratNikahId, accountId)
	if err != nil {
		return nil, fmt.Errorf("ERROR LOGIN: %s", err.Error())
	}

	return result, nil
}

func (service *SuratNikahServiceImpl) GetSuratNikah(ctx echo.Context, accountId uint) ([]domain.SuratNikah, error) {
	result, err := service.SuratNikahRepository.GetSuratNikah(accountId)
	if err != nil {
		return nil, fmt.Errorf("ERROR LOGIN: %s", err.Error())
	}

	return result, nil
}
