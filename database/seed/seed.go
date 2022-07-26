package seed

import (
	"api.sianggota.com/api/users"
	"api.sianggota.com/database"
	"github.com/bxcodec/faker/v3"
)

var funcMap = map[string]interface{}{
	"create_users": createUsers,
}

func Populate(seedName string) error {
	return funcMap[seedName].(func() error)()
}

func createUsers() error {
	u := []users.Model{}
	for i := 0; i < 100; i++ {
		email := faker.Email()
		name := faker.Name()
		u = append(u, users.Model{
			Name:     &name,
			Email:    &email,
			Password: faker.Password(),
		})
	}
	seed := database.Session().CreateInBatches(&u, 10)
	return seed.Error
}
