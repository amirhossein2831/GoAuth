package auth

import (
	driver2 "GoAuth/src/pkg/auth/driver"
	"errors"
	"os"
)

var NotValidAuthDriver = errors.New("not a valid auth driver")
var instance IAuth

type IAuth interface {
	GenerateToken() (interface{}, error)
	ValidateToken(token string) (interface{}, error)
}

// Init use to initialize the Auth Package
// TODO: should go in db to table client setting
func Init() error {
	driver := os.Getenv("AUTH_DRIVER")
	if driver == "" {
		driver = "jwt"
	}

	switch driver {
	case "jwt":
		instance = &driver2.JWT{}
	default:
		return NotValidAuthDriver
	}

	return nil
}

func GetInstance() IAuth {
	if instance == nil {
		Init()
	}
	return instance
}
