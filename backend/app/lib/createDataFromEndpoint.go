package lib

import (
	"fmt"
	"net/http"
	"strings"

	"learn.zone01dakar.sn/forum-rest-api/app/internals/core/validators"
	service "learn.zone01dakar.sn/forum-rest-api/app/service/CRUD"
)

type CreateData struct {
	ID          int
	Table       string
	Credentials interface{}
	Callback    func(response *Response, id int) error

	LookingForFields []string
	ForeignFields    []string
}

func (f *CreateData) Create(response *Response, sqlService service.DBService) {

	validators := validators.Validators{}
	if errValidator := validators.ValidatorService(f.Credentials); errValidator != nil {
		ErrorWriter(response, validators.GetValidatorErrors(errValidator), http.StatusBadRequest)
		return
	}

	fields := SlicerDBFieldsName(f.Credentials, "", false)

	var query = fmt.Sprintf("INSERT INTO %s(%s) VALUES(%s?)", f.Table, strings.Join(fields, ","), strings.Repeat("?,", len(fields)-1))
	id, insertPostErr := sqlService.Create(query, Slicer(f.Credentials, false)...)
	if insertPostErr != nil {
		message, statusCode := SqlError(insertPostErr, f.LookingForFields, f.ForeignFields)
		ErrorWriter(response, message, statusCode)
		return
	}

	f.ID = int(id)
	if f.Callback != nil {
		if err := f.Callback(response, f.ID); err != nil {
			message, statusCode := SqlError(err, []string{"post", "comment"}, []string{})
			ErrorWriter(response, message, statusCode)
			return
		}
	}
}

func (f *CreateData) CreateWithoutValidator(response *Response, sqlService service.DBService) {

	fields := SlicerDBFieldsName(f.Credentials, "", false)

	var query = fmt.Sprintf("INSERT INTO %s(%s) VALUES(%s?)", f.Table, strings.Join(fields, ","), strings.Repeat("?,", len(fields)-1))
	id, insertPostErr := sqlService.Create(query, Slicer(f.Credentials, false)...)
	if insertPostErr != nil {
		message, statusCode := SqlError(insertPostErr, f.LookingForFields, f.ForeignFields)
		ErrorWriter(response, message, statusCode)
		return
	}

	f.ID = int(id)
	if f.Callback != nil {
		if err := f.Callback(response, f.ID); err != nil {
			message, statusCode := SqlError(err, []string{"post", "comment"}, []string{})
			ErrorWriter(response, message, statusCode)
			return
		}
	}
}
