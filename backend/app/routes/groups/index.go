package groups

import (
	"fmt"
	"net/http"
	"strconv"

	internals "learn.zone01dakar.sn/forum-rest-api/app/internals/config/database"
	app "learn.zone01dakar.sn/forum-rest-api/app/internals/core/app"
	"learn.zone01dakar.sn/forum-rest-api/app/internals/core/socket"
	"learn.zone01dakar.sn/forum-rest-api/app/lib"
	"learn.zone01dakar.sn/forum-rest-api/app/models"
	service "learn.zone01dakar.sn/forum-rest-api/app/service/CRUD"
	"learn.zone01dakar.sn/forum-rest-api/app/service/socket/utils"
)

type Groups struct {
	App *app.App
}

func (g *Groups) Route(app *app.App) {
	g.App = app
	//
	app.POST("/groups", g.Create)
	app.GET("/groups", g.GetAllGroups)
	app.GET("/groups/{group_id:string}", g.GetByIDGroups)

	// groups events
	app.POST("/groups-events", g.CreateEvents)
	app.GET("/groups-events", g.GetAllGroupsEvents)                    // need the query group_id
	app.GET("/groups-events/{event_id:string}", g.GetByIDGroupsEvents) // need the query group_id

	// groups response
	app.POST("/groups-events-response", g.CreateResponseEvent)
	app.GET("/groups-events-response", g.GetAllGroupsEventsResponse) // need the query group_id, event_id

	// groups posts
	app.POST("/groups-posts", g.CreatePosts)
	app.GET("/groups-posts", g.GetAllGroupsPost)                   // need the query group_id
	app.GET("/groups-posts/{post_id:string}", g.GetByIDGroupsPost) // need the query group_id

	// groups comments
	app.POST("/groups-comments", g.CreateComments)
	app.GET("/groups-comments", g.GetAllGroupsComment)                      // need the query post_id
	app.GET("/groups-comments/{comment_id:string}", g.GetByIDGroupsComment) // need the query post_id

	// group member
	app.POST("/groups-join", g.RequestOrInviteJoinGroup)
	app.POST("/groups-members", g.Accept)             // need the query group_id, member_id, action(invite or request)
	app.GET("/groups-members", g.GetAllGroupsMembers) // need the query group_id
}

func (g *Groups) Create(w http.ResponseWriter, r *http.Request) {
	response := lib.Response{Code: 200, Message: "ok"}
	var credentials models.Groups

	if err := lib.ParseForm(&credentials, r); err != nil {
		lib.ErrorWriter(&response, "Something went wrong! Make sure you fulfill all required fields!", http.StatusBadRequest)
		lib.ResponseFormatter(w, response)
		return
	}

	userID := r.Context().Value(models.UserIDKey).(int)
	credentials.AuthorID = userID
	db, _ := r.Context().Value(models.DBInstanceKey).(lib.DB)
	var sqlService = service.SqlService(db.Instance)

	createAdminMember := func(response *lib.Response, id int) error {
		if response.Code == 200 {
			memberCredentials := models.GroupMembers{
				Status:  "accepted",
				GroupID: id,
				User:    strconv.Itoa(credentials.AuthorID),
				Role:    "admin",
			}
			members := lib.CreateData{
				Credentials:      memberCredentials,
				Table:            "groupMembers",
				ForeignFields:    []string{"user", "group"},
				LookingForFields: []string{"member"},
			}
			members.Create(response, sqlService)
		}
		return nil
	}

	groups := lib.CreateData{
		Credentials:   credentials,
		Table:         "groups",
		ForeignFields: []string{"user"},
		Callback:      createAdminMember,
	}
	groups.Create(&response, sqlService)

	client, ok := g.App.Hub.Clients[userID]
	if ok {
		payload, err := utils.GetPayloadGroups(r, *g.App.Hub, userID, &response)
		if err != nil {
			message, statusCode := lib.SqlError(err, []string{"user"}, []string{})
			lib.ErrorWriter(&response, message, statusCode)
			utils.HandleUnexpectedEvent(client, &response, []int{client.User.ID})
			return
		}

		utils.EmitToSpecificClient(client.Hub, client, socket.SocketEventResponse{
			Action:  utils.CREATEGROUP,
			Payload: lib.Response{Code: 200, Message: "ok", Data: payload},
		}, []int{client.User.ID})
	}

	lib.ResponseFormatter(w, response)
}

// Only groups where you not part in.
func (g Groups) GetAllGroups(w http.ResponseWriter, r *http.Request) {
	response := lib.Response{Code: 200, Message: "ok"}
	userID := r.Context().Value(models.UserIDKey).(int)

	params := r.URL.Query()
	limit := params.Get("limit")
	offset := params.Get("offset")

	credentials := []models.GetGroups{}
	groups := lib.GetFeed{
		Credentials:     &credentials,
		CredentialsType: models.GetGroups{},
		Query:           fmt.Sprintf(internals.QUERY_GETTING_OTHER_GROUPS, userID, userID, userID, limit, offset),
	}

	groups.GetAllFeed(r, &response)
	lib.ResponseFormatter(w, response)
}

func (g Groups) GetByIDGroups(w http.ResponseWriter, r *http.Request) {
	response := lib.Response{Code: 200, Message: "ok"}
	credentials := models.GetGroups{}

	userID := r.Context().Value(models.UserIDKey).(int)
	groupID := lib.Convert(r.URL.Query().Get("group_id"))

	groups := lib.GetFeed{
		Credentials: &credentials,
		Query:       fmt.Sprintf(internals.QUERYGETTINGGROUPS, userID, userID, userID, groupID),
	}

	groups.GetSingleFeed(r, &response)
	lib.ResponseFormatter(w, response)
}
