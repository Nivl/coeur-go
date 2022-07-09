package o11y

import (
	"context"

	"github.com/Nivl/coeur/errtracking"
	"github.com/Nivl/coeur/implementations/noop"
	"github.com/Nivl/coeur/logger"
	"github.com/Nivl/coeur/metrics"
)

// ConfigFunc defines methods uses to configure the constructor
// observability object
type ConfigFunc func(*Observability)

// Observability is a struct containing all the dependencies related
// to observability.
// Assumes it's not goroutine safe. A new child should be created
// for every goroutine using NewChild()
type Observability struct {
	Logger  logger.Logger
	Metrics metrics.Metrics
	Tracker errtracking.Tracker
	isChild bool
}

// WithLogger attaches a Logger to the object
func WithLogger(l logger.Logger) ConfigFunc {
	return func(o *Observability) {
		o.Logger = l
	}
}

// WithMetrics attaches an Metrics capturer to the object
func WithMetrics(l metrics.Metrics) ConfigFunc {
	return func(o *Observability) {
		o.Metrics = l
	}
}

// WithTracker attaches an error tracker to the object
func WithTracker(l errtracking.Tracker) ConfigFunc {
	return func(o *Observability) {
		o.Tracker = l
	}
}

// New creates a new Observability object
// Everything is a noop by default
func New(opts ...ConfigFunc) *Observability {
	o := &Observability{
		Logger:  &noop.Logger{},
		Metrics: &noop.Metrics{},
		Tracker: &noop.Tracker{},
	}
	// Apply all the options
	for _, opt := range opts {
		opt(o)
	}
	return o
}

// NewChild creates a child version of the observability.
// A new child should be created for every new go routine
func (o *Observability) NewChild() *Observability {
	return &Observability{
		Logger:  o.Logger.NewChild(),
		Tracker: o.Tracker.NewChild(),
		Metrics: o.Metrics,
		isChild: true,
	}
}

// ReportError logs an error and reports it to the error tracker
func (o *Observability) ReportError(ctx context.Context, e error, extra ...map[string]interface{}) error {
	err := o.Logger.Error(ctx, e.Error(), extra...)
	o.Tracker.ReportError(ctx, e, extra...)
	return err
}

// Close flushes and frees the resources of the all the objects
// Close should always be called.
func (o *Observability) Close() (err error) {
	err = o.Logger.Close()
	closeErr := o.Tracker.Close()
	if closeErr != nil && err == nil {
		err = closeErr
	}
	// The metrics are globals so we don't close the metrics if
	// they're a child
	if !o.isChild {
		closeErr = o.Metrics.Close()
		if closeErr != nil && err == nil {
			err = closeErr
		}
	}
	return err
}
