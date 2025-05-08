package models

type SetUrlRequest struct {
	BaseURL  string  `json:"base_url" binding:"required,url"`
	CustomID *string `json:"custom_id"`
}
