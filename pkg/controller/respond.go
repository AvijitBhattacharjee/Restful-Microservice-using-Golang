package controller

import (
	"fmt"
	"github.com/avijit/api"
	"io"
	"net/http"
)

func RespondWithSuccess(w http.ResponseWriter, requestType string) {
	w.Header().Set(api.ContentType, api.AppJsonContentType)
	w.WriteHeader(api.StatusOK)
	_, err := io.WriteString(w, requestType+api.Success)
	if err != nil {
		_ = fmt.Errorf("error while writing")
	}
}

func RespondWithError(w http.ResponseWriter, errorType string) {
	w.Header().Set(api.ContentType, api.AppJsonContentType)
	_ = fmt.Errorf("request validation failed")
	w.WriteHeader(api.StatusBadRequest)
	_, err := io.WriteString(w, errorType)
	if err != nil {
		_ = fmt.Errorf("error while writing")
	}
}
