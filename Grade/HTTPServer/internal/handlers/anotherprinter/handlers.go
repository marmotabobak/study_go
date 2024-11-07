package anotherprinter

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net/http"
)

func helloPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!\n")
}

func apiPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is \"/api\" page.\n")
}

func versionPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "vX.X.X_YYYYMMDDhhmmss\n")
}

func randomNumberPage(w http.ResponseWriter, r *http.Request) {
	randInt, _ := rand.Int(rand.Reader, big.NewInt(1000))
	fmt.Fprintf(w, "%s\n", randInt.String())
}