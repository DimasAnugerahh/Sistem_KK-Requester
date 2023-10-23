package repository

import (
	"kk-requester/model/domain"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AccountRepositoryImpl struct {
	DB *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &AccountRepositoryImpl{DB: db}
}

func (ur *AccountRepositoryImpl) GetAllAccounts() ([]domain.Account, error) {
	var Accounts = []domain.Account{}
	if err := ur.DB.Find(&Accounts).Error; err != nil {
		logrus.Error("Model: find data error", err.Error())
		return nil, err
	}
	return Accounts, nil
}

func (ur *AccountRepositoryImpl) CreateAccount(NewAccount *domain.Account) (*domain.Account, error) {
	if err := ur.DB.Create(&NewAccount).Error; err != nil {
		logrus.Error("Model: insert data error", err.Error())
		return nil, err
	}
	return NewAccount, nil
}

func (ur *AccountRepositoryImpl) AccountLogin(Account *domain.Account) (*domain.Account, error) {
	var Accounts = domain.Account{}
	if err := ur.DB.Where("email=? and password=?", Account.Email, Account.Password).First(&Accounts).Error; err != nil {
		logrus.Error("Model: find data error", err.Error())
		return nil, err
	}
	return &Accounts, nil
}

func (ur *AccountRepositoryImpl) AccountUpdate(UpdatedAccount *domain.Account, id float64) (*domain.Account, error) {
	if err := ur.DB.Model(&domain.Account{}).Where("id=?", id).Updates(domain.Account{Email: UpdatedAccount.Email, Password: UpdatedAccount.Password, Role: UpdatedAccount.Role, Name: UpdatedAccount.Name}).Error; err != nil {
		logrus.Error("Model: update data error", err.Error())
		return nil, err
	}
	return UpdatedAccount, nil
}

func (ur *AccountRepositoryImpl) AccountDelete(DeletedAccount *domain.Account, id int) (*domain.Account, error) {
	qry := ur.DB.Where("id=?", id).Delete(&domain.Account{})

	if dataCount := qry.RowsAffected; dataCount < 1 {
		logrus.Error("Model : Update error, ", "no data affected")
	}

	if err := qry.Error; err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, err.Error())

	}
	return DeletedAccount, nil
}
