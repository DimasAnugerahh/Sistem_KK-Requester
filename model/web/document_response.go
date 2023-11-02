package web

import "time"

type DocumentResponse struct {
	Id        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	File      string
	Nama      string
	AccountId uint
}
