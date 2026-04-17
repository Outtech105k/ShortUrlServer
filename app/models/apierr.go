package models

type APIError struct {
	Type    string `json:"type"`
	Message string `json:"message,omitempty"`
	Details any    `json:"details,omitempty"`
}
