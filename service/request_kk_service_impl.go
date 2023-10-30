package service

import (
	"fmt"
	"kk-requester/model/domain"
	"kk-requester/repository"

	"github.com/labstack/echo/v4"
)

type RequestKKServiceImpl struct {
	RequestKKRepository repository.RequestKKRepository
}

func NewRequestKKService(kr repository.RequestKKRepository) *RequestKKServiceImpl {
	return &RequestKKServiceImpl{
		RequestKKRepository: kr,
	}
}

func (service *RequestKKServiceImpl) CreateRequestKK(ctx echo.Context, NewkkRequest *domain.RequestKK) (*domain.RequestKK, error) {

	result, err := service.RequestKKRepository.CreateRequestKK(NewkkRequest)
	if err != nil {
		return nil, fmt.Errorf("ERROR CREATING RequestKK: %s", err.Error())
	}

	return result, nil
}

func (service *RequestKKServiceImpl) GetUserRequestKK(ctx echo.Context, page int, limit int, sortby string, order string, accountId uint) ([]domain.RequestKK, error) {
	result, err := service.RequestKKRepository.GetuserRequestKK(page, limit, sortby, order, accountId)
	if err != nil {
		return nil, fmt.Errorf("ERROR GET USERS: %s", err.Error())
	}

	return result, nil
}

func (service *RequestKKServiceImpl) GetRequestKK(ctx echo.Context, page int, limit int, sortby string, order string) ([]domain.RequestKK, error) {
	result, err := service.RequestKKRepository.GetRequestKK(page, limit, sortby, order)
	if err != nil {
		return nil, fmt.Errorf("ERROR GET USERS: %s", err.Error())
	}

	return result, nil
}

func (service *RequestKKServiceImpl) RequestKKUpdate(ctx echo.Context, UpdatedRequestKK *domain.RequestKK, RequestKKId float64) (*domain.RequestKK, error) {
	result, err := service.RequestKKRepository.RequestKKUpdate(UpdatedRequestKK, RequestKKId)
	if err != nil {
		return nil, fmt.Errorf("ERROR LOGIN: %s", err.Error())
	}

	return result, nil
}
