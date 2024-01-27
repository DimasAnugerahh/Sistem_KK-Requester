package domain

import "gorm.io/gorm"

type SuratNikah struct {
	*gorm.Model

	File_surat_nikah string       `json:"file_surat_nikah" form:"file_Surat_Nikah"`
	Nama_lengkap     string       `json:"nama_lengkap" form:"nama_lengkap"`
	AccountId        uint         `json:"account_id" form:"account_id"`
	RequestKK        []*RequestKK `gorm:"many2many:req_detail_surat_nikahs;"  json:"request,omitempty"`
}

type ReqDetailSuratNikah struct {
	Request_kk_id  int `gorm:"primaryKey"`
	Surat_nikah_id int `gorm:"primaryKey"`
}
