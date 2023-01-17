package method

import (
	"encoding/json"
	"fmt"
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
			books = append(books, ackBook(params, &book))
			fmt.Println(book.ID)
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

func reserveBook(w http.ResponseWriter, r *http.Request) {

}

func releaseBook(w http.ResponseWriter, r *http.Request) {

}
