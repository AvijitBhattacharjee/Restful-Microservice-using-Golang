package operation

import (
	"github.com/avijit/api"
	"github.com/gorilla/mux"
)

var books []api.Book

func Crud(router *mux.Router) error {
	// mock data
	books = append(books, api.Book{ID: "1", ISBN: "101", Price: 100, Author: &api.Author{
		FirstName: "Arijit", LastName: "Bhattacharjee"}, Availability: &api.Availability{
		Available: 10, Booked: 0,
	}})
	books = append(books, api.Book{ID: "2", ISBN: "102", Price: 200, Author: &api.Author{
		FirstName: "Abhishek", LastName: "Banerjee"}, Availability: &api.Availability{
		Available: 8, Booked: 0,
	}})
	books = append(books, api.Book{ID: "3", ISBN: "103", Price: 300, Author: &api.Author{
		FirstName: "Arijit", LastName: "Mukherjee"}, Availability: &api.Availability{
		Available: 10, Booked: 0,
	}})
	books = append(books, api.Book{ID: "4", ISBN: "104", Price: 400, Author: &api.Author{
		FirstName: "Abhishek", LastName: "Ganguly"}, Availability: &api.Availability{
		Available: 10, Booked: 0,
	}})

	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/add/books", addBooks).Methods("POST")
	router.HandleFunc("/api/update/books/{id}", updateBooks).Methods("PUT")
	router.HandleFunc("/api/reserve/books/{id}", reserveBooks).Methods("PUT")
	router.HandleFunc("/api/release/books/{id}", releaseBooks).Methods("PUT")
	router.HandleFunc("/api/damage/books/{id}", damageBook).Methods("DELETE")
	router.HandleFunc("/api/books/author/{id}", getBookByAuthor).Methods("GET")

	return nil
}
