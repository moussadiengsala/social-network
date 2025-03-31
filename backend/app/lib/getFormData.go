package lib

import (
	"net/http"
)

func GetFormData(r *http.Request) map[string]interface{} {
	r.ParseForm()
	datas := make(map[string]interface{}, len(r.Form))

	for key, value := range r.Form {
		if len(value) > 1 {
			datas[key] = value
		} else {
			datas[key] = value[0]
		}
	}

	return datas
}
