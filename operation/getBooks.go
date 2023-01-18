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
	w.Header().Set(config.ContentType, config.AppJsonContentType)
	params := mux.Vars(r)

	var authorBook []config.Book

	for _, item := range books {
		if item.Author.FirstName == params["id"] || item.Author.LastName == params["id"] {
			authorBook = append(authorBook, item)
		}
	}
	err := json.NewEncoder(w).Encode(authorBook)
	if err != nil {
		handler.RespondWithError(w, config.EncodingError)
		return
	}
	handler.RespondWithSuccess(w, config.GetAuthorBook)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(config.ContentType, config.AppJsonContentType)
	params := mux.Vars(r)

	for _, item := range books {
		if item.ID == params["id"] {
			err := json.NewEncoder(w).Encode(item)
			if err != nil {
				handler.RespondWithError(w, config.EncodingError)
				return
			}
			return
		}
	}
	err := json.NewEncoder(w).Encode(&config.Book{})
	if err != nil {
		handler.RespondWithError(w, config.EncodingError)
		return
	}
	handler.RespondWithSuccess(w, config.GetBook)
}

func getBooks(w http.ResponseWriter, _ *http.Request) {

	w.Header().Set(config.ContentType, config.AppJsonContentType)
	err := json.NewEncoder(w).Encode(&books)
	if err != nil {
		handler.RespondWithError(w, config.EncodingError)
		return
	}
	handler.RespondWithSuccess(w, config.GetBooks)
}

func getAvailableBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(config.ContentType, config.AppJsonContentType)
	params := mux.Vars(r)

	var authorBook []config.Book

	for _, item := range books {
		availability, err := strconv.Atoi(params["id"])
		if err != nil {
			handler.RespondWithError(w, config.EncodingError)
			return
		}
		if item.Availability.Available >= availability {
			authorBook = append(authorBook, item)
		}
	}
	err := json.NewEncoder(w).Encode(authorBook)
	if err != nil {
		handler.RespondWithError(w, config.InvalidRequest)
		return
	}
	handler.RespondWithSuccess(w, config.GetAvailableBook)
}