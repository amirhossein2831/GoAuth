package models

import (
	"GoAuth/src/hash"
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        uint           `gorm:"primarykey"`
	FirstName string         `json:"first_name" gorm:"type:varchar(64);not null"`
	LastName  string         `json:"last_name" gorm:"type:varchar(64);not null"`
	Email     string         `json:"email" gorm:"type:varchar(64);not null;unique"`
	Password  string         `json:"-" gorm:"type:varchar(255);not null"`
	Tokens    []*Token       `json:"-" gorm:"constraint:OnDelete:CASCADE;"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (user User) TableName() string {
	return "users"
}

// BeforeSave here we hash the User.Password before save
func (user *User) BeforeSave(tx *gorm.DB) error {
	if len(user.Password) > 0 {
		hashedPassword, err := hash.GetInstance().Generate([]byte(user.Password))
		if err != nil {
			return err
		}
		user.Password = string(hashedPassword)
	}
	return nil

}
