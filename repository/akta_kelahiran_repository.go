package repository

import "kk-requester/model/domain"

type AktaKelahiranRepository interface {
	GetAktaKelahiran(id uint) ([]domain.AktaKelahiran, error)
	CreateAktaKelahiran(NewAktaKelahiran *domain.AktaKelahiran, request *domain.ReqDetailAktaKelahiran) (*domain.AktaKelahiran, *domain.ReqDetailAktaKelahiran, error)
	AktaKelahiranUpdate(UpdatedAktaKelahiranAktaKelahiran *domain.AktaKelahiran, AktaKelahiranId float64, accountId uint) (*domain.AktaKelahiran, error)
}
