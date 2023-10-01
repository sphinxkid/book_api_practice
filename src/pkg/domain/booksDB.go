package domain

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type BooksDb struct {
	db *sqlx.DB
}

func NewBooksDb() *BooksDb {
	db, err := sqlx.Open("mysql", "test:test@tcp(127.0.0.1:3306)/books")
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
