package domain

import "gorm.io/gorm"

type KTP struct {
	*gorm.Model

	File_ktp      string          `json:"file_ktp" form:"file_ktp"`
	Nama_lengkap  string          `json:"nama_lengkap" form:"nama_lengkap"`
	AccountId     uint            `json:"account_id" form:"account_id"`
	RequestDetail []RequestDetail `gorm:"foreignKey:KTPId"`
}
