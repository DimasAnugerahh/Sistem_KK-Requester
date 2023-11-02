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
				"message": err.Error(),
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
				"message": err.Error(),
			})
		}

		request.Request_kk_id = idRequest

		result, _, err := kc.KTPService.CreateKTP(c, ktp, request)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		response := web.DocumentResponse{Id: result.ID, CreatedAt: result.CreatedAt, UpdatedAt: result.UpdatedAt, DeletedAt: result.DeletedAt.Time, AccountId: result.AccountId, Nama: result.Nama_lengkap, File: result.File_ktp}

		return c.JSON(http.StatusCreated, echo.Map{
			"message": "success",
			"data":    response,
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
			log.Fatal("Gagal convert id", err.Error())
		}

		updateRequest := &domain.KTP{}
		err = c.Bind(updateRequest)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"messege": err.Error(),
			})
		}

		updateRequest.File_ktp = helper.CloudinaryUpdload(c, fileheader)
		updateRequest.AccountId = uint(accountId)

		result, _ := kc.KTPService.KTPUpdate(c, updateRequest, id, uint(accountId))

		return c.JSON(http.StatusOK, echo.Map{
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
		if len(result) == 0 {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"messege": "there is no ktp",
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
					File:      result[idx].File_ktp,
				})
		}

		return c.JSON(http.StatusOK, echo.Map{
			"message": "success",
			"data":    response,
		})
	}
}
