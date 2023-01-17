package api

type Book struct {
	ID           string        `json:"id"`
	ISBN         string        `json:"isbn"`
	Price        int           `json:"Price"`
	Author       *Author       `json:"author"`
	Availability *Availability `json:"availability"`
}

type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

type Availability struct {
	Available int `json:"available"`
	Booked    int `json:"booked"`
}
