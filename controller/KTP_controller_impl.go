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

type KTPControllerImpl struct {
	KTPService service.KTPService
}

func NewKTPController(ks service.KTPService) *KTPControllerImpl {
	return &KTPControllerImpl{KTPService: ks}
}

func (kc *KTPControllerImpl) CreateKTP() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := helper.Authorization(c)
		fileheader := "file_ktp"
		ktp := &domain.KTP{}
		err := c.Bind(ktp)

		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "failed bind data",
			})
		}

		ktp.File_ktp = helper.CloudinaryUpdload(c, fileheader)
		ktp.AccountId = uint(id)

		idParam := c.Param("id")
		idRequest, _ := strconv.Atoi(idParam)

		request := &domain.ReqDetailKtp{}
		err = c.Bind(request)

		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "failed bind request data",
			})
		}

		request.Request_kk_id = idRequest

		ktpResult, _, err := kc.KTPService.CreateKTP(c, ktp, request)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "error creating ktp",
			})
		}

		return c.JSON(http.StatusCreated, echo.Map{
			"message": "success",
			"data":    ktpResult,
		})
	}

}

func (kc *KTPControllerImpl) KTPUpdate() echo.HandlerFunc {
	return func(c echo.Context) error {

		accountId, _ := helper.Authorization(c)
		fileheader := "file_ktp"
		idParam := c.Param("id")

		id, err := strconv.ParseFloat(idParam, 64)
		if err != nil {
			log.Fatal("Gagal convert id")
		}

		updateRequest := &domain.KTP{}
		err = c.Bind(updateRequest)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"messege": err.Error(),
			})
		}

		updateRequest.File_ktp = helper.CloudinaryUpdload(c, fileheader)

		result, _ := kc.KTPService.KTPUpdate(c, updateRequest, id, uint(accountId))

		return c.JSON(http.StatusCreated, echo.Map{
			"message": "success",
			"data":    result,
		})
	}
}

func (kc *KTPControllerImpl) GetKTP() echo.HandlerFunc {
	return func(c echo.Context) error {
		accountId, _ := helper.Authorization(c)
		result, err := kc.KTPService.GetKTP(c, uint(accountId))
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
