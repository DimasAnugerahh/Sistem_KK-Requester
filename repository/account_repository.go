package repository

import (
	"kk-requester/model/domain"
)

type AccountRepository interface {
	GetAllAccounts() ([]domain.Account, error)
	GetAccounts(id int) (*domain.Account, error)
	CreateAccount(NewAccount *domain.Account) (*domain.Account, error)
	AccountLogin(Account *domain.Account) (*domain.Account, error)
	AccountUpdate(UpdatedAccountAccount *domain.Account, id float64) (*domain.Account, error)
	AccountDelete(DeletedAccount *domain.Account, id int) (*domain.Account, error)
}
