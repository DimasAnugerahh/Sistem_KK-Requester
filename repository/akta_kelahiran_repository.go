package repository

import "kk-requester/model/domain"

type AktaKelahiranRepository interface {
	GetAktaKelahiran(id uint) ([]domain.AktaKelahiran, error)
	CreateAktaKelahiran(NewAccount *domain.AktaKelahiran) (*domain.AktaKelahiran, error)
	AktaKelahiranUpdate(UpdatedAktaKelahiranAktaKelahiran *domain.AktaKelahiran, AktaKelahiranId float64, accountId uint) (*domain.AktaKelahiran, error)
}
