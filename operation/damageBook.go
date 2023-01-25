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
	var flag = false

	for index, item := range books {
		if item.ID == params["id"] {
			flag = true
			books = append(books[:index], books[index+1:]...)
			break
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
	handler.RespondWithJSON(w, http.StatusOK, config.DeleteBook)
}
