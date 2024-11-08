package mux

import (
	"net/http"
	"httpserver/internal/handlers/anotherprinter"
	"httpserver/internal/handlers/versionpage"
)

func NewAnotherPrinterMux(version string) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", anotherprinter.HelloPage)
	mux.HandleFunc("/api", anotherprinter.ApiPage)
	mux.Handle("/version", versionpage.NewVersionPage(version))
	mux.HandleFunc("/randomnumber", anotherprinter.RandomNumberPage)

	return mux
}
