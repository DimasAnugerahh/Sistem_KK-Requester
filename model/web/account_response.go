package web

import (
	"time"
)

type AccountResponse struct {
	Id        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Email     string
	Name      string
	Password  string
	Role      string
}
