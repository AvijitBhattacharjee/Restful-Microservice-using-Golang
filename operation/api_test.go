package operation

import (
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCrud(t *testing.T) {

	routerTest := mux.NewRouter()
	err := Crud(routerTest)
	assert.Equal(t, nil, err)
}
