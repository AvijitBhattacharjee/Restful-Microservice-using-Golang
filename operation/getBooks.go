package operation

import (
	"encoding/json"
	"github.com/avijit/config"
	"github.com/avijit/pkg/handler"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func getBookByAuthor(w http.ResponseWriter, r *http.Request) {
	params := configParams(w, r)
	var flag = false

	var authorBook []config.Book

	for _, item := range books {
		if item.Author.FirstName == params["id"] || item.Author.LastName == params["id"] {
			flag = true
			authorBook = append(authorBook, item)
		}
	}
	if flag == false {
		handler.RespondWithError(w, http.StatusBadRequest, config.InvalidID)
		return
	}
	err := json.NewEncoder(w).Encode(authorBook)
	if err != nil {
		handler.RespondWithError(w, http.StatusBadRequest, config.EncodingError)
		return
	}
	handler.RespondWithJSON(w, http.StatusOK, config.GetAuthorBook+config.Success)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	params := configParams(w, r)
	var flag = false

	for _, item := range books {
		if item.ID == params["id"] {
			flag = true
			err := json.NewEncoder(w).Encode(item)
			if err != nil {
				handler.RespondWithError(w, http.StatusBadRequest, config.EncodingError)
				return
			}
			return
		}
	}
	if flag == false {
		handler.RespondWithError(w, http.StatusBadRequest, config.InvalidID)
		return
	}
	err := json.NewEncoder(w).Encode(&config.Book{})
	if err != nil {
		handler.RespondWithError(w, http.StatusBadRequest, config.EncodingError)
		return
	}
	handler.RespondWithJSON(w, http.StatusOK, config.GetBook+config.Success)
}

func getBooks(w http.ResponseWriter, _ *http.Request) {

	w.Header().Set(config.ContentType, config.AppJsonContentType)
	err := json.NewEncoder(w).Encode(&books)
	if err != nil {
		handler.RespondWithError(w, http.StatusBadRequest, config.EncodingError)
		return
	}
	handler.RespondWithJSON(w, http.StatusOK, config.GetBooks+config.Success)
}

func getAvailableBook(w http.ResponseWriter, r *http.Request) {
	params := configParams(w, r)

	var authorBook []config.Book

	for _, item := range books {
		availability, err := strconv.Atoi(params["id"])
		if err != nil {
			handler.RespondWithError(w, http.StatusBadRequest, config.InvalidID)
			return
		}
		if availability <= 0 {
			handler.RespondWithError(w, http.StatusBadRequest, config.NoAvailability)
			return
		}
		if item.Availability.Available >= availability {
			authorBook = append(authorBook, item)
		}
	}
	err := json.NewEncoder(w).Encode(authorBook)
	if err != nil {
		handler.RespondWithError(w, http.StatusBadRequest, config.InvalidRequest)
		return
	}
	handler.RespondWithJSON(w, http.StatusOK, config.GetAvailableBook+config.Success)
}

func configParams(w http.ResponseWriter, r *http.Request) map[string]string {
	w.Header().Set(config.ContentType, config.AppJsonContentType)
	params := mux.Vars(r)
	return params
}
