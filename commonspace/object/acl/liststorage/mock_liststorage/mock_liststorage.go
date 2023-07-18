// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/anyproto/any-sync/commonspace/object/acl/liststorage (interfaces: ListStorage)

// Package mock_liststorage is a generated GoMock package.
package mock_liststorage

import (
	context "context"
	reflect "reflect"

	consensusproto "github.com/anyproto/any-sync/consensus/consensusproto"
	gomock "go.uber.org/mock/gomock"
)

// MockListStorage is a mock of ListStorage interface.
type MockListStorage struct {
	ctrl     *gomock.Controller
	recorder *MockListStorageMockRecorder
}

// MockListStorageMockRecorder is the mock recorder for MockListStorage.
type MockListStorageMockRecorder struct {
	mock *MockListStorage
}

// NewMockListStorage creates a new mock instance.
func NewMockListStorage(ctrl *gomock.Controller) *MockListStorage {
	mock := &MockListStorage{ctrl: ctrl}
	mock.recorder = &MockListStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockListStorage) EXPECT() *MockListStorageMockRecorder {
	return m.recorder
}

// AddRawRecord mocks base method.
func (m *MockListStorage) AddRawRecord(arg0 context.Context, arg1 *consensusproto.RawRecordWithId) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRawRecord", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddRawRecord indicates an expected call of AddRawRecord.
func (mr *MockListStorageMockRecorder) AddRawRecord(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRawRecord", reflect.TypeOf((*MockListStorage)(nil).AddRawRecord), arg0, arg1)
}

// GetRawRecord mocks base method.
func (m *MockListStorage) GetRawRecord(arg0 context.Context, arg1 string) (*consensusproto.RawRecordWithId, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRawRecord", arg0, arg1)
	ret0, _ := ret[0].(*consensusproto.RawRecordWithId)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRawRecord indicates an expected call of GetRawRecord.
func (mr *MockListStorageMockRecorder) GetRawRecord(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRawRecord", reflect.TypeOf((*MockListStorage)(nil).GetRawRecord), arg0, arg1)
}

// Head mocks base method.
func (m *MockListStorage) Head() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Head")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Head indicates an expected call of Head.
func (mr *MockListStorageMockRecorder) Head() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Head", reflect.TypeOf((*MockListStorage)(nil).Head))
}

// Id mocks base method.
func (m *MockListStorage) Id() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Id")
	ret0, _ := ret[0].(string)
	return ret0
}

// Id indicates an expected call of Id.
func (mr *MockListStorageMockRecorder) Id() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Id", reflect.TypeOf((*MockListStorage)(nil).Id))
}

// Root mocks base method.
func (m *MockListStorage) Root() (*consensusproto.RawRecordWithId, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Root")
	ret0, _ := ret[0].(*consensusproto.RawRecordWithId)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Root indicates an expected call of Root.
func (mr *MockListStorageMockRecorder) Root() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Root", reflect.TypeOf((*MockListStorage)(nil).Root))
}

// SetHead mocks base method.
func (m *MockListStorage) SetHead(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetHead", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetHead indicates an expected call of SetHead.
func (mr *MockListStorageMockRecorder) SetHead(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetHead", reflect.TypeOf((*MockListStorage)(nil).SetHead), arg0)
}
