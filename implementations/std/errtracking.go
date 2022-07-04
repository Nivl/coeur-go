package std

import (
	"context"
	"log"

	"github.com/Nivl/coeur/errtracking"
)

// Tracker is an implementation of the Tracker interface that prints
// on stderr
type Tracker struct {
	extras map[string]interface{}
	user   errtracking.User
}

// we make sure the struct implements the interface
var _ errtracking.Tracker = (*Tracker)(nil)

// NewTracker returns a new Tracker using sentry as backend
func NewTracker(dsn string) (errtracking.Tracker, error) {
	return &Tracker{
		extras: map[string]interface{}{},
	}, nil
}

// ReportError asynchronously sends an error to the remote tracker
func (t *Tracker) ReportError(ctx context.Context, err error, extras ...map[string]interface{}) {
	extra := map[string]interface{}{}
	for k, v := range t.extras {
		extra[k] = v
	}
	for _, ex := range extras {
		for k, v := range ex {
			extra[k] = v
		}
	}
	log.Printf("[ERROR] %s ; user %#v ; extra %#v", err.Error(), t.user, extra)
}

// Close flushes and frees any resource allocated by the tracker.
// The tracker may not be reusable after being closed.
func (t *Tracker) Close() error {
	return nil
}

// SetUser sets the information for identifying a user associated
// with any errors or messages. Only the ID is required.
func (t *Tracker) SetUser(ctx context.Context, u errtracking.User) {
	t.user = u
}

// RemoveUser Removes any set users
func (t *Tracker) RemoveUser(ctx context.Context) {
	t.user = errtracking.User{}
}

// AddTag attaches a tag to all errors if the remote supports it.
// An extra is attached otherwise.
func (t *Tracker) AddTag(k, v string) {
	t.extras[k] = v
}

// AddTags attaches multiple tags to all errors if the remote supports it.
// Multiple extras are attached otherwise.
func (t *Tracker) AddTags(tags map[string]string) {
	for k, v := range tags {
		t.extras[k] = v
	}
}

// RemoveTag remove the given tag if the remote supports it.
// An extra is removed otherwise.
func (t *Tracker) RemoveTag(k string) {
	delete(t.extras, k)
}

// AddExtra attaches an extra information to all errors.
func (t *Tracker) AddExtra(k string, v interface{}) {
	t.extras[k] = v
}

// AddExtras attaches multiple extra information to all errors.
func (t *Tracker) AddExtras(extra map[string]interface{}) {
	for k, v := range extra {
		t.extras[k] = v
	}
}

// RemoveExtra remove an extra from being attached to errors.
func (t *Tracker) RemoveExtra(k string) {
	delete(t.extras, k)
}

// NewChild return a new Tracker that copies over the extras and tags
func (t *Tracker) NewChild() errtracking.Tracker {
	extra := map[string]interface{}{}
	for k, v := range t.extras {
		extra[k] = v
	}

	return &Tracker{
		extras: extra,
	}
}
