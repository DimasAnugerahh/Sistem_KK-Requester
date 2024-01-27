package domain

import "gorm.io/gorm"

type AktaKematian struct {
	*gorm.Model

	File_akta_Kematian string       `json:"file_akta_Kematian" form:"file_akta_Kematian"`
	Nama_lengkap       string       `json:"nama_lengkap" form:"nama_lengkap"`
	AccountId          uint         `json:"account_id" form:"account_id"`
	RequestKK          []*RequestKK `gorm:"many2many:req_detail_akta_kematians;" json:"omitempty"`
}
type ReqDetailAktaKematian struct {
	Request_kk_id    int `gorm:"primaryKey"`
	Akta_kematian_id int `gorm:"primaryKey"`
}
