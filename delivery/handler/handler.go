package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/krittawatcode/books/delivery/middleware"
	"github.com/krittawatcode/books/domain"
	"github.com/krittawatcode/books/domain/apperror"
)

type BookHandler struct {
	Router          *gin.Engine
	BookUseCase     domain.BookUseCase
	Path            string // path for book routes
	TimeoutDuration time.Duration
}

func NewBookHandler(router *gin.Engine, bu domain.BookUseCase, path string, timeout time.Duration) *BookHandler {
	handler := &BookHandler{
		Router:          router,
		BookUseCase:     bu,
		Path:            path,
		TimeoutDuration: timeout,
	}

	// Create an books group
	g := router.Group(path)
	// setup middleware
	g.Use(middleware.Timeout(timeout, apperror.NewServiceUnavailable()))
	// setup routes
	g.GET("/", handler.FetchBooks)
	g.POST("/", handler.CreateBook)
	g.GET("/:id", handler.GetBookByID)
	g.PUT("/:id", handler.UpdateBook)
	g.DELETE("/:id", handler.DeleteBook)

	return handler
}

func (h *BookHandler) FetchBooks(c *gin.Context) {
	books, err := h.BookUseCase.FetchBooks(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{response: response{Status: statusFail, Code: codeFail}, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, successResponse{response: response{Status: statusSuccess, Code: codeSuccess}, Data: books})
}

func (h *BookHandler) CreateBook(c *gin.Context) {
	var book domain.Book
	if ok := bindData(c, &book); !ok {
		return
	}

	err := h.BookUseCase.CreateBook(c.Request.Context(), &book)
	if err != nil {
		c.JSON(apperror.Status(err), errorResponse{response: response{Status: statusFail, Code: codeFail}, Error: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, successResponse{response: response{Status: statusSuccess, Code: codeSuccess}, Data: book})
}

func (h *BookHandler) GetBookByID(c *gin.Context) {
	id := c.Param("id")

	book, err := h.BookUseCase.GetBookByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(apperror.Status(err), errorResponse{response: response{Status: statusFail, Code: codeFail}, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, successResponse{response: response{Status: statusSuccess, Code: codeSuccess}, Data: book})
}

func (h *BookHandler) UpdateBook(c *gin.Context) {
	var book domain.Book
	if ok := bindData(c, &book); !ok {
		return
	}

	id := c.Param("id")
	err := h.BookUseCase.UpdateBook(c.Request.Context(), id, &book)
	if err != nil {
		c.JSON(apperror.Status(err), errorResponse{response: response{Status: statusFail, Code: codeFail}, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, successResponse{response: response{Status: statusSuccess, Code: codeSuccess}, Data: book})
}

func (h *BookHandler) DeleteBook(c *gin.Context) {
	id := c.Param("id")

	err := h.BookUseCase.DeleteBook(c.Request.Context(), id)
	if err != nil {
		c.JSON(apperror.Status(err), errorResponse{response: response{Status: statusFail, Code: codeFail}, Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, successResponse{response: response{Status: statusSuccess, Code: codeSuccess}})
}
