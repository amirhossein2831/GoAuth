package services

import (
	"GoAuth/src/database/repository"
	"GoAuth/src/models"
	"GoAuth/src/pkg/auth/driver"
	"GoAuth/src/pkg/ctx"
	"errors"
)

var TokenNotFound = errors.New("token not found")
var InvalidTokenType = errors.New("token Type is invalid")

type ITokenService interface {
	List(c ctx.CTX) ([]models.Model, error)
	ListValidToken(c ctx.CTX) ([]models.Model, error)
	Get(c ctx.CTX) (models.Model, error)
	Create(c ctx.CTX) (models.Model, error)
	Delete(c ctx.CTX) error
	DeleteByColumn(c ctx.CTX) error
}

type TokenService struct {
	Repository repository.IRepository[models.Token]
}

func NewTokenService() *TokenService {
	return &TokenService{
		Repository: repository.GetRepository[models.Token](),
	}
}

func (service *TokenService) List(c ctx.CTX) ([]models.Model, error) {
	columns := c.GetMap("columns")

	all, err := service.Repository.ListByColumn(columns)
	if err != nil {
		return nil, err
	}

	return models.ToModel(all), nil
}

func (service *TokenService) Get(c ctx.CTX) (models.Model, error) {
	columns := c.GetMap("columns")
	return service.Repository.GetByColumn(columns)
}

func (service *TokenService) Create(c ctx.CTX) (models.Model, error) {
	req := c.Get("token")
	userId := c.Get("userId").(uint)

	token := models.Token{UserId: userId}

	switch req.(type) {
	case *driver.JWT:
		jwtToken := req.(*driver.JWT)
		token.Uuid = jwtToken.Uuid
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

func (service *TokenService) Delete(c ctx.CTX) error {
	res, err := service.Get(c)
	if err != nil {
		return TokenNotFound
	}

	return service.Repository.HardDelete(*res.(*models.Token))
}

func (service *TokenService) DeleteByColumn(c ctx.CTX) error {
	token := c.Get("token").(string)

	c = ctx.New().SetMap("columns", "access_token", token)
	res, err := service.Get(c)
	if err != nil {
		return TokenNotFound
	}

	return service.Repository.Delete(*res.(*models.Token))
}

func (service *TokenService) ListValidToken(c ctx.CTX) ([]models.Model, error) {
	columns := c.GetMap("columns")
	greaterCol := c.GetMap("columns-greater-than")

	all, err := service.Repository.ListByColumnWithGreaterThan(columns, greaterCol)
	if err != nil {
		return nil, err
	}

	return models.ToModel(all), nil
}
