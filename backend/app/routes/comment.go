package routes

import (
	"fmt"
	"net/http"

	internals "learn.zone01dakar.sn/forum-rest-api/app/internals/config/database"
	core "learn.zone01dakar.sn/forum-rest-api/app/internals/core/app"
	"learn.zone01dakar.sn/forum-rest-api/app/lib"
	"learn.zone01dakar.sn/forum-rest-api/app/models"
	service "learn.zone01dakar.sn/forum-rest-api/app/service/CRUD"
)

type Comment struct{}

func (p Comment) Route(app *core.App) {
	app.POST("/comments", p.Create)
	app.GET("/comments", p.GetAllComments)
}

func (c Comment) Create(w http.ResponseWriter, r *http.Request) {
	response := lib.Response{Code: 200, Message: "ok"}

	var credentials models.CreateComments
	if errGettingCredential := lib.ParseForm(&credentials, r); errGettingCredential != nil {
		lib.ErrorWriter(&response, "Something went wrong! make sure you fulfill all required fields!", http.StatusBadRequest)
		lib.ResponseFormatter(w, response)
		return
	}

	userID := r.Context().Value(models.UserIDKey).(int)
	credentials.AuthorID = userID
	db, _ := r.Context().Value(models.DBInstanceKey).(lib.DB)
	var sqlService = service.SqlService(db.Instance)

	comment := lib.CreateData{
		Credentials: credentials,
		Table:       "Comments",
	}

	comment.Create(&response, sqlService)
	lib.ResponseFormatter(w, response)
}

func (p Comment) GetAllComments(w http.ResponseWriter, r *http.Request) {
	response := lib.Response{Code: 200, Message: "ok"}

	userID := r.Context().Value(models.UserIDKey).(int)
	params := r.URL.Query()
	postID := lib.Convert(params.Get("post_id"))
	limit := params.Get("limit")
	offset := params.Get("offset")

	comment := []models.GetComments{}
	query := fmt.Sprintf(internals.QUERY_GETTING_All_COMMENTS, userID, postID, limit, offset)

	feeds := lib.GetFeed{
		Credentials:     &comment,
		CredentialsType: models.GetComments{},
		Query:           query,
	}

	feeds.GetAllFeed(r, &response)
	lib.ResponseFormatter(w, response)
}

// func (p Comment) GetCommentsByID(w http.ResponseWriter, r *http.Request) {
// 	response := lib.Response{Code: 200, Message: "ok"}
// 	userID := r.Context().Value("user_id").(int)
// 	commentID := lib.Convert(r.URL.Query().Get("comment_id"))
// 	comment := models.GetComments{}
// 	query := fmt.Sprintf(internals.QUERYGETTINGSINGLECOMMENTS, userID, userID, commentID)
// 	feeds := lib.GetFeed{
// 		Credentials: &comment,
// 		Query:       query,
// 	}
// 	feeds.GetSingleFeed(r, &response)
// 	lib.ResponseFormatter(w, response)
// }
