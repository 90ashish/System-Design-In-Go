package appenders

import (
	"logger/internal/logging"
	"os"
	"sync"
)

// FileAppender writes logs to a file.
type FileAppender struct {
	mu   sync.Mutex
	file *os.File
}

// NewFileAppender opens (or creates) the given path.
func NewFileAppender(path string) (*FileAppender, error) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return &FileAppender{file: f}, nil
}

func (fa *FileAppender) Append(m logging.LogMessage) error {
	fa.mu.Lock()
	defer fa.mu.Unlock()
	_, err := fa.file.WriteString(m.String() + "\n")
	return err
}
