package simpleprinter

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Home Page!")
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the About Page.")
}
