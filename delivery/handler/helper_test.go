package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/krittawatcode/books/domain"
	"github.com/stretchr/testify/assert"
)

func TestBindData(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		reqBody := strings.NewReader(`{"title":"100x","author":"Prach", "publication_year":"2021"}`)
		c.Request = httptest.NewRequest(http.MethodPost, "/test", reqBody)
		c.Request.Header.Set("Content-Type", "application/json")

		var book domain.Book
		result := bindData(c, &book)

		assert.True(t, result)
		assert.Equal(t, "Prach", book.Author)
		assert.Equal(t, "100x", book.Title)
		assert.Equal(t, "2021", book.PublicationYear)
	})

	t.Run("Invalid request body - non-json", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		reqBody := strings.NewReader(`100x`) // non-json
		c.Request = httptest.NewRequest(http.MethodPost, "/test", reqBody)
		c.Request.Header.Set("Content-Type", "application/json")

		var book domain.Book
		result := bindData(c, &book)

		assert.False(t, result)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})

	t.Run("Invalid request body - incorrect json structure", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		reqBody := strings.NewReader(`{title:"100x","author":"Prach"}`) // incorrect json structure
		c.Request = httptest.NewRequest(http.MethodPost, "/test", reqBody)
		c.Request.Header.Set("Content-Type", "application/json")

		var book domain.Book
		result := bindData(c, &book)

		assert.False(t, result)
		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}
