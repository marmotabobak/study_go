package server_test

import (
	"httpserver/internal/handlers/anotherprinter"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

type TestComparisonType string

const (
	TestComparisonEqual TestComparisonType = "equal"
	TestComparisonIntType TestComparisonType = "intType"
)

func TestSimpleHTTPServer(t *testing.T) {
	version := "someVersion"
	s := httptest.NewServer(anotherprinter.NewAnotherPrinterMux(version))
	defer s.Close()

	type reqTable struct {
		path string
		result string
		testComparisonType TestComparisonType
	}

	table := []reqTable{
		{path: "/hello", result: "Hello!\n", testComparisonType: TestComparisonEqual},
		{path: "/api", result: "This is \"/api\" page.\n", testComparisonType: TestComparisonEqual},
		{path: "/version", result: "someVersion\n", testComparisonType: TestComparisonEqual},
		{path: "/randomnumber", result: "", testComparisonType: TestComparisonIntType},
	}

	for _, member := range table {
		t.Run(member.path, func(t *testing.T) {
			req, err := http.NewRequest(http.MethodGet, s.URL+member.path, nil)
			if err != nil {
				t.Fatalf("NewRequest unexpected error: %s", err.Error())
			}

			resp, err := http.DefaultClient.Do(req)
			if err != nil {
				t.Fatalf("DefaultClient.Do unexpected error: %s", err.Error())
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				t.Errorf("Status code expevted 200, got %s", http.StatusText(resp.StatusCode))
			}

			buf, err := io.ReadAll(resp.Body)
			if err != nil {
				t.Errorf("io.ReadAll error: %s", err.Error())
			}

			if len(buf) < 1 {
				t.Error("Empty string not expected")
				return
			}

			if buf[len(buf)-1] != '\n' {
				t.Error("Response expected to be finished by new string")
				return
			}

			if member.testComparisonType == TestComparisonEqual {
				if string(buf) != member.result {
					t.Errorf("Response expected: \"%s\", got: %s", member.result, buf)
				}
			}		
			
			if member.testComparisonType == TestComparisonIntType {
				_, err := strconv.Atoi(string(buf[:len(buf)-1]))
				if err != nil {
					t.Errorf("Response type expected: integer, got: %s", buf)
				}
			}		
		})
	}
}