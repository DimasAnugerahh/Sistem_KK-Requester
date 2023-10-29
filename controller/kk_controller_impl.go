package controller

import (
	"kk-requester/helper"
	"kk-requester/model/domain"
	"kk-requester/service"
	"net/http"

	"github.com/labstack/echo/v4"
)

type KkContorollerImpl struct {
	KKService service.KkService
}

func NewKKController(ks service.KkService) *KkContorollerImpl {
	return &KkContorollerImpl{KKService: ks}
}

func (kc *KkContorollerImpl) GetKK() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, role := helper.Authorization(c)
		if role == "admin" {
			result, err := kc.KKService.GetKK(c)
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]any{
					"messege": err.Error(),
				})
			}
			return c.JSON(http.StatusOK, echo.Map{
				"message": "success",
				"data":    result,
			})
		}
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "unauthorized",
		})
	}
}

func (kc *KkContorollerImpl) CreateKK() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, role := helper.Authorization(c)

		if role == "admin" {
			kk := &domain.KK{}
			err := c.Bind(kk)
			fileheader := "file_kk"

			if err != nil {
				return c.JSON(http.StatusBadRequest, echo.Map{
					"message": "failed bind request data",
				})
			}

			kk.AccountId = uint(id)
			kk.FileKK = helper.CloudinaryUpdload(c, fileheader)
			response, _ := kc.KKService.CreateKK(c, kk)

			helper.EmailSender(response.Email, response.FileKK)

			if err != nil {
				return c.JSON(http.StatusInternalServerError, echo.Map{
					"message": err.Error(),
				})
			}

			return c.JSON(http.StatusCreated, echo.Map{
				"message": "success",
				"data":    response,
			})
		}
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "unauthorized",
		})
	}
}
