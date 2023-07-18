// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/anyproto/any-sync/commonspace/peermanager (interfaces: PeerManager)

// Package mock_peermanager is a generated GoMock package.
package mock_peermanager

import (
	context "context"
	reflect "reflect"

	app "github.com/anyproto/any-sync/app"
	spacesyncproto "github.com/anyproto/any-sync/commonspace/spacesyncproto"
	peer "github.com/anyproto/any-sync/net/peer"
	gomock "go.uber.org/mock/gomock"
)

// MockPeerManager is a mock of PeerManager interface.
type MockPeerManager struct {
	ctrl     *gomock.Controller
	recorder *MockPeerManagerMockRecorder
}

// MockPeerManagerMockRecorder is the mock recorder for MockPeerManager.
type MockPeerManagerMockRecorder struct {
	mock *MockPeerManager
}

// NewMockPeerManager creates a new mock instance.
func NewMockPeerManager(ctrl *gomock.Controller) *MockPeerManager {
	mock := &MockPeerManager{ctrl: ctrl}
	mock.recorder = &MockPeerManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPeerManager) EXPECT() *MockPeerManagerMockRecorder {
	return m.recorder
}

// Broadcast mocks base method.
func (m *MockPeerManager) Broadcast(arg0 context.Context, arg1 *spacesyncproto.ObjectSyncMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Broadcast", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Broadcast indicates an expected call of Broadcast.
func (mr *MockPeerManagerMockRecorder) Broadcast(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Broadcast", reflect.TypeOf((*MockPeerManager)(nil).Broadcast), arg0, arg1)
}

// GetResponsiblePeers mocks base method.
func (m *MockPeerManager) GetResponsiblePeers(arg0 context.Context) ([]peer.Peer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetResponsiblePeers", arg0)
	ret0, _ := ret[0].([]peer.Peer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetResponsiblePeers indicates an expected call of GetResponsiblePeers.
func (mr *MockPeerManagerMockRecorder) GetResponsiblePeers(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetResponsiblePeers", reflect.TypeOf((*MockPeerManager)(nil).GetResponsiblePeers), arg0)
}

// Init mocks base method.
func (m *MockPeerManager) Init(arg0 *app.App) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Init", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init.
func (mr *MockPeerManagerMockRecorder) Init(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockPeerManager)(nil).Init), arg0)
}

// Name mocks base method.
func (m *MockPeerManager) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockPeerManagerMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockPeerManager)(nil).Name))
}

// SendPeer mocks base method.
func (m *MockPeerManager) SendPeer(arg0 context.Context, arg1 string, arg2 *spacesyncproto.ObjectSyncMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendPeer", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendPeer indicates an expected call of SendPeer.
func (mr *MockPeerManagerMockRecorder) SendPeer(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendPeer", reflect.TypeOf((*MockPeerManager)(nil).SendPeer), arg0, arg1, arg2)
}
