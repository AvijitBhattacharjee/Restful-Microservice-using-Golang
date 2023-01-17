package method

import (
	"encoding/json"
	"github.com/avijit/api"
	"github.com/avijit/pkg/controller"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"strconv"
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

func addBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(api.ContentType, api.AppJsonContentType)
	var book api.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000))
	books = append(books, book)
	err := json.NewEncoder(w).Encode(book)
	if err != nil {
		controller.RespondWithError(w, api.EncodingError)
		return
	}
	controller.RespondWithSuccess(w, api.CreateBook)
}
