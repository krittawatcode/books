package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krittawatcode/books/domain/apperror"
)

type status string

const (
	statusSuccess status = "SUCCESS"
	statusFail    status = "FAIL"
)

type code int

// represent our own service status code
const (
	codeSuccess code = 200
	codeFail    code = 500
)

type response struct {
	Status status `json:"status"`
	Code   code   `json:"code"`
}

type successResponse struct {
	response
	Data interface{} `json:"data"`
}

// ErrorResponse represents an error response
type errorResponse struct {
	response
	Error string `json:"error"`
}

// bindData is helper function, returns false if data is not bound
func bindData(c *gin.Context, req interface{}) bool {
	if c.ContentType() != "application/json" {
		msg := fmt.Sprintf("%s only accepts Content-Type application/json", c.FullPath())

		err := apperror.NewUnsupportedMediaType(msg)

		c.JSON(err.Status(), errorResponse{response: response{Status: statusFail, Code: codeFail}, Error: err.Error()})

		return false
	}

	// Bind incoming json to struct and check for validation errors
	if err := c.ShouldBind(req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{response: response{Status: statusFail, Code: codeFail}, Error: err.Error()})

		return false
	}

	return true
}
