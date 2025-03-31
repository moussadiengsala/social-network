package routes

import (
	core "learn.zone01dakar.sn/forum-rest-api/app/internals/core/app"
	"learn.zone01dakar.sn/forum-rest-api/app/routes/auth"
	"learn.zone01dakar.sn/forum-rest-api/app/routes/groups"
	message "learn.zone01dakar.sn/forum-rest-api/app/routes/messages"
	"learn.zone01dakar.sn/forum-rest-api/app/routes/notifications"
)

type Router interface {
	Route(app *core.App)
}

func Handle(app *core.App) {
	// Defining the differents endpoints
	var endpoints = []Router{&auth.Auth{}, &Reaction{}, &Post{}, &Comment{}, &Reply{}, &Follow{}, &groups.Groups{}, &User{}, &Static{}, &Verify{}, &message.Message{}, &notifications.Notifications{}, &message.Message{}}
	for _, s := range endpoints {
		s.Route(app)
	}
}
