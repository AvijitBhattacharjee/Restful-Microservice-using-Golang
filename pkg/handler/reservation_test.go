package handler

import (
	"github.com/avijit/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

var paramMockSuccess = map[string]string{"id": "1"}
var bookMockFail = config.Book{Availability: &config.Availability{
	Available: 0,
	Booked:    0,
}}

func TestReserveBookWithSuccess(t *testing.T) {
	var bookMockSuccess = config.Book{Availability: &config.Availability{
		Available: 2,
		Booked:    2,
	}}
	bookResult, msg := ReserveBook(paramMockSuccess, &bookMockSuccess)
	assert.Equal(t, config.Reserved, msg)
	assert.Equal(t, 1, bookResult.Availability.Available)
	assert.Equal(t, 3, bookResult.Availability.Booked)
}

func TestReserveBookWithError(t *testing.T) {
	_, msg := ReserveBook(paramMockSuccess, &bookMockFail)
	assert.Equal(t, config.NoReserve, msg)
}

func TestReleaseBookWithSuccess(t *testing.T) {
	var bookMockSuccess = config.Book{Availability: &config.Availability{
		Available: 2,
		Booked:    2,
	}}
	bookResult, msg := ReleaseBook(paramMockSuccess, &bookMockSuccess)
	assert.Equal(t, config.Released, msg)
	assert.Equal(t, 3, bookResult.Availability.Available)
	assert.Equal(t, 1, bookResult.Availability.Booked)
}

func TestReleaseBookWithError(t *testing.T) {
	_, msg := ReleaseBook(paramMockSuccess, &bookMockFail)
	assert.Equal(t, config.NoRelease, msg)
}
