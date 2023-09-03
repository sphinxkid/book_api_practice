package domain

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Book struct {
	ID    int
	Name  string
	Genre string
	Count int
}

type BooksDb struct {
	db *sql.DB
}

func NewBooksDb() *BooksDb {
	db, err := sql.Open("mysql", "test:test@tcp(127.0.0.1:3306)/books")
	if err != nil {
		panic(err)
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	bookDB := &BooksDb{
		db: db,
	}

	return bookDB
}

var Books = []Book{
	{ID: 1, Name: "Tree of Life", Genre: "Fantasy", Count: 5},
	{ID: 2, Name: "Reader's Perspective", Genre: "Slice of Life", Count: 5},
	{ID: 3, Name: "Born of Fire", Genre: "Fantasy", Count: 5},
}
