package model

type APIResponse struct {
	Success bool        `json:"sucess"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
}

type PaginationMetaResponse struct {
	Total      uint `json:"total"`
	TotalPages int  `json:"total_pages"`
}
