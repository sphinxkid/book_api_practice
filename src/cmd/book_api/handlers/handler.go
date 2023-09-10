package handlers

import "practice/book_api/pkg/service"

type BooksHandler struct {
	booksService *service.BookService
}

func NewBooksHandler(bookSvc *service.BookService) *BooksHandler {
	booksHandler := &BooksHandler{
		booksService: bookSvc,
	}
	return booksHandler
}
