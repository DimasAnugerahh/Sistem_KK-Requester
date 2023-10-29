package service

import (
	"fmt"
	"kk-requester/model/domain"
	"kk-requester/repository"

	"github.com/labstack/echo/v4"
)

type KTPServiceImpl struct {
	KTPRepository repository.KTPRepository
}

func NewKTPService(kr repository.KTPRepository) *KTPServiceImpl {
	return &KTPServiceImpl{
		KTPRepository: kr,
	}
}

func (service *KTPServiceImpl) CreateKTP(ctx echo.Context, NewKTP *domain.KTP, request *domain.ReqDetailKtp) (*domain.KTP, *domain.ReqDetailKtp, error) {
	ktp, req, err := service.KTPRepository.CreateKTP(NewKTP, request)
	if err != nil {
		return nil, nil, fmt.Errorf("ERROR CREATING KTP: %s", err.Error())
	}

	return ktp, req, nil
}

func (service *KTPServiceImpl) KTPUpdate(ctx echo.Context, UpdatedKTP *domain.KTP, ktpId float64, accountId uint) (*domain.KTP, error) {
	result, err := service.KTPRepository.KTPUpdate(UpdatedKTP, ktpId, accountId)
	if err != nil {
		return nil, fmt.Errorf("ERROR LOGIN: %s", err.Error())
	}

	return result, nil
}

func (service *KTPServiceImpl) GetKTP(ctx echo.Context, idAccount uint) ([]domain.KTP, error) {
	result, err := service.KTPRepository.GetKTP(idAccount)
	if err != nil {
		return nil, fmt.Errorf("ERROR LOGIN: %s", err.Error())
	}

	return result, nil
}
