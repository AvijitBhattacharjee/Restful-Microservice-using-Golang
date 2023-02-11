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

func TestDamageBookWithSuccess(t *testing.T) {
	reqGet, _ := http.NewRequest("GET", "/config/books", nil)
	rrGet := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/config/damage/books/2", nil)
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	_, err := Crud(router)
	if err != nil {
		t.Fatal("Error while routing")
	}
	router.HandleFunc("/config/damage/books/{id}", damageBook)
	router.ServeHTTP(rrGet, reqGet)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var allBooks []config.Book
	err = json.Unmarshal(rrGet.Body.Bytes(), &allBooks)

	var books []config.Book
	err = json.Unmarshal(rr.Body.Bytes(), &books)
	assert.Equal(t, len(allBooks)-1, len(books))
}

func TestDamageBookWithError(t *testing.T) {
	req, _ := http.NewRequest("DELETE", "/config/damage/books/22", nil)
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	_, err := Crud(router)
	if err != nil {
		t.Fatal("Error while routing")
	}
	router.HandleFunc("/config/damage/books/{id}", damageBook)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)

	var book config.Book
	err = json.Unmarshal(rr.Body.Bytes(), &book)
	assert.Contains(t, rr.Body.String(), config.InvalidID)
}
