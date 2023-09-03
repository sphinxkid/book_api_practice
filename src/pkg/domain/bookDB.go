package domain

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
)

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
