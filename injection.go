package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/krittawatcode/books/delivery/handler"
	"github.com/krittawatcode/books/repository"
	"github.com/krittawatcode/books/usecase"
)

func inject() (*gin.Engine, error) {
	bookRepo := repository.NewInMemoryBookRepository()
	bookUsecase := usecase.NewBookUseCase(bookRepo)

	// initialize gin.Engine
	router := gin.Default()

	// get environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	booksPath := os.Getenv("BOOKS_PATH")
	handlerTimeout := os.Getenv("HANDLER_TIMEOUT")
	ht, err := strconv.ParseInt(handlerTimeout, 0, 64)
	if err != nil {
		return nil, fmt.Errorf("could not parse HANDLER_TIMEOUT as int: %w", err)
	}
	timeout := time.Duration(time.Duration(ht) * time.Second)

	// inject dependencies
	handler.NewBookHandler(router, bookUsecase, booksPath, timeout)

	// setup health check
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "running"})
	})

	return router, nil
}
