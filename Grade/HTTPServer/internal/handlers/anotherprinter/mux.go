package anotherprinter

import "net/http"

func NewAnotherPrinterMux(version string) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", helloPage)
	mux.HandleFunc("/api", apiPage)
	mux.HandleFunc("/version", NewVersionPage(version))
	mux.HandleFunc("/randomnumber", randomNumberPage)

	return mux
}
