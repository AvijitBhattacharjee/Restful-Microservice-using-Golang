package handler

import (
	"fmt"
	"github.com/avijit/config"
	"io"
	"net/http"
)

func RespondWithSuccess(w http.ResponseWriter, requestType string) {
	w.Header().Set(config.ContentType, config.AppJsonContentType)
	w.WriteHeader(config.StatusOK)
	_, err := io.WriteString(w, requestType+config.Success)
	if err != nil {
		_ = fmt.Errorf("error while writing")
	}
}

func RespondWithError(w http.ResponseWriter, errorType string) {
	w.Header().Set(config.ContentType, config.AppJsonContentType)
	_ = fmt.Errorf("request validation failed")
	w.WriteHeader(config.StatusBadRequest)
	_, err := io.WriteString(w, errorType)
	if err != nil {
		_ = fmt.Errorf("error while writing")
	}
}
