package users

type UserCreateInput struct {
	Name                 *string `json:"name"`
	LastName             *string `json:"last_name"`
	Email                *string `json:"email"`
	Passwrod             *string `json:"password"`
	PasswordConfirmation *string `json:"password_confirmation"`
}

func UserInput(u UserCreateInput) (m Model) {
	m.Name = u.Name
	m.Email = u.Email
	m.Password = u.Passwrod
	return m
}
