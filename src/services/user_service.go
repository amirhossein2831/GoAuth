package services

import (
	"GoAuth/src/database/repository"
	"GoAuth/src/models"
	"context"
	"errors"
	"strconv"
)

var UserIdIsNotValid = errors.New("userId is not valid")
var IdShouldBeNumeric = errors.New("user Id should be numeric")
var UserNotFound = errors.New("user not found")

type IUserService interface {
	List(c context.Context) ([]models.Model, error)
	Get(c context.Context) (models.Model, error)
	Delete(c context.Context) error
}

type UserService struct {
	Repository repository.IRepository[models.User]
}

func NewUserService() *UserService {
	return &UserService{
		Repository: repository.GetRepository[models.User](),
	}
}

func (service *UserService) List(c context.Context) ([]models.Model, error) {
	all, err := service.Repository.List()
	if err != nil {
		return nil, err
	}

	return models.ToModel(all), nil
}

func (service *UserService) Get(c context.Context) (models.Model, error) {
	userId := c.Value("userId")
	if userId == nil {
		return nil, UserIdIsNotValid
	}

	id, err := strconv.Atoi(userId.(string))
	if err != nil {
		return nil, IdShouldBeNumeric
	}

	res, err := service.Repository.Get(uint(id))
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (service *UserService) Delete(c context.Context) error {
	userId := c.Value("userId")
	if userId == nil {
		return UserIdIsNotValid
	}
	id, err := strconv.Atoi(userId.(string))
	if err != nil {
		return IdShouldBeNumeric
	}

	user, err := service.Repository.Get(uint(id))
	if err != nil {
		return UserNotFound
	}

	return service.Repository.Delete(*user)
}