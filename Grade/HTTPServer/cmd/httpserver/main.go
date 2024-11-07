package main

import (
	"httpserver/internal/handlers/anotherprinter"
	"log"
	"net/http"
)

func main() {
	mux := anotherprinter.NewAnotherPrinterMux()
	log.Fatal(http.ListenAndServe(":8080", mux))
}