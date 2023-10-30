package repository

import (
	"kk-requester/model/domain"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type KTPRepositoryImpl struct {
	DB *gorm.DB
}

func NewKTPRepository(db *gorm.DB) KTPRepository {
	return &KTPRepositoryImpl{DB: db}
}

func (kr *KTPRepositoryImpl) CreateKTP(NewKTP *domain.KTP, request *domain.ReqDetailKtp) (*domain.KTP, *domain.ReqDetailKtp, error) {
	if err := kr.DB.Create(&NewKTP).Error; err != nil {
		logrus.Error("Model: insert data error", err.Error())
		return nil, nil, err
	}
	request.Ktp_id = int(NewKTP.ID)
	if err := kr.DB.Create(&request).Error; err != nil {
		logrus.Error("Model: insert data error", err.Error())
		return nil, nil, err
	}
	return NewKTP, request, nil
}

func (kr *KTPRepositoryImpl) KTPUpdate(UpdatedKTP *domain.KTP, ktpId float64, accountId uint) (*domain.KTP, error) {
	if err := kr.DB.Model(&domain.KTP{}).Where("id=? and account_id=?", ktpId, accountId).Updates(domain.KTP{Nama_lengkap: UpdatedKTP.Nama_lengkap, File_ktp: UpdatedKTP.File_ktp}).Error; err != nil {
		logrus.Error("Model: update data error", err.Error())
		return nil, err
	}
	return UpdatedKTP, nil
}

func (kr *KTPRepositoryImpl) GetKTP(idAccount uint) ([]domain.KTP, error) {
	var KTP = []domain.KTP{}
	if err := kr.DB.Where("account_id=?", idAccount).Find(&KTP).Error; err != nil {
		logrus.Error("Model: find data error", err.Error())
		return KTP, err
	}

	return KTP, nil
}
