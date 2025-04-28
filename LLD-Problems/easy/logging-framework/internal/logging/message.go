package logging

import (
	"fmt"
	"time"
)

// LogMessage carries timestamp, level & content.
type LogMessage struct {
	Timestamp time.Time
	Level     LogLevel
	Message   string
}

func (m LogMessage) String() string {
	return fmt.Sprintf("%s [%s] %s",
		m.Timestamp.Format(time.RFC3339),
		m.Level,
		m.Message)
}
