package handlers

import (
	"encoding/json"
	"net/http"
	"pagination/internal/repository/books"
)

func VewBooks(w http.ResponseWriter, r *http.Request, page int, limit int, books []books.Book) {
	startIndex := (page-1) * limit
	endIndex := startIndex + limit

	if startIndex > len(books) {
		http.Error(w, "No so many pages", http.StatusBadRequest)
		return
	}
	if endIndex > len(books) {
		endIndex = len(books)
	}

	json.NewEncoder(w).Encode(books[startIndex:endIndex])
}

func FilterBooks(author string) (filteredBooks []books.Book) {
	if author == "" {
		return books.Books
	}

	for _, book := range books.Books {
		if book.Author == author {
			filteredBooks = append(filteredBooks, book)
		}
	}
	return
}