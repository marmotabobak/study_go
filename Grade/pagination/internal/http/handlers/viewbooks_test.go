package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"pagination/internal/repository/books"
	"testing"
)

func TestViewBooks (t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/books?author=AuthorNew", nil)

	ViewBooks(w, r, 1, 1, books.Books)

	resp := w.Result()
	
	if resp.StatusCode != http.StatusOK {}

	buf, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {}
	
	expected := `[{"id":1,"title":"Book 1","author":"Author 1","price":19.99}]` + "\n"
	if string(buf) != expected {
		t.Errorf("got: %s", string(buf))
	}
}

func TestFilterBooks(t *testing.T) {
	resultJSON, _ := json.Marshal(FilterBooks("AuthorTest"))
	expectedJSON := `[{"id":51,"title":"Book 51","author":"AuthorTest","price":1.23},{"id":52,"title":"Book 52","author":"AuthorTest","price":4.56}]`

	if string(resultJSON) != expectedJSON {
		t.Errorf("got: %s", string(resultJSON))
	}
}