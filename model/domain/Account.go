package domain

import (
	"gorm.io/gorm"
)

type Account struct {
	*gorm.Model

	Email         string          `json:"email" form:"email" validate:"required,max=200,min=1"`
	Name          string          `json:"name" form:"name" validate:"required,max=50,min=1"`
	Password      string          `json:"password" form:"password" validate:"required,max=50,min=1"`
	Role          string          `json:"role" form:"role" validate:"required,max=5,min=4"`
	AktaKelahiran []AktaKelahiran `gorm:"foreignKey:AccountId"`
	KTP           []KTP           `gorm:"foreignKey:AccountId"`
	AktaKematian  []AktaKematian  `gorm:"foreignKey:AccountId"`
	SuratNikah    []SuratNikah    `gorm:"foreignKey:AccountId"`
	SuratPindah   []SuratPindah   `gorm:"foreignKey:AccountId"`
	RequestKK     []RequestKK     `gorm:"foreignKey:AccountId"`
	KK            []KK            `gorm:"foreignKey:AccountId"`
}
