package handlers

import (
	"net/http"
	"practice/book_api/pkg/domain"
	"practice/book_api/pkg/error"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (bh *BooksHandler) GetBooks(c *gin.Context) {
	books, err := bh.BooksDb.FindAllBooks()
	if err != nil {
		httpErr := error.NewHttpError("Unable to Find all Books", "", http.StatusInternalServerError)
		c.Error(httpErr)
		return
	}
	c.IndentedJSON(http.StatusOK, books)
}

func (bh *BooksHandler) PostBooks(c *gin.Context) {
	var newBook domain.Book

	err := c.BindJSON(&newBook)
	if err != nil {
		httpErr := error.NewHttpError("Unable to Bind JSON", "", http.StatusInternalServerError)
		c.Error(httpErr)
		return
	}

	domain.Books = append(domain.Books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func (bh *BooksHandler) GetBookByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		httpErr := error.NewHttpError("Unable to Parse ID to Int", "", http.StatusBadRequest)
		c.Error(httpErr)
		return
	}

	for _, b := range domain.Books {
		if b.ID == int(idInt) {
			c.IndentedJSON(http.StatusOK, b)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
