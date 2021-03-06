// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Nivl/coeur/metrics (interfaces: Metrics)

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
)

// MockMetrics is a mock of Metrics interface.
type MockMetrics struct {
	ctrl     *gomock.Controller
	recorder *MockMetricsMockRecorder
}

// MockMetricsMockRecorder is the mock recorder for MockMetrics.
type MockMetricsMockRecorder struct {
	mock *MockMetrics
}

// NewMockMetrics creates a new mock instance.
func NewMockMetrics(ctrl *gomock.Controller) *MockMetrics {
	mock := &MockMetrics{ctrl: ctrl}
	mock.recorder = &MockMetricsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMetrics) EXPECT() *MockMetricsMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockMetrics) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockMetricsMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockMetrics)(nil).Close))
}

// Count mocks base method.
func (m *MockMetrics) Count(arg0 string, arg1 int64, arg2 ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Count", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Count indicates an expected call of Count.
func (mr *MockMetricsMockRecorder) Count(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockMetrics)(nil).Count), varargs...)
}

// CountWithRate mocks base method.
func (m *MockMetrics) CountWithRate(arg0 string, arg1 int64, arg2 float64, arg3 ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CountWithRate", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// CountWithRate indicates an expected call of CountWithRate.
func (mr *MockMetricsMockRecorder) CountWithRate(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountWithRate", reflect.TypeOf((*MockMetrics)(nil).CountWithRate), varargs...)
}

// Decr mocks base method.
func (m *MockMetrics) Decr(arg0 string, arg1 ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Decr", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Decr indicates an expected call of Decr.
func (mr *MockMetricsMockRecorder) Decr(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Decr", reflect.TypeOf((*MockMetrics)(nil).Decr), varargs...)
}

// DecrWithRate mocks base method.
func (m *MockMetrics) DecrWithRate(arg0 string, arg1 float64, arg2 ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DecrWithRate", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// DecrWithRate indicates an expected call of DecrWithRate.
func (mr *MockMetricsMockRecorder) DecrWithRate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecrWithRate", reflect.TypeOf((*MockMetrics)(nil).DecrWithRate), varargs...)
}

// Gauge mocks base method.
func (m *MockMetrics) Gauge(arg0 string, arg1 float64, arg2 ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Gauge", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Gauge indicates an expected call of Gauge.
func (mr *MockMetricsMockRecorder) Gauge(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Gauge", reflect.TypeOf((*MockMetrics)(nil).Gauge), varargs...)
}

// GaugeWithRate mocks base method.
func (m *MockMetrics) GaugeWithRate(arg0 string, arg1, arg2 float64, arg3 ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GaugeWithRate", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// GaugeWithRate indicates an expected call of GaugeWithRate.
func (mr *MockMetricsMockRecorder) GaugeWithRate(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GaugeWithRate", reflect.TypeOf((*MockMetrics)(nil).GaugeWithRate), varargs...)
}

// Incr mocks base method.
func (m *MockMetrics) Incr(arg0 string, arg1 ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Incr", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Incr indicates an expected call of Incr.
func (mr *MockMetricsMockRecorder) Incr(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Incr", reflect.TypeOf((*MockMetrics)(nil).Incr), varargs...)
}

// IncrWithRate mocks base method.
func (m *MockMetrics) IncrWithRate(arg0 string, arg1 float64, arg2 ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "IncrWithRate", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// IncrWithRate indicates an expected call of IncrWithRate.
func (mr *MockMetricsMockRecorder) IncrWithRate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IncrWithRate", reflect.TypeOf((*MockMetrics)(nil).IncrWithRate), varargs...)
}

// Timing mocks base method.
func (m *MockMetrics) Timing(arg0 string, arg1 time.Duration, arg2 ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Timing", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Timing indicates an expected call of Timing.
func (mr *MockMetricsMockRecorder) Timing(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Timing", reflect.TypeOf((*MockMetrics)(nil).Timing), varargs...)
}

// TimingWithRate mocks base method.
func (m *MockMetrics) TimingWithRate(arg0 string, arg1 time.Duration, arg2 float64, arg3 ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1, arg2}
	for _, a := range arg3 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "TimingWithRate", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// TimingWithRate indicates an expected call of TimingWithRate.
func (mr *MockMetricsMockRecorder) TimingWithRate(arg0, arg1, arg2 interface{}, arg3 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1, arg2}, arg3...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TimingWithRate", reflect.TypeOf((*MockMetrics)(nil).TimingWithRate), varargs...)
}
