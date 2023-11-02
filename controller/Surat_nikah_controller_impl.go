package controller

import (
	"kk-requester/helper"
	"kk-requester/model/domain"
	"kk-requester/model/web"
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
				"message": err.Error(),
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
				"message": err.Error(),
			})
		}

		request.Request_kk_id = idRequest

		result, _, err := kc.SuratNikahService.CreateSuratNikah(c, SuratNikah, request)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}
		response := web.DocumentResponse{Id: result.ID, CreatedAt: result.CreatedAt, UpdatedAt: result.UpdatedAt, DeletedAt: result.DeletedAt.Time, AccountId: result.AccountId, Nama: result.Nama_lengkap, File: result.File_surat_nikah}

		return c.JSON(http.StatusCreated, echo.Map{
			"message": "success",
			"data":    response,
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
			log.Fatal("Gagal convert id", err.Error())
		}

		updateRequest := domain.SuratNikah{}
		err = c.Bind(&updateRequest)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"messege": err.Error(),
			})
		}

		updateRequest.File_surat_nikah = helper.CloudinaryUpdload(c, fileheader)
		updateRequest.AccountId = uint(accountId)

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

		if len(result) == 0 {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"messege": "there is no surat nikah",
			})
		}

		response := []web.DocumentResponse{}
		for idx := range result {
			response = append(response,
				web.DocumentResponse{
					Id:        result[idx].ID,
					Nama:      result[idx].Nama_lengkap,
					CreatedAt: result[idx].CreatedAt,
					UpdatedAt: result[idx].UpdatedAt,
					DeletedAt: result[idx].DeletedAt.Time,
					AccountId: result[idx].AccountId,
					File:      result[idx].File_surat_nikah,
				})
		}

		return c.JSON(http.StatusOK, echo.Map{
			"message": "success",
			"data":    response,
		})
	}
}
