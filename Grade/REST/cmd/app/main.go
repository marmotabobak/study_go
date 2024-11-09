package main

import (
	"log"
	"net/http"
	"REST/internal/rest"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			rest.GetTasks(w)
		case http.MethodPost:
			rest.AddTask(w, r)
		default:
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		}
	})
	mux.HandleFunc("/tasks/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPut:
			rest.ChangeTask(w, r)
		case http.MethodDelete:
			rest.DeleteTask(w, r)
		default:
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		}
	})
	log.Fatal(http.ListenAndServe(":8080", mux))
}
