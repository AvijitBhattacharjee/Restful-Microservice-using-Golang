package operation

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
				controller.RespondWithError(w, api.EncodingError)
				return
			}
			return
		}
	}
	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		controller.RespondWithError(w, api.EncodingError)
		return
	}
	controller.RespondWithSuccess(w, api.UpdateBook)

}

func reserveBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(api.ContentType, api.AppJsonContentType)
	params := mux.Vars(r)

	for _, item := range books {
		if item.ID == params["id"] {
			item, err := controller.ReserveBook(params, &item)
			if err != "" {
				controller.RespondWithError(w, err)
				return
			}
			err1 := json.NewEncoder(w).Encode(item)
			if err1 != nil {
				controller.RespondWithError(w, api.EncodingError)
				return
			}
			return
		}
	}
	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		controller.RespondWithError(w, api.EncodingError)
		return
	}
	controller.RespondWithSuccess(w, api.ReserveBook)

}

func releaseBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(api.ContentType, api.AppJsonContentType)
	params := mux.Vars(r)

	for _, item := range books {
		if item.ID == params["id"] {
			item, err := controller.ReleaseBook(params, &item)
			if err != "" {
				controller.RespondWithError(w, err)
				return
			}
			err1 := json.NewEncoder(w).Encode(item)
			if err1 != nil {
				controller.RespondWithError(w, api.EncodingError)
				return
			}
			return
		}
	}
	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		controller.RespondWithError(w, api.EncodingError)
		return
	}
	controller.RespondWithSuccess(w, api.ReleaseBook)
}
