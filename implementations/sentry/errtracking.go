package sentry

import (
	"context"
	"os"
	"time"

	"github.com/Nivl/coeur/errtracking"
	"github.com/Nivl/coeur/secret"
	sentry "github.com/getsentry/sentry-go"
)

// Tracker represents a sentry implementation of the Tracker interface
// You should use a different per goroutine
type Tracker struct {
	hub *sentry.Hub
}

// we make sure the struct implements the interface
var _ errtracking.Tracker = (*Tracker)(nil)

// NewTracker returns a new Tracker using sentry as backend,
// os.Getenv("SENTRY_DSN") will be used
func NewTracker() (errtracking.Tracker, error) {
	dsn := secret.NewString(os.Getenv("SENTRY_DSN"))
	return NewTrackerWithDSN(dsn)
}

// NewTrackerWithDSN returns a new Tracker using sentry as backend
func NewTrackerWithDSN(dsn secret.String) (errtracking.Tracker, error) {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: dsn.Value(),
	})
	if err != nil {
		return nil, err
	}

	return &Tracker{
		hub: sentry.CurrentHub(),
	}, nil
}

// ReportError asynchronously sends an error to the remote tracker
func (t *Tracker) ReportError(ctx context.Context, err error, extras ...map[string]interface{}) {
	t.hub.WithScope(func(scope *sentry.Scope) {
		for _, ex := range extras {
			scope.SetExtras(ex)
		}
		sentry.CaptureException(err)
	})
}

// Close flushes and frees any resource allocated by the tracker.
// The tracker may not be reusable after being closed.
func (t *Tracker) Close() error {
	if timedOut := t.hub.Flush(5 * time.Second); timedOut {
		return os.ErrDeadlineExceeded
	}
	return nil
}

// SetUser sets the information for identifying a user associated
// with any errors or messages. Only the ID is required.
func (t *Tracker) SetUser(ctx context.Context, u errtracking.User) {
	t.hub.ConfigureScope(func(scope *sentry.Scope) {
		scope.SetUser(sentry.User{
			ID:       u.ID,
			Username: u.Name,
			Email:    u.Email,
		})
	})
}

// RemoveUser Removes any set users
func (t *Tracker) RemoveUser(ctx context.Context) {
	t.hub.ConfigureScope(func(scope *sentry.Scope) {
		scope.SetUser(sentry.User{})
	})
}

// AddTag attaches a tag to all errors if the remote supports it.
// An extra is attached otherwise.
func (t *Tracker) AddTag(k, v string) {
	t.hub.ConfigureScope(func(scope *sentry.Scope) {
		scope.SetTag(k, v)
	})
}

// AddTags attaches multiple tags to all errors if the remote supports it.
// Multiple extras are attached otherwise.
func (t *Tracker) AddTags(tags map[string]string) {
	t.hub.ConfigureScope(func(scope *sentry.Scope) {
		scope.SetTags(tags)
	})
}

// RemoveTag remove the given tag if the remote supports it.
// An extra is removed otherwise.
func (t *Tracker) RemoveTag(k string) {
	t.hub.ConfigureScope(func(scope *sentry.Scope) {
		scope.RemoveTag(k)
	})
}

// AddExtra attaches an extra information to all errors.
func (t *Tracker) AddExtra(k string, v interface{}) {
	t.hub.ConfigureScope(func(scope *sentry.Scope) {
		scope.SetExtra(k, v)
	})
}

// AddExtras attaches multiple extra information to all errors.
func (t *Tracker) AddExtras(tags map[string]interface{}) {
	t.hub.ConfigureScope(func(scope *sentry.Scope) {
		scope.SetExtras(tags)
	})
}

// RemoveExtra remove an extra from being attached to errors.
func (t *Tracker) RemoveExtra(k string) {
	t.hub.ConfigureScope(func(scope *sentry.Scope) {
		scope.RemoveExtra(k)
	})
}

// NewChild return a new Tracker that copies over the extras and tags
func (t *Tracker) NewChild() errtracking.Tracker {
	return &Tracker{
		hub: t.hub.Clone(),
	}
}
