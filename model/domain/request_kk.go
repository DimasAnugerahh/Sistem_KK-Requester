package domain

import "gorm.io/gorm"

type RequestKK struct {
	*gorm.Model

	AccountId            uint             `json:"account_id" form:"account_id"`
	StatusRequest        string           `json:"status_request" form:"status_request" gorm:"default:Pending approval"`
	StatusRequestMessage string           `json:"status_request_message" form:"status_request_message" gorm:"default:File is currently being reviewed by the admin"`
	Ktp                  []*KTP           `json:"ktp,omitempty" gorm:"many2many:req_detail_ktps;"`
	AktaKelahiran        []*AktaKelahiran `json:"akta_kelahiran,omitempty" gorm:"many2many:req_detail_akta_kelahirans;"`
	AktaKematian         []*AktaKematian  `json:"akta_kematian,omitempty" gorm:"many2many:req_detail_akta_kematians;"`
	SuratNikah           []*SuratNikah    `json:"surat_nikah,omitempty" gorm:"many2many:req_detail_surat_nikahs;"`
	SuratPindah          []*SuratPindah   `json:"surat_pindah,omitempty" gorm:"many2many:req_detail_surat_pindahs;"`
}
