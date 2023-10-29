package repository

import "kk-requester/model/domain"

type SuratPindahRepository interface {
	GetSuratPindah(id uint) ([]domain.SuratPindah, error)
	CreateSuratPindah(NewSuratPindah *domain.SuratPindah, request *domain.ReqDetailSuratPindah) (*domain.SuratPindah, *domain.ReqDetailSuratPindah, error)
	SuratPindahUpdate(UpdatedSuratPindahSuratPindah *domain.SuratPindah, SuratPindahId float64, accountId uint) (*domain.SuratPindah, error)
}
