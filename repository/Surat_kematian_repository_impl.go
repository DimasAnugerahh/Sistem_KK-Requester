package repository

import (
	"kk-requester/model/domain"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type SuratKematianRepositoryImpl struct {
	DB *gorm.DB
}

func NewSuratKematianRepository(db *gorm.DB) SuratKematianRepository {
	return &SuratKematianRepositoryImpl{DB: db}
}

func (kr *SuratKematianRepositoryImpl) CreateSuratKematian(NewSuratKematian *domain.SuratKematian) (*domain.SuratKematian, error) {
	if err := kr.DB.Create(&NewSuratKematian).Error; err != nil {
		logrus.Error("Model: insert data error", err.Error())
		return nil, err
	}
	return NewSuratKematian, nil
}

func (kr *SuratKematianRepositoryImpl) SuratKematianUpdate(UpdatedSuratKematian *domain.SuratKematian, SuratKematianId float64, accountId uint) (*domain.SuratKematian, error) {
	if err := kr.DB.Model(&domain.SuratKematian{}).Where("id=? and account_id=?", SuratKematianId, accountId).Updates(domain.SuratKematian{Nama_lengkap: UpdatedSuratKematian.Nama_lengkap, File_Surat_Kematian: UpdatedSuratKematian.File_Surat_Kematian}).Error; err != nil {
		logrus.Error("Model: update data error", err.Error())
		return nil, err
	}
	return UpdatedSuratKematian, nil
}

func (kr *SuratKematianRepositoryImpl) GetSuratKematian(id uint) ([]domain.SuratKematian, error) {
	var SuratKematian = []domain.SuratKematian{}
	if err := kr.DB.Where("account_id=?", id).Find(&SuratKematian).Error; err != nil {
		logrus.Error("Model: find data error", err.Error())
		return SuratKematian, err
	}
	return SuratKematian, nil
}
