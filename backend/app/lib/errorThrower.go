package lib


func ErrorWriter(response *Response, message string, statusCode int) {
	response.Message = message
	response.Code = statusCode
}
