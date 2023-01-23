package handler

import (
	"encoding/json"
	"fmt"
	"github.com/avijit/config"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, message string) {
	payload := map[string]string{"error": message}
	response, _ := json.Marshal(payload)

	// set headers and write response
	w.Header().Set(config.ContentType, config.AppJsonContentType)
	w.WriteHeader(code)
	fmt.Println(response)
	_, err := w.Write(response)
	if err != nil {
		return
	}
}

// RespondWithJSON Called for responses to encode and send json data
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	//encode payload to json
	response, _ := json.Marshal(payload)

	// set headers and write response
	w.Header().Set(config.ContentType, config.AppJsonContentType)
	w.WriteHeader(code)
	fmt.Println(response)
	//w.Write(response)
}
