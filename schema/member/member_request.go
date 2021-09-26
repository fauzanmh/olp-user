package member

type RegisterRequest struct {
	Name                 string `json:"name" validate:"required"`
	Email                string `json:"email" validate:"required"`
	Address              string `json:"address" validate:"required"`
	Username             string `json:"username" validate:"required"`
	Password             string `json:"password" validate:"required,gte=8,eqfield=PasswordConfirmation"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required"`
}
