package handler

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRespond(t *testing.T) {
	tests := []struct {
		name           string
		inputCode      int
		inputPayload   string
		wantErr        bool
		expectedStatus int
		expectedResult string
	}{
		{
			name:           "Testing Response With Success",
			inputCode:      200,
			inputPayload:   "Success Case",
			wantErr:        false,
			expectedStatus: http.StatusOK,
			expectedResult: "OK",
		},
		{
			name:           "Testing Response With Error",
			inputCode:      405,
			inputPayload:   "Error Case",
			wantErr:        true,
			expectedStatus: http.StatusMethodNotAllowed,
			expectedResult: "Method Not Allowed",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			if tt.wantErr == true {
				RespondWithError(rr, 405, tt.inputPayload)
				assert.Equal(t, true, strings.Contains(rr.Result().Status, tt.expectedResult))
				assert.Equal(t, tt.expectedStatus, rr.Code)
			} else {
				RespondWithJSON(rr, 200, tt.inputPayload)
				assert.Equal(t, true, strings.Contains(rr.Result().Status, tt.expectedResult))
				assert.Equal(t, tt.expectedStatus, rr.Code)
			}
		})
	}
}
