package models

import (
	"gorm.io/gorm"
	"time"
)

type Token struct {
	gorm.Model
	AccessToken           string    `json:"access_token"`
	RefreshToken          string    `json:"refresh_token"`
	AccessTokenExpiresAt  time.Time `json:"access_token_expires_at"`
	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
	UserId                uint      `json:"user_id"`
}

func (Token) TableName() string {
	return "tokens"
}
