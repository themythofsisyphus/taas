package dto

type TagRequest struct {
	Name string `json:"name" binding:"required"`
}

type TagResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
