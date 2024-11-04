package simpleprinter

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_NewSimplePrinterServeMux(t *testing.T) {
	mux := NewSimplePrinterServeMux()

	t.Run("/home", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/home", nil)

		mux.ServeHTTP(w, r)

		if w.Result().StatusCode == http.StatusNotFound {
			t.Errorf("endpoint %s not found in mux", w.Result().Request.URL)
		}
	})

	t.Run("/about", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/about", nil)

		mux.ServeHTTP(w, r)

		if w.Result().StatusCode == http.StatusNotFound {
			t.Errorf("endpoint %s not found in mux", w.Result().Request.URL)
		}
	})

}