package models

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Error  *Error      `json:"error"`
}
