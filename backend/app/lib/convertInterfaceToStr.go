package lib

func ConvertToString(value interface{}) string {
	var stringValue string
	switch v := value.(type) {
	case string:
		stringValue = v
	case *string:
		if v != nil {
			stringValue = *v
		}
	}

	return stringValue
}
