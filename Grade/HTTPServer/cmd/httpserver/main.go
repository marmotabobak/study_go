package main

import (
	"httpserver/internal/handlers/simpleprinter"
	"log"
	"net/http"
)

func main() {
	mux := simpleprinter.NewSimplePrinterServeMux()
	log.Fatal(http.ListenAndServe(":8080", mux))
}