package anotherprinter

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_helloPage(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/hello", nil)

	HelloPage(w, r)

	res := w.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Status code expected 200, got %s", http.StatusText(res.StatusCode))
	}

	buf, err := io.ReadAll(res.Body)
	if err != nil {
		t.Errorf("io.ReadAll error: %s", err.Error())
	}

	if string(buf) != "Hello!\n" {
		t.Errorf("Response expected: \"Hello!\", got: %s", buf)
	}
}
