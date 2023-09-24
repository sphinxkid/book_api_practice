package handlers

import (
	"fmt"
	"net/http"
	"practice/book_api/pkg/domain"
	"practice/book_api/pkg/error"
	"practice/book_api/pkg/logger"
	"practice/book_api/pkg/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BooksHandler struct {
	booksService *service.BookService
}

func NewBooksHandler(bookSvc *service.BookService) *BooksHandler {
	booksHandler := &BooksHandler{
		booksService: bookSvc,
	}
	return booksHandler
}

func (bh *BooksHandler) GetBooks(c *gin.Context) {
	books, appErr := bh.booksService.GetAllBooks()
	if appErr != nil {
		errorString := fmt.Sprintf("booksHandler.GetBooks. Unable to Get All Books")
		logger.Error(errorString)
		c.Error(*appErr)
		return
	}
	c.IndentedJSON(http.StatusOK, books)
}

func (bh *BooksHandler) PostBooks(c *gin.Context) {
	var newBook domain.Book

	err := c.BindJSON(&newBook)
	if err != nil {
		errorString := fmt.Sprintf("booksHandler.PostBooks. Unable to bind book json: %+v", newBook)
		logger.Error(errorString)
		c.Error(*error.JsonBindingError.Wrap(err))
		return
	}

	book, appErr := bh.booksService.CreateABook(newBook)
	if appErr != nil {
		errorString := fmt.Sprintf("booksHandler.PostBooks. Unable to create a book: %+v", newBook)
		logger.Error(errorString)
		c.Error(*appErr)
		return
	}
	c.IndentedJSON(http.StatusCreated, book)
}

func (bh *BooksHandler) GetBookByID(c *gin.Context) {
	id := c.Param("id")
	idInt, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		errorString := fmt.Sprintf("booksHandler.GetBookById. Unable to Parse ID: %s", id)
		logger.Error(errorString)
		c.Error(*error.ParamParseError.Wrap(err))
		return
	}

	book, appErr := bh.booksService.GetBookByID(int(idInt))
	if appErr != nil {
		errorString := fmt.Sprintf("booksHandler.GetBookByID. Unable to Get Book with id: %d", idInt)
		logger.Error(errorString)
		c.Error(*appErr)
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}
