package main

import (
	"log"
	"net/http"
	"REST/internal/mux"
)

func main() {
	mux := mux.NewMux()
	log.Fatal(http.ListenAndServe(":8080", mux))
}
