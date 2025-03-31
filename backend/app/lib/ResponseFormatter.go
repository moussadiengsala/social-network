package lib

import (
	"encoding/json"
	"net/http"
)

// ResponseFormatter is a utility function for formatting HTTP responses in JSON format.
// It takes an http.ResponseWriter and a custom Response struct as parameters. The function
// sets the Content-Type header to "application/json", writes the provided HTTP status code
// to the response header, and encodes the Response struct as JSON to the response body.

func ResponseFormatter(w http.ResponseWriter,response Response){

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Code)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}