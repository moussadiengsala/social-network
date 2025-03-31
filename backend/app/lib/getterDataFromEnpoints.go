package lib

import (
	"fmt"
	"net/http"

	"learn.zone01dakar.sn/forum-rest-api/app/models"
	service "learn.zone01dakar.sn/forum-rest-api/app/service/CRUD"
)

type GetFeed struct {
	Credentials     interface{}
	CredentialsType interface{}
	Query           string
	Conditions      []interface{}

	LookingForFields []string
	ForeignFields    []string
}

func (g GetFeed) GetAllFeed(r *http.Request, response *Response) {
	db, _ := r.Context().Value(models.DBInstanceKey).(DB)
	var sqlService = service.SqlService(db.Instance)

	if err := sqlService.SelectAllForStruct(g.Query, g.Conditions, g.Credentials, SlicerFieldsName(g.CredentialsType, true)); err != nil {
		fmt.Println(err)
		message, code := SqlError(err, g.LookingForFields, g.ForeignFields)
		ErrorWriter(response, message, code)
		return
	}

	response.Data = g.Credentials
}

func (g GetFeed) GetSingleFeed(r *http.Request, response *Response) {
	db, _ := r.Context().Value(models.DBInstanceKey).(DB)
	var sqlService = service.SqlService(db.Instance)

	if err := sqlService.SelectSingle(g.Query, g.Conditions, SlicerReferenceFields(g.Credentials, true)...); err != nil {
		fmt.Println(err)
		message, code := SqlError(err, g.LookingForFields, g.ForeignFields)
		ErrorWriter(response, message, code)
		return
	}

	response.Data = g.Credentials
}
