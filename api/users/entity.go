package users

type UserCreateInput struct {
	Name                 string `json:"name"`
	Email                string `json:"email" validate:"required|email"`
	Password             string `json:"password" validate:"required"`
	PasswordConfirmation string `json:"password_confirmation" validate:"required|eqField:password"`
}

func UserInput(u UserCreateInput) (m Model) {
	m.Name = &u.Name
	m.Email = &u.Email
	m.Password = &u.Password
	return m
}
