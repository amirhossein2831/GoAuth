package models

import (
	"gorm.io/gorm"
	"time"
)

type UserType string

var (
	SuperAdmin UserType = "super-admin"
	Admin      UserType = "admin"
	SimpleUser UserType = "user"
)

type User struct {
	ID        uint           `gorm:"primarykey"`
	FirstName string         `json:"first_name" gorm:"type:varchar(64);not null"`
	LastName  string         `json:"last_name" gorm:"type:varchar(64);not null"`
	Email     string         `json:"email" gorm:"type:varchar(64);not null;unique"`
	Password  string         `json:"-" gorm:"type:varchar(255);not null"`
	Type      UserType       `json:"type" gorm:"type:varchar(255);not null;default:user"`
	Tokens    []*Token       `json:"-" gorm:"constraint:OnDelete:CASCADE;"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (user User) TableName() string {
	return "users"
}
