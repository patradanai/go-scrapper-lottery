package httperror

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	ErrBadRequest          = "Bad request"
	ErrNotFound            = "Not Found"
	ErrUnauthorized        = "Unauthorized"
	ErrRequestTimeout      = "Request Timeout"
	ErrInvalidEmail        = "Invalid email"
	ErrInvalidPassword     = "Invalid password"
	ErrInvalidField        = "Invalid field"
	ErrInternalServerError = "Internal Server Error"
)

type responseError struct {
	Success      bool
	ErrorStatus  int
	ErrorCode    string
	ErrorMessage interface{}
	TimeStamp    time.Time
}

func NewInternalServerError(ctx *gin.Context, message interface{}) {
	responseError := &responseError{
		Success:      false,
		ErrorStatus:  http.StatusInternalServerError,
		ErrorCode:    ErrInternalServerError,
		ErrorMessage: message,
		TimeStamp:    time.Now().UTC(),
	}

	ctx.JSON(http.StatusInternalServerError, responseError)
}

func NewBadRequest(ctx *gin.Context, message interface{}) {
	responseError := &responseError{
		Success:      false,
		ErrorStatus:  http.StatusBadRequest,
		ErrorCode:    ErrBadRequest,
		ErrorMessage: message,
		TimeStamp:    time.Now().UTC(),
	}

	ctx.JSON(http.StatusBadRequest, responseError)
}

func NewRequestNotFound(ctx *gin.Context, message interface{}) {
	responseError := &responseError{
		Success:      false,
		ErrorStatus:  http.StatusNotFound,
		ErrorCode:    ErrNotFound,
		ErrorMessage: message,
		TimeStamp:    time.Now().UTC(),
	}

	ctx.JSON(http.StatusNotFound, responseError)
}

func NewUnauthorize(ctx *gin.Context, message interface{}) {
	responseError := &responseError{
		Success:      false,
		ErrorStatus:  http.StatusUnauthorized,
		ErrorCode:    ErrUnauthorized,
		ErrorMessage: message,
		TimeStamp:    time.Now().UTC(),
	}

	ctx.JSON(http.StatusUnauthorized, responseError)
}

func NewCtxResponse() {

}
