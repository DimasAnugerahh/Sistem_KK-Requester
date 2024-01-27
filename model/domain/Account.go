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
	AktaKelahiran []AktaKelahiran `gorm:"foreignKey:AccountId" json:"akta_kelahiran,omitempty"`
	KTP           []KTP           `gorm:"foreignKey:AccountId" json:"ktp,omitempty"`
	AktaKematian  []AktaKematian  `gorm:"foreignKey:AccountId" json:"akta_kematian,omitempty"`
	SuratNikah    []SuratNikah    `gorm:"foreignKey:AccountId" json:"surat_nikah,omitempty"`
	SuratPindah   []SuratPindah   `gorm:"foreignKey:AccountId" json:"surat_pindah,omitempty"`
	RequestKK     []RequestKK     `gorm:"foreignKey:AccountId" json:"request,omitempty"`
	KK            []KK            `gorm:"foreignKey:AccountId" json:"kk,omitempty"`
}
