package repository

import (
	"kk-requester/model/domain"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type AktaKelahiranRepositoryImpl struct {
	DB *gorm.DB
}

func NewAktaKelahiranRepository(db *gorm.DB) AktaKelahiranRepository {
	return &AktaKelahiranRepositoryImpl{DB: db}
}

func (kr *AktaKelahiranRepositoryImpl) CreateAktaKelahiran(NewAktaKelahiran *domain.AktaKelahiran) (*domain.AktaKelahiran, error) {
	if err := kr.DB.Create(&NewAktaKelahiran).Error; err != nil {
		logrus.Error("Model: insert data error", err.Error())
		return nil, err
	}
	return NewAktaKelahiran, nil
}

func (kr *AktaKelahiranRepositoryImpl) AktaKelahiranUpdate(UpdatedAktaKelahiran *domain.AktaKelahiran, AktaKelahiranId float64, accountId uint) (*domain.AktaKelahiran, error) {
	if err := kr.DB.Model(&domain.AktaKelahiran{}).Where("id=? and account_id=?", AktaKelahiranId, accountId).Updates(domain.AktaKelahiran{Nama_lengkap: UpdatedAktaKelahiran.Nama_lengkap, File_Akta_kelahiran: UpdatedAktaKelahiran.File_Akta_kelahiran}).Error; err != nil {
		logrus.Error("Model: update data error", err.Error())
		return nil, err
	}
	return UpdatedAktaKelahiran, nil
}

func (kr *AktaKelahiranRepositoryImpl) GetAktaKelahiran(id uint) ([]domain.AktaKelahiran, error) {
	var AktaKelahiran = []domain.AktaKelahiran{}
	if err := kr.DB.Where("account_id=?", id).Find(&AktaKelahiran).Error; err != nil {
		logrus.Error("Model: find data error", err.Error())
		return AktaKelahiran, err
	}
	return AktaKelahiran, nil
}
