package model

type Response struct {
	Message *string `json:"message"`
	Status  *int    `json:"status"`
}
