package auth

import (
	"fmt"
	"net/http"
	"strings"

	validator "learn.zone01dakar.sn/forum-rest-api/app/internals/core/validators"
	"learn.zone01dakar.sn/forum-rest-api/app/lib"
	"learn.zone01dakar.sn/forum-rest-api/app/models"
	service "learn.zone01dakar.sn/forum-rest-api/app/service/CRUD"
)

func (a Auth) Register(w http.ResponseWriter, r *http.Request) {

	response := lib.Response{Code: 200, Message: "ok"}
	var userData models.User
	if err := lib.ParseForm(&userData, r); err != nil {
		lib.ErrorWriter(&response, "Something went wrong! make sure you fulfill all required fields!", http.StatusBadRequest)
		lib.ResponseFormatter(w, response)
		return
	}

	validators := validator.Validators{}
	if errValidator := validators.ValidatorService(userData); len(errValidator) != 0 {
		lib.ErrorWriter(&response, validators.GetValidatorErrors(errValidator), http.StatusBadRequest)
		lib.ResponseFormatter(w, response)
		return
	}
	userData.Password = lib.PasswordEncrypter(userData.Password)

	db, _ := r.Context().Value(models.DBInstanceKey).(lib.DB)
	sqlService := service.SqlService(db.Instance)

	fields := lib.SlicerDBFieldsName(userData, "", false)
	var query = fmt.Sprintf("INSERT INTO User(%s) VALUES(%s?)", strings.Join(fields, ","), strings.Repeat("?,", len(fields)-1))

	if _, err := sqlService.Create(query, lib.Slicer(userData, false)...); err != nil {
		message, code := lib.SqlError(err, []string{"email", "username"}, []string{})
		lib.ErrorWriter(&response, message, code)
	}

	lib.ResponseFormatter(w, response)
}
