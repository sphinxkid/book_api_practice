package domain

import (
	"practice/book_api/pkg/dto"
	"practice/book_api/pkg/error"
	"practice/book_api/pkg/logger"
)

type Book struct {
	ID    int    `db:"book_id"`
	Name  string `db:"book_name"`
	Genre string `db:"genre"`
	Count int    `db:"count"`
}

func (b Book) ToDto() dto.BookResponse {
	response := dto.BookResponse{
		ID:    b.ID,
		Name:  b.Name,
		Genre: b.Genre,
		Count: b.Count,
	}
	return response
}

// FindAllBooks queries all the books in the db
func (b *BooksDb) FindAllBooks() ([]Book, *error.Error) {

	findAllSql := "select book_id, book_name, genre, count from books"
	books := make([]Book, 0)
	err := b.db.Select(&books, findAllSql)
	if err != nil {
		logger.Error("booksDB.FindAllBooks. Error while querying books table " + err.Error())
		return nil, error.BookDBError.Wrap(err)
	}
	return books, nil
}

// CreateBook inserts a book in the db
func (b *BooksDb) CreateBook(book Book) (*int, *error.Error) {

	insertStatement := `
		INSERT INTO books(book_name, genre, count)
		VALUES(?, ?, ?)
	`

	result, err := b.db.Exec(insertStatement, book.Name, book.Genre, book.Count)
	if err != nil {
		logger.Error("booksDB.CreateBook. Error while Inserting Book " + err.Error())
		return nil, error.BookDBError.Wrap(err)
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		logger.Error("booksDB.CreateBook. Error while Getting Last Insert ID " + err.Error())
		return nil, error.BookDBError.Wrap(err)
	}
	lastInsertIdInt := int(lastInsertId)

	return &lastInsertIdInt, nil
}

// GetBookByID gets a book in the db using the book_id
func (b *BooksDb) GetBookByID(id int) (*Book, *error.Error) {
	getStatement := `
		SELECT book_name, genre, count FROM books WHERE book_id = ?
	`

	var book Book
	err := b.db.Get(&book, getStatement, id)
	if err != nil {
		logger.Error("booksDB.GetBookByID. Unable to Get Book By ID")
		return nil, error.BookDBError.Wrap(err)
	}
	if book.Name == "" {
		logger.Error("booksDB.GetBookByID. Error No Book Found")
		return nil, error.NoBookError.New()
	}
	book.ID = id

	return &book, nil
}

var Books = []Book{
	{ID: 1, Name: "Tree of Life", Genre: "Fantasy", Count: 5},
	{ID: 2, Name: "Reader's Perspective", Genre: "Slice of Life", Count: 5},
	{ID: 3, Name: "Born of Fire", Genre: "Fantasy", Count: 5},
}
