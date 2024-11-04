package server_test

import (
	"httpserver/internal/handlers/simpleprinter"
	"net/http"
	"net/http/httptest"
	"io"
	"testing"
)

func TestHTTPServer(t *testing.T) {
	s := httptest.NewServer(simpleprinter.NewSimplePrinterServeMux())
	defer s.Close()

	type reqTable struct {
		path string
		result string
	}

	table := []reqTable{
		{path: "/home", result: "Welcome to the Home Page!\n"},
		{path: "/about", result: "This is the About Page.\n"},
	}

	for _, member := range table {
		t.Run(member.path, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, s.URL+member.path, nil)
			if err != nil {
				t.Fatalf("NewRequest unexpected error: %s", err.Error())
			}

			resp, err := http.DefaultClient.Do(req)
			defer resp.Body.Close()
			if err != nil {
				t.Fatalf("DefaultClient.Do unexpected error: %s", err.Error())
			}

			if resp.StatusCode != http.StatusOK {
				t.Errorf("Status code expevted 200, got %s", http.StatusText(resp.StatusCode))
			}

			buf, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("io.ReadAll error: %s", err.Error())
			}

			if string(buf) != member.result {
				t.Errorf("Response expected: \"%s\", got: %s", member.result, buf)
			}		
		})
	}
}