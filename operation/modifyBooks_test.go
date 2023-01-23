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

func TestReserveBookWithSuccess(t *testing.T) {
	req, _ := http.NewRequest("GET", "/config/reserve/books/1", nil)
	rr, _ := commonResponseModifyBook(t, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var book config.Book
	err := json.Unmarshal(rr.Body.Bytes(), &book)
	if err != nil {
		t.Fatal("failed to unmarshall data")
	}
	assert.Equal(t, book.Availability.Available, 1)
	assert.Equal(t, book.Availability.Booked, 1)
}

func TestReserveBookWithError(t *testing.T) {
	req, _ := http.NewRequest("GET", "/config/reserve/books/2", nil)
	rr, _ := commonResponseModifyBook(t, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)

	var book config.Book
	err := json.Unmarshal(rr.Body.Bytes(), &book)
	if err != nil {
		t.Fatal("failed to unmarshall data")
	}
	assert.Equal(t, config.EmptyString, book.ID)
}

func TestReleaseBookWithSuccess(t *testing.T) {
	req, _ := http.NewRequest("GET", "/config/reserve/books/3", nil)
	rr, _ := commonResponseModifyBook(t, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var book config.Book
	err := json.Unmarshal(rr.Body.Bytes(), &book)
	if err != nil {
		t.Fatal("Error while unmarshalling")
	}
	assert.Equal(t, book.Availability.Available, 9)
	assert.Equal(t, book.Availability.Booked, 1)
}

func TestReleaseBookWithError(t *testing.T) {
	req, _ := http.NewRequest("GET", "/config/release/books/4", nil)
	rr, _ := commonResponseModifyBook(t, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)

	var book config.Book
	err := json.Unmarshal(rr.Body.Bytes(), &book)
	if err != nil {
		t.Fatal("failed to unmarshall data")
	}
	assert.Equal(t, config.EmptyString, book.ID)
}

func TestUpdateBook(t *testing.T) {
	var jsonStr = []byte(`{   
    "isbn": "ISBN05",
    "Price": 6999,
    "author": {"firstname": "Ankit", "lastname": "Mondal"},
    "availability": {"available": 10,"booked": 0}
}`)
	req, _ := http.NewRequest("PUT", "/config/update/books/5", bytes.NewBuffer(jsonStr))
	rr, _ := commonResponseModifyBook(t, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var book config.Book
	err := json.Unmarshal(rr.Body.Bytes(), &book)
	if err != nil {
		t.Fatal("failed to unmarshall data")
	}
	assert.Equal(t, "ISBN05", book.ISBN)
	assert.Equal(t, 6999, book.Price)
	assert.Equal(t, 10, book.Availability.Available)
	assert.Equal(t, "Ankit", book.Author.FirstName)
}

func commonResponseModifyBook(t *testing.T, req *http.Request) (*httptest.ResponseRecorder, error) {
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	_, err := Crud(router)
	if err != nil {
		t.Fatal("Error while routing")
	}
	router.ServeHTTP(rr, req)
	return rr, err
}
