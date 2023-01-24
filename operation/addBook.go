package operation

import (
	"encoding/json"
	"github.com/avijit/config"
	"github.com/avijit/pkg/handler"
	"math/rand"
	"net/http"
	"strconv"
)

func addBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(config.ContentType, config.AppJsonContentType)
	var book config.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(1000))
	book, err := handler.ValidateBook(&book)
	if err != "book input is valid" {
		handler.RespondWithError(w, http.StatusNonAuthoritativeInfo, config.EncodingError)
		return
	}
	books = append(books, book)
	err1 := json.NewEncoder(w).Encode(book)
	if err1 != nil {
		handler.RespondWithError(w, http.StatusBadRequest, config.EncodingError)
		return
	}
	handler.RespondWithJSON(w, http.StatusOK, config.CreateBook+config.Success)
}
