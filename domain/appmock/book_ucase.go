package appmock

import (
	"context"

	"github.com/krittawatcode/books/domain"
	"github.com/stretchr/testify/mock"
)

type MockBookUseCase struct {
	mock.Mock
}

func (m *MockBookUseCase) FetchBooks(ctx context.Context) (*[]domain.Book, error) {
	args := m.Called(ctx)
	return args.Get(0).(*[]domain.Book), args.Error(1)
}

func (m *MockBookUseCase) GetBookByID(ctx context.Context, id string) (*domain.Book, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*domain.Book), args.Error(1)
}

func (m *MockBookUseCase) CreateBook(ctx context.Context, book *domain.Book) error {
	args := m.Called(ctx, book)
	return args.Error(0)
}

func (m *MockBookUseCase) UpdateBook(ctx context.Context, id string, book *domain.Book) error {
	args := m.Called(ctx, id, book)
	return args.Error(0)
}

func (m *MockBookUseCase) DeleteBook(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
