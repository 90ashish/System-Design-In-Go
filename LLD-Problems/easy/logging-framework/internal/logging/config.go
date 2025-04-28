package logging

// LoggerConfig holds level & subscribed appenders.
type LoggerConfig struct {
	Level     LogLevel
	Appenders []Appender
}

// ConfigBuilder implements the Builder pattern.
type ConfigBuilder struct {
	cfg LoggerConfig
}

// NewConfigBuilder starts with INFO and no appenders.
func NewConfigBuilder() *ConfigBuilder {
	return &ConfigBuilder{cfg: LoggerConfig{Level: INFO}}
}

func (b *ConfigBuilder) Level(l LogLevel) *ConfigBuilder {
	b.cfg.Level = l
	return b
}

func (b *ConfigBuilder) AddAppender(a Appender) *ConfigBuilder {
	b.cfg.Appenders = append(b.cfg.Appenders, a)
	return b
}

func (b *ConfigBuilder) Build() LoggerConfig {
	return b.cfg
}
