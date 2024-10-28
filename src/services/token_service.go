package services

import (
	"GoAuth/src/database/repository"
	"GoAuth/src/models"
	"GoAuth/src/pkg/auth/driver"
	"context"
)

type ITokenService interface {
	Create(c context.Context) (models.Model, error)
}

type TokenService struct {
	Repository repository.IRepository[models.Token]
}

func NewTokenService() *TokenService {
	return &TokenService{
		Repository: repository.GetRepository[models.Token](),
	}
}

func (service *TokenService) Create(c context.Context) (models.Model, error) {
	req := c.Value("token")
	userId := c.Value("userId").(uint)

	token := models.Token{UserId: userId}

	switch req.(type) {
	case *driver.JWT:
		jwtToken := req.(*driver.JWT)
		token.AccessToken = jwtToken.AccessTokenString
		token.RefreshToken = jwtToken.RefreshTokenString
		token.AccessTokenExpiresAt = jwtToken.AccessTokenExpiresAt
		token.RefreshTokenExpiresAt = jwtToken.RefreshTokenExpiresAt
	}

	res, err := service.Repository.Create(token)
	if err != nil {
		return nil, err
	}

	return res, nil
}
