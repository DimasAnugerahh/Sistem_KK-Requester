package web

import "time"

type DocumentResponse struct {
	Id        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
	File      string    `json:"file"`
	Nama      string    `json:"nama"`
	AccountId uint      `json:"account_id"`
}
