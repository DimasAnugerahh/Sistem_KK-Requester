package domain

import "gorm.io/gorm"

type AktaKelahiran struct {
	*gorm.Model

	File_Akta_kelahiran string          `json:"file_Akta_kelahiran" form:"file_Akta_kelahiran"`
	Nama_lengkap        string          `json:"nama_lengkap" form:"nama_lengkap"`
	AccountId           uint            `json:"account_id" form:"account_id"`
	RequestDetail       []RequestDetail `gorm:"foreignKey:AktaKelahiranId"`
}
