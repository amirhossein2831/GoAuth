package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Token struct {
	ID                    uint           `gorm:"primarykey"`
	Uuid                  uuid.UUID      `gorm:"type:uuid;"`
	AccessToken           string         `json:"access_token"`
	RefreshToken          string         `json:"refresh_token"`
	AccessTokenExpiresAt  time.Time      `json:"access_token_expires_at"`
	RefreshTokenExpiresAt time.Time      `json:"refresh_token_expires_at"`
	UserId                uint           `json:"user_id"`
	CreatedAt             time.Time      `json:"created_at"`
	UpdatedAt             time.Time      `json:"updated_at"`
	DeletedAt             gorm.DeletedAt `json:"-" gorm:"index"`
}

func (Token) TableName() string {
	return "tokens"
}
