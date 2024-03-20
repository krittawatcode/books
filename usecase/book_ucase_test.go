package usecase

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/krittawatcode/books/domain"
	"github.com/krittawatcode/books/domain/appmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFetchBooks(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockBookRepo := new(appmock.MockBookRepository)
		mockBooks := &[]domain.Book{
			{Title: "Test Book 1", Author: "Test Author 1", PublicationYear: "2021"},
			{Title: "Test Book 2", Author: "Test Author 2", PublicationYear: "2022"},
		}

		mockBookRepo.On("FetchBooks", mock.Anything).Return(mockBooks, nil).Once()

		u := NewBookUseCase(mockBookRepo)

		// Call the FetchBooks method on the use case
		books, err := u.FetchBooks(context.Background())

		assert.NoError(t, err)
		assert.NotNil(t, books)
		assert.Equal(t, len(*mockBooks), len(*books))
		mockBookRepo.AssertExpectations(t)
	})
}

func TestCreateBook(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockBookRepo := new(appmock.MockBookRepository)
		mockBook := &domain.Book{Title: "Test Book", Author: "Test Author", PublicationYear: "2021"}

		mockBookRepo.On("CreateBook", mock.Anything, mock.AnythingOfType("*domain.Book")).Return(nil).Once()

		u := NewBookUseCase(mockBookRepo)

		// Call the CreateBook method on the use case
		err := u.CreateBook(context.Background(), mockBook)

		assert.NoError(t, err)
		mockBookRepo.AssertExpectations(t)
	})
}

func TestGetBookByID(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockBookRepo := new(appmock.MockBookRepository)
		id := uuid.New()
		mockBook := &domain.Book{ID: id, Title: "Test Book", Author: "Test Author", PublicationYear: "2021"}

		mockBookRepo.On("GetBookByID", mock.Anything, mock.Anything).Return(mockBook, nil)

		u := NewBookUseCase(mockBookRepo)

		// Call the GetBookByID method on the use case
		book, err := u.GetBookByID(context.Background(), id.String())

		assert.NoError(t, err)
		assert.NotNil(t, book)
		assert.Equal(t, mockBook.ID, book.ID)
		mockBookRepo.AssertExpectations(t)
	})
}

func TestUpdateBook(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockBookRepo := new(appmock.MockBookRepository)
		id := uuid.New()
		mockBook := &domain.Book{ID: id, Title: "Updated Book", Author: "Updated Author", PublicationYear: "2022"}

		mockBookRepo.On("UpdateBook", mock.Anything, mock.AnythingOfType("*domain.Book")).Return(nil).Once()

		u := NewBookUseCase(mockBookRepo)

		// Call the UpdateBook method on the use case
		err := u.UpdateBook(context.Background(), id.String(), mockBook)

		assert.NoError(t, err)
		mockBookRepo.AssertExpectations(t)
	})
}

func TestDeleteBook(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		mockBookRepo := new(appmock.MockBookRepository)
		mockBookID := "1"

		mockBookRepo.On("DeleteBook", mock.Anything, mockBookID).Return(nil).Once()

		u := NewBookUseCase(mockBookRepo)

		// Call the DeleteBook method on the use case
		err := u.DeleteBook(context.Background(), mockBookID)

		assert.NoError(t, err)
		mockBookRepo.AssertExpectations(t)
	})
}
