package users

import (
	"api.sianggota.com/database"
	"github.com/neko-neko/echo-logrus/v2/log"
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
	log.Info(m)
	return m, nil
}
