package metrics

import (
	"os"
	"time"

	"github.com/DataDog/datadog-go/v5/statsd"
)

type metrics struct {
	c *statsd.Client
}

// we make sure the struct implements the interface
var _ Metrics = (*metrics)(nil)

// New creates a new statsd client that sends the data to the
// address specified in os.Getenv("STASTD_HOST")
func New(addr string) (Metrics, error) {
	return NewWithAddr("")
}

// NewWithAddr creates a new statsd client that sends the data to the
// given address. if addr is empty, os.Getenv("STATSD_HOST") will be used
func NewWithAddr(addr string) (Metrics, error) {
	if addr == "" {
		addr = os.Getenv("STASTD_HOST")
	}
	c, err := statsd.New(addr)
	if err != nil {
		return nil, err
	}
	return &metrics{
		c: c,
	}, nil
}

// Count tracks how many times something happened per second.
func (m *metrics) Count(name string, value int64, tags ...string) error {
	return m.CountWithRate(name, value, 1, tags...)
}

// CounWithRate tracks how many times something happened per second.
func (m *metrics) CountWithRate(name string, value int64, rate float64, tags ...string) error {
	return m.c.Count(name, value, tags, rate)
}

// Incr is just Count of 1
func (m *metrics) Incr(name string, tags ...string) error {
	return m.IncrWithRate(name, 1, tags...)
}

// IncrWithRate is just Count of 1
func (m *metrics) IncrWithRate(name string, rate float64, tags ...string) error {
	return m.c.Incr(name, tags, rate)
}

// Decr is just Count of -1
func (m *metrics) Decr(name string, tags ...string) error {
	return m.DecrWithRate(name, 1, tags...)
}

// DecrWithRate is just Count of -1
func (m *metrics) DecrWithRate(name string, rate float64, tags ...string) error {
	return m.c.Decr(name, tags, rate)
}

// Gauge measures the value of a metric at a particular time.
func (m *metrics) Gauge(name string, value float64, tags ...string) error {
	return m.GaugeWithRate(name, value, 1, tags...)
}

// GaugeWithRate measures the value of a metric at a particular time.
func (m *metrics) GaugeWithRate(name string, value, rate float64, tags ...string) error {
	return m.c.Gauge(name, value, tags, rate)
}

// Timing sends timing information. The value is expected to be in nanosecond
func (m *metrics) Timing(name string, value time.Duration, tags ...string) error {
	return m.TimingWithRate(name, value, 1, tags...)
}

// TimingWithRate sends timing information. The value is expected to be in nanosecond
func (m *metrics) TimingWithRate(name string, value time.Duration, rate float64, tags ...string) error {
	return m.c.Timing(name, value, tags, rate)
}

// Close flushes and frees any resource allocated by the logger.
// The logger may not be reusable after being closed
func (m *metrics) Close() error {
	return m.c.Close()
}
