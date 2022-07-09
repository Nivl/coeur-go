package noop

import (
	"context"

	"github.com/Nivl/coeur/errtracking"
)

// Tracker is an implementation of the Tracker interface that does nothing
type Tracker struct{}

// we make sure the struct implements the interface
var _ errtracking.Tracker = (*Tracker)(nil)

// NewTracker returns a new Tracker that does nothing
func NewTracker(dsn string) errtracking.Tracker {
	return &Tracker{}
}

// ReportError asynchronously sends an error to the remote tracker
func (t *Tracker) ReportError(ctx context.Context, err error, extras ...map[string]interface{}) {}

// Close does nothing
func (t *Tracker) Close() error {
	return nil
}

// SetUser does nothing
func (t *Tracker) SetUser(ctx context.Context, u errtracking.User) {
}

// RemoveUser does nothing
func (t *Tracker) RemoveUser(ctx context.Context) {
}

// AddTag does nothing
func (t *Tracker) AddTag(k, v string) {
}

// AddTags does nothing
func (t *Tracker) AddTags(tags map[string]string) {
}

// RemoveTag does nothing
func (t *Tracker) RemoveTag(k string) {
}

// AddExtra does nothing
func (t *Tracker) AddExtra(k string, v interface{}) {
}

// AddExtras does nothing
func (t *Tracker) AddExtras(extra map[string]interface{}) {
}

// RemoveExtra does nothing
func (t *Tracker) RemoveExtra(k string) {
}

// NewChild return the same noop tracker
func (t *Tracker) NewChild() errtracking.Tracker {
	return t
}
