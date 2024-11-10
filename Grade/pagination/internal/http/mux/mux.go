package mux

import (
	"net/http"
	"pagination/internal/http/handlers"
	"strconv"
)

func NewBooksMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request){
		var pageInt, limitInt int
		var err error
		r.ParseForm()
		
		page := r.FormValue("page")
		if page != "" {
			if pageInt, err = strconv.Atoi(page); err != nil || pageInt <= 0 {
				http.Error(w, "Incorrect page value", http.StatusBadRequest)
				return
			}

		limit := r.FormValue("limit")
		if limit != "" {
			if limitInt, err = strconv.Atoi(limit); err != nil || limitInt <= 0 {
				http.Error(w, "Incorrect limit value", http.StatusBadRequest)
				return
			}
		}
		
		author := r.FormValue("author")

		handlers.VewBooks(w, r, pageInt, limitInt, handlers.FilterBooks(author))

		}
	})
	return mux
}