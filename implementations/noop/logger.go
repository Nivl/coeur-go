package noop

import (
	"context"

	"github.com/Nivl/coeur/logger"
)

// Logger is a logger that does nothing
type Logger struct{}

// NewLogger returns a new Logger that does nothing
func NewLogger() *Logger {
	return &Logger{}
}

// We make sure the struct implements the interface.
var _ logger.Logger = (*Logger)(nil)

// Debug does nothing
func (l *Logger) Debug(ctx context.Context, msg string, extras ...map[string]interface{}) error {
	return nil
}

// Info does nothing
func (l *Logger) Info(ctx context.Context, msg string, extras ...map[string]interface{}) error {
	return nil
}

// Warning does nothing
func (l *Logger) Warning(ctx context.Context, msg string, extras ...map[string]interface{}) error {
	return nil
}

// Error does nothing
func (l *Logger) Error(ctx context.Context, msg string, extras ...map[string]interface{}) error {
	return nil
}

// Close does nothing
func (l *Logger) Close() error {
	return nil
}

// SetExtra does nothing
func (l *Logger) SetExtra(extras map[string]interface{}) {
}

// AddExtra does nothing
func (l *Logger) AddExtra(k string, v interface{}) {
}

// RemoveExtra does nothing
func (l *Logger) RemoveExtra(k string) {
}

// NewChild return a the same noop logger
func (l *Logger) NewChild() logger.Logger {
	return l
}
