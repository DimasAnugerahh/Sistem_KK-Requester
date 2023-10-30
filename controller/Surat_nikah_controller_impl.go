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

type SuratNikahControllerImpl struct {
	SuratNikahService service.SuratNikahService
}

func NewSuratNikahController(ks service.SuratNikahService) *SuratNikahControllerImpl {
	return &SuratNikahControllerImpl{SuratNikahService: ks}
}

func (kc *SuratNikahControllerImpl) CreateSuratNikah() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := helper.Authorization(c)
		fileheader := "file_surat_nikah"
		SuratNikah := &domain.SuratNikah{}
		err := c.Bind(SuratNikah)

		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "failed bind data",
			})
		}

		SuratNikah.File_surat_nikah = helper.CloudinaryUpdload(c, fileheader)
		SuratNikah.AccountId = uint(id)

		idParam := c.Param("id")
		idRequest, _ := strconv.Atoi(idParam)

		request := &domain.ReqDetailSuratNikah{}
		err = c.Bind(request)

		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "failed bind request data",
			})
		}

		request.Request_kk_id = idRequest

		suratNikah, _, err := kc.SuratNikahService.CreateSuratNikah(c, SuratNikah, request)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "error creating SuratNikah",
			})
		}

		return c.JSON(http.StatusCreated, echo.Map{
			"message": "success",
			"data":    suratNikah,
		})
	}
}
func (kc *SuratNikahControllerImpl) SuratNikahUpdate() echo.HandlerFunc {
	return func(c echo.Context) error {

		accountId, _ := helper.Authorization(c)
		fileheader := "file_surat_nikah"
		idParam := c.Param("id")

		id, err := strconv.ParseFloat(idParam, 64)
		if err != nil {
			log.Fatal("Gagal convert id")
		}

		updateRequest := domain.SuratNikah{}
		err = c.Bind(&updateRequest)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"messege": err.Error(),
			})
		}

		updateRequest.File_surat_nikah = helper.CloudinaryUpdload(c, fileheader)

		result, _ := kc.SuratNikahService.SuratNikahUpdate(c, &updateRequest, id, uint(accountId))
		result.AccountId = uint(accountId)
		return c.JSON(http.StatusOK, echo.Map{
			"message": "success",
			"data":    result,
		})
	}
}

func (kc *SuratNikahControllerImpl) GetSuratNikah() echo.HandlerFunc {
	return func(c echo.Context) error {
		accountId, _ := helper.Authorization(c)
		result, err := kc.SuratNikahService.GetSuratNikah(c, uint(accountId))
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
