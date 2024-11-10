package main

import (
	"log"
	"net/http"
	"query/internal/calculate/calculatemux"
)


func main() {
	mux := calculatemux.NewCalculateMux()
	log.Fatal(http.ListenAndServe(":8080", mux))
}