// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/anyproto/any-sync/consensus/consensusclient (interfaces: Service)
//
// Generated by this command:
//
//	mockgen -destination mock_consensusclient/mock_consensusclient.go github.com/anyproto/any-sync/consensus/consensusclient Service
//
// Package mock_consensusclient is a generated GoMock package.
package mock_consensusclient

import (
	context "context"
	reflect "reflect"

	app "github.com/anyproto/any-sync/app"
	consensusclient "github.com/anyproto/any-sync/consensus/consensusclient"
	consensusproto "github.com/anyproto/any-sync/consensus/consensusproto"
	gomock "go.uber.org/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// AddLog mocks base method.
func (m *MockService) AddLog(arg0 context.Context, arg1 string, arg2 *consensusproto.RawRecordWithId) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddLog", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddLog indicates an expected call of AddLog.
func (mr *MockServiceMockRecorder) AddLog(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddLog", reflect.TypeOf((*MockService)(nil).AddLog), arg0, arg1, arg2)
}

// AddRecord mocks base method.
func (m *MockService) AddRecord(arg0 context.Context, arg1 string, arg2 *consensusproto.RawRecord) (*consensusproto.RawRecordWithId, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRecord", arg0, arg1, arg2)
	ret0, _ := ret[0].(*consensusproto.RawRecordWithId)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddRecord indicates an expected call of AddRecord.
func (mr *MockServiceMockRecorder) AddRecord(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRecord", reflect.TypeOf((*MockService)(nil).AddRecord), arg0, arg1, arg2)
}

// Close mocks base method.
func (m *MockService) Close(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockServiceMockRecorder) Close(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockService)(nil).Close), arg0)
}

// DeleteLog mocks base method.
func (m *MockService) DeleteLog(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteLog", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteLog indicates an expected call of DeleteLog.
func (mr *MockServiceMockRecorder) DeleteLog(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLog", reflect.TypeOf((*MockService)(nil).DeleteLog), arg0, arg1)
}

// Init mocks base method.
func (m *MockService) Init(arg0 *app.App) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockServiceMockRecorder) Init(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockService)(nil).Init), arg0)
}

// Name mocks base method.
func (m *MockService) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockServiceMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockService)(nil).Name))
}

// Run mocks base method.
func (m *MockService) Run(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Run", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockServiceMockRecorder) Run(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockService)(nil).Run), arg0)
}

// UnWatch mocks base method.
func (m *MockService) UnWatch(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnWatch", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnWatch indicates an expected call of UnWatch.
func (mr *MockServiceMockRecorder) UnWatch(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnWatch", reflect.TypeOf((*MockService)(nil).UnWatch), arg0)
}

// Watch mocks base method.
func (m *MockService) Watch(arg0 string, arg1 consensusclient.Watcher) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Watch", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Watch indicates an expected call of Watch.
func (mr *MockServiceMockRecorder) Watch(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Watch", reflect.TypeOf((*MockService)(nil).Watch), arg0, arg1)
}
