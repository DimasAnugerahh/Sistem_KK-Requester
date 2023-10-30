package domain

import "gorm.io/gorm"

type RequestKK struct {
	*gorm.Model

	AccountId            uint             `json:"account_id" form:"account_id"`
	StatusRequest        string           `json:"status_request" form:"status_request" gorm:"default:Pending approval"`
	StatusRequestMessage string           `json:"status_request_message" form:"status_request_message" gorm:"default:File is currently being reviewed by the admin"`
	Ktp                  []*KTP           `gorm:"many2many:req_detail_ktps;"`
	AktaKelahiran        []*AktaKelahiran `gorm:"many2many:req_detail_akta_kelahirans;"`
	AktaKematian         []*AktaKematian  `gorm:"many2many:req_detail_akta_kematians;"`
	SuratNikah           []*SuratNikah    `gorm:"many2many:req_detail_surat_nikahs;"`
	SuratPindah          []*SuratPindah   `gorm:"many2many:req_detail_surat_pindahs;"`
}
