package domain

import "gorm.io/gorm"

type KTP struct {
	*gorm.Model
	// Id           uint         `json:"id_ktp" form:"id_ktp"`
	File_ktp     string       `json:"file_ktp" form:"file_ktp"`
	Nama_lengkap string       `json:"nama_lengkap" form:"nama_lengkap"`
	AccountId    uint         `json:"account_id" form:"account_id"`
	RequestKK    []*RequestKK `gorm:"many2many:req_detail_ktps;"`
}

type ReqDetailKtp struct {
	Request_kk_id int `gorm:"primaryKey"`
	Ktp_id        int `gorm:"primaryKey"`
}
