package handlers

import (
	"fmt"
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

	id, err := bh.BooksDb.CreateBook(newBook)
	if err != nil {
		httpErr := error.NewHttpError("Unable to Create Book", "", http.StatusInternalServerError)
		c.Error(httpErr)
		return
	}
	newBook.ID = *id
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

	book, err := bh.BooksDb.GetBookByID(int(idInt))
	if err != nil {
		errorMsg := fmt.Sprintf("Unable to Get Book with ID:%d", idInt)
		httpErr := error.NewHttpError(errorMsg, "", http.StatusNotFound)
		c.Error(httpErr)
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}
