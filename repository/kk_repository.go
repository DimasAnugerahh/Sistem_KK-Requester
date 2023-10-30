package repository

import "kk-requester/model/domain"

type KkRepository interface {
	GetKK() ([]domain.KK, error)
	CreateKK(NewKK *domain.KK) (*domain.KK, error)
}
