package domain

import "gorm.io/gorm"

type SuratPindah struct {
	*gorm.Model

	File_surat_pindah string          `json:"file_surat_pindah" form:"file_surat_pindah"`
	Nama_lengkap      string          `json:"nama_lengkap" form:"nama_lengkap"`
	AccountId         uint            `json:"account_id" form:"account_id"`
	RequestDetail     []RequestDetail `gorm:"foreignKey:SuratPindahId"`
}
