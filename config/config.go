package config

import (
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/gommon/log"
)

type App struct {
	Name string `env:"APP_NAME" env-default:"sianggota"`
	Host string `env:"APP_HOST" env-default:"localhost"`
	Port int    `env:"APP_PORT" env-default:"3000"`
}
type Database struct {
	Name     string `env:"DB_DB_NAME" env-default:"sianggota"`
	Host     string `env:"DB_HOST" env-default:"localhost"`
	Port     int    `env:"DB_PORT" env-default:"5432"`
	Username string `env:"DB_USERNAME" env-default:"postgres"`
	Password string `env:"DB_PASSWORD" env-default:"root"`
	SSLMode  string `env:"DB_SSL_MODE" env-default:"disable"`
	Logger   bool   `env:"DB_LOGGER" env-default:"true"`
	TimeZone string `env:"DB_TIMEZONE" env-default:"Asia/Jakarta"`
}

type Config struct {
	App
	Database
}

var lock = &sync.Mutex{}
var config *Config
var cfg Config

func GetConfig() *Config {
	lock.Lock()
	defer lock.Unlock()
	if config == nil {
		err := cleanenv.ReadEnv(&cfg)
		if err != nil {
			log.Error(err)
		}
		config = &cfg
	}
	return config
}
