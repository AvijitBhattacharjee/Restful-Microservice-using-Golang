package operation

import (
	"encoding/json"
	"github.com/avijit/config"
	"github.com/avijit/pkg/handler"
	"github.com/gorilla/mux"
	"net/http"
)

func damageBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(config.ContentType, config.AppJsonContentType)
	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		handler.RespondWithError(w, config.EncodingError)
		return
	}
	handler.RespondWithSuccess(w, config.DeleteBook)
}
