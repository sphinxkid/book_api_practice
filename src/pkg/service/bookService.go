package service

import (
	"practice/book_api/pkg/domain"
	"practice/book_api/pkg/error"
)

type BookService struct {
	booksDb *domain.BooksDb
}

func NewBookService(db *domain.BooksDb) *BookService {
	newBS := &BookService{
		booksDb: db,
	}
	return newBS
}

func (bs *BookService) GetAllBooks() ([]domain.Book, *error.Error) {
	books, appErr := bs.booksDb.FindAllBooks()
	if appErr != nil {
		return nil, appErr
	}
	return books, nil
}

func (bs *BookService) CreateABook(newBook domain.Book) (*domain.Book, *error.Error) {
	if newBook.Name == "" || newBook.Genre == "" || newBook.Count < 0 {
		return nil, error.IncorrectNewBookError.New()
	}

	id, appErr := bs.booksDb.CreateBook(newBook)
	if appErr != nil {
		return nil, appErr
	}
	newBook.ID = *id
	return &newBook, nil
}

func (bs *BookService) GetBookByID(id int) (*domain.Book, *error.Error) {
	book, appErr := bs.booksDb.GetBookByID(id)
	if appErr != nil {
		return nil, appErr
	}
	return book, nil
}
