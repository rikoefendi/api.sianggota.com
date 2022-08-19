package users

type UserCreateInput struct {
	Name     string `json:"name" validate:"required|minLength:5|maxLength:25"`
	Email    string `json:"email" validate:"required|email|unique:users,email"`
	Phone    string `json:"phone" validate:"maxLength:16|unique:users,phone"`
	Password string `json:"password" validate:"required"`
}

type UserUpdateInput struct {
	Name     string `json:"name"`
	Username string `json:"username"`
	Phone    string `json:"phone"`
}

type UserUpdateEmail struct {
	Email string          `json:"email" validate:"required|email"`
	OTP   int             `json:"otp" validate:"required"`
	Type  EmailUpdateType `json:"type" validate:"in:VERIFY,CHANGE"`
	Ref   string          `json:"ref"`
}

type EmailUpdateType string

const (
	VERIFY EmailUpdateType = "VERIFY"
	CHANGE EmailUpdateType = "CHANGE"
)

type UserUpdatePassword struct {
	OldPassword string             `json:"old_password"`
	Password    string             `json:"password" valdiate:"required"`
	Email       string             `json:"email" validate:"required|email"`
	OTP         int                `json:"otp" validate:"required"`
	Type        PasswordUpdateType `json:"type" validate:"in:FORGOT,CHANGE"`
}

type PasswordUpdateType string

const (
	ForgotPassword PasswordUpdateType = "FORGOT"
	ChangePassword PasswordUpdateType = "CHANGE"
)
