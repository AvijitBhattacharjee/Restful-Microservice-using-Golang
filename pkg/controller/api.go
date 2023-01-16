package controller

import (
	"encoding/json"
	"fmt"
	"github.com/avijit/api"
	"io"
	"net/http"
)

func RespondWithSuccess(w http.ResponseWriter) {
	w.Header().Set(api.ContentType, api.AppJsonContentType)
	w.WriteHeader(api.StatusOK)
	_, err := io.WriteString(w, api.Success)
	if err != nil {
		_ = fmt.Errorf("error while writing")
	}
}

func RespondWithError(w http.ResponseWriter, err error) {
	w.Header().Set(api.ContentType, api.AppJsonContentType)
	_ = fmt.Errorf("request validation failed")
	w.WriteHeader(api.StatusBadRequest)
	message, _ := json.Marshal(err)
	w.Write(message)
}
