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
