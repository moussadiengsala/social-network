package routes

import (
	"net/http"

	core "learn.zone01dakar.sn/forum-rest-api/app/internals/core/app"
	"learn.zone01dakar.sn/forum-rest-api/app/lib"
	"learn.zone01dakar.sn/forum-rest-api/app/models"
	service "learn.zone01dakar.sn/forum-rest-api/app/service/CRUD"
)

type Reply struct{}

func (rp Reply) Route(app *core.App) {
	app.POST("/reply", rp.Create)
	// app.GET("/comments/", p.GetCommentByID)
}

func (rp Reply) Create(w http.ResponseWriter, r *http.Request) {
	response := lib.Response{Code: 200, Message: "ok"}

	var credentials models.CreateReply
	if errGettingCredential := lib.ParseForm(&credentials, r); errGettingCredential != nil {
		lib.ErrorWriter(&response, "Something went wrong! make sure you fulfill all required fields!", http.StatusBadRequest)
		lib.ResponseFormatter(w, response)
		return
	}

	db, _ := r.Context().Value(models.DBInstanceKey).(lib.DB)
	var sqlService = service.SqlService(db.Instance)

	reply := lib.CreateData{
		Credentials: credentials,
		Table:       "reply",
	}

	reply.Create(&response, sqlService)
	lib.ResponseFormatter(w, response)
}
