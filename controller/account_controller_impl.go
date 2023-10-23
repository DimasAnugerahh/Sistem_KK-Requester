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

	return helper.RoleAuth(func(c echo.Context) error {
		response, err := uc.AccountService.GetAllAccounts(c)

		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": err,
			})
		}

		return c.JSON(http.StatusOK, echo.Map{
			"message": "success",
			"data":    response,
		})
	})

}

func (uc *AccountControllerImpl) CreateAccount() echo.HandlerFunc {
	return func(c echo.Context) error {
		Account := &domain.Account{}
		err := c.Bind(Account)

		if err != nil {
			if err != nil {
				return c.JSON(http.StatusBadRequest, echo.Map{
					"message": err,
				})
			}
		}

		response, err := uc.AccountService.CreateAccount(c, Account)

		if err != nil {
			if err != nil {
				return c.JSON(http.StatusBadRequest, echo.Map{
					"message": err,
				})
			}
		}

		return c.JSON(http.StatusOK, echo.Map{
			"message": "success",
			"data":    response,
		})
	}
}

func (uc *AccountControllerImpl) AccountLogin() echo.HandlerFunc {

	return func(c echo.Context) error {
		Account := &domain.Account{}
		err := c.Bind(Account)

		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": err,
			})
		}

		response, err := uc.AccountService.AccountLogin(c, Account)
		if err != nil {
			if err != nil {
				return c.JSON(http.StatusBadRequest, echo.Map{
					"message": err,
				})
			}
		}

		token, err := helper.GenerateToken(Account, int(response.ID), response.Role, response.Email)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": err,
			})
		}

		AccountResponse := web.LoginResponse{Email: response.Email, Nama: response.Name, Role: response.Role, Token: token}

		return c.JSON(http.StatusOK, echo.Map{
			"message": "success",
			"data":    AccountResponse,
		})
	}
}

func (uc *AccountControllerImpl) AccountUpdate() echo.HandlerFunc {
	return helper.IdAuth(func(c echo.Context) error {

		idParam := c.Param("id")
		id, err := strconv.ParseFloat(idParam, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"messege": "invalid id",
			})
		}

		updateRequest := domain.Account{}
		err = c.Bind(&updateRequest)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"messege": err.Error(),
			})
		}

		response, _ := uc.AccountService.AccountUpdate(c, &updateRequest, id)

		return c.JSON(http.StatusOK, echo.Map{
			"message": "success",
			"data":    response,
		})

	})
}

func (uc *AccountControllerImpl) AccountDelete() echo.HandlerFunc {
	return helper.RoleAuth(func(c echo.Context) error {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"messege": "invalid id",
			})
		}

		updateRequest := domain.Account{}
		err = c.Bind(&updateRequest)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"messege": err.Error(),
			})
		}

		_, err = uc.AccountService.AccountDelete(c, &updateRequest, id)

		return c.JSON(http.StatusOK, echo.Map{
			"message": "deleted success",
		})
	})
}
