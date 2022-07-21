package migration

import (
	"api.sianggota.com/api/users"
	"api.sianggota.com/database"
)

func Migrate() {
	db := database.Session()
	db.AutoMigrate(&users.Model{})
}
