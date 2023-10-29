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

type RequestKKControllerImpl struct {
	RequestKKService service.RequestKKService
}

func NewRequestKKController(ks service.RequestKKService) *RequestKKControllerImpl {
	return &RequestKKControllerImpl{RequestKKService: ks}
}

func (kc *RequestKKControllerImpl) CreateRequestKK() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := helper.Authorization(c)
		RequestKK := &domain.RequestKK{}
		err := c.Bind(RequestKK)
		RequestKK.AccountId = uint(id)

		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "failed bind req kk",
			})
		}

		result, err := kc.RequestKKService.CreateRequestKK(c, RequestKK)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": "error creating RequestKK",
			})
		}

		response := web.KKRequestResponse{
			Id:                   result.ID,
			AccountId:            result.AccountId,
			StatusRequest:        result.StatusRequest,
			StatusRequestMessage: result.StatusRequestMessage,
		}

		return c.JSON(http.StatusCreated, echo.Map{
			"message": "success",
			"data":    response,
		})
	}

}

func (kc *RequestKKControllerImpl) RequestKKUpdate() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, role := helper.Authorization(c)

		if role == "admin" {
			idParam := c.Param("id")
			id, err := strconv.ParseFloat(idParam, 64)
			if err != nil {
				log.Fatal("Gagal convert id")
			}

			updateRequest := &domain.RequestKK{}
			err = c.Bind(updateRequest)
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]any{
					"messege": err.Error(),
				})
			}

			result, _ := kc.RequestKKService.RequestKKUpdate(c, updateRequest, id)

			return c.JSON(http.StatusCreated, echo.Map{
				"message": "success",
				"data":    result,
			})
		}
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "unauthorized",
		})
	}
}

func (kc *RequestKKControllerImpl) GetRequestKK() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, role := helper.Authorization(c)

		if role == "admin" {

			page := 1
			limit := 10

			pageParam := c.QueryParam("page")
			SortByParam := c.QueryParam("SortBy")
			orderParam := c.QueryParam("order")
			limitParam := c.QueryParam("limit")

			if pageParam != "" {
				page, _ = strconv.Atoi(pageParam)
			}
			if limitParam != "" {
				limit, _ = strconv.Atoi(limitParam)
			}
			if SortByParam == "" {
				SortByParam = "Updated_At"
			}
			if orderParam == "" {
				orderParam = "asc"
			}

			result, err := kc.RequestKKService.GetRequestKK(c, page, limit, SortByParam, orderParam)
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
