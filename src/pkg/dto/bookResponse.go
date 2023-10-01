package dto

type BookResponse struct {
	ID    int    `json:"book_id"`
	Name  string `json:"book_name"`
	Genre string `json:"genre"`
	Count int    `json:"count"`
}
