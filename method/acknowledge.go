package method

import "github.com/avijit/api"

func ackBook(params map[string]string, book *api.Book) api.Book {
	book.ID = params["id"]
	book.Price = 200
	book.ISBN = "booked"
	return *book
}

func filterBook(books *[]api.Book) []api.Book {
	var newBooks []api.Book
	for _, item := range *books {
		if item.ISBN != "booked" {
			newBooks = append(newBooks, item)
		}
	}
	return newBooks
}