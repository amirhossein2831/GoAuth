package seed

import (
	"GoAuth/src/database"
	"GoAuth/src/models"
)

func AdminSeeder() error {
	user := models.User{
		FirstName: "admin",
		LastName:  "admin",
		Email:     "admin@admin.com",
		Password:  "admin",
		Type:      models.SuperAdmin,
	}

	res := database.GetInstance().GetClient().Create(&user)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
