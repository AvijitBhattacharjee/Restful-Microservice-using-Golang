package handler

import "github.com/avijit/config"

func ValidateBook(book *config.Book) (config.Book, string) {
	if book.Availability.Available <= 0 || book.Availability.Booked >= 1 {
		return config.Book{}, config.NoAvailability
	}
	if book.Price <= 0 || book.ISBN == "" {
		return config.Book{}, config.InvalidBook
	}
	if book.Author.FirstName == "" || book.Author.LastName == "" {
		return config.Book{}, config.NoAuthor
	}
	return *book, config.ValidBook
}
