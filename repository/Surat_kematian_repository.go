package repository

import "kk-requester/model/domain"

type AktaKematianRepository interface {
	GetAktaKematian(id uint) ([]domain.AktaKematian, error)
	CreateAktaKematian(NewAccount *domain.AktaKematian, request *domain.ReqDetailAktaKematian) (*domain.AktaKematian, *domain.ReqDetailAktaKematian, error)
	AktaKematianUpdate(UpdatedAktaKematianAktaKematian *domain.AktaKematian, AktaKematianId float64, accountId uint) (*domain.AktaKematian, error)
}
