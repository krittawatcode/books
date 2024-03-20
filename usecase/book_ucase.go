package usecase

import (
	"context"

	"github.com/krittawatcode/books/domain"
)

type bookUseCase struct {
	bookRepository domain.BookRepository
}

func NewBookUseCase(bookRepository domain.BookRepository) domain.BookUseCase {
	return &bookUseCase{
		bookRepository: bookRepository,
	}
}

func (b *bookUseCase) FetchBooks(ctx context.Context) (*[]domain.Book, error) {
	return b.bookRepository.FetchBooks(ctx)
}

func (b *bookUseCase) GetBookByID(ctx context.Context, id string) (*domain.Book, error) {
	return b.bookRepository.GetBookByID(ctx, id)
}

func (b *bookUseCase) CreateBook(ctx context.Context, book *domain.Book) error {
	return b.bookRepository.CreateBook(ctx, book)
}

func (b *bookUseCase) UpdateBook(ctx context.Context, id string, book *domain.Book) error {
	return b.bookRepository.UpdateBook(ctx, id, book)
}

func (b *bookUseCase) DeleteBook(ctx context.Context, id string) error {
	return b.bookRepository.DeleteBook(ctx, id)
}
