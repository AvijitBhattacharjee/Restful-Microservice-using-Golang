package operation

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/avijit/config"
	"github.com/avijit/pkg/handler"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

var books []config.Book

func Crud(router *mux.Router) ([]config.Book, error) {

	// mock data
	books = append(books, config.Book{ID: "1", ISBN: "101", Price: 100, Author: &config.Author{
		FirstName: "Arijit", LastName: "Bhattacharjee"}, Availability: &config.Availability{
		Available: 2, Booked: 0,
	}})
	books = append(books, config.Book{ID: "2", ISBN: "102", Price: 200, Author: &config.Author{
		FirstName: "Abhishek", LastName: "Banerjee"}, Availability: &config.Availability{
		Available: 0, Booked: 2,
	}})
	books = append(books, config.Book{ID: "3", ISBN: "103", Price: 300, Author: &config.Author{
		FirstName: "Arijit", LastName: "Mukherjee"}, Availability: &config.Availability{
		Available: 10, Booked: 0,
	}})
	books = append(books, config.Book{ID: "4", ISBN: "104", Price: 400, Author: &config.Author{
		FirstName: "Abhishek", LastName: "Ganguly"}, Availability: &config.Availability{
		Available: 10, Booked: 0,
	}})

	books = append(books, config.Book{ID: "5", ISBN: "1104", Price: 7400, Author: &config.Author{
		FirstName: "Anindya", LastName: "Shankar"}, Availability: &config.Availability{
		Available: 8, Booked: 2,
	}})

	router.HandleFunc("/config/books", getBooks).Methods("GET")
	router.HandleFunc("/config/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/config/books/author/{id}", getBookByAuthor).Methods("GET")
	router.HandleFunc("/config/books/available/{id}", getAvailableBook).Methods("GET")
	router.HandleFunc("/config/add/books", addBooks).Methods("POST")
	router.HandleFunc("/config/update/books/{id}", updateBooks).Methods("PUT")
	router.HandleFunc("/config/reserve/books/{id}", reserveBooks).Methods("GET")
	router.HandleFunc("/config/release/books/{id}", releaseBooks).Methods("GET")
	router.HandleFunc("/config/damage/books/{id}", damageBook).Methods("DELETE")
	router.HandleFunc("/config/download/books", downloadBook).Methods("GET")

	return books, nil
}

func downloadBook(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("http://localhost:8080/config/books/author/Arijit")
	if err != nil {
		handler.RespondWithError(w, http.StatusInternalServerError, "Somehow host could not be reached.")
		return
	}
	data, _ := ioutil.ReadAll(response.Body)
	var books []config.Book
	err = json.Unmarshal(data, &books)
	if err != nil {
		handler.RespondWithError(w, http.StatusInternalServerError, "Unable to unmarshal data.")
		return
	}
	fmt.Println(books)
	b := &bytes.Buffer{}
	t := time.Now().Unix()
	fileName := "address-book-" + strconv.FormatInt(t, 10) + ".csv"
	writer := csv.NewWriter(b)
	heading := []string{"id", "first_name", "last_name", "email_address", "phone_number"}
	writer.Write(heading)
	for _, eachEntry := range books {
		var record []string
		record = append(record, eachEntry.ID)
		record = append(record, eachEntry.ISBN)
		record = append(record, eachEntry.Author.LastName)
		record = append(record, eachEntry.Author.FirstName)
		writer.Write(record)
	}
	writer.Flush()
	w.Header().Set("Content-Type", "text/csv") // setting the content type header to text/csv
	w.Header().Set("Content-Disposition", "attachment;filename="+fileName)
	w.WriteHeader(http.StatusOK)
	w.Write(b.Bytes())
	return
}
