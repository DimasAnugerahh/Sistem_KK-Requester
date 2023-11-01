package controller

import (
	"kk-requester/helper"
	"kk-requester/model/domain"
	"kk-requester/model/web"
	"strconv"

	"kk-requester/service"
	"net/http"

	echo "github.com/labstack/echo/v4"
)

type AccountControllerImpl struct {
	AccountService service.AccountService
}

func NewAccountController(us service.AccountService) *AccountControllerImpl {
	return &AccountControllerImpl{AccountService: us}
}

func (uc *AccountControllerImpl) GetAllAccounts() echo.HandlerFunc {

	return func(c echo.Context) error {
		_, role := helper.Authorization(c)
		if role == "admin" {
			Result, err := uc.AccountService.GetAllAccounts(c)

			if len(Result) == 0 {
				return c.JSON(http.StatusOK, echo.Map{
					"message": "there is no account",
				})
			}

			if err != nil {
				return c.JSON(http.StatusBadRequest, echo.Map{
					"message": err.Error(),
				})
			}

			return c.JSON(http.StatusOK, echo.Map{
				"message": "success",
				"data":    Result,
			})
		}
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})

	}

}

func (uc *AccountControllerImpl) GetAccounts() echo.HandlerFunc {

	return func(c echo.Context) error {
		id, _ := helper.Authorization(c)

		Result, err := uc.AccountService.GetAccounts(c, int(id))

		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		response := web.AccountResponse{Email: Result.Email, Name: Result.Name, Password: Result.Password, Role: Result.Role}

		return c.JSON(http.StatusOK, echo.Map{
			"message": "success",
			"data":    response,
		})
	}

}

func (uc *AccountControllerImpl) CreateAccount() echo.HandlerFunc {
	return func(c echo.Context) error {
		Account := &domain.Account{}
		err := c.Bind(Account)

		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		result, err := uc.AccountService.CreateAccount(c, Account)

		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error()})
		}

		Result := web.CreateAccountResponse{CreatedAt: result.CreatedAt, UpdatedAt: result.UpdatedAt, DeletedAt: result.DeletedAt.Time, Email: result.Email, Name: result.Name, Role: result.Role, Model: result.Model}

		return c.JSON(http.StatusCreated, echo.Map{
			"message": "success",
			"data":    Result,
		})
	}
}

func (uc *AccountControllerImpl) AccountLogin() echo.HandlerFunc {

	return func(c echo.Context) error {
		Account := &domain.Account{}
		err := c.Bind(Account)

		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error()})
		}

		Result, err := uc.AccountService.AccountLogin(c, Account)
		if Result == nil {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"message": "account not found",
			})
		}

		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error()})
		}

		token, err := helper.GenerateToken(Account, int(Result.ID), Result.Role, Result.Email)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error()})
		}

		AccountResult := web.LoginResponse{Email: Result.Email, Nama: Result.Name, Role: Result.Role, Token: token}

		return c.JSON(http.StatusOK, echo.Map{
			"message": "success",
			"data":    AccountResult,
		})
	}
}

func (uc *AccountControllerImpl) AccountUpdate() echo.HandlerFunc {
	return func(c echo.Context) error {

		id, _ := helper.Authorization(c)

		updateRequest := &domain.Account{}
		err := c.Bind(updateRequest)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"messege": err.Error(),
			})
		}

		Result, _ := uc.AccountService.AccountUpdate(c, updateRequest, id)

		updateAccountResult := web.AccountResponse{Email: Result.Email, Name: Result.Name, Password: Result.Password, Role: Result.Role}

		return c.JSON(http.StatusOK, echo.Map{
			"message": "success",
			"data":    updateAccountResult,
		})

	}
}

func (uc *AccountControllerImpl) AccountDelete() echo.HandlerFunc {
	return func(c echo.Context) error {

		id, role := helper.Authorization(c)
		updateRequest := &domain.Account{}
		if role == "admin" {
			idParam := c.Param("id")
			id, _ := strconv.Atoi(idParam)
			uc.AccountService.AccountDelete(c, updateRequest, int(id))
			return c.JSON(http.StatusOK, echo.Map{
				"message": "deleted success",
			})
		}

		uc.AccountService.AccountDelete(c, updateRequest, int(id))

		return c.JSON(http.StatusOK, echo.Map{
			"message": "deleted success",
		})
	}
}
