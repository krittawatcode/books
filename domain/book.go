package domain

import (
	"context"

	"github.com/google/uuid"
)

type Book struct {
	ID              uuid.UUID `json:"id"`
	Title           string    `binding:"required" json:"title"`
	Author          string    `binding:"required" json:"author"`
	PublicationYear string    `binding:"required" json:"publication_year"`
}

type BookUseCase interface {
	FetchBooks(ctx context.Context) (*[]Book, error)
	GetBookByID(ctx context.Context, id string) (*Book, error)
	CreateBook(ctx context.Context, book *Book) error
	UpdateBook(ctx context.Context, id string, book *Book) error
	DeleteBook(ctx context.Context, id string) error
}

type BookRepository interface {
	FetchBooks(ctx context.Context) (*[]Book, error)
	GetBookByID(ctx context.Context, id string) (*Book, error)
	CreateBook(ctx context.Context, book *Book) error
	UpdateBook(ctx context.Context, id string, book *Book) error
	DeleteBook(ctx context.Context, id string) error
}
