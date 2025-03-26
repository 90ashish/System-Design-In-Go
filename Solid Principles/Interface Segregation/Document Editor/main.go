package main

import "fmt"

// ------------------ Segregated Interfaces ------------------

type Editable interface {
	Edit(content string)
}

type Printable interface {
	Print()
}

type Scannable interface {
	Scan()
}

// ------------------ Concrete Document Types ------------------

// TextDocument supports only editing.
type TextDocument struct {
	Content string
}

func (t *TextDocument) Edit(content string) {
	t.Content = content
	fmt.Println("TextDocument edited:", t.Content)
}

// PrintableDocument supports only printing.
type PrintableDocument struct {
	Content string
}

func (p *PrintableDocument) Print() {
	fmt.Println("PrintableDocument printed:", p.Content)
}

// ScannableDocument supports only scanning.
type ScannableDocument struct{}

func (s *ScannableDocument) Scan() {
	fmt.Println("ScannableDocument scanned")
}

// MultiFunctionalDocument supports editing, printing, and scanning.
type MultiFunctionalDocument struct {
	Content string
}

func (m *MultiFunctionalDocument) Edit(content string) {
	m.Content = content
	fmt.Println("MultiFunctionalDocument edited:", m.Content)
}

func (m *MultiFunctionalDocument) Print() {
	fmt.Println("MultiFunctionalDocument printed:", m.Content)
}

func (m *MultiFunctionalDocument) Scan() {
	fmt.Println("MultiFunctionalDocument scanned")
}

// ------------------ Main Demonstration ------------------

func main() {
	// Use TextDocument with Editable interface
	var editable Editable = &TextDocument{}
	editable.Edit("New article draft")

	// Use PrintableDocument with Printable interface
	var printable Printable = &PrintableDocument{Content: "Invoice PDF"}
	printable.Print()

	// Use ScannableDocument with Scannable interface
	var scannable Scannable = &ScannableDocument{}
	scannable.Scan()

	// Use MultiFunctionalDocument
	multiDoc := &MultiFunctionalDocument{}
	multiDoc.Edit("Annual Report")
	multiDoc.Print()
	multiDoc.Scan()
}
