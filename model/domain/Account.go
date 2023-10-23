package domain

import "gorm.io/gorm"

type Account struct {
	*gorm.Model

	Email    string `json:"email" form:"email" validate:"required,max=200,min=1"`
	Name     string `json:"name" form:"name" validate:"required,max=50,min=1"`
	Password string `json:"password" form:"password" validate:"required,max=50,min=1"`
	Role     string `json:"role" form:"role" validate:"required,max=5,min=4"`
}
