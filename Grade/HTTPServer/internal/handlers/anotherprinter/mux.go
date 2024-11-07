package anotherprinter

import "net/http"

func NewAnotherPrinterMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", helloPage)
	mux.HandleFunc("/api", apiPage)
	mux.HandleFunc("/version", versionPage)
	mux.HandleFunc("/randomnumber", randomNumberPage)

	return mux
}
