package handler

import (
	"encoding/json"
	"github.com/avijit/config"
	"net/http"
)

func RespondWithError(w http.ResponseWriter, code int, message string) {
	RespondWithJSON(w, code, map[string]string{"error": message})
}

// RespondWithJSON Called for responses to encode and send json data
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	//encode payload to json
	response, _ := json.Marshal(payload)

	// set headers and write response
	w.Header().Set(config.ContentType, config.AppJsonContentType)
	w.WriteHeader(code)
	w.Write(response)
}
