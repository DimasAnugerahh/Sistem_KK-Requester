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

		RequestKK.AccountId = uint(id)

		result, err := kc.RequestKKService.CreateRequestKK(c, RequestKK)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"message": err.Error(),
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
				log.Fatal("Gagal convert id", err.Error())
			}

			updateRequest := &domain.RequestKK{}
			err = c.Bind(updateRequest)
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]any{
					"messege": err.Error(),
				})
			}

			result, _ := kc.RequestKKService.RequestKKUpdate(c, updateRequest, id)

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

func (kc *RequestKKControllerImpl) GetRequestKK() echo.HandlerFunc {
	return func(c echo.Context) error {
		_, role := helper.Authorization(c)

		if role == "admin" {

			page := 1
			limit := 10
			previousPage := page - 1
			nextpages := page + 1

			pageParam := c.QueryParam("page")
			SortByParam := c.QueryParam("SortBy")
			orderParam := c.QueryParam("order")
			limitParam := c.QueryParam("limit")

			if pageParam != "" {
				page, _ = strconv.Atoi(pageParam)
				previousPage = page - 1

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
			if page == 1 {
				previousPage = 1
			}

			result, totalPages, err := kc.RequestKKService.GetRequestKK(c, page, limit, SortByParam, orderParam)
			if err != nil {
				return c.JSON(http.StatusBadRequest, map[string]any{
					"messege": err.Error(),
				})
			}

			if len(result) == 0 {
				return c.JSON(http.StatusBadRequest, map[string]any{
					"messege": "there is no request",
				})
			}
			if page == totalPages {
				nextpages = page
			}

			return c.JSON(http.StatusOK, echo.Map{
				"message": "success",
				"data":    result,
				"pagination": []map[string]any{
					{
						"previousPage": previousPage,
						"current_page": page,
						"next_page":    nextpages,
						"total_pages":  totalPages,
					},
				},
			})
		}
		return c.JSON(http.StatusUnauthorized, echo.Map{
			"message": "unauthorized",
		})

	}
}

func (kc *RequestKKControllerImpl) GetUserRequestKK() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := helper.Authorization(c)

		page := 1
		limit := 10
		previousPage := page - 1
		nextpages := page + 1

		pageParam := c.QueryParam("page")
		SortByParam := c.QueryParam("SortBy")
		orderParam := c.QueryParam("order")
		limitParam := c.QueryParam("limit")

		if pageParam != "" {
			page, _ = strconv.Atoi(pageParam)
			previousPage = page - 1
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
		if page == 1 {
			previousPage = 1
		}

		result, totalPages, err := kc.RequestKKService.GetUserRequestKK(c, page, limit, SortByParam, orderParam, uint(id))
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"messege": err.Error(),
			})
		}

		if page == totalPages {
			nextpages = page
		}

		if len(result) == 0 {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"messege": "there is no request",
			})
		}
		return c.JSON(http.StatusOK, echo.Map{
			"message": "success",
			"data":    result,
			"pagination": []map[string]any{
				{
					"previousPage": previousPage,
					"current_page": page,
					"next_page":    nextpages,
					"total_pages":  totalPages,
				},
			},
		})
	}

}
