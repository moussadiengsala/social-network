package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	internals "learn.zone01dakar.sn/forum-rest-api/app/internals/config/database"
	tokenjwt "learn.zone01dakar.sn/forum-rest-api/app/internals/config/session"

	validator "learn.zone01dakar.sn/forum-rest-api/app/internals/core/validators"
	"learn.zone01dakar.sn/forum-rest-api/app/lib"
	"learn.zone01dakar.sn/forum-rest-api/app/models"
	service "learn.zone01dakar.sn/forum-rest-api/app/service/CRUD"
)

func (a Auth) SignIn(w http.ResponseWriter, r *http.Request) {
	response := lib.Response{Code: 200, Message: "ok"}
	var credentials models.Credentials

	errGettingCredential := json.NewDecoder(r.Body).Decode(&credentials)
	if errGettingCredential != nil {
		lib.ErrorWriter(&response, "Something went wrong! make sure you fulfill all required fields!", http.StatusBadRequest)
		lib.ResponseFormatter(w, response)
		return
	}

	validators := validator.Validators{}
	if errValidator := validators.ValidatorService(credentials); errValidator != nil {
		lib.ErrorWriter(&response, validators.GetValidatorErrors(errValidator), http.StatusBadRequest)
		lib.ResponseFormatter(w, response)
		return
	}

	var userData models.GetUser
	db, _ := r.Context().Value(models.DBInstanceKey).(lib.DB)
	var sqlService = service.SqlService(db.Instance)
	fields := lib.SlicerDBFieldsName(models.User{}, "u.", true)
	var query = fmt.Sprintf(internals.QUERY_GETTING_USER, strings.Join(fields, ","), credentials.Identifiers, credentials.Identifiers, strings.Join(fields, ","))

	err := sqlService.SelectSingle(query, []interface{}{}, lib.SlicerReferenceFields(&userData, true)...)
	if err != nil {
		message, statuscode := lib.SqlError(err, []string{"email", "username"}, []string{})
		lib.ErrorWriter(&response, message, statuscode)
		lib.ResponseFormatter(w, response)
		return
	}

	if !lib.PasswordDecrypter(userData.Password, credentials.Password) {
		lib.ErrorWriter(&response, "*Incorrect password!", http.StatusBadRequest)
		lib.ResponseFormatter(w, response)
		return
	}

	userData.Password = ""
	token := tokenjwt.JwtToken{}
	token.Create(w, r, &response, userData)

	lib.ResponseFormatter(w, response)
}
