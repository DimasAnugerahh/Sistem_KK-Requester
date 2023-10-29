package repository

import (
	"kk-requester/model/domain"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type SuratNikahRepositoryImpl struct {
	DB *gorm.DB
}

func NewSuratNikahRepository(db *gorm.DB) SuratNikahRepository {
	return &SuratNikahRepositoryImpl{DB: db}
}

func (kr *SuratNikahRepositoryImpl) CreateSuratNikah(NewSuratNikah *domain.SuratNikah, request *domain.ReqDetailSuratNikah) (*domain.SuratNikah, *domain.ReqDetailSuratNikah, error) {
	if err := kr.DB.Create(&NewSuratNikah).Error; err != nil {
		logrus.Error("Model: insert data error", err.Error())
		return nil, nil, err
	}
	request.Surat_nikah_id = int(NewSuratNikah.ID)
	if err := kr.DB.Create(&request).Error; err != nil {
		logrus.Error("Model: insert data error", err.Error())
		return nil, nil, err
	}
	return NewSuratNikah, request, nil
}

func (kr *SuratNikahRepositoryImpl) SuratNikahUpdate(UpdatedSuratNikah *domain.SuratNikah, SuratNikahId float64, accountId uint) (*domain.SuratNikah, error) {
	if err := kr.DB.Model(&domain.SuratNikah{}).Where("id=? and account_id=?", SuratNikahId, accountId).Updates(domain.SuratNikah{Nama_lengkap: UpdatedSuratNikah.Nama_lengkap, File_surat_nikah: UpdatedSuratNikah.File_surat_nikah}).Error; err != nil {
		logrus.Error("Model: update data error", err.Error())
		return nil, err
	}
	return UpdatedSuratNikah, nil
}

func (kr *SuratNikahRepositoryImpl) GetSuratNikah(id uint) ([]domain.SuratNikah, error) {
	var SuratNikah = []domain.SuratNikah{}
	if err := kr.DB.Where("account_id=?", id).Find(&SuratNikah).Error; err != nil {
		logrus.Error("Model: find data error", err.Error())
		return SuratNikah, err
	}
	return SuratNikah, nil
}
