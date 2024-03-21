package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/krittawatcode/books/domain/apperror"
	"github.com/stretchr/testify/assert"
)

func TestTimeout(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("Success", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		// Initialize the Request of gin.Context
		c.Request, _ = http.NewRequest("GET", "/books", nil)

		h := Timeout(1*time.Second, apperror.NewServiceUnavailable())

		h(c)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Timeout", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		// Initialize the Request of gin.Context
		c.Request, _ = http.NewRequest("GET", "/books", nil)

		h := Timeout(1*time.Nanosecond, apperror.NewServiceUnavailable())

		h(c)

		assert.Equal(t, http.StatusServiceUnavailable, w.Code)
	})
}
