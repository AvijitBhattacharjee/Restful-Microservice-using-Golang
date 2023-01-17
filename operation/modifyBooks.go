package operation

import (
	"encoding/json"
	"github.com/avijit/config"
	"github.com/avijit/pkg/handler"
	"github.com/gorilla/mux"
	"net/http"
)

func updateBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(config.ContentType, config.AppJsonContentType)
	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book config.Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			err := json.NewEncoder(w).Encode(book)
			if err != nil {
				handler.RespondWithError(w, config.EncodingError)
				return
			}
			return
		}
	}
	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		handler.RespondWithError(w, config.EncodingError)
		return
	}
	handler.RespondWithSuccess(w, config.UpdateBook)

}

func reserveBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(config.ContentType, config.AppJsonContentType)
	params := mux.Vars(r)

	for _, item := range books {
		if item.ID == params["id"] {
			item, err := handler.ReserveBook(params, &item)
			if err != "" {
				handler.RespondWithError(w, err)
				return
			}
			err1 := json.NewEncoder(w).Encode(item)
			if err1 != nil {
				handler.RespondWithError(w, config.EncodingError)
				return
			}
			return
		}
	}
	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		handler.RespondWithError(w, config.EncodingError)
		return
	}
	handler.RespondWithSuccess(w, config.ReserveBook)

}

func releaseBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(config.ContentType, config.AppJsonContentType)
	params := mux.Vars(r)

	for _, item := range books {
		if item.ID == params["id"] {
			item, err := handler.ReleaseBook(params, &item)
			if err != "" {
				handler.RespondWithError(w, err)
				return
			}
			err1 := json.NewEncoder(w).Encode(item)
			if err1 != nil {
				handler.RespondWithError(w, config.EncodingError)
				return
			}
			return
		}
	}
	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		handler.RespondWithError(w, config.EncodingError)
		return
	}
	handler.RespondWithSuccess(w, config.ReleaseBook)
}
