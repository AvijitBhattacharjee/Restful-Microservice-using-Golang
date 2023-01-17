package method

import (
	"github.com/avijit/api"
)

func reserveBook(params map[string]string, book *api.Book) (api.Book, string) {
	book.ID = params["id"]
	if book.Availability.Available <= 0 {
		return api.Book{}, "this book is not available for booking"
	}
	book.Availability.Available--
	book.Availability.Booked++
	return *book, ""
}

func releaseBook(params map[string]string, book *api.Book) (api.Book, string) {
	book.ID = params["id"]
	if book.Availability.Booked <= 0 {
		return api.Book{}, "this book is in full stock"
	}
	book.Availability.Available++
	book.Availability.Booked--
	return *book, ""
}
