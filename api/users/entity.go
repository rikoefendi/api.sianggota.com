package users

type UserCreateInput struct {
	Name                 string `json:"name"`
	Email                string `json:"email" validate:"required|email|unique:users,email"`
	Password             string `json:"password" validate:"required"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required|eqField:password"`
}

type UserUpdateInput struct {
	Name string `json:"name"`
}
