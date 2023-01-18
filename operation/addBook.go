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
	books = append(books, book)
	err := json.NewEncoder(w).Encode(book)
	if err != nil {
		handler.RespondWithError(w, config.EncodingError)
		return
	}
	handler.RespondWithSuccess(w, config.CreateBook)
}