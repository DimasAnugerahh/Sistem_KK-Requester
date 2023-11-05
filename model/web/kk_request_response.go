package web

type KKRequestResponse struct {
	Id                   uint   `json:"id"`
	AccountId            uint   `json:"account_id"`
	StatusRequest        string `json:"status_request"`
	StatusRequestMessage string `json:"status_request_message"`
}
