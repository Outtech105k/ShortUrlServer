package models

type SetUrlRequest struct {
	BaseURL      string    `json:"base_url" binding:"required,url"`
	CustomID     *string   `json:"custom_id"`
	UseUppercase *bool     `json:"use_uppercase"`
	UseLowercase *bool     `json:"use_lowercase"`
	UseNumbers   *bool     `json:"use_numbers"`
	IDLength     *uint32   `json:"id_length"`
	ExpireIn     *Duration `json:"expire_in"`
}
