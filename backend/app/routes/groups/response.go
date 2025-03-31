package groups

import (
	"fmt"
	"net/http"
	"strings"

	internals "learn.zone01dakar.sn/forum-rest-api/app/internals/config/database"
	"learn.zone01dakar.sn/forum-rest-api/app/internals/core/validators"
	"learn.zone01dakar.sn/forum-rest-api/app/lib"
	"learn.zone01dakar.sn/forum-rest-api/app/models"
	service "learn.zone01dakar.sn/forum-rest-api/app/service/CRUD"
)

func (g Groups) CreateResponseEvent(w http.ResponseWriter, r *http.Request) {
	response := lib.Response{Code: 200, Message: "ok"}
	var credentials models.GroupEventResponses

	params := r.URL.Query()
	credentials.EventID = lib.Convert(params.Get("event_id"))
	credentials.GroupID = lib.Convert(params.Get("group_id"))
	credentials.Response = strings.Join(strings.Split(params.Get("response"), "_"), " ")
	userID := r.Context().Value(models.UserIDKey).(int)
	credentials.UserID = userID

	validators := validators.Validators{}
	if errValidator := validators.ValidatorService(credentials); errValidator != nil {
		lib.ErrorWriter(&response, validators.GetValidatorErrors(errValidator), http.StatusBadRequest)
		lib.ResponseFormatter(w, response)
		return
	}

	if isAuthorUserMember := lib.IsMembersExistAndAccepted(r, credentials.GroupID, userID); !isAuthorUserMember {
		lib.ErrorWriter(&response, "the group is not found or you not part of it to perform this operation!", http.StatusBadRequest)
		lib.ResponseFormatter(w, response)
		return
	}

	db, _ := r.Context().Value(models.DBInstanceKey).(lib.DB)
	var sqlService = service.SqlService(db.Instance)

	responseEvent := lib.CreateData{
		Credentials:   credentials,
		Table:         "groupEventResponses",
		ForeignFields: []string{"events", "user", "group"},
	}

	responseEvent.CreateWithoutValidator(&response, sqlService)
	g.ChangingResponseEvent(credentials, &response, sqlService)

	lib.ResponseFormatter(w, response)
}

func (g Groups) ChangingResponseEvent(credentials models.GroupEventResponses, response *lib.Response, sqlService service.DBService) {
	if response.Code == 409 {
		query := fmt.Sprintf(`UPDATE groupEventResponses SET response = '%s' WHERE event_id = %d AND group_id = %d AND user_id = %d;`, credentials.Response, credentials.EventID, credentials.GroupID, credentials.UserID)
		_, insertPostErr := sqlService.Update(query)

		if insertPostErr != nil {
			message, statusCode := lib.SqlError(insertPostErr, []string{"response"}, []string{})
			lib.ErrorWriter(response, message, statusCode)
			return
		}
		lib.ErrorWriter(response, "ok", http.StatusOK)
	}
}

func (g Groups) GetAllGroupsEventsResponse(w http.ResponseWriter, r *http.Request) {
	response := lib.Response{Code: 200, Message: "ok"}

	params := r.URL.Query()
	groupID := lib.Convert(params.Get("group_id"))
	eventID := lib.Convert(params.Get("event_id"))

	credentials := []models.GetGroupEventResponses{}

	events := lib.GetFeed{
		Credentials:     &credentials,
		CredentialsType: models.GetGroupEventResponses{},
		Query:           fmt.Sprintf(internals.QUERYGETTINGEVENTSGROUPSRESPONSE, groupID, eventID),
	}

	events.GetAllFeed(r, &response)
	lib.ResponseFormatter(w, response)
}
