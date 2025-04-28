package appenders

import (
	"fmt"
	"logger/internal/logging"
	"sync"
)

// ConsoleAppender writes to stdout.
type ConsoleAppender struct {
	mu sync.Mutex
}

func NewConsoleAppender() *ConsoleAppender {
	return &ConsoleAppender{}
}

func (c *ConsoleAppender) Append(m logging.LogMessage) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	fmt.Println(m.String())
	return nil
}
