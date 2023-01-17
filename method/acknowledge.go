package method

import (
	"github.com/avijit/api"
)

func reserveBook(params map[string]string, book *api.Book) api.Book {
	book.ID = params["id"]
	book.Availability.Available--
	book.Availability.Booked++
	return *book
}

func releaseBook(params map[string]string, book *api.Book) api.Book {
	book.ID = params["id"]
	book.Availability.Available++
	book.Availability.Booked--
	return *book
}
