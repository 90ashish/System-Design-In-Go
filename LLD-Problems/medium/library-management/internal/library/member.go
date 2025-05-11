package library

import "sync"

// Member represents a library member.
// It contains member details and a list of borrowed books.
// The borrowed books are stored in a map for O(1) lookups.
// The map key is the book's ISBN, and the value is a pointer to the Book struct.
type Member struct {
	MemberID      string
	Name          string
	ContactInfo   string
	borrowedBooks map[string]*Book // Using map for O(1) lookups
	mu            sync.RWMutex
}

// NewMember creates a new member instance.
func NewMember(memberID, name, contactInfo string) *Member {
	return &Member{
		MemberID:      memberID,
		Name:          name,
		ContactInfo:   contactInfo,
		borrowedBooks: make(map[string]*Book),
	}
}

// BorrowBook adds a book to the borrowed books list.
func (m *Member) BorrowBook(book *Book) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.borrowedBooks[book.ISBN] = book
}

// ReturnBook removes a book from the borrowed books list.
func (m *Member) ReturnBook(book *Book) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.borrowedBooks, book.ISBN)
}

// GetBorrowedBooks returns a slice of borrowed books.
func (m *Member) GetBorrowedBooks() []*Book {
	m.mu.RLock()
	defer m.mu.RUnlock()
	books := make([]*Book, 0, len(m.borrowedBooks))
	for _, book := range m.borrowedBooks {
		books = append(books, book)
	}
	return books
}
