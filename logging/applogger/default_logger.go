package applogger

import "log"

// A Level is a logging priority. Higher levels are more important.
type Level int

const (
	// DebugLevel logs are typically voluminous, and are usually disabled in
	// production.
	DebugLevel Level = iota - 1
	// InfoLevel is the default logging priority.
	InfoLevel
	// WarnLevel logs are more important than Info, but don't need individual
	// human review.
	WarnLevel
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel
	// PanicLevel logs a message, then panics.
	PanicLevel
)

// DefaultLogger is the implementation for a Logger using golang log.
type DefaultLogger struct {
	level Level
}

func (l *DefaultLogger) Debugf(s string, i ...interface{}) {
	if l.level <= DebugLevel {
		log.Printf(s, i...)
	}
}

func (l *DefaultLogger) Infof(s string, i ...interface{}) {
	if l.level <= InfoLevel {
		log.Printf(s, i...)
	}
}

func (l *DefaultLogger) Warnf(s string, i ...interface{}) {
	if l.level <= WarnLevel {
		log.Printf(s, i...)
	}
}

func (l *DefaultLogger) Errorf(s string, i ...interface{}) {
	if l.level <= ErrorLevel {
		log.Printf(s, i...)
	}
}

func (l *DefaultLogger) Panicf(s string, i ...interface{}) {
	if l.level <= PanicLevel {
		log.Printf(s, i...)
	}
}
