package operation

import (
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCrud(t *testing.T) {

	routerTest := mux.NewRouter()
	books, err := Crud(routerTest)
	assert.NotEmpty(t, books)
	assert.Equal(t, nil, err)
}
