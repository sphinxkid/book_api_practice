package handlers

import (
	"practice/book_api/pkg/domain"
)

type BooksHandler struct {
	BooksDb *domain.BooksDb
}
