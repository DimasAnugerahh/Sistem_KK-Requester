package domain

import "gorm.io/gorm"

type AktaKelahiran struct {
	*gorm.Model

	File_Akta_kelahiran string       `json:"file_Akta_kelahiran" form:"file_Akta_kelahiran"`
	Nama_lengkap        string       `json:"nama_lengkap" form:"nama_lengkap"`
	AccountId           uint         `json:"account_id" form:"account_id"`
	RequestKK           []*RequestKK `gorm:"many2many:req_detail_akta_kelahirans;"`
}
type ReqDetailAktaKelahiran struct {
	Request_kk_id     int `gorm:"primaryKey"`
	Akta_kelahiran_id int `gorm:"primaryKey"`
}
