package anotherprinter

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_NewSimplePrinterServeMux(t *testing.T) {
	mux := NewAnotherPrinterMux()

	t.Run("/hello", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/hello", nil)

		mux.ServeHTTP(w, r)

		if w.Result().StatusCode == http.StatusNotFound {
			t.Error("endpoint /hello not found in mux")
		}
	})

	t.Run("/api", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/api", nil)

		mux.ServeHTTP(w, r)

		if w.Result().StatusCode == http.StatusNotFound {
			t.Error("endpoint /api not found in mux")
		}
	})

	t.Run("/version", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/version", nil)

		mux.ServeHTTP(w, r)

		if w.Result().StatusCode == http.StatusNotFound {
			t.Error("endpoint /version not found in mux")
		}
	})

	t.Run("/randomnumber", func(t *testing.T) {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/randomnumber", nil)

		mux.ServeHTTP(w, r)

		if w.Result().StatusCode == http.StatusNotFound {
			t.Error("endpoint /randomnumber not found in mux")
		}
	})

}