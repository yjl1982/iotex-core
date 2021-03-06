// Code generated by MockGen. DO NOT EDIT.
// Source: ./trie/trie.go

// Package mock_trie is a generated GoMock package.
package mock_trie

import (
	gomock "github.com/golang/mock/gomock"
	common "github.com/iotexproject/iotex-core/common"
	reflect "reflect"
)

// MockTrie is a mock of Trie interface
type MockTrie struct {
	ctrl     *gomock.Controller
	recorder *MockTrieMockRecorder
}

// MockTrieMockRecorder is the mock recorder for MockTrie
type MockTrieMockRecorder struct {
	mock *MockTrie
}

// NewMockTrie creates a new mock instance
func NewMockTrie(ctrl *gomock.Controller) *MockTrie {
	mock := &MockTrie{ctrl: ctrl}
	mock.recorder = &MockTrieMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTrie) EXPECT() *MockTrieMockRecorder {
	return m.recorder
}

// Upsert mocks base method
func (m *MockTrie) Upsert(key, value []byte) error {
	ret := m.ctrl.Call(m, "Upsert", key, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert
func (mr *MockTrieMockRecorder) Upsert(key, value interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockTrie)(nil).Upsert), key, value)
}

// Get mocks base method
func (m *MockTrie) Get(key []byte) ([]byte, error) {
	ret := m.ctrl.Call(m, "Get", key)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockTrieMockRecorder) Get(key interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTrie)(nil).Get), key)
}

// Delete mocks base method
func (m *MockTrie) Delete(key []byte) error {
	ret := m.ctrl.Call(m, "Delete", key)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockTrieMockRecorder) Delete(key interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockTrie)(nil).Delete), key)
}

// Close mocks base method
func (m *MockTrie) Close() error {
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockTrieMockRecorder) Close() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockTrie)(nil).Close))
}

// RootHash mocks base method
func (m *MockTrie) RootHash() common.Hash32B {
	ret := m.ctrl.Call(m, "RootHash")
	ret0, _ := ret[0].(common.Hash32B)
	return ret0
}

// RootHash indicates an expected call of RootHash
func (mr *MockTrieMockRecorder) RootHash() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RootHash", reflect.TypeOf((*MockTrie)(nil).RootHash))
}
