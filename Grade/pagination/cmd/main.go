package main

import (
	"log"
	"net/http"
	"pagination/internal/http/mux"
)

func main() {
	mux := mux.NewBooksMux()
	log.Fatal(http.ListenAndServe(":8080", mux))
}