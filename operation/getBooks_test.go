package operation

import (
	"encoding/json"
	"fmt"
	"github.com/avijit/config"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetSingleBookWithSuccess(t *testing.T) {
	req, _ := http.NewRequest("GET", "/config/books/2", nil)
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	err := Crud(router)
	if err != nil {
		t.Fatal("Error while routing")
	}
	router.HandleFunc("/config/books/{id}", getBook)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var book config.Book
	err = json.Unmarshal(rr.Body.Bytes(), &book)
	assert.Nil(t, err)
	assert.Equal(t, 200, book.Price)
	assert.Equal(t, "102", book.ISBN)
}

func TestGetSingleBookWithWrongID(t *testing.T) {
	req, _ := http.NewRequest("GET", "/config/books/21", nil)
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	err := Crud(router)
	if err != nil {
		t.Fatal("Error while routing")
	}
	router.HandleFunc("/config/books/{id}", getBook)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var book config.Book
	err = json.Unmarshal(rr.Body.Bytes(), &book)
	assert.NotNil(t, err)
	assert.Equal(t, "", book.ID)

}

func TestGetAllBookWithSuccess(t *testing.T) {
	req, _ := http.NewRequest("GET", "/config/books", nil)
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	err := Crud(router)
	if err != nil {
		t.Fatal("Error while routing")
	}
	router.HandleFunc("/config/books", getBooks)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var books []config.Book
	err = json.Unmarshal(rr.Body.Bytes(), &books)
	fmt.Println(rr.Body.String())
	var expected = `[{"id":"1","isbn":"101","Price":100,"author":{"firstname":"Arijit","lastname":"Bhattacharjee"},"availability":{"available":2,"booked":0}},{"id":"2","isbn":"102","Price":200,"author":{"firstname":"Abhishek","lastname":"Banerjee"},"availability":{"available":8,"booked":0}},{"id":"3","isbn":"103","Price":300,"author":{"firstname":"Arijit","lastname":"Mukherjee"},"availability":{"available":10,"booked":0}},{"id":"4","isbn":"104","Price":400,"author":{"firstname":"Abhishek","lastname":"Ganguly"},"availability":{"available":10,"booked":0}}]
					"Getting all book details request is successful"`
	assert.Equal(t, expected, rr.Body.String())
	//assert.Equal(t, 200, book.Price)
	//assert.Equal(t, "102", book.ISBN)
}
