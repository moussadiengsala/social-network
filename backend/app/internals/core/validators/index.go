package validators

import (
	"fmt"
	"reflect"
	"strings"
)

type Validators struct{}

func (v Validators) ValidatorService(submitedData interface{}) []RulesOfField {
	var validationResults []RulesOfField

	tempStruct := map[string]reflect.Value{}

	// pattern
	patternPassword := []string{`(.*[a-z])`, `(.*[A-Z])`, `(.*\d)`}
	patternEmail := []string{`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`}
	patternUsername := []string{`^[a-zA-Z0-9_.-]+$`}

	valueOfData := reflect.ValueOf(submitedData)

	typeOf := valueOfData.Type()

	for i := 0; i < valueOfData.NumField(); i++ {

		fieldName := typeOf.Field(i).Name
		fieldValue := valueOfData.Field(i)

		tempStruct[fieldName] = fieldValue

		fieldTag := typeOf.Field(i).Tag.Get("validate")
		tags := strings.Split(fieldTag, ",")

		fieldValidator := RulesOfField{Field: fieldName, Value: fieldValue}

		for _, tag := range tags {
			tagParts := strings.Split(tag, "=")
			tagName := tagParts[0]
			tagValue := ""
			if len(tagParts) > 1 {
				tagValue = tagParts[1]
			}

			switch tagName {
			case "required":
				fieldValidator.Required()
			case "password":
				fieldValidator.ValidFieldsByRegex(patternPassword, ErrInvalidPassword)
			case "email":
				fieldValidator.ValidFieldsByRegex(patternEmail, ErrInvalidEmail)
			case "username":
				fieldValidator.ValidFieldsByRegex(patternUsername, ErrInvalidUsername)
			case "min":
				fieldValidator.Min(tagValue)
			case "max":
				fieldValidator.Max(tagValue)
			case "value":
				fieldValidator.Values(tagValue, tagName)
			case "date":
				fieldValidator.Date()
			case "identifiers":
				values := strings.Split(tagValue, "|")
				min := "4"
				max := "16"
				if len(values) == 2 {
					min = values[0]
					max = values[1]
				}

				f := fieldValidator
				isUsername := !f.ValidFieldsByRegex(patternUsername, ErrInvalidUsername) && !f.Min(min) && !f.Max(max)
				isEmail := !f.ValidFieldsByRegex(patternEmail, ErrInvalidEmail)
				if !isEmail && !isUsername {
					fieldValidator.Error = append(fieldValidator.Error, fmt.Sprint(ErrInvalidIdentifier))
				}
			case "selected_user":
				fieldValidator.RequiredIfPrivacyIsAlmostPrivate(tempStruct)
			case "numeric":
				fieldValidator.ValidFieldsByRegex([]string{`^\d+$`}, ErrInvalidNumeric, fieldName)
			}

		}

		if len(fieldValidator.Error) != 0 {
			validationResults = append(validationResults, fieldValidator)
		}
	}

	return validationResults
}

func (v Validators) GetValidatorErrors(validationResults []RulesOfField) string {
	info := []string{}
	for _, field := range validationResults {
		info = append(info, field.Error...)
	}
	return strings.Join(info, " * ")
}
