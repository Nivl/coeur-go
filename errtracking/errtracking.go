// Package errtracking contains struct and interfaces to track errors
package errtracking

import "context"

// gomock interface, requires mockgen
// Update with "go generate github.com/Nivl/coeur/errtracking"
//go:generate mockgen -destination ../implementations/mock/errtracking.go -package mock github.com/Nivl/coeur/errtracking Tracker

// User represents a user
type User struct {
	ID    string
	Email string
	Name  string
}

// Tracker is an interface used to track errors
// You should use a different Tacker per goroutine by using NewChild()
type Tracker interface {
	// ReportError asynchronously sends an error to the remote tracker
	ReportError(ctx context.Context, err error, extras ...map[string]interface{})
	// Close flushes and frees any resource allocated by the tracker.
	// The tracker may not be reusable after being closed.
	Close() error

	// SetUser sets the information for identifying a user associated
	// with any errors or messages. Only the ID is required.
	SetUser(context.Context, User)
	// RemoveUser Remove any set users
	RemoveUser(context.Context)

	// AddTag attaches a tag to all errors if the remote supports it.
	// An extra is attached otherwise.
	AddTag(key, value string)
	// AddTags attaches multiple tags to all errors if the remote supports it.
	// Multiple extras are attached otherwise.
	AddTags(tags map[string]string)
	// RemoveTag remove the given tag if the remote supports it.
	// An extra is removed otherwise.
	RemoveTag(key string)

	// AddExtra attaches an extra information to all errors.
	AddExtra(key string, value interface{})
	// AddExtras attaches multiple extra information to all errors.
	AddExtras(tags map[string]interface{})
	// RemoveExtra remove an extra from being attached to errors.
	RemoveExtra(key string)

	// NewChild return a new Tracker that copies over the extras and tags
	NewChild() Tracker
}
