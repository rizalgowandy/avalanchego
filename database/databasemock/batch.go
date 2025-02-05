// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/ava-labs/avalanchego/database (interfaces: Batch)
//
// Generated by this command:
//
//	mockgen -package=databasemock -destination=databasemock/batch.go -mock_names=Batch=Batch . Batch
//

// Package databasemock is a generated GoMock package.
package databasemock

import (
	reflect "reflect"

	database "github.com/ava-labs/avalanchego/database"
	gomock "go.uber.org/mock/gomock"
)

// Batch is a mock of Batch interface.
type Batch struct {
	ctrl     *gomock.Controller
	recorder *BatchMockRecorder
	isgomock struct{}
}

// BatchMockRecorder is the mock recorder for Batch.
type BatchMockRecorder struct {
	mock *Batch
}

// NewBatch creates a new mock instance.
func NewBatch(ctrl *gomock.Controller) *Batch {
	mock := &Batch{ctrl: ctrl}
	mock.recorder = &BatchMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Batch) EXPECT() *BatchMockRecorder {
	return m.recorder
}

// Delete mocks base method.
func (m *Batch) Delete(key []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *BatchMockRecorder) Delete(key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*Batch)(nil).Delete), key)
}

// Inner mocks base method.
func (m *Batch) Inner() database.Batch {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Inner")
	ret0, _ := ret[0].(database.Batch)
	return ret0
}

// Inner indicates an expected call of Inner.
func (mr *BatchMockRecorder) Inner() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Inner", reflect.TypeOf((*Batch)(nil).Inner))
}

// Put mocks base method.
func (m *Batch) Put(key, value []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Put", key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Put indicates an expected call of Put.
func (mr *BatchMockRecorder) Put(key, value any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Put", reflect.TypeOf((*Batch)(nil).Put), key, value)
}

// Replay mocks base method.
func (m *Batch) Replay(w database.KeyValueWriterDeleter) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Replay", w)
	ret0, _ := ret[0].(error)
	return ret0
}

// Replay indicates an expected call of Replay.
func (mr *BatchMockRecorder) Replay(w any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Replay", reflect.TypeOf((*Batch)(nil).Replay), w)
}

// Reset mocks base method.
func (m *Batch) Reset() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Reset")
}

// Reset indicates an expected call of Reset.
func (mr *BatchMockRecorder) Reset() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reset", reflect.TypeOf((*Batch)(nil).Reset))
}

// Size mocks base method.
func (m *Batch) Size() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(int)
	return ret0
}

// Size indicates an expected call of Size.
func (mr *BatchMockRecorder) Size() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*Batch)(nil).Size))
}

// Write mocks base method.
func (m *Batch) Write() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write")
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write.
func (mr *BatchMockRecorder) Write() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*Batch)(nil).Write))
}
