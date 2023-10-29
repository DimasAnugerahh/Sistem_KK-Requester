package repository

import "kk-requester/model/domain"

type KTPRepository interface {
	GetKTP(idAccount uint) ([]domain.KTP, error)
	CreateKTP(NewAccount *domain.KTP, request *domain.ReqDetailKtp) (*domain.KTP, *domain.ReqDetailKtp, error)
	KTPUpdate(UpdatedKTPKTP *domain.KTP, ktpId float64, accountId uint) (*domain.KTP, error)
}
