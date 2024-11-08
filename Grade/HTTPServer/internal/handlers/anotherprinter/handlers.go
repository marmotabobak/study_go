package anotherprinter

import (
	"fmt"
	"net/http"
	"httpserver/internal/randomprovider"
)

func HelloPage(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Hello!\n")
}

func ApiPage(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "This is \"/api\" page.\n")
}

// func NewVersionPage(version string) func(w http.ResponseWriter, r *http.Request) {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "%s\n", version)
// 	}
// }

func RandomNumberPage(w http.ResponseWriter, _ *http.Request) {
	rp := randomprovider.NewRandomNumberProvider()
	fmt.Fprintf(w, "%d\n", rp.GetRandomInt())
}