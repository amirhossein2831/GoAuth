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
	TokenIsNotExits    = errors.New("token is not exits")
)

type IAuthService interface {
	Login(ctx context.Context) (interface{}, error)
	Register(ctx context.Context) (models.Model, error)
	Profile(ctx context.Context) (models.Model, error)
	Verify(ctx context.Context) (interface{}, error)
	Logout(ctx context.Context) error
	ChangePassword(ctx context.Context) error
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

// Login send token and let the user login
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

func (service *AuthService) Register(ctx context.Context) (models.Model, error) {
	return service.UserService.Create(ctx)
}

func (service *AuthService) Logout(ctx context.Context) error {
	return service.TokenService.DeleteByColumn(ctx)
}

func (service *AuthService) Verify(ctx context.Context) (interface{}, error) {
	token := ctx.Value("token").(string)
	ctx = context.WithValue(ctx, "columns", map[string]any{"access_token": token})

	_, err := service.TokenService.GetByColumn(ctx)
	if err != nil {
		return nil, TokenIsNotExits
	}

	return authenticator.GetInstance().ValidateToken(token)
}

func (service *AuthService) Profile(ctx context.Context) (models.Model, error) {
	token := ctx.Value("token").(string)
	ctx = context.WithValue(ctx, "columns", map[string]any{"access_token": token})

	tokenModel, err := service.TokenService.GetByColumn(ctx)
	if err != nil {
		return nil, TokenIsNotExits
	}

	_, err = authenticator.GetInstance().ValidateToken(token)
	if err != nil {
		return nil, err
	}

	ctx = context.WithValue(ctx, "userId", tokenModel.(*models.Token).UserId)
	return service.UserService.Get(ctx)
}

func (service *AuthService) ChangePassword(ctx context.Context) error {
	user, err := service.Profile(ctx)
	if err != nil {
		return err
	}

	req := ctx.Value("req").(*auth.ChangePasswordRequest)
	userModel := user.(*models.User)

	ok, err := hash.VerifyStoredHash([]byte(userModel.Password), req.OldPassword)
	if err != nil {
		return AuthenticateFailed
	}

	if !ok {
		return PasswordMismatch
	}

	ctx = context.WithValue(ctx, "user", userModel)
	ctx = context.WithValue(ctx, "new_password", req.NewPassword)
	_, err = service.UserService.UpdatePassword(ctx)
	return err
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
