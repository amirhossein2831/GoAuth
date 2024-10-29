package user

type ChangePasswordRequest struct {
	NewPassword        string `json:"new_password" validate:"required,min=4,max=16"`
	ConfirmNewPassword string `json:"confirm_new_password" validate:"required,eqfield=NewPassword"`
}
