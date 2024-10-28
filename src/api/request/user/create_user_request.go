package user

type CreateUserRequest struct {
	FirstName string `json:"first_name" validate:"required,max=64"`
	LastName  string `json:"last_name" validate:"required,max=64"`
	Email     string `json:"email" validate:"required,email,max=64"`
	Password  string `json:"password" validate:"required,min=4,max=16"`
}
