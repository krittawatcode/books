package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/krittawatcode/books/domain"
	"github.com/krittawatcode/books/domain/apperror"
	"github.com/krittawatcode/books/domain/appmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestBookHandler_FetchBooks(t *testing.T) {
	// setup
	gin.SetMode(gin.TestMode)

	t.Run("Error", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		// Initialize the Request of gin.Context
		c.Request, _ = http.NewRequest("GET", "/books", nil)

		mockBookUseCase := new(appmock.MockBookUseCase)
		mockBookUseCase.On("FetchBooks", mock.Anything).Return(&[]domain.Book{}, apperror.NewNotFound("Book", "ID", ""))

		h := &BookHandler{
			BookUseCase: mockBookUseCase,
		}

		h.FetchBooks(c)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("Success", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		// Initialize the Request of gin.Context
		c.Request, _ = http.NewRequest("GET", "/books", nil)

		mockBooks := []domain.Book{
			{ID: uuid.New(), Title: "Book 1", Author: "Author 1", PublicationYear: "2021"},
			{ID: uuid.New(), Title: "Book 2", Author: "Author 2", PublicationYear: "2022"},
			{ID: uuid.New(), Title: "Book 3", Author: "Author 3", PublicationYear: "2023"},
		}
		mockBookUseCase := new(appmock.MockBookUseCase)
		mockBookUseCase.On("FetchBooks", mock.Anything).Return(&mockBooks, nil)

		h := &BookHandler{
			BookUseCase: mockBookUseCase,
		}

		h.FetchBooks(c)

		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func TestBookHandler_CreateBook(t *testing.T) {
	// setup
	gin.SetMode(gin.TestMode)

	t.Run("Error", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		// Create a new book
		newBook := &domain.Book{
			Title:           "New Book",
			Author:          "New Author",
			PublicationYear: "2022",
		}

		// Convert the new book to JSON
		newBookJSON, _ := json.Marshal(newBook)

		// Initialize the Request of gin.Context
		c.Request, _ = http.NewRequest("POST", "/books", bytes.NewBuffer(newBookJSON))
		c.Request.Header.Set("Content-Type", "application/json")

		mockBookUseCase := new(appmock.MockBookUseCase)
		mockBookUseCase.On("CreateBook", mock.Anything, mock.AnythingOfType("*domain.Book")).Return(apperror.NewNotFound("Book", "ID", ""))

		h := &BookHandler{
			BookUseCase: mockBookUseCase,
		}

		h.CreateBook(c)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("Success", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		// Create a new book
		newBook := &domain.Book{
			Title:           "New Book",
			Author:          "New Author",
			PublicationYear: "2022",
		}

		// Convert the new book to JSON
		newBookJSON, _ := json.Marshal(newBook)

		// Initialize the Request of gin.Context
		c.Request, _ = http.NewRequest("POST", "/books", bytes.NewBuffer(newBookJSON))
		c.Request.Header.Set("Content-Type", "application/json")

		mockBookUseCase := new(appmock.MockBookUseCase)
		mockBookUseCase.On("CreateBook", mock.Anything, mock.AnythingOfType("*domain.Book")).Return(nil)

		h := &BookHandler{
			BookUseCase: mockBookUseCase,
		}

		h.CreateBook(c)

		assert.Equal(t, http.StatusCreated, w.Code)
	})
}

func TestBookHandler_GetBookByID(t *testing.T) {
	// setup
	gin.SetMode(gin.TestMode)

	t.Run("Error", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		// Initialize the Request of gin.Context
		c.Request, _ = http.NewRequest("GET", "/books/1", nil)

		mockBookUseCase := new(appmock.MockBookUseCase)
		mockBookUseCase.On("GetBookByID", mock.Anything, mock.AnythingOfType("string")).Return(&domain.Book{}, apperror.NewNotFound("Book", "ID", "1"))

		h := &BookHandler{
			BookUseCase: mockBookUseCase,
		}

		h.GetBookByID(c)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("Success", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		// Initialize the Request of gin.Context
		c.Request, _ = http.NewRequest("GET", "/books/1", nil)

		mockBook := &domain.Book{
			ID:              uuid.New(),
			Title:           "Book 1",
			Author:          "Author 1",
			PublicationYear: "2021",
		}

		mockBookUseCase := new(appmock.MockBookUseCase)
		mockBookUseCase.On("GetBookByID", mock.Anything, mock.AnythingOfType("string")).Return(mockBook, nil)

		h := &BookHandler{
			BookUseCase: mockBookUseCase,
		}

		h.GetBookByID(c)

		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func TestBookHandler_UpdateBook(t *testing.T) {
	// setup
	gin.SetMode(gin.TestMode)

	t.Run("Error", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		// Create a book to update
		book := &domain.Book{
			ID:              uuid.New(),
			Title:           "Updated Book",
			Author:          "Updated Author",
			PublicationYear: "2022",
		}

		// Convert the book to JSON
		bookBytes, _ := json.Marshal(book)

		// Initialize the Request of gin.Context with the JSON book
		c.Request, _ = http.NewRequest("PUT", "/books/1", bytes.NewBuffer(bookBytes))
		c.Request.Header.Set("Content-Type", "application/json")

		mockBookUseCase := new(appmock.MockBookUseCase)
		mockBookUseCase.On("UpdateBook", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("*domain.Book")).Return(apperror.NewNotFound("Book", "ID", "1"))

		h := &BookHandler{
			BookUseCase: mockBookUseCase,
		}

		h.UpdateBook(c)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("Success", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		// Create a book to update
		book := &domain.Book{
			ID:              uuid.New(),
			Title:           "Updated Book",
			Author:          "Updated Author",
			PublicationYear: "2022",
		}

		// Convert the book to JSON
		bookBytes, _ := json.Marshal(book)

		// Initialize the Request of gin.Context with the JSON book
		c.Request, _ = http.NewRequest("PUT", "/books/1", bytes.NewBuffer(bookBytes))
		c.Request.Header.Set("Content-Type", "application/json")

		mockBookUseCase := new(appmock.MockBookUseCase)
		mockBookUseCase.On("UpdateBook", mock.Anything, mock.AnythingOfType("string"), mock.AnythingOfType("*domain.Book")).Return(nil)

		h := &BookHandler{
			BookUseCase: mockBookUseCase,
		}

		h.UpdateBook(c)

		assert.Equal(t, http.StatusOK, w.Code)
	})
}

func TestBookHandler_DeleteBook(t *testing.T) {
	// setup
	gin.SetMode(gin.TestMode)

	t.Run("Error", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		// Initialize the Request of gin.Context
		c.Request, _ = http.NewRequest("DELETE", "/books/1", nil)

		mockBookUseCase := new(appmock.MockBookUseCase)
		mockBookUseCase.On("DeleteBook", mock.Anything, mock.AnythingOfType("string")).Return(apperror.NewNotFound("Book", "ID", "1"))

		h := &BookHandler{
			BookUseCase: mockBookUseCase,
		}

		h.DeleteBook(c)

		assert.Equal(t, http.StatusNotFound, w.Code)
	})

	t.Run("Success", func(t *testing.T) {
		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		// Initialize the Request of gin.Context
		c.Request, _ = http.NewRequest("DELETE", "/books/1", nil)

		mockBookUseCase := new(appmock.MockBookUseCase)
		mockBookUseCase.On("DeleteBook", mock.Anything, mock.AnythingOfType("string")).Return(nil)

		h := &BookHandler{
			BookUseCase: mockBookUseCase,
		}

		h.DeleteBook(c)

		assert.Equal(t, http.StatusOK, w.Code)
	})
}
