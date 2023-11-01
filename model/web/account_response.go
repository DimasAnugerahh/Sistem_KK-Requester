package web

import (
	"time"

	"gorm.io/gorm"
)

type AccountResponse struct {
	*gorm.Model

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Email     string
	Name      string
	Password  string
	Role      string
}

type CreateAccountResponse struct {
	*gorm.Model

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Email     string
	Name      string
	Role      string
}
