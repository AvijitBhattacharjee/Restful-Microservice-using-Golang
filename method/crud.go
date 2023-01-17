package method

import (
	"encoding/json"
	"fmt"
	"github.com/avijit/api"
	"github.com/avijit/pkg/controller"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"strconv"
)

var books []api.Book

func Crud(router *mux.Router) error {
	// mock data
	books = append(books, api.Book{ID: "1", ISBN: "101", Price: 100, Author: &api.Author{
		FirstName: "Arijit", LastName: "Bhattacharjee"}})
	books = append(books, api.Book{ID: "2", ISBN: "102", Price: 200, Author: &api.Author{
		FirstName: "Abhishek", LastName: "Banerjee"}})
	books = append(books, api.Book{ID: "1", ISBN: "103", Price: 300, Author: &api.Author{
		FirstName: "Arijit", LastName: "Mukherjee"}})
	books = append(books, api.Book{ID: "4", ISBN: "104", Price: 400, Author: &api.Author{
		FirstName: "Abhishek", LastName: "Ganguly"}})

	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBooks).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBooks).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBooks).Methods("DELETE")
	router.HandleFunc("/api/books/author/{id}", getBookByAuthor).Methods("GET")

	return nil
}

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
		controller.RespondWithError(w, err)
		return
	}
	controller.RespondWithSuccess(w, api.GetAuthorBook)
}

func getBooks(w http.ResponseWriter, _ *http.Request) {

	w.Header().Set(api.ContentType, api.AppJsonContentType)
	err := json.NewEncoder(w).Encode(filterBook(&books))
	if err != nil {
		controller.RespondWithError(w, err)
		return
	}
	controller.RespondWithSuccess(w, api.GetBooks)
}

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

func createBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(api.ContentType, api.AppJsonContentType)
	var book api.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000))
	books = append(books, book)
	err := json.NewEncoder(w).Encode(book)
	if err != nil {
		controller.RespondWithError(w, err)
		return
	}
	controller.RespondWithSuccess(w, api.CreateBook)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(api.ContentType, api.AppJsonContentType)
	params := mux.Vars(r)

	for _, item := range books {
		if item.ID == params["id"] {
			err := json.NewEncoder(w).Encode(item)
			if err != nil {
				controller.RespondWithError(w, err)
				return
			}
			return
		}
	}
	err := json.NewEncoder(w).Encode(&api.Book{})
	if err != nil {
		controller.RespondWithError(w, err)
		return
	}
	controller.RespondWithSuccess(w, api.GetBook)
}

func deleteBooks(w http.ResponseWriter, r *http.Request) {
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
		controller.RespondWithError(w, err)
		return
	}
	controller.RespondWithSuccess(w, api.DeleteBook)
}
