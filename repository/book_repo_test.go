package repository

import (
	"context"
	"sync"
	"testing"

	"github.com/krittawatcode/books/domain"
	"github.com/stretchr/testify/assert"
)

func TestFetchBooks(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		repo := NewInMemoryBookRepository()

		// Add a book to the repository
		book := &domain.Book{Title: "Test Book", Author: "Test Author"}
		err := repo.CreateBook(context.Background(), book)
		assert.Nil(t, err)

		// Fetch books from the repository
		books, err := repo.FetchBooks(context.Background())
		assert.Nil(t, err)

		// Check that the fetched books include the added book
		assert.Contains(t, *books, *book)
	})

	t.Run("Success - Empty arr", func(t *testing.T) {
		repo := NewInMemoryBookRepository()

		// Fetch books from the repository
		books, err := repo.FetchBooks(context.Background())

		// Check that an error is returned and that no books are fetched
		assert.Error(t, err)
		assert.Nil(t, books)
	})
}

func TestGetBookByID(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		repo := NewInMemoryBookRepository()
		book := &domain.Book{Title: "Test Book", Author: "Test Author"}
		err := repo.CreateBook(context.Background(), book)
		assert.Nil(t, err)

		fetchedBook, err := repo.GetBookByID(context.Background(), book.ID.String())
		assert.Nil(t, err)
		assert.Equal(t, book, fetchedBook)
	})

	t.Run("Failure - Book not found", func(t *testing.T) {
		repo := NewInMemoryBookRepository()

		// Try to fetch a book with an ID that doesn't exist in the repository
		fetchedBook, err := repo.GetBookByID(context.Background(), "nonexistent-id")
		assert.Error(t, err)
		assert.Nil(t, fetchedBook)
	})
}

func TestCreateBook(t *testing.T) {
	t.Run("Failure - Book already exists", func(t *testing.T) {
		repo := NewInMemoryBookRepository()
		book := &domain.Book{Title: "Test Book", Author: "Test Author", PublicationYear: "2021"}
		err := repo.CreateBook(context.Background(), book)
		assert.Nil(t, err)

		// Try to create the same book again
		err = repo.CreateBook(context.Background(), book)
		assert.Error(t, err)
	})
	t.Run("Concurrent calls", func(t *testing.T) {
		repo := NewInMemoryBookRepository()
		book := &domain.Book{Title: "Test Book", Author: "Test Author", PublicationYear: "2021"}

		var wg sync.WaitGroup
		errs := make(chan error)

		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				err := repo.CreateBook(context.Background(), book)
				errs <- err
			}()
		}

		go func() {
			wg.Wait()
			close(errs)
		}()

		var errCount int
		for err := range errs {
			if err != nil {
				errCount++
			}
		}

		// Since the same book cannot be created twice, there should be 9 errors
		assert.Equal(t, 9, errCount)
	})

	t.Run("Success", func(t *testing.T) {
		repo := NewInMemoryBookRepository()
		book := &domain.Book{Title: "Test Book", Author: "Test Author", PublicationYear: "2021"}
		err := repo.CreateBook(context.Background(), book)
		assert.Nil(t, err)

		createdBook, err := repo.GetBookByID(context.Background(), book.ID.String())
		assert.Nil(t, err)
		assert.Equal(t, book, createdBook)
	})
}

func TestUpdateBook(t *testing.T) {
	t.Run("Failure - Book not found", func(t *testing.T) {
		repo := NewInMemoryBookRepository()

		// Try to update a book with an ID that doesn't exist in the repository
		book := &domain.Book{Title: "Nonexistent Book", Author: "Test Author"}
		err := repo.UpdateBook(context.Background(), "nonexistent-id", book)
		assert.Error(t, err)
	})

	t.Run("Success - Concurrent calls", func(t *testing.T) {
		repo := NewInMemoryBookRepository()
		book := &domain.Book{Title: "Test Book", Author: "Test Author"}
		err := repo.CreateBook(context.Background(), book)
		assert.Nil(t, err)

		var wg sync.WaitGroup
		errs := make(chan error)

		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				book.Title = "Updated Test Book"
				err := repo.UpdateBook(context.Background(), book.ID.String(), book)
				errs <- err
			}()
		}

		go func() {
			wg.Wait()
			close(errs)
		}()

		for err := range errs {
			assert.Nil(t, err)
		}

		updatedBook, err := repo.GetBookByID(context.Background(), book.ID.String())
		assert.Nil(t, err)
		assert.Equal(t, "Updated Test Book", updatedBook.Title)
	})

	t.Run("Success", func(t *testing.T) {
		repo := NewInMemoryBookRepository()
		book := &domain.Book{Title: "Test Book", Author: "Test Author"}
		err := repo.CreateBook(context.Background(), book)
		assert.Nil(t, err)

		book.Title = "Updated Test Book"
		err = repo.UpdateBook(context.Background(), book.ID.String(), book)
		assert.Nil(t, err)

		updatedBook, err := repo.GetBookByID(context.Background(), book.ID.String())
		assert.Nil(t, err)
		assert.Equal(t, "Updated Test Book", updatedBook.Title)
	})

}

func TestDeleteBook(t *testing.T) {
	t.Run("Failure - Book not found", func(t *testing.T) {
		repo := NewInMemoryBookRepository()

		// Try to delete a book with an ID that doesn't exist in the repository
		err := repo.DeleteBook(context.Background(), "nonexistent-id")
		assert.Error(t, err)
	})

	t.Run("Success - Concurrent calls", func(t *testing.T) {
		repo := NewInMemoryBookRepository()
		book := &domain.Book{Title: "Test Book", Author: "Test Author"}
		err := repo.CreateBook(context.Background(), book)
		assert.Nil(t, err)

		var wg sync.WaitGroup
		errs := make(chan error)

		for i := 0; i < 10; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				err := repo.DeleteBook(context.Background(), book.ID.String())
				errs <- err
			}()
		}

		go func() {
			wg.Wait()
			close(errs)
		}()

		for err := range errs {
			if err != nil {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		}
	})

	t.Run("Success", func(t *testing.T) {
		repo := NewInMemoryBookRepository()
		book := &domain.Book{Title: "Test Book", Author: "Test Author"}
		err := repo.CreateBook(context.Background(), book)
		assert.Nil(t, err)

		err = repo.DeleteBook(context.Background(), book.ID.String())
		assert.Nil(t, err)

		deletedBook, err := repo.GetBookByID(context.Background(), book.ID.String())
		assert.Nil(t, deletedBook)
		assert.Error(t, err)
	})
}
