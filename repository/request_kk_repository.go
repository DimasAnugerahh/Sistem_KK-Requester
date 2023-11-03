package repository

import "kk-requester/model/domain"

type RequestKKRepository interface {
	GetRequestKK(page int, limit int, sortby string, order string) ([]domain.RequestKK, int, error)
	GetuserRequestKK(page int, limit int, sortby string, order string, accountId uint) ([]domain.RequestKK, int, error)

	RequestKKUpdate(UpdatedRequestKKRequestKK *domain.RequestKK, RequestKKId float64) (*domain.RequestKK, error)
	CreateRequestKK(NewkkRequest *domain.RequestKK) (*domain.RequestKK, error)
}
