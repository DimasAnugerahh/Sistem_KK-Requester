package repository

import "kk-requester/model/domain"

type SuratPindahRepository interface {
	GetSuratPindah(id uint) ([]domain.SuratPindah, error)
	CreateSuratPindah(NewAccount *domain.SuratPindah) (*domain.SuratPindah, error)
	SuratPindahUpdate(UpdatedSuratPindahSuratPindah *domain.SuratPindah, SuratPindahId float64, accountId uint) (*domain.SuratPindah, error)
}
