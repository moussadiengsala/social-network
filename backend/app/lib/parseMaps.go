package lib

import (
	"reflect"
	"strconv"
)

func ParseMap(data interface{}, dataMap map[string]interface{}) error {
	value := reflect.ValueOf(data).Elem()

	for i := 0; i < value.NumField(); i++ {
		// Get the field and its tag
		field := value.Field(i)
		tag := value.Type().Field(i).Tag.Get("json")

		// Get value from the map based on the JSON tag
		mapValue, ok := dataMap[tag]
		if !ok {
			continue // Skip if the key is not found in the map
		}

		// Convert mapValue to the appropriate type and set it to the field
		switch field.Kind() {
		case reflect.String:
			if strVal, ok := mapValue.(string); ok {
				field.SetString(strVal)
			}
		case reflect.Int:
			if intVal, ok := mapValue.(int); ok {
				field.SetInt(int64(intVal))
			} else if strVal, ok := mapValue.(string); ok {
				// Convert string to int
				intVal, err := strconv.Atoi(strVal)
				if err != nil {
					return err
				}
				field.SetInt(int64(intVal))
			}
		}
	}
	return nil
}
