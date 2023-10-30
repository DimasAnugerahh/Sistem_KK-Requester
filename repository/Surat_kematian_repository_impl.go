package repository

import (
	"kk-requester/model/domain"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AktaKematianRepositoryImpl struct {
	DB *gorm.DB
}

func NewAktaKematianRepository(db *gorm.DB) AktaKematianRepository {
	return &AktaKematianRepositoryImpl{DB: db}
}

func (kr *AktaKematianRepositoryImpl) CreateAktaKematian(NewAktaKematian *domain.AktaKematian, request *domain.ReqDetailAktaKematian) (*domain.AktaKematian, *domain.ReqDetailAktaKematian, error) {
	if err := kr.DB.Create(&NewAktaKematian).Error; err != nil {
		logrus.Error("Model: insert data error", err.Error())
		return nil, nil, err
	}
	request.Akta_kematian_id = int(NewAktaKematian.ID)
	if err := kr.DB.Create(&request).Error; err != nil {
		logrus.Error("Model: insert data error", err.Error())
		return nil, nil, err
	}
	return NewAktaKematian, request, nil
}

func (kr *AktaKematianRepositoryImpl) AktaKematianUpdate(UpdatedAktaKematian *domain.AktaKematian, AktaKematianId float64, accountId uint) (*domain.AktaKematian, error) {
	if err := kr.DB.Model(&domain.AktaKematian{}).Where("id=? and account_id=?", AktaKematianId, accountId).Updates(domain.AktaKematian{Nama_lengkap: UpdatedAktaKematian.Nama_lengkap, File_akta_Kematian: UpdatedAktaKematian.File_akta_Kematian}).Error; err != nil {
		logrus.Error("Model: update data error", err.Error())
		return nil, err
	}
	return UpdatedAktaKematian, nil
}

func (kr *AktaKematianRepositoryImpl) GetAktaKematian(id uint) ([]domain.AktaKematian, error) {
	var AktaKematian = []domain.AktaKematian{}
	if err := kr.DB.Where("account_id=?", id).Find(&AktaKematian).Error; err != nil {
		logrus.Error("Model: find data error", err.Error())
		return AktaKematian, err
	}
	return AktaKematian, nil
}
