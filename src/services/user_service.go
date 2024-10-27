package services

import (
	"GoAuth/src/database/repository"
	"GoAuth/src/models"
	"context"
)

type IUserService interface {
	Get(c context.Context) ([]models.Model, error)
}

type UserService struct {
	Repository repository.IRepository[models.User]
}

func NewUserService() *UserService {
	return &UserService{
		Repository: repository.GetRepository[models.User](),
	}
}

func (service *UserService) Get(c context.Context) ([]models.Model, error) {
	all, err := service.Repository.List()
	if err != nil {
		return nil, err
	}

	return models.ToModel(all), nil
}
