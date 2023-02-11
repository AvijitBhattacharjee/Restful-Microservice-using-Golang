package handler

import (
	"fmt"
	"github.com/avijit/config"
)

func ValidateBook(book *config.Book) error {
	if book == nil {
		return fmt.Errorf("nil book")
	}
	if book.Availability.Available <= 0 {
		return fmt.Errorf(config.NoAvailability)
	}
	if book.Price <= 0 || book.ISBN == "" {
		return fmt.Errorf(config.InvalidBook)
	}
	if book.Author.FirstName == "" || book.Author.LastName == "" {
		return fmt.Errorf(config.NoAuthor)
	}
	return nil
}
