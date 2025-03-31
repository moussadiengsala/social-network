package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"learn.zone01dakar.sn/forum-rest-api/app/internals/core/app"
	"learn.zone01dakar.sn/forum-rest-api/app/lib"
	"learn.zone01dakar.sn/forum-rest-api/app/models"
	service "learn.zone01dakar.sn/forum-rest-api/app/service/CRUD"
)

type Reaction struct {
	Action map[string][]string
}

func (re *Reaction) SetAction() {
	re.Action = make(map[string][]string)
	re.Action["post_like"] = []string{"post_like", "post_dislike"}
	re.Action["post_dislike"] = []string{"post_dislike", "post_like"}
	re.Action["comment_like"] = []string{"comment_like", "comment_dislike"}
	re.Action["comment_dislike"] = []string{"comment_dislike", "comment_like"}
}

func (r *Reaction) Route(app *app.App) {
	r.SetAction()
	app.POST("/reactions", r.Create)
}

func (re Reaction) Create(w http.ResponseWriter, r *http.Request) {
	response := lib.Response{Code: 200, Message: "ok"}

	var credentials models.Reaction
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		fmt.Println(err)
		lib.ErrorWriter(&response, "Error getting informations to perform this operation!!!", http.StatusBadRequest)
		lib.ResponseFormatter(w, response)
		return
	}

	tables := re.Action[credentials.Action]
	db, _ := r.Context().Value(models.DBInstanceKey).(lib.DB)
	var sqlService = service.SqlService(db.Instance)
credentials.AuthorID = r.Context().Value(models.UserIDKey).(int)
	err := re.CheckAndDeleteReaction(sqlService, tables[0], []interface{}{credentials.AuthorID, credentials.EntriesID}, &credentials.ID)
	_, statusCode := lib.SqlError(err, []string{}, []string{})
	fmt.Println("status", statusCode, err)
	fmt.Println(credentials.AuthorID,credentials.EntriesID)
	if statusCode == 404 {
		q := fmt.Sprintf("INSERT INTO %s (author_id,entries_id) VALUES(?,?);", tables[0])
		fmt.Println("hi")
		if _, errCreate := sqlService.Create(q, credentials.AuthorID, credentials.EntriesID); errCreate != nil {
			fmt.Println(errCreate, q, credentials.AuthorID, credentials.EntriesID)
			message, statusCode := lib.SqlError(errCreate, []string{}, []string{})
			lib.ErrorWriter(&response, message, statusCode)
			lib.ResponseFormatter(w, response)
			return
		}

		errI := re.CheckAndDeleteReaction(sqlService, tables[1], []interface{}{credentials.AuthorID, credentials.EntriesID}, &credentials.ID)
		message, statusCode := lib.SqlError(errI, []string{}, []string{})
		if errI != nil && statusCode != 404 {
			lib.ErrorWriter(&response, message, statusCode)
			lib.ResponseFormatter(w, response)
			return
		}
	}

	response.Data = credentials
	lib.ResponseFormatter(w, response)
}

func (re Reaction) CheckAndDeleteReaction(sqlService service.DBService, table string, values []interface{}, dest *int) error {
	checkQuery := fmt.Sprintf("SELECT id FROM %s WHERE author_id = ? AND entries_id = ?", table)
	deletequery := fmt.Sprintf("DELETE FROM %s WHERE author_id = ? AND entries_id = ?;", table)

	err := sqlService.SelectSingle(checkQuery, values, dest)
	if err == nil {
		err = sqlService.Delete(deletequery, values...)
	}
	return err
}
