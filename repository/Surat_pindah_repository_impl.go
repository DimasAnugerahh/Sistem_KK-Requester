package repository

import (
	"kk-requester/model/domain"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type SuratPindahRepositoryImpl struct {
	DB *gorm.DB
}

func NewSuratPindahRepository(db *gorm.DB) SuratPindahRepository {
	return &SuratPindahRepositoryImpl{DB: db}
}

func (kr *SuratPindahRepositoryImpl) CreateSuratPindah(NewSuratPindah *domain.SuratPindah) (*domain.SuratPindah, error) {
	if err := kr.DB.Create(&NewSuratPindah).Error; err != nil {
		logrus.Error("Model: insert data error", err.Error())
		return nil, err
	}
	return NewSuratPindah, nil
}

func (kr *SuratPindahRepositoryImpl) SuratPindahUpdate(UpdatedSuratPindah *domain.SuratPindah, SuratPindahId float64, accountId uint) (*domain.SuratPindah, error) {
	if err := kr.DB.Model(&domain.SuratPindah{}).Where("id=? and account_id=?", SuratPindahId, accountId).Updates(domain.SuratPindah{Nama_lengkap: UpdatedSuratPindah.Nama_lengkap, File_surat_pindah: UpdatedSuratPindah.File_surat_pindah}).Error; err != nil {
		logrus.Error("Model: update data error", err.Error())
		return nil, err
	}
	return UpdatedSuratPindah, nil
}

func (kr *SuratPindahRepositoryImpl) GetSuratPindah(id uint) ([]domain.SuratPindah, error) {
	var SuratPindah = []domain.SuratPindah{}
	if err := kr.DB.Where("account_id=?", id).Find(&SuratPindah).Error; err != nil {
		logrus.Error("Model: find data error", err.Error())
		return SuratPindah, err
	}
	return SuratPindah, nil
}
