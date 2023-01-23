package operation

import (
	"bytes"
	"encoding/json"
	"github.com/avijit/config"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAddBook(t *testing.T) {
	var jsonStr = []byte(`{   
    "isbn": "isbnUpdated",
    "Price": 2999,
    "author": {"firstname": "Sourav", "lastname": "Ganguly"},
    "availability": {"available": 10,"booked": 0}
}`)
	req, _ := http.NewRequest("POST", "/config/add/books", bytes.NewBuffer(jsonStr))
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	_, err := Crud(router)
	if err != nil {
		t.Fatal("Error while routing")
	}
	router.HandleFunc("/config/add/books", addBooks)
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var book config.Book
	err = json.Unmarshal(rr.Body.Bytes(), &book)
	assert.Equal(t, "isbnUpdated", book.ISBN)
	assert.Equal(t, 10, book.Availability.Available)
	assert.Equal(t, "Sourav", book.Author.FirstName)
}
