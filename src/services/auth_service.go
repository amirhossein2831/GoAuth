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
	UserService  IUserService
	TokenService ITokenService
}

func NewAuthService() *AuthService {
	return &AuthService{
		UserService:  NewUserService(),
		TokenService: NewTokenService(),
	}
}

func (service *AuthService) Login(ctx context.Context) (interface{}, error) {
	req := ctx.Value("req").(*auth.LoginRequest)

	ctx = context.WithValue(ctx, "columns", map[string]any{
		"email": req.Email,
	})
	user, err := service.UserService.GetByColumn(ctx)
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

	// Generate Token
	token, err := authenticator.GetInstance().GenerateToken(userModel.Email)
	if err != nil {
		return nil, err
	}

	ctx = context.WithValue(ctx, "token", token)
	ctx = context.WithValue(ctx, "userId", userModel.ID)

	go service.createTokenAsync(ctx)
	return token, nil
}

// createTokenAsync handles token creation in a separate goroutine
func (service *AuthService) createTokenAsync(ctx context.Context) {
	for {
		_, err := service.TokenService.Create(ctx)
		if err == nil {
			break
		}
	}
}
