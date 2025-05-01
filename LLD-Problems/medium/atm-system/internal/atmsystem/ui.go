package atmsystem

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// UI defines how we interact with the user (Strategy Pattern).
type UI interface {
	Display(message string)
	Prompt(prompt string) (string, error)
}

// ConsoleUI is a simple terminal‐based UI implementation.
type ConsoleUI struct {
	reader *bufio.Reader
}

// NewConsoleUI constructs a ConsoleUI (Factory Pattern).
func NewConsoleUI() UI {
	return &ConsoleUI{reader: bufio.NewReader(os.Stdin)}
}

// Display prints a message to the console.
func (c *ConsoleUI) Display(message string) {
	fmt.Println(message)
}

// Prompt shows a prompt and reads one line of input.
func (c *ConsoleUI) Prompt(prompt string) (string, error) {
	fmt.Print(prompt)
	input, err := c.reader.ReadString('\n')
	return strings.TrimSpace(input), err
}
