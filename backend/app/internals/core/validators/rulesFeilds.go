package validators

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type RulesOfField struct {
	Field string
	Value reflect.Value
	Tag   string
	Error []string
}

func (r *RulesOfField) Required() bool {
	strRequired := r.Value.Kind() == reflect.String && len(cleanInput(r.Value.String())) == 0
	intRequired := r.Value.Kind() == reflect.Int && r.Value.Int() == 0
	if strRequired {
		r.Error = append(r.Error, fmt.Sprintf(ErrInvalidFieldString, r.Field))
		return true
	} else if intRequired {
		r.Error = append(r.Error, fmt.Sprintf(ErrInvalidFieldInt, r.Field))
		return true
	}
	return false
}

func (r *RulesOfField) ValidFieldsByRegex(pattern []string, Err string, arg ...interface{}) bool {
	// if r.Value.Kind() != reflect.String {
	// 	r.Error = append(r.Error, fmt.Sprintf("Invalid %s, check the type of this!", r.Field))
	// 	return true
	// }

	for _, p := range pattern {
		re := regexp.MustCompile(p)
		if r.Value.Kind() == reflect.String {
			value := r.Value.String()
			if !re.MatchString(value) {
				r.Error = append(r.Error, fmt.Sprintf(Err, arg...))
				return true
			}
		} else if r.Value.Kind() == reflect.Int {
			value := r.Value.Int()
			if value <= 0 {
				r.Error = append(r.Error, fmt.Sprintf(Err, arg...))
				return true
			}
		}
	}
	return false
}

func (r *RulesOfField) Min(value string) bool {
	minLen, err := strconv.Atoi(value)
	isThereAnError := false
	if err != nil {
		r.Error = append(r.Error, "Error getting the value of the tag min.")
		return true
	}

	minFieldStr := r.Value.Kind() == reflect.String && len(r.Value.String()) < minLen
	minFieldInt := r.Value.Kind() == reflect.Int && int(r.Value.Int()) < minLen

	if minFieldStr || minFieldInt {
		r.Error = append(r.Error, fmt.Sprintf(ErrMinSizeRequired, r.Field, minLen))
		isThereAnError = true
	}
	return isThereAnError
}

func (r *RulesOfField) Max(value string) bool {
	maxLen, err := strconv.Atoi(value)
	isThereAnError := false

	if err != nil {
		r.Error = append(r.Error, "Error getting the value of the tag max")
		return true
	}

	maxFieldStr := r.Value.Kind() == reflect.String && len(r.Value.String()) > maxLen
	maxFieldInt := r.Value.Kind() == reflect.Int && int(r.Value.Int()) > maxLen

	if maxFieldStr || maxFieldInt {
		r.Error = append(r.Error, fmt.Sprintf(ErrMinSizeRequired, r.Field, maxLen))
		isThereAnError = true
	}
	return isThereAnError
}

func (r *RulesOfField) Values(value, tag string) bool {
	values := strings.Split(value, "|")
	isThereAnError := false
	if len(values) <= 0 {
		r.Error = append(r.Error, fmt.Sprintf("Error getting the value of the tag %s", tag))
		return true
	}

	if !Contains(values, r.Value.String()) {
		r.Error = append(r.Error, fmt.Sprintf(ErrInvalidAction, r.Field, r.Value.String()))
		isThereAnError = true
	}
	return isThereAnError
}

func (r *RulesOfField) Date() bool {
	if r.Value.Kind() != reflect.String {
		r.Error = append(r.Error, "Invalid type for the date, expected string")
		return true
	}

	value := r.Value.String()

	layout1 := "2006-01-02"
	_, err1 := time.Parse(layout1, value)

	layout2 := "2006-01-02T15:04"
	_, err2 := time.Parse(layout2, value)

	layout3 := "2006-01-02T15:04:02"
	_, err3 := time.Parse(layout3, value)

	if err1 != nil && err2 != nil && err3 != nil {
		r.Error = append(r.Error, ErrInvalidDateFormate)
		return true
	}

	return false
}

func (r *RulesOfField) RequiredIfPrivacyIsAlmostPrivate(tempStruct map[string]reflect.Value) bool {
	privacy, ok := tempStruct["Privacy"]
	isThereAnError := false

	if ok && privacy.Kind() == reflect.String && privacy.String() == "almost_private" {
		isThereAnError = r.Required()
		pattern := `^\d+(,\d+)*$`
		if !regexp.MustCompile(pattern).MatchString(r.Value.String()) {
			r.Error = append(r.Error, ErrInvalidSelectedUsers)
			isThereAnError = true
		}
	}
	return isThereAnError
}
