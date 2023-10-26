package repository

import "kk-requester/model/domain"

type KTPRepository interface {
	GetKTP(id uint) ([]domain.KTP, error)
	CreateKTP(NewAccount *domain.KTP) (*domain.KTP, error)
	KTPUpdate(UpdatedKTPKTP *domain.KTP, ktpId float64, accountId uint) (*domain.KTP, error)
}
