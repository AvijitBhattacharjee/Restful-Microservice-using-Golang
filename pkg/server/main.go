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
	books, err := operation.Crud(router)
	if err != nil {
		_ = fmt.Errorf("error while handling endpoints")
	}
	if books == nil {
		log.Fatal("no books are there")
	}
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
