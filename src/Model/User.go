package Model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string `json:"first_name" gorm:"type:varchar(20);not null"`
	LastName  string `json:"last_name" gorm:"type:varchar(20);not null"`
	Email     string `json:"email" gorm:"type:varchar(20);not null"`
	Password  string `json:"password" gorm:"type:varchar(255);not null"`
}
