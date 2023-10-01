package service

import (
	"practice/book_api/pkg/domain"
	"practice/book_api/pkg/dto"
	"practice/book_api/pkg/error"
	"practice/book_api/pkg/logger"
)

// BookService is the struct that contains the service for books
type BookService struct {
	booksDb *domain.BooksDb
}

// Creates a new Book Service
func NewBookService(db *domain.BooksDb) *BookService {
	newBS := &BookService{
		booksDb: db,
	}
	return newBS
}

// GetAllBooks gets all the books
func (bs *BookService) GetAllBooks() ([]dto.BookResponse, *error.Error) {
	books, appErr := bs.booksDb.FindAllBooks()
	if appErr != nil {
		logger.Error("bookService.GetAllBooks. Unable to get all books")
		return nil, appErr
	}
	bookResponses := make([]dto.BookResponse, 0)
	for _, book := range books {
		bookDto := book.ToDto()
		bookResponses = append(bookResponses, bookDto)
	}
	return bookResponses, nil
}

// CreateABook creates a new book
func (bs *BookService) CreateABook(newBook dto.BookResponse) (*dto.BookResponse, *error.Error) {
	// Checks if the book has correct fields
	if newBook.Name == "" || newBook.Genre == "" || newBook.Count < 0 {
		logger.Error("bookService.CreateABook. New Book has incorrect fields")
		return nil, error.IncorrectNewBookError.New()
	}
	domainNewBook := domain.Book{
		ID:    newBook.ID,
		Name:  newBook.Name,
		Genre: newBook.Genre,
		Count: newBook.Count,
	}

	id, appErr := bs.booksDb.CreateBook(domainNewBook)
	if appErr != nil {
		logger.Error("bookService.CreateABook. Unable to create book")
		return nil, appErr
	}
	newBook.ID = *id
	return &newBook, nil
}

// GetBookByID gets a book using the ID
func (bs *BookService) GetBookByID(id int) (*dto.BookResponse, *error.Error) {
	book, appErr := bs.booksDb.GetBookByID(id)
	if appErr != nil {
		logger.Error("bookService.GetBookByID. Unable to get book")
		return nil, appErr
	}
	response := book.ToDto()
	return &response, nil
}
