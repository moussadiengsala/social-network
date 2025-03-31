package routes

import (
	"fmt"
	"net/http"
	"strings"

	internals "learn.zone01dakar.sn/forum-rest-api/app/internals/config/database"
	core "learn.zone01dakar.sn/forum-rest-api/app/internals/core/app"
	"learn.zone01dakar.sn/forum-rest-api/app/internals/core/validators"
	"learn.zone01dakar.sn/forum-rest-api/app/lib"
	"learn.zone01dakar.sn/forum-rest-api/app/models"
	service "learn.zone01dakar.sn/forum-rest-api/app/service/CRUD"
)

type Post struct{}

func (p Post) Route(app *core.App) {
	app.POST("/posts", p.Create)
	app.GET("/posts", p.GetAllPost)
	app.GET("/feeds", p.GetUserPostIntractions)
}

func (p *Post) Create(w http.ResponseWriter, r *http.Request) {
	response := lib.Response{Code: 200, Message: "ok"}
	var credentials models.CreatePost

	if errGettingCredential := lib.ParseForm(&credentials, r); errGettingCredential != nil {
		lib.ErrorWriter(&response, "Something went wrong! make sure you fulfill all required fields!", http.StatusBadRequest)
		lib.ResponseFormatter(w, response)
		return
	}

	credentials.AuthorID = r.Context().Value(models.UserIDKey).(int)
	validators := validators.Validators{}
	if errValidator := validators.ValidatorService(credentials); errValidator != nil {
		lib.ErrorWriter(&response, validators.GetValidatorErrors(errValidator), http.StatusBadRequest)
		lib.ResponseFormatter(w, response)
		return
	}

	db, _ := r.Context().Value(models.DBInstanceKey).(lib.DB)
	var sqlService = service.SqlService(db.Instance)

	fields := lib.SlicerDBFieldsName(credentials, "", false)
	var query = fmt.Sprintf("INSERT INTO post(%s) VALUES(%s?)", strings.Join(fields, ","), strings.Repeat("?,", len(fields)-1))

	id, insertPostErr := sqlService.Create(query, lib.Slicer(credentials, false)...)
	if insertPostErr != nil {
		message, statusCode := lib.SqlError(insertPostErr, []string{"post", "comment"}, []string{"user"})
		lib.ErrorWriter(&response, message, statusCode)
		lib.ResponseFormatter(w, response)
		return
	}

	if credentials.Privacy == "almost_private" {
		values := []string{}
		for _, user := range strings.Split(credentials.SelectedUsers, ",") {
			values = append(values, fmt.Sprintf("(%v,%s)", id, user))
		}
		query = fmt.Sprintf("INSERT INTO post_visibility(post_id,user_id) VALUES %s", strings.Join(values, ", "))
		_, insertPostErr := sqlService.Create(query)
		if insertPostErr != nil {
			message, statusCode := lib.SqlError(insertPostErr, []string{"post", "comment"}, []string{"user", "post"})
			lib.ErrorWriter(&response, message, statusCode)
			lib.ResponseFormatter(w, response)
			return
		}
	}

	lib.ResponseFormatter(w, response)
}

func (p Post) GetAllPost(w http.ResponseWriter, r *http.Request) {
	response := lib.Response{Code: 200, Message: "ok"}

	posts := []models.GetPost{}
	queryParameters := r.URL.Query()

	userID := r.Context().Value(models.UserIDKey).(int)
	limit := queryParameters.Get("limit")
	offset := queryParameters.Get("offset")

	query := fmt.Sprintf(internals.QUERY_GETTING_All_POST, userID, userID, userID, userID, userID, userID, userID, limit, offset)
	feeds := lib.GetFeed{
		Credentials:     &posts,
		CredentialsType: models.GetPost{},
		Query:           query,
	}

	feeds.GetAllFeed(r, &response)
	lib.ResponseFormatter(w, response)
}

func (p Post) GetUserPostIntractions(w http.ResponseWriter, r *http.Request) {
	posts := []models.GetPost{}
	response := lib.Response{Code: 200, Message: "ok", Data: posts}
	userID := r.Context().Value(models.UserIDKey).(int)

	queryParameters := r.URL.Query()
	profileID := lib.Convert(queryParameters.Get("user_id"))
	section := queryParameters.Get("section")
	limit := queryParameters.Get("limit")
	offset := queryParameters.Get("offset")
	query := ""

	if userID != profileID {
		isPrivateAccount := lib.IsPrivateAccount(r, &response, profileID)
		isFreinds := lib.IsFollowsUser(r, profileID, userID)
		if !isFreinds && isPrivateAccount {
			lib.ErrorWriter(&response, "", http.StatusOK)
			lib.ResponseFormatter(w, response)
			return
		}
	}

	switch section {
	case "owned_posts":
		query = fmt.Sprintf(internals.QUERY_GETTING_OWN_POST, userID, userID, userID, userID, profileID, profileID, userID, limit, offset)
	case "commented_posts":
		query = fmt.Sprintf(internals.QUERY_GETTING_COMMENTED_POST, userID, userID, userID, userID, profileID, profileID, userID, userID, limit, offset)
	case "liked_posts":
		query = fmt.Sprintf(internals.QUERY_GETTING_LIKED_POST, userID, userID, userID, userID, profileID, profileID, userID, userID, limit, offset)
	default:
		lib.ResponseFormatter(w, response)
		return
	}

	feeds := lib.GetFeed{
		Credentials:     &posts,
		CredentialsType: models.GetPost{},
		Query:           query,
	}
	feeds.GetAllFeed(r, &response)
	lib.ResponseFormatter(w, response)
}

// func (p Post) GetPostByID(w http.ResponseWriter, r *http.Request) {
// 	response := lib.Response{Code: 200, Message: "ok"}
// 	userID := r.Context().Value(models.UserIDKey).(int)
// 	postID := lib.Convert(r.URL.Query().Get("post_id"))
// 	posts := models.GetPost{}
// 	query := fmt.Sprintf(internals.QUERY_GETTING_SINGLE_POST, userID, userID, userID, postID, userID, userID, userID)
// 	feeds := lib.GetFeed{
// 		Credentials: &posts,
// 		Query:       query,
// 	}
// 	feeds.GetSingleFeed(r, &response)
// 	lib.ResponseFormatter(w, response)
// }
