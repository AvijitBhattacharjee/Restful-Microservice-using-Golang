package method

import (
	"encoding/json"
	"github.com/avijit/api"
	"github.com/avijit/pkg/controller"
	"github.com/gorilla/mux"
	"net/http"
)

func getBookByAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(api.ContentType, api.AppJsonContentType)
	params := mux.Vars(r)

	var authorBook []api.Book

	for _, item := range books {
		if item.Author.FirstName == params["id"] || item.Author.LastName == params["id"] {
			authorBook = append(authorBook, item)
		}
	}
	err := json.NewEncoder(w).Encode(authorBook)
	if err != nil {
		controller.RespondWithError(w, api.EncodingError)
		return
	}
	controller.RespondWithSuccess(w, api.GetAuthorBook)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(api.ContentType, api.AppJsonContentType)
	params := mux.Vars(r)

	for _, item := range books {
		if item.ID == params["id"] {
			err := json.NewEncoder(w).Encode(item)
			if err != nil {
				controller.RespondWithError(w, api.EncodingError)
				return
			}
			return
		}
	}
	err := json.NewEncoder(w).Encode(&api.Book{})
	if err != nil {
		controller.RespondWithError(w, api.EncodingError)
		return
	}
	controller.RespondWithSuccess(w, api.GetBook)
}

func getBooks(w http.ResponseWriter, _ *http.Request) {

	w.Header().Set(api.ContentType, api.AppJsonContentType)
	err := json.NewEncoder(w).Encode(&books)
	if err != nil {
		controller.RespondWithError(w, api.EncodingError)
		return
	}
	controller.RespondWithSuccess(w, api.GetBooks)
}
