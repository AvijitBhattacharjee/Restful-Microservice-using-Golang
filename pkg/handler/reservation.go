package handler

import "github.com/avijit/config"

func ReserveBook(params map[string]string, book *config.Book) (config.Book, string) {
	book.ID = params["id"]
	if book.Availability.Available <= 0 {
		return config.Book{}, config.NoReserve
	}
	book.Availability.Available--
	book.Availability.Booked++
	return *book, config.Reserved
}

func ReleaseBook(params map[string]string, book *config.Book) (config.Book, string) {
	book.ID = params["id"]
	if book.Availability.Booked <= 0 {
		return config.Book{}, config.NoRelease
	}
	book.Availability.Available++
	book.Availability.Booked--
	return *book, config.Released
}
