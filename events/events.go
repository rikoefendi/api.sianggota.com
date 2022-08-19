package events

import "github.com/gookit/event"

func New() {
	event.On("users.create", UserCreateListener())
}
