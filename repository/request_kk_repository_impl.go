package repository

import (
	"kk-requester/model/domain"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type RequestKKRepositoryImpl struct {
	DB *gorm.DB
}

func NewRequestKKRepository(db *gorm.DB) RequestKKRepository {
	return &RequestKKRepositoryImpl{DB: db}
}

func (kr *RequestKKRepositoryImpl) CreateRequestKK(

	NewkkRequest *domain.RequestKK,

) (*domain.RequestKK, error) {

	if err := kr.DB.Save(&NewkkRequest).Error; err != nil {
		logrus.Error("Model: insert kk req error", err.Error())
	}

	return NewkkRequest, nil
}

func (kr *RequestKKRepositoryImpl) GetRequestKK(page int, limit int, sortby string, order string) ([]domain.RequestKK, error) {
	var RequestKK = []domain.RequestKK{}
	offset := (page - 1) * limit
	if err := kr.DB.Preload("Ktp").Preload("AktaKelahiran").Preload("AktaKematian").Preload("SuratNikah").Preload("SuratPindah").Limit(limit).Offset(offset).Order(sortby + " " + order).Find(&RequestKK).Error; err != nil {
		logrus.Error("Model: find data error", err.Error())
		return RequestKK, err
	}
	return RequestKK, nil
}

func (kr *RequestKKRepositoryImpl) RequestKKUpdate(UpdatedRequestKK *domain.RequestKK, RequestKKId float64) (*domain.RequestKK, error) {
	if err := kr.DB.Model(&domain.RequestKK{}).Where("id=?", RequestKKId).Updates(domain.RequestKK{StatusRequest: UpdatedRequestKK.StatusRequest, StatusRequestMessage: UpdatedRequestKK.StatusRequestMessage}).Error; err != nil {
		logrus.Error("Model: update data error", err.Error())
		return nil, err
	}
	return UpdatedRequestKK, nil
}
