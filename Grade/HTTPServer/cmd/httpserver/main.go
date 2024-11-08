package main

import (
	"fmt"
	"log"
	"net/http"
	"httpserver/internal/mux"
)

var (
	Version = "v0.0.0_YYYYMMDDhhmmss"
)

func main() {
	fmt.Printf("Started with version: %s\n\n", Version)

	mux := mux.NewAnotherPrinterMux(Version)
	log.Fatal(http.ListenAndServe(":8080", mux))
}