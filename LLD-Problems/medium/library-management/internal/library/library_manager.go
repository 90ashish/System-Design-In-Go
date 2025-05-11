package library

import (
	"fmt"
	"strings"
	"sync"
)

const (
	maxBooksPerMember = 5
	loanDurationDays  = 14
)

// LibraryManager is a singleton that manages the library's catalog and members.
type LibraryManager struct {
	catalog map[string]*Book
	members map[string]*Member
	mu      sync.RWMutex
}

// Singleton instance of LibraryManager.
var (
	instance *LibraryManager
	once     sync.Once
)

// GetLibraryManager returns the singleton instance of LibraryManager.
func GetLibraryManager() *LibraryManager {
	once.Do(func() {
		instance = &LibraryManager{
			catalog: make(map[string]*Book),
			members: make(map[string]*Member),
		}
	})
	return instance
}

// AddBook adds a book to the library's catalog.
func (lm *LibraryManager) AddBook(book *Book) {
	lm.mu.Lock()
	defer lm.mu.Unlock()
	lm.catalog[book.ISBN] = book
}

// RemoveBook removes a book from the library's catalog.
func (lm *LibraryManager) RemoveBook(isbn string) {
	lm.mu.Lock()
	defer lm.mu.Unlock()
	delete(lm.catalog, isbn)
}

// GetBook retrieves a book from the library's catalog by its ISBN.
func (lm *LibraryManager) GetBook(isbn string) *Book {
	lm.mu.RLock()
	defer lm.mu.RUnlock()
	return lm.catalog[isbn]
}

// RegisterMember registers a new member in the library.
func (lm *LibraryManager) RegisterMember(member *Member) {
	lm.mu.Lock()
	defer lm.mu.Unlock()
	lm.members[member.MemberID] = member
}

// UnregisterMember removes a member from the library.
func (lm *LibraryManager) UnregisterMember(memberID string) {
	lm.mu.Lock()
	defer lm.mu.Unlock()
	delete(lm.members, memberID)
}

// GetMember retrieves a member from the library by their ID.
func (lm *LibraryManager) GetMember(memberID string) *Member {
	lm.mu.RLock()
	defer lm.mu.RUnlock()
	return lm.members[memberID]
}

// BorrowBook allows a member to borrow a book from the library.
func (lm *LibraryManager) BorrowBook(memberID, isbn string) error {
	lm.mu.Lock()
	defer lm.mu.Unlock()

	member := lm.members[memberID]
	book := lm.catalog[isbn]

	if member == nil || book == nil {
		return fmt.Errorf("book or member not found")
	}

	if !book.IsAvailable() {
		return fmt.Errorf("book is not available")
	}

	if len(member.GetBorrowedBooks()) >= maxBooksPerMember {
		return fmt.Errorf("member %s has reached the maximum number of borrowed books", member.Name)
	}

	member.BorrowBook(book)
	book.SetAvailable(false)
	fmt.Printf("Book borrowed: %s by %s\n", book.Title, member.Name)
	return nil
}

// ReturnBook allows a member to return a borrowed book to the library.
func (lm *LibraryManager) ReturnBook(memberID, isbn string) error {
	lm.mu.Lock()
	defer lm.mu.Unlock()

	member := lm.members[memberID]
	book := lm.catalog[isbn]

	if member == nil || book == nil {
		return fmt.Errorf("book or member not found")
	}

	member.ReturnBook(book)
	book.SetAvailable(true)
	fmt.Printf("Book returned: %s by %s\n", book.Title, member.Name)
	return nil
}

// SearchBooks searches for books in the library's catalog by title or author.
func (lm *LibraryManager) SearchBooks(keyword string) []*Book {
	lm.mu.RLock()
	defer lm.mu.RUnlock()

	keyword = strings.ToLower(keyword)
	matchingBooks := make([]*Book, 0)

	for _, book := range lm.catalog {
		if strings.Contains(strings.ToLower(book.Title), keyword) ||
			strings.Contains(strings.ToLower(book.Author), keyword) {
			matchingBooks = append(matchingBooks, book)
		}
	}
	return matchingBooks
}
