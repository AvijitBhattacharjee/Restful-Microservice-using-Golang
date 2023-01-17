package main

import (
	"fmt"
	"github.com/avijit/operation"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	// handlers endpoints
	err := operation.Crud(router)
	if err != nil {
		_ = fmt.Errorf("error while handling endpoints")
	}
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
