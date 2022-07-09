package noop

import (
	"time"

	"github.com/Nivl/coeur/metrics"
)

// Metrics is a noop implementing the Metrics interface
type Metrics struct{}

// we make sure the struct implements the interface
var _ metrics.Metrics = (*Metrics)(nil)

// New creates a new Metrics struct that does nothing
func New() metrics.Metrics {
	return &Metrics{}
}

// Count is a noop
func (m *Metrics) Count(name string, value int64, tags ...string) error {
	return nil
}

// CountWithRate is a noop
func (m *Metrics) CountWithRate(name string, value int64, rate float64, tags ...string) error {
	return nil
}

// Incr is a noop
func (m *Metrics) Incr(name string, tags ...string) error {
	return nil
}

// IncrWithRate is a noop
func (m *Metrics) IncrWithRate(name string, rate float64, tags ...string) error {
	return nil
}

// Decr is a noop
func (m *Metrics) Decr(name string, tags ...string) error {
	return nil
}

// DecrWithRate is a noop
func (m *Metrics) DecrWithRate(name string, rate float64, tags ...string) error {
	return nil
}

// Gauge is a noop
func (m *Metrics) Gauge(name string, value float64, tags ...string) error {
	return nil
}

// GaugeWithRate is a noop
func (m *Metrics) GaugeWithRate(name string, value, rate float64, tags ...string) error {
	return nil
}

// Timing is a noop
func (m *Metrics) Timing(name string, value time.Duration, tags ...string) error {
	return nil
}

// TimingWithRate is a noop
func (m *Metrics) TimingWithRate(name string, value time.Duration, rate float64, tags ...string) error {
	return nil
}

// Close is a noop
func (m *Metrics) Close() error {
	return nil
}
