package seed

import (
	"time"

	"api.sianggota.com/api/users"
	"api.sianggota.com/database"
	"github.com/bxcodec/faker/v3"
	"github.com/cheggaaa/pb/v3"
	"github.com/neko-neko/echo-logrus/v2/log"
)

var funcMap = map[string]interface{}{
	"create_users": createUsers,
	"delete_users": deleteUsers,
}
var count int = 10
var cc int = 50

func Populate(seedName string, c int) error {
	if c != 0 {
		count = c
	}
	log.Info(count)
	if seedName == "" {
		var err error
		for _, f := range funcMap {
			err = f.(func() error)()
		}
		if err != nil {
			return err
		}
		return nil
	}
	return funcMap[seedName].(func() error)()
}

func createUsers() (err error) {
	var cf int = 1
	bar := pb.StartNew(count)
	if count > cc {
		cf = count / cc
	}
	for i := 0; i < cf; i++ {
		u := []users.Model{}
		for i := 0; i < cc; i++ {
			email := faker.Email()
			name := faker.Name()
			u = append(u, users.Model{
				Name:     &name,
				Email:    &email,
				Password: faker.Password(),
			})
		}
		seed := database.Session().CreateInBatches(&u, 100)
		err = seed.Error
		bar.Add(cc)
		time.Sleep(time.Microsecond)
	}
	bar.Finish()

	return
}

func deleteUsers() (err error) {
	for i := 0; i < count; i++ {
		err = database.Session().Exec("update users set deleted_at = transaction_timestamp()  where id = (SELECT id FROM users order by random() limit 1);").Error
	}
	return err
}
