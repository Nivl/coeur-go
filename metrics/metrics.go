// Package metrics contains interfaces and structs to send metrics
// to statd backend
package metrics

import "time"

// gomock interface, requires mockgen
// Update with "go generate github.com/Nivl/coeur/metrics"
//go:generate mockgen -destination ../implementations/mock/metrics.go -package mock github.com/Nivl/coeur/metrics Metrics

// Metrics is an interface used to log data
type Metrics interface {
	// Count tracks how many times something happened per second.
	Count(name string, value int64, tags ...string) error
	// CounWithRate tracks how many times something happened per second.
	CountWithRate(name string, value int64, rate float64, tags ...string) error

	// Incr is just Count of 1
	Incr(name string, tags ...string) error
	// IncrWithRate is just Count of 1
	IncrWithRate(name string, rate float64, tags ...string) error

	// Decr is just Count of -1
	Decr(name string, tags ...string) error
	// DecrWithRate is just Count of -1
	DecrWithRate(name string, rate float64, tags ...string) error

	// Gauge measures the value of a metric at a particular time.
	Gauge(name string, value float64, tags ...string) error
	// GaugeWithRate measures the value of a metric at a particular time.
	GaugeWithRate(name string, value float64, rate float64, tags ...string) error

	// Timing sends timing information. The value is expected to be in nanosecond
	Timing(name string, value time.Duration, tags ...string) error
	// TimingWithRate sends timing information. The value is expected to be in nanosecond
	TimingWithRate(name string, value time.Duration, rate float64, tags ...string) error

	// Close flushes and frees any resource allocated by the logger.
	// The logger may not be reusable after being closed
	Close() error
}
