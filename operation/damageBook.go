package operation

import (
	"encoding/json"
	"github.com/avijit/api"
	"github.com/avijit/pkg/controller"
	"github.com/gorilla/mux"
	"net/http"
)

func damageBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(api.ContentType, api.AppJsonContentType)
	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		controller.RespondWithError(w, api.EncodingError)
		return
	}
	controller.RespondWithSuccess(w, api.DeleteBook)
}
