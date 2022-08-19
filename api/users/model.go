package users

import (
	"time"

	"api.sianggota.com/database"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Model struct {
	database.ID
	Name            *string    `json:"name" gorm:"type:varchar;size:225"`
	Username        *string    `json:"usernmae" gorm:"type:varchar;size:50;unique;index;null"`
	Email           *string    `json:"email" gorm:"type:varchar;size:50;unique;index;not null"`
	EmailVerifiedAt *time.Time `json:"email_verified_at"`
	Phone           *string    `json:"phone" gorm:"size:20;null"`
	Password        string     `json:"-" gorm:"type:varchar;size:225;not null"`
	Status          *float64   `json:"status" gorm:"default:0"`
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
