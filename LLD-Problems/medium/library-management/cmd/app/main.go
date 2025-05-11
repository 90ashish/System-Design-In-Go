package main

import (
	"fmt"
	"library-management/internal/library"
)

func main() {
	libraryManager := library.GetLibraryManager()

	// Add books to the catalog
	book1 := library.NewBook("ISBN1", "Book 1", "Author 1", 2020)
	book2 := library.NewBook("ISBN2", "Book 2", "Author 2", 2019)
	book3 := library.NewBook("ISBN3", "Book 3", "Author 3", 2021)

	libraryManager.AddBook(book1)
	libraryManager.AddBook(book2)
	libraryManager.AddBook(book3)

	// Register members
	member1 := library.NewMember("M1", "John Doe", "john@example.com")
	member2 := library.NewMember("M2", "Jane Smith", "jane@example.com")

	libraryManager.RegisterMember(member1)
	libraryManager.RegisterMember(member2)

	// Borrow books
	if err := libraryManager.BorrowBook("M1", "ISBN1"); err != nil {
		fmt.Printf("Error borrowing book: %v\n", err)
	}

	if err := libraryManager.BorrowBook("M2", "ISBN2"); err != nil {
		fmt.Printf("Error borrowing book: %v\n", err)
	}

	// Return books
	if err := libraryManager.ReturnBook("M1", "ISBN1"); err != nil {
		fmt.Printf("Error returning book: %v\n", err)
	}

	// Search books
	searchResults := libraryManager.SearchBooks("Book")
	fmt.Println("\nSearch Results:")
	for _, book := range searchResults {
		fmt.Printf("%s by %s\n", book.Title, book.Author)
	}
}
