package calculatehandlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)


type ResultParameters struct {
	Number1 int `json:"number1"`
	Number2 int `json:"number2"`
}
type Result struct {
	Result int `json:"result"`
	Error string `json:"error"`
	Parameters ResultParameters `json:"parameters"`
}

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	var number1, number2 int
	var err1, err2 error
	var errors []string
	
	query := r.URL.Query()
	
	if number1string, found := query["number1"]; found && len(number1string) > 0 {
		number1, err1 = strconv.Atoi(number1string[0])
	} else {
		errors = append(errors,"number 1 absent")
	}
	if number2string, found := query["number2"]; found && len(number2string) > 0 {
		number2, err2 = strconv.Atoi(number2string[0])
	} else {
		errors = append(errors,"number 2 absent")
	}
	if err1 != nil || err2 != nil {
		errors = append(errors,"numbers should be int")
	}

	errorString := strings.Join(errors, ", ")

	result := Result{
		Result: number1 + number2,
		Error: errorString,
		Parameters: ResultParameters{
			Number1: number1,
			Number2: number2,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(result); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
