package domain

import "gorm.io/gorm"

type KTP struct {
	*gorm.Model

	File_ktp     string       `json:"file_ktp" form:"file_ktp"`
	Nama_lengkap string       `json:"nama_lengkap" form:"nama_lengkap"`
	AccountId    uint         `json:"account_id" form:"account_id"`
	RequestKK    []*RequestKK `gorm:"many2many:req_detail_ktps;" json:"request,omitempty"`
}

type ReqDetailKtp struct {
	Request_kk_id int `gorm:"primaryKey"`
	Ktp_id        int `gorm:"primaryKey"`
}
