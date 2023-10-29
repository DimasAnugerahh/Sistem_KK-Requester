package controller

import (
	"kk-requester/helper"
	"kk-requester/model/domain"
	"kk-requester/service"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AktaKematianControllerImpl struct {
	AktaKematianService service.AktaKematianService
}

func NewAktaKematianController(ks service.AktaKematianService) *AktaKematianControllerImpl {
	return &AktaKematianControllerImpl{AktaKematianService: ks}
}

func (kc *AktaKematianControllerImpl) CreateAktaKematian() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := helper.Authorization(c)
		fileheader := "file_akta_Kematian"
		AktaKematian := &domain.AktaKematian{}
		err := c.Bind(AktaKematian)

		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "failed bind data",
			})
		}

		AktaKematian.File_akta_Kematian = helper.CloudinaryUpdload(c, fileheader)
		AktaKematian.AccountId = uint(id)

		idParam := c.Param("id")
		idRequest, _ := strconv.Atoi(idParam)

		request := &domain.ReqDetailAktaKematian{}
		err = c.Bind(request)

		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "failed bind request data",
			})
		}

		request.Request_kk_id = idRequest

		result, _, err := kc.AktaKematianService.CreateAktaKematian(c, AktaKematian, request)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "error creating AktaKematian",
			})
		}
		return c.JSON(http.StatusCreated, echo.Map{
			"message": "success",
			"data":    result,
		})
	}

}

func (kc *AktaKematianControllerImpl) AktaKematianUpdate() echo.HandlerFunc {
	return func(c echo.Context) error {

		accountId, _ := helper.Authorization(c)
		fileheader := "file_akta_Kematian"

		idParam := c.Param("id")

		id, err := strconv.ParseFloat(idParam, 64)
		if err != nil {
			log.Fatal("Gagal convert id")
		}

		updateRequest := domain.AktaKematian{}
		err = c.Bind(&updateRequest)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"messege": err.Error(),
			})
		}

		updateRequest.File_akta_Kematian = helper.CloudinaryUpdload(c, fileheader)

		result, _ := kc.AktaKematianService.AktaKematianUpdate(c, &updateRequest, id, uint(accountId))
		result.AccountId = uint(accountId)
		return c.JSON(http.StatusCreated, echo.Map{
			"message": "success",
			"data":    result,
		})
	}
}

func (kc *AktaKematianControllerImpl) GetAktaKematian() echo.HandlerFunc {
	return func(c echo.Context) error {
		accountId, _ := helper.Authorization(c)
		result, err := kc.AktaKematianService.GetAktaKematian(c, uint(accountId))
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
}
