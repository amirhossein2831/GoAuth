package services

import (
	"GoAuth/src/api/request/auth"
	"GoAuth/src/hash"
	"GoAuth/src/models"
	authenticator "GoAuth/src/pkg/auth"
	"context"
	"errors"
)

var (
	AuthenticateFailed = errors.New("authentication failed")
	PasswordMismatch   = errors.New("password does not match")
)

type IAuthService interface {
	Login(ctx context.Context) (interface{}, error)
}

type AuthService struct {
	UserService IUserService
}

func NewAuthService() *AuthService {
	return &AuthService{
		UserService: NewUserService(),
	}
}

func (service *AuthService) Login(ctx context.Context) (interface{}, error) {
	req := ctx.Value("req").(*auth.LoginRequest)

	c := context.WithValue(ctx, "columns", map[string]any{
		"email": req.Email,
	})
	user, err := service.UserService.GetByColumn(c)
	if err != nil {
		return nil, err
	}

	userModel := user.(*models.User)
	ok, err := hash.VerifyStoredHash([]byte(userModel.Password), req.Password)
	if err != nil {
		return nil, AuthenticateFailed
	}

	if !ok {
		return nil, PasswordMismatch
	}

	token, err := authenticator.GetInstance().GenerateToken(userModel.Email)
	if err != nil {
		return nil, err
	}

	return token, nil
}
