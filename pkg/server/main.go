package main

import (
	"fmt"
	"github.com/avijit/method"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	// handlers endpoints
	err := method.Crud(router)
	if err != nil {
		fmt.Errorf("error while handling endpoints")
	}
	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
