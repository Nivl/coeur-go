// Package std implements a logger using the standard error and output
package std

import (
	"context"
	"log"
	"sync"

	"github.com/Nivl/coeur/logger"
)

// Logger is a logger that uses the standard error to print the
// message.
type Logger struct {
	globalExtras map[string]interface{}
	mu           sync.RWMutex
}

// NewLogger returns a new Logger that uses Go's standard logger
func NewLogger() *Logger {
	return &Logger{ //nolint:exhaustruct // mu shouldn't be set manually
		globalExtras: map[string]interface{}{},
	}
}

// We make sure the struct implements the interface.
var _ logger.Logger = (*Logger)(nil)

// Debug logs a message that is intended for use in a development
// environment while actively debugging your subsystem, not in shipping
// software
func (l *Logger) Debug(ctx context.Context, msg string, extras ...map[string]interface{}) error {
	return l.log(ctx, "DEBUG", msg, extras)
}

// Info logs a message that may be helpful, but isn’t essential,
// for troubleshooting
func (l *Logger) Info(ctx context.Context, msg string, extras ...map[string]interface{}) error {
	return l.log(ctx, "INFO", msg, extras)
}

// Warning logs a message that warns that something potentially wrong
// or suspect may be happening.
func (l *Logger) Warning(ctx context.Context, msg string, extras ...map[string]interface{}) error {
	return l.log(ctx, "WARNING", msg, extras)
}

// Error logs an error message
func (l *Logger) Error(ctx context.Context, msg string, extras ...map[string]interface{}) error {
	return l.log(ctx, "ERROR", msg, extras)
}

func (l *Logger) log(ctx context.Context, tag, msg string, extras []map[string]interface{}) error {
	extra := l.mergeExtras(extras)
	log.Printf("[%s] %s %#v", tag, msg, extra)
	return nil
}

// Close is a noop
func (l *Logger) Close() error {
	return nil
}

// SetExtra replaces all the global extras with the provided one
func (l *Logger) SetExtra(extras map[string]interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.globalExtras = extras
}

// AddExtra adds a new k/v to the extra list
func (l *Logger) AddExtra(k string, v interface{}) {
	l.mu.Lock()
	defer l.mu.Unlock()

	l.globalExtras[k] = v
}

// RemoveExtra removes a key from the extras
func (l *Logger) RemoveExtra(k string) {
	l.mu.Lock()
	defer l.mu.Unlock()

	delete(l.globalExtras, k)
}

// NewChild return a new logger that copies over the global extras
func (l *Logger) NewChild() logger.Logger {
	l.mu.RLock()
	defer l.mu.RUnlock()

	newLogger := NewLogger()
	for k, v := range l.globalExtras {
		newLogger.globalExtras[k] = v
	}

	return newLogger
}

func (l *Logger) mergeExtras(extras []map[string]interface{}) map[string]interface{} {
	l.mu.RLock()
	extra := map[string]interface{}{}
	for k, v := range l.globalExtras {
		extra[k] = v
	}
	l.mu.RUnlock()

	for _, ex := range extras {
		for k, v := range ex {
			extra[k] = v
		}
	}
	return extra
}
