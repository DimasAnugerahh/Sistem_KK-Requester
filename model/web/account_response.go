package web

import (
	"time"
)

type AccountResponse struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Email     string
	Name      string
	Password  string
	Role      string
}
