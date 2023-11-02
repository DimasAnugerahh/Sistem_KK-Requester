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
		response := web.DocumentResponse{Id: result.ID, CreatedAt: result.CreatedAt, UpdatedAt: result.UpdatedAt, DeletedAt: result.DeletedAt.Time, AccountId: result.AccountId, Nama: result.Nama_lengkap, File: result.File_akta_Kematian}

		return c.JSON(http.StatusCreated, echo.Map{
			"message": "success",
			"data":    response,
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
		updateRequest.AccountId = uint(accountId)

		result, _ := kc.AktaKematianService.AktaKematianUpdate(c, &updateRequest, id, uint(accountId))
		result.AccountId = uint(accountId)
		return c.JSON(http.StatusOK, echo.Map{
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
		if len(result) == 0 {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"messege": "there is no surat kematian",
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
					File:      result[idx].File_akta_Kematian,
				})
		}
		return c.JSON(http.StatusOK, echo.Map{
			"message": "success",
			"data":    response,
		})
	}
}
