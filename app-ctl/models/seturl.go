package models

type SetUrlRequest struct {
	BaseURL  string  `json:"base_url" binding:"required"`
	CustomID *string `json:"custom_id"`
}
