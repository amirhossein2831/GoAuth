package services

import (
	"GoAuth/src/database/repository"
	"GoAuth/src/models"
	"GoAuth/src/pkg/auth/driver"
	"context"
	"errors"
)

var TokenNotFound = errors.New("token not found")
var InvalidTokenType = errors.New("token Type is invalid")

type ITokenService interface {
	List(c context.Context) ([]models.Model, error)
	ListByColumn(c context.Context) ([]models.Model, error)
	Get(c context.Context) (models.Model, error)
	GetByColumn(c context.Context) (models.Model, error)
	Create(c context.Context) (models.Model, error)
	Delete(c context.Context) error
	DeleteByColumn(c context.Context) error
}

type TokenService struct {
	Repository repository.IRepository[models.Token]
}

func NewTokenService() *TokenService {
	return &TokenService{
		Repository: repository.GetRepository[models.Token](),
	}
}

func (service *TokenService) List(c context.Context) ([]models.Model, error) {
	all, err := service.Repository.List()
	if err != nil {
		return nil, err
	}

	return models.ToModel(all), nil
}

func (service *TokenService) ListByColumn(c context.Context) ([]models.Model, error) {
	columns := c.Value("columns").(map[string]any)

	all, err := service.Repository.ListByColumn(columns)
	if err != nil {
		return nil, err
	}

	return models.ToModel(all), nil
}

func (service *TokenService) Get(c context.Context) (models.Model, error) {
	id := c.Value("tokenId").(uint)

	res, err := service.Repository.Get(id)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (service *TokenService) GetByColumn(c context.Context) (models.Model, error) {
	columns := c.Value("columns").(map[string]any)
	return service.Repository.GetByColumn(columns)
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
	default:
		return nil, InvalidTokenType
	}

	res, err := service.Repository.Create(token)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (service *TokenService) Delete(c context.Context) error {
	res, err := service.Get(c)
	if err != nil {
		return TokenNotFound
	}

	return service.Repository.HardDelete(*res.(*models.Token))
}

func (service *TokenService) DeleteByColumn(c context.Context) error {
	token := c.Value("token").(string)
	c = context.WithValue(c, "columns", map[string]any{"access_token": token})

	res, err := service.GetByColumn(c)
	if err != nil {
		return TokenNotFound
	}

	return service.Repository.HardDelete(*res.(*models.Token))
}
