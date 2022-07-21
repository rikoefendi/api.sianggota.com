package users

import (
	"time"

	"api.sianggota.com/database"
)

type Model struct {
	database.ID
	Name            *string    `json:"name" gorm:"type:varchar;size:225"`
	Username        *string    `json:"usernmae" gorm:"type:varchar;size:50;unique;index;null"`
	Email           *string    `json:"email" gorm:"type:varchar;size:50;unique;index;not null"`
	Password        *string    `json:"password" gorm:"type:varchar;size:225;not null"`
	EmailVerifiedAt *time.Time `json:"email_verified_at"`
	Status          *float64   `json:"status" gorm:"default:0"`
	database.TimeStamp
}

func (Model) TableName() string {
	return "users"
}
