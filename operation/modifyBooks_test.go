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

func TestReserveBookWithSuccess(t *testing.T) {
	req, _ := http.NewRequest("GET", "/config/reserve/books/1", nil)
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	_, err := Crud(router)
	if err != nil {
		t.Fatal("Error while routing")
	}
	router.HandleFunc("/config/reserve/books/{id}", reserveBooks)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var book config.Book
	err = json.Unmarshal(rr.Body.Bytes(), &book)
	assert.Equal(t, book.Availability.Available, 1)
	assert.Equal(t, book.Availability.Booked, 1)
}

func TestReserveBookWithError(t *testing.T) {
	req, _ := http.NewRequest("GET", "/config/reserve/books/2", nil)
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	_, err := Crud(router)
	if err != nil {
		t.Fatal("Error while routing")
	}
	router.HandleFunc("/config/reserve/books/{id}", reserveBooks)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)
	fmt.Println(rr.Result())

	var book config.Book
	err = json.Unmarshal(rr.Body.Bytes(), &book)
	fmt.Println(book.Availability)
	assert.Equal(t, config.EmptyString, book.ID)
}

func TestReleaseBookWithSuccess(t *testing.T) {
	req, _ := http.NewRequest("GET", "/config/reserve/books/3", nil)
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	_, err := Crud(router)
	if err != nil {
		t.Fatal("Error while routing")
	}
	router.HandleFunc("/config/reserve/books/{id}", releaseBooks)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var book config.Book
	err = json.Unmarshal(rr.Body.Bytes(), &book)
	if err != nil {
		t.Fatal("Error while unmarshalling")
	}
	assert.Equal(t, book.Availability.Available, 9)
	assert.Equal(t, book.Availability.Booked, 1)
}

func TestReleaseBookWithError(t *testing.T) {
	req, _ := http.NewRequest("GET", "/config/release/books/4", nil)
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	_, err := Crud(router)
	if err != nil {
		t.Fatal("Error while routing")
	}
	router.HandleFunc("/config/release/books/{id}", releaseBooks)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)

	var book config.Book
	err = json.Unmarshal(rr.Body.Bytes(), &book)
	assert.Equal(t, config.EmptyString, book.ID)
}
