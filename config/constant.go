package config

const (
	StatusOK           = 200
	StatusBadRequest   = 400
	ContentType        = "Content-Type"
	AppJsonContentType = "application/json"
	Success            = " request is successful"
	GetBook            = "Getting a book details"
	GetBooks           = "Getting all book details"
	UpdateBook         = "Updating one book"
	DeleteBook         = "Deleting one book"
	CreateBook         = "Creating one book"
	GetAuthorBook      = "Getting an Author's Book"
	GetAvailableBook   = "Getting available books"
	EncodingError      = "Error while encoding"
	InvalidRequest     = "Requested number is not valid"
	ReleaseBook        = "Releasing of this Book"
	ReserveBook        = "Reserving of this Book"
	Reserved           = "Book got reserved"
	Released           = "Book got released"
	NoRelease          = "this book is in full stock, can not release"
	NoReserve          = "This book is not available for booking"
)
