package repository

import (
	"kk-requester/model/domain"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type KkRepositoryImpl struct {
	DB *gorm.DB
}

func NewKkRepository(db *gorm.DB) KkRepository {
	return &KkRepositoryImpl{DB: db}
}

func (kr *KkRepositoryImpl) CreateKK(NewKK *domain.KK) (*domain.KK, error) {
	if err := kr.DB.Create(&NewKK).Error; err != nil {
		logrus.Error("Model: insert data error", err.Error())
		return nil, err
	}
	return NewKK, nil
}

func (kr *KkRepositoryImpl) GetKK() ([]domain.KK, error) {
	var KK = []domain.KK{}
	if err := kr.DB.Find(&KK).Error; err != nil {
		logrus.Error("Model: find data error", err.Error())
		return KK, err
	}
	return KK, nil
}
