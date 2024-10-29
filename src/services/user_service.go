package services

import (
	"GoAuth/src/api/request/user"
	"GoAuth/src/database/repository"
	"GoAuth/src/models"
	"context"
	"errors"
)

var UserNotFound = errors.New("user not found")
var EmailShouldBeUnique = errors.New("email should be unique")

type IUserService interface {
	List(c context.Context) ([]models.Model, error)
	Get(c context.Context) (models.Model, error)
	GetByColumn(c context.Context) (models.Model, error)
	Create(c context.Context) (models.Model, error)
	Update(c context.Context) (models.Model, error)
	UpdatePassword(c context.Context) (models.Model, error)
	ChangePassword(c context.Context) (models.Model, error)
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
	userId := c.Value("userId").(uint)

	res, err := service.Repository.Get(userId)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (service *UserService) GetByColumn(c context.Context) (models.Model, error) {
	columns := c.Value("columns").(map[string]any)
	return service.Repository.GetByColumn(columns)
}

func (service *UserService) Create(c context.Context) (models.Model, error) {
	req := c.Value("req").(*user.CreateUserRequest)

	_, err := service.Repository.GetByColumn(map[string]any{
		"email": req.Email,
	})
	if err == nil {
		return nil, EmailShouldBeUnique
	}

	entity := models.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
	}

	res, err := service.Repository.Create(entity)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (service *UserService) Update(c context.Context) (models.Model, error) {
	req := c.Value("req").(*user.UpdateUserRequest)

	userToUpdate, err := service.Get(c)
	if err != nil {
		return nil, err
	}

	userModel := userToUpdate.(*models.User)
	if req.FirstName != nil && *req.FirstName != "" {
		userModel.FirstName = *req.FirstName
	}

	if req.LastName != nil && *req.LastName != "" {
		userModel.LastName = *req.LastName
	}

	if req.Email != nil && *req.Email != "" {
		userModel.LastName = *req.Email
	}

	res, err := service.Repository.Update(*userModel)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (service *UserService) UpdatePassword(ctx context.Context) (models.Model, error) {
	userModel := ctx.Value("user").(*models.User)
	newPassword := ctx.Value("new_password").(string)

	userModel.Password = newPassword

	return service.Repository.Update(*userModel)
}

func (service *UserService) ChangePassword(c context.Context) (models.Model, error) {
	req := c.Value("req").(*user.ChangePasswordRequest)
	user, err := service.Get(c)
	if err != nil {
		return nil, UserNotFound
	}

	userModel := user.(*models.User)
	userModel.Password = req.NewPassword

	return service.Repository.Update(*userModel)
}

func (service *UserService) Delete(c context.Context) error {
	res, err := service.Get(c)
	if err != nil {
		return UserNotFound
	}

	return service.Repository.Delete(*res.(*models.User))
}
