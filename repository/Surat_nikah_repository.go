package repository

import "kk-requester/model/domain"

type SuratNikahRepository interface {
	GetSuratNikah(id uint) ([]domain.SuratNikah, error)
	CreateSuratNikah(NewSuratNikah *domain.SuratNikah, request *domain.ReqDetailSuratNikah) (*domain.SuratNikah, *domain.ReqDetailSuratNikah, error)
	SuratNikahUpdate(UpdatedSuratNikahSuratNikah *domain.SuratNikah, SuratNikahId float64, accountId uint) (*domain.SuratNikah, error)
}
