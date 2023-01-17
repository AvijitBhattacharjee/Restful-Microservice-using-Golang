package method

import (
	"encoding/json"
	"github.com/avijit/api"
	"github.com/avijit/pkg/controller"
	"github.com/gorilla/mux"
	"net/http"
)

func updateBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(api.ContentType, api.AppJsonContentType)
	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book api.Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			err := json.NewEncoder(w).Encode(book)
			if err != nil {
				controller.RespondWithError(w, err)
				return
			}
			return
		}
	}
	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		controller.RespondWithError(w, err)
		return
	}
	controller.RespondWithSuccess(w, api.UpdateBook)

}

func reserveBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(api.ContentType, api.AppJsonContentType)
	params := mux.Vars(r)

	for _, item := range books {
		if item.ID == params["id"] {
			item = reserveBook(params, &item)
			err := json.NewEncoder(w).Encode(item)
			if err != nil {
				controller.RespondWithError(w, err)
				return
			}
			return
		}
	}
	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		controller.RespondWithError(w, err)
		return
	}
	controller.RespondWithSuccess(w, api.UpdateBook)

}

func releaseBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(api.ContentType, api.AppJsonContentType)
	params := mux.Vars(r)

	for _, item := range books {
		if item.ID == params["id"] {
			item = releaseBook(params, &item)
			err := json.NewEncoder(w).Encode(item)
			if err != nil {
				controller.RespondWithError(w, err)
				return
			}
			return
		}
	}
	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		controller.RespondWithError(w, err)
		return
	}
	controller.RespondWithSuccess(w, api.UpdateBook)
}
