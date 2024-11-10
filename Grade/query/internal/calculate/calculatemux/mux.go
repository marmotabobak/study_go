package calculatemux

import (
	"net/http"
	"query/internal/calculate/calculatehandlers"
)

func NewCalculateMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/calculate", calculatehandlers.CalculateHandler)
	return mux
}