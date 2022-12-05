package errors

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

var (
	ErrAlreadyExists = errors.New("Already Exists")
)

func GetErrorCode(err error) int {
	switch {
	case errors.Is(err, sql.ErrNoRows):
		return http.StatusNoContent
	case errors.Is(err, ErrAlreadyExists):
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
