package services

import (
	"GoAuth/src/api/request/user"
	"GoAuth/src/database/repository"
	"GoAuth/src/hash"
	"GoAuth/src/models"
	"GoAuth/src/pkg/ctx"
	"errors"
)

var UserNotFound = errors.New("user not found")
var EmailShouldBeUnique = errors.New("email should be unique")

type IUserService interface {
	List(c ctx.CTX) ([]models.Model, error)
	Get(c ctx.CTX) (models.Model, error)
	Create(c ctx.CTX) (models.Model, error)
	Update(c ctx.CTX) (models.Model, error)
	UpdatePassword(c ctx.CTX) (models.Model, error)
	ChangePassword(c ctx.CTX) (models.Model, error)
	Delete(c ctx.CTX) error
}

type UserService struct {
	Repository repository.IRepository[models.User]
}

func NewUserService() *UserService {
	return &UserService{
		Repository: repository.GetRepository[models.User](),
	}
}

func (service *UserService) List(c ctx.CTX) ([]models.Model, error) {
	columns := c.GetMap("columns")
	all, err := service.Repository.ListByColumn(columns)
	if err != nil {
		return nil, err
	}

	return models.ToModel(all), nil
}

func (service *UserService) Get(c ctx.CTX) (models.Model, error) {
	columns := c.GetMap("columns")
	return service.Repository.GetByColumn(columns)
}

func (service *UserService) Create(c ctx.CTX) (models.Model, error) {
	req := c.Get("req").(*user.CreateUserRequest)

	_, err := service.Repository.GetByColumn(map[string]any{
		"email": req.Email,
	})
	if err == nil {
		return nil, EmailShouldBeUnique
	}

	userPass, err := hash.GetInstance().Generate([]byte(req.Password))
	if err != nil {
		return nil, err
	}

	entity := models.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  string(userPass),
	}

	res, err := service.Repository.Create(entity)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (service *UserService) Update(c ctx.CTX) (models.Model, error) {
	req := c.Get("req").(*user.UpdateUserRequest)

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

func (service *UserService) UpdatePassword(ctx ctx.CTX) (models.Model, error) {
	userModel := ctx.Get("user").(*models.User)
	newPassword := ctx.Get("new_password").(string)

	userPass, err := hash.GetInstance().Generate([]byte(newPassword))
	if err != nil {
		return nil, err
	}

	userModel.Password = string(userPass)

	return service.Repository.Update(*userModel)
}

func (service *UserService) ChangePassword(c ctx.CTX) (models.Model, error) {
	req := c.Get("req").(*user.ChangePasswordRequest)
	user, err := service.Get(c)
	if err != nil {
		return nil, UserNotFound
	}

	userPass, err := hash.GetInstance().Generate([]byte(req.NewPassword))
	if err != nil {
		return nil, err
	}

	userModel := user.(*models.User)
	userModel.Password = string(userPass)

	return service.Repository.Update(*userModel)
}

func (service *UserService) Delete(c ctx.CTX) error {
	res, err := service.Get(c)
	if err != nil {
		return UserNotFound
	}

	return service.Repository.Delete(*res.(*models.User))
}
