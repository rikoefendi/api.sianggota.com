package users

import (
	"api.sianggota.com/database"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func New() *Repository {
	db := database.Session()
	r := &Repository{db}
	return r
}

func (r *Repository) Create(u UserCreateInput) (m Model, err error) {
	m = UserInput(u)
	result := r.db.Create(&m)
	if result.Error != nil {
		return m, result.Error
	}
	return m, nil
}
