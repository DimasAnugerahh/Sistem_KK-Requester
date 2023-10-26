package service

import (
	"fmt"
	"kk-requester/model/domain"
	"kk-requester/repository"

	"github.com/labstack/echo/v4"
)

type AktaKelahiranServiceImpl struct {
	AktaKelahiranRepository repository.AktaKelahiranRepository
}

func NewAktaKelahiranService(kr repository.AktaKelahiranRepository) *AktaKelahiranServiceImpl {
	return &AktaKelahiranServiceImpl{
		AktaKelahiranRepository: kr,
	}
}

func (service *AktaKelahiranServiceImpl) CreateAktaKelahiran(ctx echo.Context, NewAktaKelahiran *domain.AktaKelahiran) (*domain.AktaKelahiran, error) {
	result, err := service.AktaKelahiranRepository.CreateAktaKelahiran(NewAktaKelahiran)
	if err != nil {
		return nil, fmt.Errorf("ERROR CREATING AktaKelahiran: %s", err.Error())
	}

	return result, nil
}

func (service *AktaKelahiranServiceImpl) AktaKelahiranUpdate(ctx echo.Context, UpdatedAktaKelahiran *domain.AktaKelahiran, AktaKelahiranId float64, accountId uint) (*domain.AktaKelahiran, error) {
	result, err := service.AktaKelahiranRepository.AktaKelahiranUpdate(UpdatedAktaKelahiran, AktaKelahiranId, accountId)
	if err != nil {
		return nil, fmt.Errorf("ERROR LOGIN: %s", err.Error())
	}

	return result, nil
}

func (service *AktaKelahiranServiceImpl) GetAktaKelahiran(ctx echo.Context, accountId uint) ([]domain.AktaKelahiran, error) {
	result, err := service.AktaKelahiranRepository.GetAktaKelahiran(accountId)
	if err != nil {
		return nil, fmt.Errorf("ERROR LOGIN: %s", err.Error())
	}

	return result, nil
}
