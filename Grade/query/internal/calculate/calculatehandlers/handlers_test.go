package calculatehandlers

import (
	// "encoding/json"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"

	// "strconv"
	// "strings"
	"testing"
)

func TestCalculateHandlerErrors(t *testing.T) {
	errorTestCases := map[string]string{
		"": "numbers absent",
		"number1=1": "numbers absent",
		"number2=1": "numbers absent",
		"number1=1a": "numbers absent",
		"number2=1a": "numbers absent",
		"number1=1a&number2=1": "numbers should be int",
		"number1=1&number2=1a": "numbers should be int",
		"number1=1a&number2=1a": "numbers should be int",
	}

	for testCase, expectedResult := range errorTestCases {
		w := httptest.NewRecorder()
		r := httptest.NewRequest(http.MethodGet, "/calculate?" + testCase, nil)
		
		CalculateHandler(w, r)
	
		resp := w.Result()
		defer resp.Body.Close()
	
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Status code expected 200, got %s", http.StatusText(resp.StatusCode))
		}
	
		buf, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("io.ReadAll error: %s", err.Error())
		}
		var bufJSON Result
		err = json.Unmarshal(buf, &bufJSON)
		if err != nil {
			t.Errorf("Error while unmarshalling JSON answer: %s", err.Error())
		}
	
		if bufJSON.Error != expectedResult {
			t.Errorf("Test case: %s, got: %s", testCase, bufJSON.Error)
		}
	}
}

func TestCalculateHandlerSuccess(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/calculate?number1=10&number2=15", nil)
	
	CalculateHandler(w, r)

	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status code expected 200, got %s", http.StatusText(resp.StatusCode))
	}

	buf, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("io.ReadAll error: %s", err.Error())
	}

	expected := `{"result":25,"error":"","parameters":{"number1":10,"number2":15}}` + "\n"
	if string(buf) != expected {
		t.Errorf("Test case: %s, got: %s", expected, string(buf))
	}
}