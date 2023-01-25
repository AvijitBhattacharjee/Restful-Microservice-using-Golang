package handler

import (
	"github.com/avijit/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

var book1 = config.Book{
	ID:           "1",
	ISBN:         "isbn1",
	Price:        120,
	Author:       &config.Author{FirstName: "Avijit", LastName: "Bhattacharjee"},
	Availability: &config.Availability{Available: 0, Booked: 2},
}

var book2 = config.Book{
	ID:           "1",
	ISBN:         "",
	Price:        0,
	Author:       &config.Author{FirstName: "Avijit", LastName: "Bhattacharjee"},
	Availability: &config.Availability{Available: 10, Booked: 0},
}

var book3 = config.Book{
	ID:           "1",
	ISBN:         "isbn1",
	Price:        120,
	Author:       &config.Author{FirstName: "", LastName: ""},
	Availability: &config.Availability{Available: 10, Booked: 0},
}

var book4 = config.Book{
	ID:           "1",
	ISBN:         "isbn1",
	Price:        120,
	Author:       &config.Author{FirstName: "Avijit", LastName: "Bhattacharjee"},
	Availability: &config.Availability{Available: 10, Booked: 0},
}

func TestValidateBook(t *testing.T) {
	tests := []struct {
		name           string
		inputPayload   config.Book
		expectedOutput string
	}{
		{
			name:           "Book with zero availability",
			inputPayload:   book1,
			expectedOutput: config.NoAvailability,
		},
		{
			name:           "Book with empty ISBN and price 0",
			inputPayload:   book2,
			expectedOutput: config.InvalidBook,
		},
		{
			name:           "Book with empty author name",
			inputPayload:   book3,
			expectedOutput: config.NoAuthor,
		},
		{
			name:           "Valid Book Input",
			inputPayload:   book4,
			expectedOutput: config.ValidBook,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(*testing.T) {
			_, err := ValidateBook(&tt.inputPayload)
			assert.Equal(t, tt.expectedOutput, err)
		})
	}
}
