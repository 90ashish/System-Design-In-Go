package atmsystem

// CardReader abstracts reading card number and PIN (Strategy Pattern).
type CardReader interface {
	ReadCard() (string, error)
	ReadPIN() (string, error)
}

// ConsoleCardReader reads from the console.
type ConsoleCardReader struct {
	ui UI
}

// NewConsoleCardReader injects UI into the reader (DIP).
func NewConsoleCardReader(ui UI) CardReader {
	return &ConsoleCardReader{ui}
}

// ReadCard prompts for a card number.
func (cr *ConsoleCardReader) ReadCard() (string, error) {
	return cr.ui.Prompt("Insert card (enter card number): ")
}

// ReadPIN prompts for the PIN.
func (cr *ConsoleCardReader) ReadPIN() (string, error) {
	return cr.ui.Prompt("Enter PIN: ")
}
