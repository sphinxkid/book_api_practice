package domain

import "log"

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

var Books = []Book{
	{ID: 1, Name: "Tree of Life", Genre: "Fantasy", Count: 5},
	{ID: 2, Name: "Reader's Perspective", Genre: "Slice of Life", Count: 5},
	{ID: 3, Name: "Born of Fire", Genre: "Fantasy", Count: 5},
}
