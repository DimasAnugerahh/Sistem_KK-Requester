package repository

import "kk-requester/model/domain"

type SuratNikahRepository interface {
	GetSuratNikah(id uint) ([]domain.SuratNikah, error)
	CreateSuratNikah(NewAccount *domain.SuratNikah) (*domain.SuratNikah, error)
	SuratNikahUpdate(UpdatedSuratNikahSuratNikah *domain.SuratNikah, SuratNikahId float64, accountId uint) (*domain.SuratNikah, error)
}
