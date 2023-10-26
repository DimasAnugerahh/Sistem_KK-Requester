package repository

import "kk-requester/model/domain"

type SuratKematianRepository interface {
	GetSuratKematian(id uint) ([]domain.SuratKematian, error)
	CreateSuratKematian(NewAccount *domain.SuratKematian) (*domain.SuratKematian, error)
	SuratKematianUpdate(UpdatedSuratKematianSuratKematian *domain.SuratKematian, SuratKematianId float64, accountId uint) (*domain.SuratKematian, error)
}
