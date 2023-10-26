package domain

import "gorm.io/gorm"

type SuratKematian struct {
	*gorm.Model

	File_Surat_Kematian string          `json:"file_Surat_Kematian" form:"file_Surat_Kematian"`
	Nama_lengkap        string          `json:"nama_lengkap" form:"nama_lengkap"`
	AccountId           uint            `json:"account_id" form:"account_id"`
	RequestDetail       []RequestDetail `gorm:"foreignKey:SuratKematianId"`
}
