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

type SuratPindahControllerImpl struct {
	SuratPindahService service.SuratPindahService
}

func NewSuratPindahController(ks service.SuratPindahService) *SuratPindahControllerImpl {
	return &SuratPindahControllerImpl{SuratPindahService: ks}
}

func (kc *SuratPindahControllerImpl) CreateSuratPindah() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := helper.Authorization(c)
		fileheader := "file_surat_pindah"
		SuratPindah := &domain.SuratPindah{}
		err := c.Bind(SuratPindah)

		if err != nil {
			if err != nil {
				return c.JSON(http.StatusBadRequest, echo.Map{
					"message": "failed bind data",
				})
			}
		}

		SuratPindah.File_surat_pindah = helper.CloudinaryUpdload(c, fileheader)
		SuratPindah.AccountId = uint(id)

		result, err := kc.SuratPindahService.CreateSuratPindah(c, SuratPindah)
		if err != nil {
			if err != nil {
				return c.JSON(http.StatusBadRequest, echo.Map{
					"message": "error creating SuratPindah",
				})
			}
		}
		return c.JSON(http.StatusCreated, echo.Map{
			"message": "success",
			"data":    result,
		})
	}

}

func (kc *SuratPindahControllerImpl) SuratPindahUpdate() echo.HandlerFunc {
	return func(c echo.Context) error {

		accountId, _ := helper.Authorization(c)
		fileheader := "file_surat_pindah"
		idParam := c.Param("id")

		id, err := strconv.ParseFloat(idParam, 64)
		if err != nil {
			log.Fatal("Gagal convert id")
		}

		updateRequest := domain.SuratPindah{}
		err = c.Bind(&updateRequest)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"messege": err.Error(),
			})
		}

		updateRequest.File_surat_pindah = helper.CloudinaryUpdload(c, fileheader)

		result, _ := kc.SuratPindahService.SuratPindahUpdate(c, &updateRequest, id, uint(accountId))

		return c.JSON(http.StatusCreated, echo.Map{
			"message": "success",
			"data":    result,
		})
	}
}

func (kc *SuratPindahControllerImpl) GetSuratPindah() echo.HandlerFunc {
	return func(c echo.Context) error {
		accountId, _ := helper.Authorization(c)
		result, err := kc.SuratPindahService.GetSuratPindah(c, uint(accountId))
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
