package service

import (
	"fmt"
	"kk-requester/model/domain"
	"kk-requester/repository"

	"github.com/labstack/echo/v4"
)

type SuratPindahServiceImpl struct {
	SuratPindahRepository repository.SuratPindahRepository
}

func NewSuratPindahService(kr repository.SuratPindahRepository) *SuratPindahServiceImpl {
	return &SuratPindahServiceImpl{
		SuratPindahRepository: kr,
	}
}

func (service *SuratPindahServiceImpl) CreateSuratPindah(ctx echo.Context, NewSuratPindah *domain.SuratPindah, request *domain.ReqDetailSuratPindah) (*domain.SuratPindah, *domain.ReqDetailSuratPindah, error) {
	suratPindah, req, err := service.SuratPindahRepository.CreateSuratPindah(NewSuratPindah, request)
	if err != nil {
		return nil, nil, fmt.Errorf("ERROR CREATING SuratPindah: %s", err.Error())
	}

	return suratPindah, req, nil
}

func (service *SuratPindahServiceImpl) SuratPindahUpdate(ctx echo.Context, UpdatedSuratPindah *domain.SuratPindah, SuratPindahId float64, accountId uint) (*domain.SuratPindah, error) {
	result, err := service.SuratPindahRepository.SuratPindahUpdate(UpdatedSuratPindah, SuratPindahId, accountId)
	if err != nil {
		return nil, fmt.Errorf("ERROR LOGIN: %s", err.Error())
	}

	return result, nil
}

func (service *SuratPindahServiceImpl) GetSuratPindah(ctx echo.Context, accountId uint) ([]domain.SuratPindah, error) {
	result, err := service.SuratPindahRepository.GetSuratPindah(accountId)
	if err != nil {
		return nil, fmt.Errorf("ERROR LOGIN: %s", err.Error())
	}

	return result, nil
}
