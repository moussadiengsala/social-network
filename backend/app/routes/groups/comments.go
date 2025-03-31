package groups

import (
	"fmt"
	"net/http"

	internals "learn.zone01dakar.sn/forum-rest-api/app/internals/config/database"
	"learn.zone01dakar.sn/forum-rest-api/app/internals/core/validators"
	"learn.zone01dakar.sn/forum-rest-api/app/lib"
	"learn.zone01dakar.sn/forum-rest-api/app/models"
	service "learn.zone01dakar.sn/forum-rest-api/app/service/CRUD"
)

func (g Groups) CreateComments(w http.ResponseWriter, r *http.Request) {
	response := lib.Response{Code: 200, Message: "ok"}
	var credentials models.GroupComments

	if errGettingCredential := lib.ParseForm(&credentials, r); errGettingCredential != nil {
		fmt.Println(errGettingCredential)
		lib.ErrorWriter(&response, "Something went wrong! make sure you fulfill all required fields!", http.StatusBadRequest)
		lib.ResponseFormatter(w, response)
		return
	}

	userID := r.Context().Value(models.UserIDKey).(int)
	credentials.AuthorID = userID
	validators := validators.Validators{}
	if errValidator := validators.ValidatorService(credentials); errValidator != nil {
		lib.ErrorWriter(&response, validators.GetValidatorErrors(errValidator), http.StatusBadRequest)
		lib.ResponseFormatter(w, response)
		return
	}

	if isMember := lib.IsMembersExistAndAccepted(r, credentials.GroupID, credentials.AuthorID); !isMember {
		lib.ErrorWriter(&response, "you need to be part of this group to perform this operation.", http.StatusBadRequest)
		lib.ResponseFormatter(w, response)
		return
	}

	db, _ := r.Context().Value(models.DBInstanceKey).(lib.DB)
	var sqlService = service.SqlService(db.Instance)

	comment := lib.CreateData{
		Credentials:   credentials,
		Table:         "groupComments",
		ForeignFields: []string{"user", "group", "post"},
	}

	comment.CreateWithoutValidator(&response, sqlService)
	lib.ResponseFormatter(w, response)
}

func (g Groups) GetAllGroupsComment(w http.ResponseWriter, r *http.Request) {
	response := lib.Response{Code: 200, Message: "ok"}
	credentials := []models.GetGroupComment{}

	params := r.URL.Query()
	postID := lib.Convert(params.Get("post_id"))
	limit := params.Get("limit")
	offset := params.Get("offset")

	posts := lib.GetFeed{
		Credentials:     &credentials,
		CredentialsType: models.GetGroupComment{},
		Query:           fmt.Sprintf(internals.QUERY_GETTING_ALL_GROUP_COMMENTS, postID, limit, offset),
	}

	posts.GetAllFeed(r, &response)
	lib.ResponseFormatter(w, response)
}

func (g Groups) GetByIDGroupsComment(w http.ResponseWriter, r *http.Request) {
	response := lib.Response{Code: 200, Message: "ok"}
	commentID := lib.Convert(r.URL.Query().Get("comment_id"))

	credentials := models.GetGroupComment{}
	post := lib.GetFeed{
		Credentials: &credentials,
		Query:       fmt.Sprintf(internals.QUERY_GETTING_GROUP_COMMENTS, commentID),
	}

	post.GetSingleFeed(r, &response)
	lib.ResponseFormatter(w, response)
}
