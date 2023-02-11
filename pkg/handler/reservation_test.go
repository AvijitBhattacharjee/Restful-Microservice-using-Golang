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
	err := ReserveBook(paramMockSuccess, &bookMockSuccess)
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, bookMockSuccess.Availability.Available)
	assert.Equal(t, 3, bookMockSuccess.Availability.Booked)
}

func TestReserveBookWithError(t *testing.T) {
	err := ReserveBook(paramMockSuccess, &bookMockFail)
	assert.Contains(t, err.Error(), config.NoReserve)
}

func TestReleaseBookWithSuccess(t *testing.T) {
	var bookMockSuccess = config.Book{Availability: &config.Availability{
		Available: 2,
		Booked:    2,
	}}
	err := ReleaseBook(paramMockSuccess, &bookMockSuccess)
	assert.Equal(t, nil, err)
	assert.Equal(t, 3, bookMockSuccess.Availability.Available)
	assert.Equal(t, 1, bookMockSuccess.Availability.Booked)
}

func TestReleaseBookWithError(t *testing.T) {
	err := ReleaseBook(paramMockSuccess, &bookMockFail)
	assert.Contains(t, err.Error(), config.NoRelease)
}
