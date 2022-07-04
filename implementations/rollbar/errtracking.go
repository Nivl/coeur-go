package rollbar

import (
	"context"
	"os"

	"github.com/Nivl/coeur/errtracking"
	"github.com/Nivl/coeur/secret"
	"github.com/rollbar/rollbar-go"
)

// Tracker represents a sentry implementation of the Tracker interface
type Tracker struct {
	c      *rollbar.Client
	cfg    *Config
	extras map[string]interface{}
	token  secret.String
}

// we make sure the struct implements the interface
var _ errtracking.Tracker = (*Tracker)(nil)

// Config represents all the option settings for Rollbar
type Config struct {
	// Token represents the Rollbar token.
	// Defaults to os.Getenv("ROLLBAR_TOKEN")
	Token secret.String
	// CodeVersion is the  Git hash/branch/tag (required for
	// GitHub integration)
	CodeVersion string
	// defaults to os.Getenv("environment") || os.Getenv("env")
	Environment string
	// ServerHost is the hostname sent with each item
	// defaults to rollbar's default (os.Hostname)
	ServerHost string
	// ServerRoot is the path to the application code root, not
	// including the final slash. This is used to collapse non-project
	// code when displaying tracebacks.
	// defaults to /
	ServerRoot string
}

// NewTracker returns a new Tracker using sentry as backend
// You should use a different tracker per goroutine
func NewTracker(cfg Config) (errtracking.Tracker, error) {
	if cfg.Token.Value() == "" {
		cfg.Token = secret.NewString(os.Getenv("ROLLBAR_TOKEN"))
	}
	if cfg.Environment == "" {
		env := os.Getenv("environment")
		if env == "" {
			env = os.Getenv("env")
		}
		if env != "" {
			cfg.Environment = env
		}
	}
	if cfg.ServerRoot == "" {
		cfg.ServerRoot = "/"
	}

	c := rollbar.New(cfg.Token.Value(), cfg.Environment, cfg.CodeVersion, cfg.ServerHost, cfg.ServerRoot)
	c.SetEnabled(true)
	return &Tracker{
		c:      c,
		cfg:    &cfg,
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
	t.c.ErrorWithExtras("error", err, extra)
}

// Close flushes and frees any resource allocated by the tracker.
// The tracker may not be reusable after being closed.
func (t *Tracker) Close() error {
	return t.c.Close()
}

// SetUser sets the information for identifying a user associated
// with any errors or messages. Only the ID is required.
func (t *Tracker) SetUser(ctx context.Context, u errtracking.User) {
	t.c.SetPerson(u.ID, u.Name, u.Email)
}

// RemoveUser Removes any set users
func (t *Tracker) RemoveUser(ctx context.Context) {
	t.c.ClearPerson()
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
func (t *Tracker) AddExtras(extras map[string]interface{}) {
	for k, v := range extras {
		t.extras[k] = v
	}
}

// RemoveExtra remove an extra from being attached to errors.
func (t *Tracker) RemoveExtra(k string) {
	delete(t.extras, k)
}

// NewChild return a new Tracker that copies over the extras and tags.
func (t *Tracker) NewChild() errtracking.Tracker {
	extra := map[string]interface{}{}
	for k, v := range t.extras {
		extra[k] = v
	}

	c := rollbar.New(t.token.Value(), t.cfg.Environment, t.cfg.CodeVersion, t.cfg.ServerHost, t.cfg.ServerRoot)
	c.SetEnabled(true)
	return &Tracker{
		c:      c,
		cfg:    t.cfg,
		token:  t.token,
		extras: extra,
	}
}
