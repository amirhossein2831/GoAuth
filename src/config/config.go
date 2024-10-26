package config

import (
	"github.com/joho/godotenv"
	"sync"
)

var once sync.Once

func Init() error {
	var err error
	once.Do(func() {
		err = godotenv.Load()
		if err != nil {
			return
		}
	})
	return err
}
