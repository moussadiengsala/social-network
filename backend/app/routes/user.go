package routes

import (
	"fmt"
	"net/http"
	"strings"

	internals "learn.zone01dakar.sn/forum-rest-api/app/internals/config/database"
	core "learn.zone01dakar.sn/forum-rest-api/app/internals/core/app"
	"learn.zone01dakar.sn/forum-rest-api/app/lib"
	"learn.zone01dakar.sn/forum-rest-api/app/models"
	service "learn.zone01dakar.sn/forum-rest-api/app/service/CRUD"
)

type User struct{}

func (u User) Route(app *core.App) {
	app.GET("/users/{user_id:string}", u.UserProfile)
	app.GET("/users", u.GetAllUsers)

	app.GET("/users/{user_id:string}/follows", u.GetUsersFollows)
	app.POST("/change-privacy-status", u.SettingsProfile)
}

func (u *User) UserProfile(w http.ResponseWriter, r *http.Request) {
	response := lib.Response{Code: 200, Message: "ok"}

	targetID := lib.Convert(r.URL.Query().Get("user_id"))
	userID := r.Context().Value(models.UserIDKey).(int)

	credentials := models.GetUser{}
	fields := lib.SlicerDBFieldsName(models.User{}, "u.", true)
	var query = fmt.Sprintf(internals.QUERY_GETTING_USER_PROFILE, strings.Join(fields, ","), userID, userID, targetID, strings.Join(fields, ","))

	user := lib.GetFeed{
		Credentials: &credentials,
		Query:       query,
	}

	user.GetSingleFeed(r, &response)
	lib.ResponseFormatter(w, response)
}

// Only users who you not friends with.
func (u *User) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	response := lib.Response{Code: 200, Message: "ok"}
	params := r.URL.Query()
	userID := r.Context().Value(models.UserIDKey).(int)
	credentials := []models.GetUser{}
	fields := lib.SlicerDBFieldsName(models.User{}, "u.", true)
	limit := params.Get("limit")
	offset := params.Get("offset")

	var query = fmt.Sprintf(internals.QUERY_GETTING_All_USERS, strings.Join(fields, ","), userID, userID, userID, userID, userID, strings.Join(fields, ","), limit, offset)
	user := lib.GetFeed{
		Credentials:     &credentials,
		CredentialsType: models.GetUser{},
		Query:           query,
	}

	user.GetAllFeed(r, &response)
	lib.ResponseFormatter(w, response)
}

func (u User) GetUsersFollows(w http.ResponseWriter, r *http.Request) {
	response := lib.Response{Code: 200, Message: "ok"}

	credentials := []models.GetFollow{}
	params := r.URL.Query()
	userID := lib.Convert(params.Get("user_id"))
	entry := params.Get("entry")
	query := ""

	switch entry {
	case "followers":
		query = fmt.Sprintf(internals.QUERY_GETTING_ALL_FOLLOWERS, userID)
	case "followings":
		query = fmt.Sprintf(internals.QUERY_GETTING_ALL_FOLLOWING, userID)
	default:
		lib.ErrorWriter(&response, "invalid query.", http.StatusBadRequest)
		lib.ResponseFormatter(w, response)
		return
	}

	follower := lib.GetFeed{
		Credentials:     &credentials,
		CredentialsType: models.GetFollow{},
		Query:           query,
	}
	follower.GetAllFeed(r, &response)
	lib.ResponseFormatter(w, response)
}

func (u User) SettingsProfile(w http.ResponseWriter, r *http.Request) {
	response := lib.Response{Code: 200, Message: "ok"}
	userID := r.Context().Value(models.UserIDKey).(int)

	db, _ := r.Context().Value(models.DBInstanceKey).(lib.DB)
	var sqlService = service.SqlService(db.Instance)
	query := fmt.Sprintf(`UPDATE user SET privacy = 'private' WHERE id = %d;`, userID)

	if isPrivate := lib.IsPrivateAccount(r, &response, userID); isPrivate {
		query = fmt.Sprintf(`UPDATE user SET privacy = 'public' WHERE id = %d;`, userID)
	}

	if _, insertErr := sqlService.Update(query); insertErr != nil {
		message, statusCode := lib.SqlError(insertErr, []string{"user"}, []string{})
		lib.ErrorWriter(&response, message, statusCode)
	}

	lib.ResponseFormatter(w, response)
}

// func (u User) SettingsProfile(w http.ResponseWriter, r *http.Request) {
// 	response := lib.Response{Code: 200, Message: "ok"}
// 	userID := r.Context().Value(models.UserIDKey).(int)
// 	credentilas := models.SettingsUserProfile{}
// 	type Items struct {
// 		Item      string
// 		Value     string
// 		Condition bool
// 	}
// 	if err := lib.ParseForm(credentilas, r); err != nil {
// 		lib.ErrorWriter(&response, "Something went wrong! make sure you fulfill all required fields!", http.StatusBadRequest)
// 		lib.ResponseFormatter(w, response)
// 		return
// 	}
// 	L := []Items{
// 		{Item: "avatar", Value: credentilas.Avatar, Condition: credentilas.Avatar != ""},
// 		{Item: "bio", Value: credentilas.Bio, Condition: credentilas.Bio != ""},
// 		{Item: "privacy", Value: credentilas.Privacy, Condition: credentilas.Privacy != "" && (credentilas.Privacy == "public" || credentilas.Privacy == "private")},
// 	}
// 	db, _ := r.Context().Value(models.DBInstanceKey).(lib.DB)
// 	var sqlService = service.SqlService(db.Instance)
// 	chamgeItems := func(items, value string) error {
// 		query := fmt.Sprintf(`UPDATE user SET %s = %s WHERE id = %d;`, items, value, userID)
// 		if _, insertErr := sqlService.Update(query); insertErr != nil {
// 			message, statusCode := lib.SqlError(insertErr, []string{"follower"}, []string{"user"})
// 			lib.ErrorWriter(&response, message, statusCode)
// 			return insertErr
// 		}
// 		return nil
// 	}
// 	for _, v := range L {
// 		if v.Condition{
// 			chamgeItems(v.Item, v.Value)
// 		}
// 	}
// 	lib.ResponseFormatter(w, response)
// }
