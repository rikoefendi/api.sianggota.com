package command

import (
	"flag"
	"os"

	"api.sianggota.com/database/migration"
	"api.sianggota.com/database/seed"
)

func New() {
	if len(os.Args) < 2 {
		return
	}
	seed := flag.NewFlagSet("database seeder", flag.ExitOnError)
	nameFile := seed.String("name", "", "fucntion name to seed")
	countCmd := seed.Int("count", 0, "")

	migrate := flag.NewFlagSet("database migrator", flag.ExitOnError)
	nameMigrate := migrate.String("name", "", "function name to migrate")
	refresh := migrate.Bool("refresh", false, "to refresh database")
	init := migrate.Bool("init", false, "to init database from sql file")
	switch os.Args[1] {
	case "seed":
		seed.Parse(os.Args[2:])
		err := seeder(string(*nameFile), *countCmd)
		if err != nil {
			panic(err)
		}
		os.Exit(1)
	case "migrate":
		migrate.Parse(os.Args[2:])
		name := string(*nameMigrate)
		err := migrator(name, *refresh, *init)
		if err != nil {
			panic(err)
		}
		os.Exit(1)
	default:
		os.Exit(1)
	}
}

func seeder(seedName string, count int) error {
	return seed.Populate(seedName, count)
}
func migrator(name string, refresh bool, init bool) error {
	return migration.Migrate(name, refresh, init)
}
