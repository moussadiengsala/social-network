package validators

import (
	"regexp"
	"strings"
)

const (
	ErrInvalidFieldString   = "Fields %s length must be more than zero and must not contain only spaces, tabs, or newlines"
	ErrInvalidFieldInt      = "Fields %s value must be valid and greater than zero"
	ErrInvalidPassword      = "Fields password must contain at least one lowercase, one uppercase and one number."
	ErrMaxSizeExceeded      = "You have exceeded the maximum character size allowed for the %s field!"
	ErrInvalidUsername      = "the username must contain only letters, digits, dots, underscores, and dashes"
	ErrMinSizeRequired      = "The length or value of the field %s must be at least %d."
	ErrInvalidEmail         = "The email address is invalid!"
	ErrInvalidAction        = "Invalid value. The %s field cannot contain %s value."
	ErrInvalidIdentifier    = "Invalid username or email address!"
	ErrInvalidDateFormate   = "The date format must be YYYY-MM-DD or YYYY-MM-DDTHH:MM:SS"
	ErrInvalidSelectedUsers = "Invalid selected users format"
	ErrInvalidNumeric       = "This %s fields should be numeric!"
)

func cleanInput(input string) string {
	pattern := regexp.MustCompile(`\s+`)
	cleaned := strings.TrimSpace(input)
	cleaned = pattern.ReplaceAllString(cleaned, " ")
	return cleaned
}

func Contains(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}
