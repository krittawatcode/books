package repository

import (
	"context"
	"sync"

	"github.com/google/uuid"
	"github.com/krittawatcode/books/domain"
	"github.com/krittawatcode/books/domain/apperror"
)

type InMemoryBookRepository struct {
	books []domain.Book
	mu    sync.Mutex
}

func NewInMemoryBookRepository() domain.BookRepository {
	return &InMemoryBookRepository{
		books: []domain.Book{},
	}
}

func (r *InMemoryBookRepository) FetchBooks(ctx context.Context) (*[]domain.Book, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if len(r.books) == 0 {
		return nil, apperror.NewNotFound("Book", "ID", "")
	}

	return &r.books, nil
}

func (r *InMemoryBookRepository) GetBookByID(ctx context.Context, id string) (*domain.Book, error) {
	for _, book := range r.books {
		if book.ID.String() == id {
			return &book, nil
		}
	}

	return nil, apperror.NewNotFound("Book", "ID", id)
}

func (r *InMemoryBookRepository) CreateBook(ctx context.Context, book *domain.Book) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	// check if book already exists
	for _, b := range r.books {
		if b.Title == book.Title && b.Author == book.Author && b.PublicationYear == book.PublicationYear {
			return apperror.NewConflict("book", "title, author, and publication year")
		}
	}

	book.ID = uuid.New()
	r.books = append(r.books, *book)

	return nil
}

func (r *InMemoryBookRepository) UpdateBook(ctx context.Context, id string, book *domain.Book) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, b := range r.books {
		if b.ID.String() == id {
			book.ID = b.ID
			r.books[i] = *book
			return nil
		}
	}

	return apperror.NewNotFound("Book", "ID", id)
}

func (r *InMemoryBookRepository) DeleteBook(ctx context.Context, id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i, book := range r.books {
		if book.ID.String() == id {
			r.books = append(r.books[:i], r.books[i+1:]...)
			return nil
		}
	}

	return apperror.NewNotFound("Book", "ID", id)
}
