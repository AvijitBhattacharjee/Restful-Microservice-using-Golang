package handler

import (
	"github.com/avijit/config"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRespondWithError(t *testing.T) {
	rr := httptest.NewRecorder()
	RespondWithError(rr, 405, "error message")
	assert.Equal(t, true, strings.Contains(rr.Result().Status, "Method Not Allowed"))
	assert.Equal(t, 405, rr.Code)

}

func TestRespondWithSuccess(t *testing.T) {
	rr := httptest.NewRecorder()
	RespondWithJSON(rr, 200, "Hi")
	assert.Equal(t, config.AppJsonContentType, rr.Header().Get(config.ContentType))
	assert.Equal(t, 200, rr.Code)
}
