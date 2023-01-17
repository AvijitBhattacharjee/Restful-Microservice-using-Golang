package controller

import "github.com/avijit/api"

func ReserveBook(params map[string]string, book *api.Book) (api.Book, string) {
	book.ID = params["id"]
	if book.Availability.Available <= 0 {
		return api.Book{}, api.NoReserve
	}
	book.Availability.Available--
	book.Availability.Booked++
	return *book, ""
}

func ReleaseBook(params map[string]string, book *api.Book) (api.Book, string) {
	book.ID = params["id"]
	if book.Availability.Booked <= 0 {
		return api.Book{}, api.NoRelease
	}
	book.Availability.Available++
	book.Availability.Booked--
	return *book, ""
}
