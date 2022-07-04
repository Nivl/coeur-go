// Package logger contain a logger interface and implementation
package logger

import "context"

// gomock interface, requires mockgen
// Update with "go generate github.com/Nivl/coeur/logger"
//go:generate mockgen -destination ../implementations/mock/logger.go -package mock github.com/Nivl/coeur/logger Logger

// Logger is an interface used to log data
type Logger interface {
	// Debug logs a message that is intended for use in a development
	// environment while actively debugging your subsystem, not in shipping
	// software
	Debug(ctx context.Context, msg string, extras ...map[string]interface{}) error
	// Info logs a message that may be helpful, but isn’t essential,
	// for troubleshooting
	Info(ctx context.Context, msg string, extras ...map[string]interface{}) error
	// Warning logs a message that warns that something potentially wrong
	// or suspect may be happening.
	Warning(ctx context.Context, msg string, extras ...map[string]interface{}) error
	// Error logs an error message
	Error(ctx context.Context, msg string, extras ...map[string]interface{}) error

	// Close flushes and frees any resource allocated by the logger.
	// The logger may not be reusable after being closed
	Close() error

	// SetExtra replaces all the global extras with the provided one
	SetExtra(extras map[string]interface{})
	// AddExtra adds a new k/v to the extras
	AddExtra(k string, v interface{})
	// RemoveExtra removes a key from the extras
	RemoveExtra(k string)

	// NewChild return a new logger that copies over the global extras
	NewChild() Logger
}
