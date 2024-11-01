package auth

import (
	driver2 "GoAuth/src/pkg/auth/driver"
	"errors"
	"os"
	"strconv"
)

var instance IAuth

var NotValidAuthDriver = errors.New("not a valid auth driver")

type IAuth interface {
	GenerateToken(email string) (interface{}, error)
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

func ActiveTokenNumber() int {
	numStr := os.Getenv("ACTIVE_TOKEN_NUMBER")
	num, err := strconv.Atoi(numStr)
	if err != nil {
		num = 5
	}
	return num
}
