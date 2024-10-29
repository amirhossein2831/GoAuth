package dto

import (
	"time"
)

type TokenDto struct {
	AccessTokenString     string    `json:"access_token_string"`
	RefreshTokenString    string    `json:"refresh_token_string"`
	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
	AccessTokenExpiresAt  time.Time `json:"access_token_expires_at"`
}
