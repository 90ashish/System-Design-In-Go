package library

import "sync"

// Book represents a book in the library.
type Book struct {
	ISBN            string
	Title           string
	Author          string
	PublicationYear int
	available       bool
	mu              sync.RWMutex
}

// NewBook creates a new book instance.
func NewBook(isbn, title, author string, publicationYear int) *Book {
	return &Book{
		ISBN:            isbn,
		Title:           title,
		Author:          author,
		PublicationYear: publicationYear,
		available:       true,
	}
}

// IsAvailable checks if the book is available for borrowing.
func (b *Book) IsAvailable() bool {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.available
}

// SetAvailable sets the availability of the book.
func (b *Book) SetAvailable(available bool) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.available = available
}
