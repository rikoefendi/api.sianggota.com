package events

import (
	"api.sianggota.com/api/users"
	"github.com/gookit/event"
	"github.com/neko-neko/echo-logrus/v2/log"
)

func UserCreateListener() event.ListenerFunc {
	return func(e event.Event) error {
		user := e.Data()["users"].(users.Model)
		log.Info("user event create fired", *user.Email)
		return nil
	}
}
