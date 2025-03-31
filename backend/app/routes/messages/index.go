package message

import (
	"net/http"

	core "learn.zone01dakar.sn/forum-rest-api/app/internals/core/app"
	"learn.zone01dakar.sn/forum-rest-api/app/lib"
	"learn.zone01dakar.sn/forum-rest-api/app/models"
	service "learn.zone01dakar.sn/forum-rest-api/app/service/CRUD"
)

type Message struct {
	App *core.App
}

func (pm *Message) Route(app *core.App) {
	pm.App = app

	app.POST("/group-message", pm.CreateGroupMessage)
	app.GET("/message/{message_id:string}", pm.GetByIDPrivateMessage)
	app.GET("/messages/{target_id:string}", pm.GetALLPrivateMessages)

	app.POST("/private-message", pm.CreatePrivateMessage)
	app.GET("/groups/message/{message_id:string}", pm.GetByIDGroupsMessage)
	app.GET("/groups/messages/{group_id:string}", pm.GetALLGroupsMessages)
}

func (pm Message) MarkAsReadedPrivateMessage(r *http.Request, response *lib.Response, query string) {
	db, _ := r.Context().Value(models.DBInstanceKey).(lib.DB)
	var sqlService = service.SqlService(db.Instance)

	if _, insertErr := sqlService.Update(query); insertErr != nil {
		message, statusCode := lib.SqlError(insertErr, []string{"follower"}, []string{"user"})
		lib.ErrorWriter(response, message, statusCode)
		return
	}
}
