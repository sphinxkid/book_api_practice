package domain

import (
	"fmt"
	"log"
)

type Book struct {
	ID    int
	Name  string
	Genre string
	Count int
}

func (b *BooksDb) FindAllBooks() ([]Book, error) {

	findAllSql := "select book_id, book_name, genre, count from books"

	rows, err := b.db.Query(findAllSql)
	if err != nil {
		log.Println("Error while querying books table " + err.Error())
		return nil, err
	}

	books := make([]Book, 0)
	for rows.Next() {
		var b Book
		err := rows.Scan(&b.ID, &b.Name, &b.Genre, &b.Count)
		if err != nil {
			log.Println("Error while scanning books " + err.Error())
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}

func (b *BooksDb) CreateBook(book Book) (*int, error) {

	insertStatement := `
		INSERT INTO books(book_name, genre, count)
		VALUES(?, ?, ?)
	`

	result, err := b.db.Exec(insertStatement, book.Name, book.Genre, book.Count)
	if err != nil {
		log.Println("Error while Inserting Book " + err.Error())
		return nil, err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		log.Println("Error while Getting Last Insert ID " + err.Error())
		return nil, err
	}
	lastInsertIdInt := int(lastInsertId)

	return &lastInsertIdInt, nil
}

func (b *BooksDb) GetBookByID(id int) (*Book, error) {
	getStatement := `
		SELECT book_name, genre, count FROM books WHERE book_id = ?
	`
	row := b.db.QueryRow(getStatement, id)
	if row == nil {
		log.Println("Error while Getting Book")
		return nil, fmt.Errorf("Error while Getting Book")
	}

	var book Book
	row.Scan(&book.Name, &book.Genre, &book.Count)
	book.ID = id

	return &book, nil
}

var Books = []Book{
	{ID: 1, Name: "Tree of Life", Genre: "Fantasy", Count: 5},
	{ID: 2, Name: "Reader's Perspective", Genre: "Slice of Life", Count: 5},
	{ID: 3, Name: "Born of Fire", Genre: "Fantasy", Count: 5},
}
