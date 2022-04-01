// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/xiaotianfork/quic-go (interfaces: EarlySession)

// Package mockquic is a generated GoMock package.
package mockquic

import (
	context "context"
	net "net"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	quic "github.com/xiaotianfork/quic-go"
	qerr "github.com/xiaotianfork/quic-go/internal/qerr"
)

// MockEarlySession is a mock of EarlySession interface.
type MockEarlySession struct {
	ctrl     *gomock.Controller
	recorder *MockEarlySessionMockRecorder
}

// MockEarlySessionMockRecorder is the mock recorder for MockEarlySession.
type MockEarlySessionMockRecorder struct {
	mock *MockEarlySession
}

// NewMockEarlySession creates a new mock instance.
func NewMockEarlySession(ctrl *gomock.Controller) *MockEarlySession {
	mock := &MockEarlySession{ctrl: ctrl}
	mock.recorder = &MockEarlySessionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEarlySession) EXPECT() *MockEarlySessionMockRecorder {
	return m.recorder
}

// AcceptStream mocks base method.
func (m *MockEarlySession) AcceptStream(arg0 context.Context) (quic.Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcceptStream", arg0)
	ret0, _ := ret[0].(quic.Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AcceptStream indicates an expected call of AcceptStream.
func (mr *MockEarlySessionMockRecorder) AcceptStream(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptStream", reflect.TypeOf((*MockEarlySession)(nil).AcceptStream), arg0)
}

// AcceptUniStream mocks base method.
func (m *MockEarlySession) AcceptUniStream(arg0 context.Context) (quic.ReceiveStream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AcceptUniStream", arg0)
	ret0, _ := ret[0].(quic.ReceiveStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AcceptUniStream indicates an expected call of AcceptUniStream.
func (mr *MockEarlySessionMockRecorder) AcceptUniStream(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AcceptUniStream", reflect.TypeOf((*MockEarlySession)(nil).AcceptUniStream), arg0)
}

// CloseWithError mocks base method.
func (m *MockEarlySession) CloseWithError(arg0 qerr.ApplicationErrorCode, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseWithError", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseWithError indicates an expected call of CloseWithError.
func (mr *MockEarlySessionMockRecorder) CloseWithError(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseWithError", reflect.TypeOf((*MockEarlySession)(nil).CloseWithError), arg0, arg1)
}

// ConnectionState mocks base method.
func (m *MockEarlySession) ConnectionState() quic.ConnectionState {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConnectionState")
	ret0, _ := ret[0].(quic.ConnectionState)
	return ret0
}

// ConnectionState indicates an expected call of ConnectionState.
func (mr *MockEarlySessionMockRecorder) ConnectionState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConnectionState", reflect.TypeOf((*MockEarlySession)(nil).ConnectionState))
}

// Context mocks base method.
func (m *MockEarlySession) Context() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Context")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// Context indicates an expected call of Context.
func (mr *MockEarlySessionMockRecorder) Context() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Context", reflect.TypeOf((*MockEarlySession)(nil).Context))
}

// HandshakeComplete mocks base method.
func (m *MockEarlySession) HandshakeComplete() context.Context {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HandshakeComplete")
	ret0, _ := ret[0].(context.Context)
	return ret0
}

// HandshakeComplete indicates an expected call of HandshakeComplete.
func (mr *MockEarlySessionMockRecorder) HandshakeComplete() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HandshakeComplete", reflect.TypeOf((*MockEarlySession)(nil).HandshakeComplete))
}

// LocalAddr mocks base method.
func (m *MockEarlySession) LocalAddr() net.Addr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LocalAddr")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// LocalAddr indicates an expected call of LocalAddr.
func (mr *MockEarlySessionMockRecorder) LocalAddr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocalAddr", reflect.TypeOf((*MockEarlySession)(nil).LocalAddr))
}

// NextSession mocks base method.
func (m *MockEarlySession) NextSession() quic.Session {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NextSession")
	ret0, _ := ret[0].(quic.Session)
	return ret0
}

// NextSession indicates an expected call of NextSession.
func (mr *MockEarlySessionMockRecorder) NextSession() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NextSession", reflect.TypeOf((*MockEarlySession)(nil).NextSession))
}

// OpenStream mocks base method.
func (m *MockEarlySession) OpenStream() (quic.Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenStream")
	ret0, _ := ret[0].(quic.Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenStream indicates an expected call of OpenStream.
func (mr *MockEarlySessionMockRecorder) OpenStream() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenStream", reflect.TypeOf((*MockEarlySession)(nil).OpenStream))
}

// OpenStreamSync mocks base method.
func (m *MockEarlySession) OpenStreamSync(arg0 context.Context) (quic.Stream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenStreamSync", arg0)
	ret0, _ := ret[0].(quic.Stream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenStreamSync indicates an expected call of OpenStreamSync.
func (mr *MockEarlySessionMockRecorder) OpenStreamSync(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenStreamSync", reflect.TypeOf((*MockEarlySession)(nil).OpenStreamSync), arg0)
}

// OpenUniStream mocks base method.
func (m *MockEarlySession) OpenUniStream() (quic.SendStream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenUniStream")
	ret0, _ := ret[0].(quic.SendStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenUniStream indicates an expected call of OpenUniStream.
func (mr *MockEarlySessionMockRecorder) OpenUniStream() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenUniStream", reflect.TypeOf((*MockEarlySession)(nil).OpenUniStream))
}

// OpenUniStreamSync mocks base method.
func (m *MockEarlySession) OpenUniStreamSync(arg0 context.Context) (quic.SendStream, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenUniStreamSync", arg0)
	ret0, _ := ret[0].(quic.SendStream)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenUniStreamSync indicates an expected call of OpenUniStreamSync.
func (mr *MockEarlySessionMockRecorder) OpenUniStreamSync(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenUniStreamSync", reflect.TypeOf((*MockEarlySession)(nil).OpenUniStreamSync), arg0)
}

// ReceiveMessage mocks base method.
func (m *MockEarlySession) ReceiveMessage() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReceiveMessage")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReceiveMessage indicates an expected call of ReceiveMessage.
func (mr *MockEarlySessionMockRecorder) ReceiveMessage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReceiveMessage", reflect.TypeOf((*MockEarlySession)(nil).ReceiveMessage))
}

// RemoteAddr mocks base method.
func (m *MockEarlySession) RemoteAddr() net.Addr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoteAddr")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// RemoteAddr indicates an expected call of RemoteAddr.
func (mr *MockEarlySessionMockRecorder) RemoteAddr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoteAddr", reflect.TypeOf((*MockEarlySession)(nil).RemoteAddr))
}

// SendMessage mocks base method.
func (m *MockEarlySession) SendMessage(arg0 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendMessage", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SendMessage indicates an expected call of SendMessage.
func (mr *MockEarlySessionMockRecorder) SendMessage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendMessage", reflect.TypeOf((*MockEarlySession)(nil).SendMessage), arg0)
}
