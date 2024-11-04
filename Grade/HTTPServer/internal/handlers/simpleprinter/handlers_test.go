package simpleprinter

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_homePage(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/api", nil)

	homePage(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Status code expevted 200, got %s", http.StatusText(res.StatusCode))
	}

	buf, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("io.ReadAll error: %s", err.Error())
	}

	if string(buf) != "Welcome to the Home Page!\n" {
		t.Errorf("Response expected: \"Welcome to the Home Page!\", got: %s", buf)
	}
}

func Test_aboutPage(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/about", nil)

	aboutPage(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Status code expevted 200, got %s", http.StatusText(res.StatusCode))
	}

	buf, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("io.ReadAll error: %s", err.Error())
	}

	if string(buf) != "This is the About Page.\n" {
		t.Errorf("Response expected: \"This is the About Page.\", got: %s", buf)
	}
}