package lib

import "strconv"

func Convert(valueStr string) int {
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return 0
	}
	return value
}
