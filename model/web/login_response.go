package web

type LoginResponse struct {
	Email string `json:"email" form:"email"`
	Nama  string `json:"nama" form:"nama"`
	Role  string `json:"role" form:"role"`
	Token string `json:"token" form:"token"`
}
