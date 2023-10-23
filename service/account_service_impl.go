package service

import (
	"fmt"
	"kk-requester/model/domain"
	"kk-requester/repository"

	"github.com/labstack/echo/v4"
)

type AccountServiceImpl struct {
	AccountRepository repository.AccountRepository
}

func NewAccountService(ur repository.AccountRepository) *AccountServiceImpl {
	return &AccountServiceImpl{
		AccountRepository: ur,
	}
}

func (service *AccountServiceImpl) CreateAccount(ctx echo.Context, NewAccount *domain.Account) (*domain.Account, error) {
	result, err := service.AccountRepository.CreateAccount(NewAccount)
	if err != nil {
		return nil, fmt.Errorf("ERROR CREATING Account: %s", err.Error())
	}

	return result, nil
}

func (service *AccountServiceImpl) GetAllAccounts(ctx echo.Context) ([]domain.Account, error) {
	result, err := service.AccountRepository.GetAllAccounts()
	if err != nil {
		return nil, fmt.Errorf("ERROR CREATING Account: %s", err.Error())
	}

	return result, nil
}

func (service *AccountServiceImpl) AccountLogin(ctx echo.Context, Account *domain.Account) (*domain.Account, error) {
	result, err := service.AccountRepository.AccountLogin(Account)
	if err != nil {
		return nil, fmt.Errorf("ERROR LOGIN: %s", err.Error())
	}

	return result, nil
}

func (service *AccountServiceImpl) AccountUpdate(ctx echo.Context, UpdatedAccount *domain.Account, id float64) (*domain.Account, error) {
	result, err := service.AccountRepository.AccountUpdate(UpdatedAccount, id)
	if err != nil {
		return nil, fmt.Errorf("ERROR LOGIN: %s", err.Error())
	}

	return result, nil
}

func (service *AccountServiceImpl) AccountDelete(ctx echo.Context, DeletedAccount *domain.Account, id int) (*domain.Account, error) {
	_, err := service.AccountRepository.AccountDelete(DeletedAccount, id)
	if err != nil {
		return nil, fmt.Errorf("ERROR DELETE: %s", err.Error())
	}

	return DeletedAccount, nil
}
