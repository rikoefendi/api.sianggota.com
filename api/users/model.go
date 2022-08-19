package users

import (
	"time"

	"api.sianggota.com/database"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Model struct {
	database.ID
	Name            *string    `json:"name"`
	Username        *string    `json:"usernmae"`
	Email           *string    `json:"email"`
	EmailVerifiedAt *time.Time `json:"email_verified_at"`
	Phone           *string    `json:"phone"`
	PhoneVerifiedAt *time.Time `json:"phone_verified_at"`
	Password        string     `json:"-"`
	Status          int        `json:"status"`
	database.TimeStamp
}

func (Model) TableName() string {
	return "users"
}

func (m *Model) BeforeCreate(tx *gorm.DB) (err error) {
	if m.Password == "" {
		return
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(m.Password), bcrypt.MinCost)
	m.Password = string(hashPassword)
	// m.Password
	return
}

func (m *Model) VerifyPassword(pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(m.Password), []byte(pass))
	return err == nil
}
