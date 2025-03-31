package routes

import (
	"net/http"

	core "learn.zone01dakar.sn/forum-rest-api/app/internals/core/app"
	"learn.zone01dakar.sn/forum-rest-api/app/lib"
)

type Verify struct{}

func (g Verify) Route(app *core.App) {
	app.GET("/verify", g.Init)
}

func (g Verify) Init(w http.ResponseWriter, r *http.Request) {
	response := lib.Response{Code: 200, Message: "ok"}
	lib.ResponseFormatter(w, response)
}
