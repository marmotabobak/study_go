package main

import (
	"log"
	"net/http"
	"REST/internal/app/mux"	
)

func main() {
	mux := mux.NewMux()
	log.Fatal(http.ListenAndServe(":8080", mux))
}
