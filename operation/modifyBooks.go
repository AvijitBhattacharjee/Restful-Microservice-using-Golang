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
	var flag = false

	for index, item := range books {
		if item.ID == params["id"] {
			flag = true
			books = append(books[:index], books[index+1:]...)
			var book config.Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			err := handler.ValidateBook(&book)
			if err != nil {
				handler.RespondWithError(w, http.StatusNonAuthoritativeInfo, err.Error())
				return
			}
			books = append(books, book)
			err1 := json.NewEncoder(w).Encode(book)
			if err1 != nil {
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
	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		handler.RespondWithError(w, http.StatusBadRequest, config.EncodingError)
		return
	}
	handler.RespondWithJSON(w, http.StatusOK, config.UpdateBook+config.Success)

}

func reserveBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(config.ContentType, config.AppJsonContentType)
	params := mux.Vars(r)
	var flag = false

	for _, item := range books {
		if item.ID == params["id"] {
			flag = true
			err := handler.ReserveBook(params, &item)
			if err != nil {
				handler.RespondWithError(w, http.StatusBadRequest, err.Error())
				return
			}
			err1 := json.NewEncoder(w).Encode(item)
			if err1 != nil {
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
	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		handler.RespondWithError(w, http.StatusBadRequest, config.EncodingError)
		return
	}
	handler.RespondWithJSON(w, http.StatusOK, config.ReserveBook+config.Success)

}

func releaseBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(config.ContentType, config.AppJsonContentType)
	params := mux.Vars(r)
	var flag = false

	for _, item := range books {
		if item.ID == params["id"] {
			flag = true
			err := handler.ReleaseBook(params, &item)
			if err != nil {
				handler.RespondWithError(w, http.StatusBadRequest, err.Error())
				return
			}
			err1 := json.NewEncoder(w).Encode(item)
			if err1 != nil {
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
	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		handler.RespondWithError(w, http.StatusBadRequest, config.EncodingError)
		return
	}
	handler.RespondWithJSON(w, http.StatusOK, config.ReleaseBook+config.Success)
}
