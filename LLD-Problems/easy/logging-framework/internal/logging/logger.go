package logging

import (
	"os"
	"sync"
	"time"
)

// Logger is a thread-safe singleton.
type Logger struct {
	mu     sync.RWMutex
	config LoggerConfig
}

var (
	instance *Logger
	once     sync.Once
)

// GetLogger returns the singleton instance.
func GetLogger() *Logger {
	once.Do(func() {
		instance = &Logger{config: NewConfigBuilder().Build()}
	})
	return instance
}

// Configure replaces the logger’s config.
func (l *Logger) Configure(cfg LoggerConfig) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.config = cfg
}

// Log checks level, builds a LogMessage, then dispatches to each Appender.
func (l *Logger) Log(level LogLevel, msg string) {
	l.mu.RLock()
	cfg := l.config
	l.mu.RUnlock()

	if level < cfg.Level {
		return
	}
	logMsg := LogMessage{Timestamp: time.Now(), Level: level, Message: msg}
	for _, app := range cfg.Appenders {
		// synchronous dispatch for ordering & simplicity
		app.Append(logMsg)
	}
}

// Convenience methods:
func (l *Logger) Debug(msg string)   { l.Log(DEBUG, msg) }
func (l *Logger) Info(msg string)    { l.Log(INFO, msg) }
func (l *Logger) Warning(msg string) { l.Log(WARNING, msg) }
func (l *Logger) Error(msg string)   { l.Log(ERROR, msg) }
func (l *Logger) Fatal(msg string) {
	l.Log(FATAL, msg)
	os.Exit(1)
}
