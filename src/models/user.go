package models

import (
	"GoAuth/src/hash"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string `json:"first_name" gorm:"type:varchar(64);not null"`
	LastName  string `json:"last_name" gorm:"type:varchar(64);not null"`
	Email     string `json:"email" gorm:"type:varchar(64);not null"`
	Password  string `json:"-" gorm:"type:varchar(255);not null"`
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
