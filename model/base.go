package model

type APIResponse struct {
	Success bool        `json:"sucess"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}
