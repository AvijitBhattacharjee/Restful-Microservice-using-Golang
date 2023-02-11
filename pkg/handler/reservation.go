package handler

import (
	"fmt"
	"github.com/avijit/config"
)

func ReserveBook(params map[string]string, book *config.Book) error {
	book.ID = params["id"]
	if book.Availability.Available <= 0 {
		return fmt.Errorf(config.NoReserve)
	}
	book.Availability.Available--
	book.Availability.Booked++
	return nil
}

func ReleaseBook(params map[string]string, book *config.Book) error {
	book.ID = params["id"]
	if book.Availability.Booked <= 0 {
		return fmt.Errorf(config.NoRelease)
	}
	book.Availability.Available++
	book.Availability.Booked--
	return nil
}
