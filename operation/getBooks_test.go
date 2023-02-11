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
	rr, _ := commonResponseGetBook(t, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var books []config.Book
	err := json.Unmarshal(rr.Body.Bytes(), &books)
	if err != nil {
		t.Fatal("Failed to unmarshall data")
	}
	for index, _ := range books {
		assert.GreaterOrEqual(t, books[index].Availability.Available, 9)
	}
}

func TestGetAvailabilityBookWithError(t *testing.T) {
	req, _ := http.NewRequest("GET", "/config/books/available/99", nil)
	rr, _ := commonResponseGetBook(t, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var book config.Book
	err := json.Unmarshal(rr.Body.Bytes(), &book)
	if err != nil {
		t.Fatal("Failed to unmarshall data")
	}
	assert.Equal(t, "", book.ID)
}

func TestGetAllBookWithSuccess(t *testing.T) {
	req, _ := http.NewRequest("GET", "/config/books", nil)
	rr, _ := commonResponseGetBook(t, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var books []config.Book
	err := json.Unmarshal(rr.Body.Bytes(), &books)
	if err != nil {
		t.Fatal("Failed to unmarshall data")
	}
	assert.NotEmpty(t, books)
}

func TestGetSingleBookWithSuccess(t *testing.T) {
	req, _ := http.NewRequest("GET", "/config/books/2", nil)
	rr, _ := commonResponseGetBook(t, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var book config.Book
	err := json.Unmarshal(rr.Body.Bytes(), &book)
	if err != nil {
		t.Fatal("Failed to unmarshall data")
	}
	assert.Nil(t, err)
	assert.Equal(t, 200, book.Price)
	assert.Equal(t, "102", book.ISBN)
}

func TestGetSingleBookWithWrongID(t *testing.T) {
	req, _ := http.NewRequest("GET", "/config/books/21", nil)
	rr, _ := commonResponseGetBook(t, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)

	var book config.Book
	err := json.Unmarshal(rr.Body.Bytes(), &book)
	if err != nil {
		t.Fatal("Failed to unmarshall data")
	}
	assert.Contains(t, rr.Body.String(), config.InvalidID)
	assert.Equal(t, config.EmptyString, book.ID)
}

func TestGetAuthorBookWithSuccess(t *testing.T) {
	req, _ := http.NewRequest("GET", "/config/books/author/Arijit", nil)
	rr, _ := commonResponseGetBook(t, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var books []config.Book
	err := json.Unmarshal(rr.Body.Bytes(), &books)
	if err != nil {
		t.Fatal("Failed to unmarshall data")
	}
	for index, _ := range books {
		assert.Equal(t, "Arijit", books[index].Author.FirstName)
	}
}

func TestGetAuthorBookWithError(t *testing.T) {
	req, _ := http.NewRequest("GET", "/config/books/author/Arijita", nil)
	rr, _ := commonResponseGetBook(t, req)

	assert.Equal(t, http.StatusBadRequest, rr.Code)

	var book config.Book
	err := json.Unmarshal(rr.Body.Bytes(), &book)
	if err != nil {
		t.Fatal("Failed to unmarshall data")
	}
	assert.Contains(t, rr.Body.String(), config.InvalidID)
	assert.Equal(t, config.EmptyString, book.ID)
}

func commonResponseGetBook(t *testing.T, req *http.Request) (*httptest.ResponseRecorder, error) {
	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	_, err := Crud(router)
	if err != nil {
		t.Fatal("Error while routing")
	}
	router.ServeHTTP(rr, req)
	return rr, err
}
