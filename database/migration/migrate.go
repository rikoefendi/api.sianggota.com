package migration

import (
	"api.sianggota.com/api/users"
	"api.sianggota.com/database"
	"gorm.io/gorm"
)

var db *gorm.DB

func Migrate(name string, refresh bool) {
	db = database.Session()
	if refresh {
		drop()
	}
	db.AutoMigrate(&users.Model{})
}

func drop() {
	db.Migrator().DropTable(&users.Model{})
}
