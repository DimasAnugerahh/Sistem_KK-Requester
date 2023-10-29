package web

import "gorm.io/gorm"

type AccountResponse struct {
	*gorm.Model

	Email    string `json:"email" form:"email"`
	Nama     string `json:"nama" form:"nama"`
	Password string `json:"password" form:"password"`
	Role     string `json:"role" form:"role"`
}
