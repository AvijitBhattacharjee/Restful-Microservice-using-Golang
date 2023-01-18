package operation

import (
	"github.com/avijit/config"
	"github.com/gorilla/mux"
)

var books []config.Book

func Crud(router *mux.Router) error {
	// mock data
	books = append(books, config.Book{ID: "1", ISBN: "101", Price: 100, Author: &config.Author{
		FirstName: "Arijit", LastName: "Bhattacharjee"}, Availability: &config.Availability{
		Available: 2, Booked: 0,
	}})
	books = append(books, config.Book{ID: "2", ISBN: "102", Price: 200, Author: &config.Author{
		FirstName: "Abhishek", LastName: "Banerjee"}, Availability: &config.Availability{
		Available: 8, Booked: 0,
	}})
	books = append(books, config.Book{ID: "3", ISBN: "103", Price: 300, Author: &config.Author{
		FirstName: "Arijit", LastName: "Mukherjee"}, Availability: &config.Availability{
		Available: 10, Booked: 0,
	}})
	books = append(books, config.Book{ID: "4", ISBN: "104", Price: 400, Author: &config.Author{
		FirstName: "Abhishek", LastName: "Ganguly"}, Availability: &config.Availability{
		Available: 10, Booked: 0,
	}})

	router.HandleFunc("/config/books", getBooks).Methods("GET")
	router.HandleFunc("/config/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/config/books/author/{id}", getBookByAuthor).Methods("GET")
	router.HandleFunc("/config/books/available/{id}", getAvailableBook).Methods("GET")
	router.HandleFunc("/config/add/books", addBooks).Methods("POST")
	router.HandleFunc("/config/update/books/{id}", updateBooks).Methods("PUT")
	router.HandleFunc("/config/reserve/books/{id}", reserveBooks).Methods("PUT")
	router.HandleFunc("/config/release/books/{id}", releaseBooks).Methods("PUT")
	router.HandleFunc("/config/damage/books/{id}", damageBook).Methods("DELETE")

	return nil
}
