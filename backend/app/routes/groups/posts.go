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

func (g Groups) CreatePosts(w http.ResponseWriter, r *http.Request) {
	response := lib.Response{Code: 200, Message: "ok"}
	var credentials models.GroupPost

	if errGettingCredential := lib.ParseForm(&credentials, r); errGettingCredential != nil {
		fmt.Println(errGettingCredential)
		lib.ErrorWriter(&response, "Something went wrong! make sure you fulfill all required fields!", http.StatusBadRequest)
		lib.ResponseFormatter(w, response)
		return
	}

	userID := r.Context().Value(models.UserIDKey).(int)
	credentials.AuthorID = userID
	db, _ := r.Context().Value(models.DBInstanceKey).(lib.DB)
	var sqlService = service.SqlService(db.Instance)

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

	post := lib.CreateData{
		Credentials:   credentials,
		Table:         "groupPosts",
		ForeignFields: []string{"user", "group"},
	}

	post.CreateWithoutValidator(&response, sqlService)
	lib.ResponseFormatter(w, response)
}

func (g Groups) GetAllGroupsPost(w http.ResponseWriter, r *http.Request) {
	response := lib.Response{Code: 200, Message: "ok"}
	credentials := []models.GetGroupPost{}
	// userID := r.Context().Value(models.UserIDKey).(int)

	params := r.URL.Query()
	groupID := lib.Convert(params.Get("group_id"))
	limit := params.Get("limit")
	offset := params.Get("offset")

	posts := lib.GetFeed{
		Credentials:     &credentials,
		CredentialsType: models.GetGroupPost{},
		Query:           fmt.Sprintf(internals.QUERY_GETTING_ALL_GROUP_POSTS, groupID, limit, offset),
	}

	posts.GetAllFeed(r, &response)
	lib.ResponseFormatter(w, response)
}

func (g Groups) GetByIDGroupsPost(w http.ResponseWriter, r *http.Request) {
	response := lib.Response{Code: 200, Message: "ok"}

	postID := r.URL.Query().Get("post_id")
	credentials := models.GetGroupPost{}
	post := lib.GetFeed{
		Credentials: &credentials,
		Query:       fmt.Sprintf(internals.QUERY_GETTING_GROUP_POSTS),
		Conditions:  []interface{}{postID},
	}

	post.GetSingleFeed(r, &response)
	lib.ResponseFormatter(w, response)
}
