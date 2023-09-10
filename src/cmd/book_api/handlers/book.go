package handlers

import (
	"net/http"
	"practice/book_api/pkg/domain"
	"practice/book_api/pkg/error"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (bh *BooksHandler) GetBooks(c *gin.Context) {
	books, appErr := bh.BooksDb.FindAllBooks()
	if appErr != nil {
		c.Error(*appErr)
		return
	}
	c.IndentedJSON(http.StatusOK, books)
}

func (bh *BooksHandler) PostBooks(c *gin.Context) {
	var newBook domain.Book

	err := c.BindJSON(&newBook)
	if err != nil {
		c.Error(*error.JsonBindingError.Wrap(err))
		return
	}

	id, appErr := bh.BooksDb.CreateBook(newBook)
	if appErr != nil {
		c.Error(*appErr)
		return
	}
	newBook.ID = *id
	c.IndentedJSON(http.StatusCreated, newBook)
}

func (bh *BooksHandler) GetBookByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		c.Error(*error.ParamParseError.Wrap(err))
		return
	}

	book, appErr := bh.BooksDb.GetBookByID(int(idInt))
	if appErr != nil {
		c.Error(*appErr)
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}
