package user

type UpdateUserRequest struct {
	FirstName *string `json:"first_name" validate:"omitempty,max=64"`
	LastName  *string `json:"last_name" validate:"omitempty,max=64"`
	Email     *string `json:"email" validate:"omitempty,email,max=64"`
}
