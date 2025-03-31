package lib

import (
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

func ParseForm(data interface{}, r *http.Request) error {
	value := reflect.ValueOf(data).Elem()

	for i := 0; i < value.NumField(); i++ {
		// Get the field and its tag
		field := value.Field(i)
		tag := value.Type().Field(i).Tag.Get("json")

		if tag == "file" || tag == "avatar" || tag == "cover" ||  tag == "photo" {
			f, err := FileUploader(r, tag)
			if err != nil {
				return err
			}
			field.SetString(f)
			continue
		}

		formValue := r.FormValue(tag)

		// Convert formValue to the appropriate type and set it to the field
		if field.CanSet() {
			switch field.Kind() {
			case reflect.String:
				field.SetString(formValue)
			case reflect.Int:
				formValue = strings.Trim(formValue, `"`)
				a, err := strconv.Atoi(formValue)
				if err != nil {
					continue
				}
				field.SetInt(int64(a))
			}
		}
	}
	return nil
}
