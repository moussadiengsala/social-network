package lib

import (
	"fmt"
	"html"
	"reflect"
)

func Slicer(data interface{}, isIDFieldNeeded bool) []interface{} {
	val := reflect.ValueOf(data)
	var result []interface{}
	for i := 0; i < val.NumField(); i++ {
		fieldName := val.Type().Field(i).Name
		fieldValue := val.Field(i).Interface()
		if !isIDFieldNeeded && fieldName == "ID" {
			continue
		}
		if strValue, ok := fieldValue.(string); ok {
			fieldValue = html.EscapeString(strValue)
		}
		result = append(result, fieldValue)
	}

	return result
}

func SlicerFieldsName(data interface{}, isIDFieldNeeded bool) []string {
	val := reflect.ValueOf(data)
	var result []string
	for i := 0; i < val.NumField(); i++ {
		fieldName := val.Type().Field(i).Name
		if !isIDFieldNeeded && fieldName == "ID" || fieldName == "" {
			continue
		}
		result = append(result, fieldName)
	}
	return result
}

func SlicerDBFieldsName(data interface{}, prefix string, isIDFieldNeeded bool) []string {
	val := reflect.ValueOf(data)
	typeOf := val.Type()

	var fields []string
	for i := 0; i < val.NumField(); i++ {
		fieldTag := typeOf.Field(i).Tag.Get("db")
		if !isIDFieldNeeded && fieldTag == "id" || fieldTag == "" || fieldTag == "-" {
			continue
		}
		fieldTag = fmt.Sprintf("%s%s", prefix, fieldTag)
		fields = append(fields, fieldTag)
	}

	return fields
}

func SlicerReferenceFields(data interface{}, isIDFieldNeeded bool) []interface{} {
	// Get the type of the struct
	valueType := reflect.TypeOf(data).Elem()

	// Get the value of the struct
	value := reflect.ValueOf(data).Elem()

	// Create a slice to store field values
	fieldValues := []interface{}{}

	// Iterate over the fields of the struct
	for i := 0; i < valueType.NumField(); i++ {
		fieldName := valueType.Field(i).Name

		if !isIDFieldNeeded && fieldName == "ID" {
			continue
		}
		// Get the value of the field
		fieldValue := value.Field(i)

		// Store the field value in the slice
		fieldValues = append(fieldValues, fieldValue.Addr().Interface())
	}

	return fieldValues
}
