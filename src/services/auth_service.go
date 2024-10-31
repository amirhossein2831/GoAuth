package services

import (
	"GoAuth/src/api/dto"
	"GoAuth/src/api/request/auth"
	"GoAuth/src/hash"
	"GoAuth/src/models"
	authenticator "GoAuth/src/pkg/auth"
	"GoAuth/src/pkg/ctx"
	ctx2 "GoAuth/src/pkg/ctx"
	"GoAuth/src/pkg/utils"
	"errors"
	"time"
)

var (
	AuthenticateFailed  = errors.New("authentication failed")
	PasswordMismatch    = errors.New("password does not match")
	TokenIsNotExits     = errors.New("token is not exits")
	RefreshTokenExpired = errors.New("refresh token expired")
)

type IAuthService interface {
	Login(ctx ctx.CTX) (interface{}, error)
	RefreshToken(ctx ctx.CTX) (interface{}, error)
	Register(ctx ctx.CTX) (models.Model, error)
	Update(ctx ctx.CTX) (models.Model, error)
	Profile(ctx ctx.CTX) (models.Model, error)
	Verify(ctx ctx.CTX) (interface{}, error)
	Logout(ctx ctx.CTX) error
	ChangePassword(ctx ctx.CTX) error
	TokenList(ctx ctx.CTX) ([]models.Model, error)
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
func (service *AuthService) Login(ctx ctx.CTX) (interface{}, error) {
	req := ctx.Get("req").(*auth.LoginRequest)

	// Get user from with email
	ctx = ctx2.New().SetMap("columns", "email", req.Email)
	user, err := service.UserService.Get(ctx)
	if err != nil {
		return nil, err
	}

	// check the user pass is valid
	userModel := user.(*models.User)
	ok, err := hash.VerifyStoredHash([]byte(userModel.Password), req.Password)
	if err != nil {
		return nil, AuthenticateFailed
	}

	if !ok {
		return nil, PasswordMismatch
	}

	// get the number of active token
	ctx = ctx2.New().SetMap("columns-greater-than", "access_token_expires_at", time.Now()).
		SetMap("columns", "user_id", userModel.ID)
	tokens, err := service.TokenService.ListValidToken(ctx)
	if err != nil {
		return nil, err
	}

	// send one of active tokens if user try to get token more that limit number
	if len(tokens) > authenticator.ActiveTokenNumber() {
		token := tokens[utils.RandomInRange(0, 4)].(*models.Token)
		return dto.TokenDto{
			AccessTokenString:     token.AccessToken,
			RefreshTokenString:    token.RefreshToken,
			AccessTokenExpiresAt:  token.AccessTokenExpiresAt,
			RefreshTokenExpiresAt: token.RefreshTokenExpiresAt,
		}, nil
	}

	// Generate New Token
	token, err := authenticator.GetInstance().GenerateToken(userModel.Email)
	if err != nil {
		return nil, err
	}

	ctx = ctx2.New().Set("token", token).Set("userId", userModel.ID)
	go service.createTokenAsync(ctx)
	return token, nil
}

func (service *AuthService) RefreshToken(ctx ctx.CTX) (interface{}, error) {
	req := ctx.Get("req").(*auth.RefreshTokenRequest)

	ctx = ctx2.New().SetMap("columns", "refresh_token", req.RefreshToken)
	token, err := service.TokenService.Get(ctx)
	if err != nil {
		return nil, err
	}

	// Check if the refresh token is still valid
	if token.(*models.Token).RefreshTokenExpiresAt.Before(time.Now()) {
		return nil, RefreshTokenExpired
	}

	// Get user from with email
	ctx = ctx2.New().SetMap("columns", "id", token.(*models.Token).UserId)
	user, err := service.UserService.Get(ctx)
	if err != nil {
		return nil, err
	}

	// get the number of active token
	userModel := user.(*models.User)
	ctx = ctx2.New().SetMap("columns-greater-than", "access_token_expires_at", time.Now()).
		SetMap("columns", "user_id", userModel.ID)
	tokens, err := service.TokenService.ListValidToken(ctx)
	if err != nil {
		return nil, err
	}

	// send one of active tokens if user try to get token more that limit number
	if len(tokens) > authenticator.ActiveTokenNumber() {
		token := tokens[utils.RandomInRange(0, authenticator.ActiveTokenNumber()-1)].(*models.Token)
		return dto.TokenDto{
			AccessTokenString:     token.AccessToken,
			RefreshTokenString:    token.RefreshToken,
			AccessTokenExpiresAt:  token.AccessTokenExpiresAt,
			RefreshTokenExpiresAt: token.RefreshTokenExpiresAt,
		}, nil
	}

	// Generate New Token
	newToken, err := authenticator.GetInstance().GenerateToken(userModel.Email)
	if err != nil {
		return nil, err
	}

	ctx = ctx2.New().Set("token", newToken).Set("userId", userModel.ID)
	// Add token to list of user token
	go service.createTokenAsync(ctx)
	return newToken, nil
}

func (service *AuthService) TokenList(ctx ctx.CTX) ([]models.Model, error) {
	token := ctx.Get("token").(string)

	ctx = ctx2.New().SetMap("columns", "access_token", token)
	res, err := service.TokenService.Get(ctx)
	if err != nil {
		return nil, TokenIsNotExits
	}

	ctx = ctx2.New().SetMap("columns", "user_id", res.(*models.Token).UserId)
	return service.TokenService.List(ctx)
}

func (service *AuthService) Register(ctx ctx.CTX) (models.Model, error) {
	return service.UserService.Create(ctx)
}

func (service *AuthService) Update(ctx ctx.CTX) (models.Model, error) {
	return service.UserService.Update(ctx)
}

func (service *AuthService) Profile(ctx ctx.CTX) (models.Model, error) {
	token := ctx.Get("token").(string)
	ctx = ctx2.New().SetMap("columns", "access_token", token)

	tokenModel, err := service.TokenService.Get(ctx)
	if err != nil {
		return nil, TokenIsNotExits
	}

	_, err = authenticator.GetInstance().ValidateToken(token)
	if err != nil {
		return nil, err
	}

	ctx = ctx2.New().SetMap("columns", "id", tokenModel.(*models.Token).UserId)
	return service.UserService.Get(ctx)
}

func (service *AuthService) Verify(ctx ctx.CTX) (interface{}, error) {
	token := ctx.Get("token").(string)

	ctx = ctx2.New().SetMap("columns", "access_token", token)
	_, err := service.TokenService.Get(ctx)
	if err != nil {
		return nil, TokenIsNotExits
	}

	return authenticator.GetInstance().ValidateToken(token)
}

func (service *AuthService) Logout(ctx ctx.CTX) error {
	return service.TokenService.DeleteByColumn(ctx)
}

func (service *AuthService) ChangePassword(ctx ctx.CTX) error {
	user, err := service.Profile(ctx)
	if err != nil {
		return err
	}

	req := ctx.Get("req").(*auth.ChangePasswordRequest)
	userModel := user.(*models.User)

	ok, err := hash.VerifyStoredHash([]byte(userModel.Password), req.OldPassword)
	if err != nil {
		return AuthenticateFailed
	}

	if !ok {
		return PasswordMismatch
	}

	ctx = ctx2.New().Set("user", userModel).Set("new_password", req.NewPassword)
	_, err = service.UserService.UpdatePassword(ctx)
	return err
}

// createTokenAsync handles token creation in a separate goroutine
func (service *AuthService) createTokenAsync(ctx ctx.CTX) {
	for {
		_, err := service.TokenService.Create(ctx)
		if err == nil {
			break
		}
	}
}
