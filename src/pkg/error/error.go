package error

import (
	"fmt"
	"net/http"
)

// DB Errors
var (
	BookDBError     = NewError(1001, http.StatusInternalServerError, "Error in Books DB")
	RowParsingError = NewError(1002, http.StatusInternalServerError, "Error in Parsing Rows")
	NoBookError     = NewError(1003, http.StatusInternalServerError, "No Book Found")
)

// Http Errors
var (
	JsonBindingError = NewError(2001, http.StatusInternalServerError, "Error in Binding JSON")
	ParamParseError  = NewError(2002, http.StatusBadRequest, "Unable to Parse URL Param")
)

// Service Errors
var (
	IncorrectNewBookError = NewError(3001, http.StatusBadRequest, "Book has incorrect fields")
)

type Error struct {
	Code        int    `json:"code"`
	StatusCode  int    `json:"status_code"`
	Description string `json:"description,omitempty"`
	Inner       *Error `json:"inner,omitempty"`
}

func (e Error) Error() string {
	return fmt.Sprintf("Code: %d, Description: %s", e.Code, e.Description)
}

func NewError(Code int, StatusCode int, Description string) *Error {
	newError := &Error{
		Code:        Code,
		StatusCode:  StatusCode,
		Description: Description,
	}
	return newError
}

func (e Error) New() *Error {
	newError := &Error{
		Code:        e.Code,
		StatusCode:  e.StatusCode,
		Description: e.Description,
	}
	return newError
}

func (e Error) Copy() *Error {
	newError := &Error{
		Code:        e.Code,
		StatusCode:  e.StatusCode,
		Description: e.Description,
		Inner:       e.Inner,
	}
	return newError
}

func (e Error) Wrap(err error) *Error {
	newError := &Error{
		Code:        e.Code,
		StatusCode:  e.StatusCode,
		Description: err.Error(),
	}
	e.Inner = newError
	return e.Copy()
}
