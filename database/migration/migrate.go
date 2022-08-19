package migration

import (
	"os"

	"api.sianggota.com/api/users"
	"api.sianggota.com/database"
	"gorm.io/gorm"
)

var db *gorm.DB

func Migrate(name string, refresh bool, init bool) error {
	db = database.Session()
	if init {
		return InitDatabase(name)
	}
	if refresh {
		drop()
	}
	return db.Error
}

func InitDatabase(fileName string) error {
	dir, err := os.Getwd()
	if err != nil {
		return err
	}
	sqlByte, err := os.ReadFile(dir + "/" + fileName)
	if err != nil {
		return err
	}
	tx := db.Exec(string(sqlByte))
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
func drop() {
	db.Migrator().DropTable(&users.Model{})
}
