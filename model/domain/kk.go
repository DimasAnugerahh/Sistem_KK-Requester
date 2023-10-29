package domain

import "gorm.io/gorm"

type KK struct {
	*gorm.Model

	Email     string `json:"email" form:"email"`
	FileKK    string `json:"file_kk" form:"file_kk"`
	AccountId uint   `json:"account_id" form:"account_id"`
}
