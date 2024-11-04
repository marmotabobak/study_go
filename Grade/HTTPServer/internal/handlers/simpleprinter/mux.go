package simpleprinter

import (
	"net/http"
)

func NewSimplePrinterServeMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/home", homePage)
	mux.HandleFunc("/about", aboutPage)

	return mux
}
