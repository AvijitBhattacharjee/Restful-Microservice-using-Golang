package operation

import (
	"encoding/json"
	"github.com/avijit/config"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAvailabilityBookWithSuccess(t *testing.T) {
	req, _ := http.NewRequest("GET", "/config/books/available/9", nil)
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	_, err := Crud(router)
	if err != nil {
		t.Fatal("Error while routing")
	}
	//router.HandleFunc("/config/books", getBooks)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var books []config.Book
	err = json.Unmarshal(rr.Body.Bytes(), &books)
	for index, _ := range books {
		assert.GreaterOrEqual(t, books[index].Availability.Available, 9)
	}

}

func TestGetAllBookWithSuccess(t *testing.T) {
	req, _ := http.NewRequest("GET", "/config/books", nil)
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	_, err := Crud(router)
	if err != nil {
		t.Fatal("Error while routing")
	}
	//router.HandleFunc("/config/books", getBooks)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var books []config.Book
	err = json.Unmarshal(rr.Body.Bytes(), &books)
	assert.NotEmpty(t, books)

}

func TestGetSingleBookWithSuccess(t *testing.T) {
	req, _ := http.NewRequest("GET", "/config/books/2", nil)
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	_, err := Crud(router)
	if err != nil {
		t.Fatal("Error while routing")
	}
	//router.HandleFunc("/config/books/{id}", getBook)
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
	_, err := Crud(router)
	if err != nil {
		t.Fatal("Error while routing")
	}
	//router.HandleFunc("/config/books/{id}", getBook)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var book config.Book
	err = json.Unmarshal(rr.Body.Bytes(), &book)
	assert.Equal(t, "", book.ID)

}

func TestGetAuthorBookWithSuccess(t *testing.T) {
	req, _ := http.NewRequest("GET", "/config/books/author/Arijit", nil)
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	_, err := Crud(router)
	if err != nil {
		t.Fatal("Error while routing")
	}
	router.HandleFunc("/config/books/author/{id}", getBooks)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var books []config.Book
	err = json.Unmarshal(rr.Body.Bytes(), &books)
	for index, _ := range books {
		assert.Equal(t, "Arijit", books[index].Author.FirstName)
	}
}
