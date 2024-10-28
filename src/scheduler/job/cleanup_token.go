package job

import (
	"GoAuth/src/database"
	"GoAuth/src/models"
	"fmt"
	"github.com/robfig/cron/v3"
	"log"
	"strconv"
	"time"
)

func CleanUpToken() error {
	delay := "100"
	duration, err := strconv.Atoi(delay)
	if err != nil {
		duration = 86400
	}

	c := cron.New()
	_, err = c.AddFunc(fmt.Sprintf("@every %vs", duration), func() {
		result := database.GetInstance().GetClient().Where("refresh_token_expires_at < ?", time.Now()).Unscoped().Delete(&models.Token{})
		if result.Error != nil {
			// Todo: put the log in a file and use a logger module
			log.Printf("Error cleaning up token expiration: %v", result.Error)
		}
	})
	if err != nil {
		return err
	}

	c.Start()
	return nil
}
