package seed

import (
	"time"

	"api.sianggota.com/api/users"
	"api.sianggota.com/database"
	"github.com/bxcodec/faker/v3"
	"github.com/cheggaaa/pb/v3"
)

var funcMap = map[string]interface{}{
	"create_users": createUsers,
}
var count int = 10
var cc int = 50

func Populate(seedName string, c int) error {
	count = c
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
