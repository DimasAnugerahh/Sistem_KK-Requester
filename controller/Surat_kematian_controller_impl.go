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

type SuratKematianControllerImpl struct {
	SuratKematianService service.SuratKematianService
}

func NewSuratKematianController(ks service.SuratKematianService) *SuratKematianControllerImpl {
	return &SuratKematianControllerImpl{SuratKematianService: ks}
}

func (kc *SuratKematianControllerImpl) CreateSuratKematian() echo.HandlerFunc {
	return func(c echo.Context) error {
		fileheader := "file_Surat_Kematian"
		id, _ := helper.Authorization(c)
		SuratKematian := &domain.SuratKematian{}
		err := c.Bind(SuratKematian)

		if err != nil {
			if err != nil {
				return c.JSON(http.StatusBadRequest, echo.Map{
					"message": "failed bind data",
				})
			}
		}

		SuratKematian.File_Surat_Kematian = helper.CloudinaryUpdload(c, fileheader)
		SuratKematian.AccountId = uint(id)

		result, err := kc.SuratKematianService.CreateSuratKematian(c, SuratKematian)
		if err != nil {
			if err != nil {
				return c.JSON(http.StatusBadRequest, echo.Map{
					"message": "error creating SuratKematian",
				})
			}
		}
		return c.JSON(http.StatusCreated, echo.Map{
			"message": "success",
			"data":    result,
		})
	}

}

func (kc *SuratKematianControllerImpl) SuratKematianUpdate() echo.HandlerFunc {
	return func(c echo.Context) error {

		accountId, _ := helper.Authorization(c)
		fileheader := "file_Surat_Kematian"

		idParam := c.Param("id")

		id, err := strconv.ParseFloat(idParam, 64)
		if err != nil {
			log.Fatal("Gagal convert id")
		}

		updateRequest := domain.SuratKematian{}
		err = c.Bind(&updateRequest)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"messege": err.Error(),
			})
		}

		updateRequest.File_Surat_Kematian = helper.CloudinaryUpdload(c, fileheader)

		result, _ := kc.SuratKematianService.SuratKematianUpdate(c, &updateRequest, id, uint(accountId))

		return c.JSON(http.StatusCreated, echo.Map{
			"message": "success",
			"data":    result,
		})
	}
}

func (kc *SuratKematianControllerImpl) GetSuratKematian() echo.HandlerFunc {
	return func(c echo.Context) error {
		accountId, _ := helper.Authorization(c)
		result, err := kc.SuratKematianService.GetSuratKematian(c, uint(accountId))
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
