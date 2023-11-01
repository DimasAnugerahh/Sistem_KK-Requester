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

type AktaKelahiranControllerImpl struct {
	AktaKelahiranService service.AktaKelahiranService
}

func NewAktaKelahiranController(ks service.AktaKelahiranService) *AktaKelahiranControllerImpl {
	return &AktaKelahiranControllerImpl{AktaKelahiranService: ks}
}

func (kc *AktaKelahiranControllerImpl) CreateAktaKelahiran() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := helper.Authorization(c)
		fileheader := "file_Akta_kelahiran"
		AktaKelahiran := &domain.AktaKelahiran{}
		err := c.Bind(AktaKelahiran)

		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		AktaKelahiran.File_Akta_kelahiran = helper.CloudinaryUpdload(c, fileheader)
		AktaKelahiran.AccountId = uint(id)

		idParam := c.Param("id")
		idRequest, _ := strconv.Atoi(idParam)

		request := &domain.ReqDetailAktaKelahiran{}
		err = c.Bind(request)

		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		request.Request_kk_id = idRequest

		result, _, err := kc.AktaKelahiranService.CreateAktaKelahiran(c, AktaKelahiran, request)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
			})
		}

		response := web.AktaKelahiranResponse{CreatedAt: result.CreatedAt, UpdatedAt: result.UpdatedAt, DeletedAt: result.DeletedAt.Time, Nama_lengkap: result.Nama_lengkap, AccountId: result.AccountId, File_Akta_kelahiran: result.File_Akta_kelahiran}

		return c.JSON(http.StatusCreated, echo.Map{
			"message": "success",
			"data":    response,
		})
	}

}

func (kc *AktaKelahiranControllerImpl) AktaKelahiranUpdate() echo.HandlerFunc {
	return func(c echo.Context) error {

		accountId, _ := helper.Authorization(c)
		fileheader := "file_Akta_kelahiran"

		idParam := c.Param("id")

		id, err := strconv.ParseFloat(idParam, 64)
		if err != nil {
			log.Fatal("Gagal convert id")
		}

		updateRequest := &domain.AktaKelahiran{}
		err = c.Bind(updateRequest)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"messege": err.Error(),
			})
		}

		updateRequest.File_Akta_kelahiran = helper.CloudinaryUpdload(c, fileheader)

		result, _ := kc.AktaKelahiranService.AktaKelahiranUpdate(c, updateRequest, id, uint(accountId))

		response := web.AktaKelahiranResponse{CreatedAt: result.CreatedAt, UpdatedAt: result.UpdatedAt, DeletedAt: result.DeletedAt.Time, Nama_lengkap: result.Nama_lengkap, AccountId: result.AccountId, File_Akta_kelahiran: result.File_Akta_kelahiran}

		return c.JSON(http.StatusOK, echo.Map{
			"message": "success",
			"data":    response,
		})
	}
}

func (kc *AktaKelahiranControllerImpl) GetAktaKelahiran() echo.HandlerFunc {
	return func(c echo.Context) error {
		accountId, _ := helper.Authorization(c)
		result, err := kc.AktaKelahiranService.GetAktaKelahiran(c, uint(accountId))

		if len(result) == 0 {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"messege": "there is no akta kelahiran",
			})
		}
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
