package domain

import "time"

type LoginResponse struct {
	User  map[string]interface{} `json:"user"`
	Token string                 `json:"access_token"`
}

type Tokens struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresAt    time.Time `json:"expires_at"`
}