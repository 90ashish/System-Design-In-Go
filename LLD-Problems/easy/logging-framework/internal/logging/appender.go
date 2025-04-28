package logging

// Appender is the Strategy / Observer interface.
type Appender interface {
	Append(LogMessage) error
}
