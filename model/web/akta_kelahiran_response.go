package web

import "time"

type AktaKelahiranResponse struct {
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           time.Time
	File_Akta_kelahiran string
	Nama_lengkap        string
	AccountId           uint
}
