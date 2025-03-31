package auth

import (
	app "learn.zone01dakar.sn/forum-rest-api/app/internals/core/app"
)

type Auth struct{}

func (a Auth) Route(app *app.App) {
	app.POST("/auth/signin", a.SignIn)
	app.POST("/auth/signup", a.Register)
}
