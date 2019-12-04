// Code generated by MockGen. DO NOT EDIT.
// Source: ethereumapis/eth/v1alpha1/node.pb.go

// Package internal is a generated GoMock package.
package internal

import (
	context "context"
	reflect "reflect"

	types "github.com/gogo/protobuf/types"
	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/prysmaticlabs/ethereumapis/eth/v1alpha1"
	grpc "google.golang.org/grpc"
)

// MockNodeClient is a mock of NodeClient interface
type MockNodeClient struct {
	ctrl     *gomock.Controller
	recorder *MockNodeClientMockRecorder
}

// MockNodeClientMockRecorder is the mock recorder for MockNodeClient
type MockNodeClientMockRecorder struct {
	mock *MockNodeClient
}

// NewMockNodeClient creates a new mock instance
func NewMockNodeClient(ctrl *gomock.Controller) *MockNodeClient {
	mock := &MockNodeClient{ctrl: ctrl}
	mock.recorder = &MockNodeClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNodeClient) EXPECT() *MockNodeClientMockRecorder {
	return m.recorder
}

// GetSyncStatus mocks base method
func (m *MockNodeClient) GetSyncStatus(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*v1alpha1.SyncStatus, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetSyncStatus", varargs...)
	ret0, _ := ret[0].(*v1alpha1.SyncStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSyncStatus indicates an expected call of GetSyncStatus
func (mr *MockNodeClientMockRecorder) GetSyncStatus(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSyncStatus", reflect.TypeOf((*MockNodeClient)(nil).GetSyncStatus), varargs...)
}

// GetGenesis mocks base method
func (m *MockNodeClient) GetGenesis(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*v1alpha1.Genesis, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetGenesis", varargs...)
	ret0, _ := ret[0].(*v1alpha1.Genesis)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGenesis indicates an expected call of GetGenesis
func (mr *MockNodeClientMockRecorder) GetGenesis(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGenesis", reflect.TypeOf((*MockNodeClient)(nil).GetGenesis), varargs...)
}

// GetVersion mocks base method
func (m *MockNodeClient) GetVersion(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*v1alpha1.Version, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetVersion", varargs...)
	ret0, _ := ret[0].(*v1alpha1.Version)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVersion indicates an expected call of GetVersion
func (mr *MockNodeClientMockRecorder) GetVersion(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVersion", reflect.TypeOf((*MockNodeClient)(nil).GetVersion), varargs...)
}

// ListImplementedServices mocks base method
func (m *MockNodeClient) ListImplementedServices(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*v1alpha1.ImplementedServices, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListImplementedServices", varargs...)
	ret0, _ := ret[0].(*v1alpha1.ImplementedServices)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListImplementedServices indicates an expected call of ListImplementedServices
func (mr *MockNodeClientMockRecorder) ListImplementedServices(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListImplementedServices", reflect.TypeOf((*MockNodeClient)(nil).ListImplementedServices), varargs...)
}

// ListPeers mocks base method
func (m *MockNodeClient) ListPeers(ctx context.Context, in *types.Empty, opts ...grpc.CallOption) (*v1alpha1.Peers, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListPeers", varargs...)
	ret0, _ := ret[0].(*v1alpha1.Peers)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPeers indicates an expected call of ListPeers
func (mr *MockNodeClientMockRecorder) ListPeers(ctx, in interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPeers", reflect.TypeOf((*MockNodeClient)(nil).ListPeers), varargs...)
}

// MockNodeServer is a mock of NodeServer interface
type MockNodeServer struct {
	ctrl     *gomock.Controller
	recorder *MockNodeServerMockRecorder
}

// MockNodeServerMockRecorder is the mock recorder for MockNodeServer
type MockNodeServerMockRecorder struct {
	mock *MockNodeServer
}

// NewMockNodeServer creates a new mock instance
func NewMockNodeServer(ctrl *gomock.Controller) *MockNodeServer {
	mock := &MockNodeServer{ctrl: ctrl}
	mock.recorder = &MockNodeServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNodeServer) EXPECT() *MockNodeServerMockRecorder {
	return m.recorder
}

// GetSyncStatus mocks base method
func (m *MockNodeServer) GetSyncStatus(arg0 context.Context, arg1 *types.Empty) (*v1alpha1.SyncStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSyncStatus", arg0, arg1)
	ret0, _ := ret[0].(*v1alpha1.SyncStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSyncStatus indicates an expected call of GetSyncStatus
func (mr *MockNodeServerMockRecorder) GetSyncStatus(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSyncStatus", reflect.TypeOf((*MockNodeServer)(nil).GetSyncStatus), arg0, arg1)
}

// GetGenesis mocks base method
func (m *MockNodeServer) GetGenesis(arg0 context.Context, arg1 *types.Empty) (*v1alpha1.Genesis, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGenesis", arg0, arg1)
	ret0, _ := ret[0].(*v1alpha1.Genesis)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGenesis indicates an expected call of GetGenesis
func (mr *MockNodeServerMockRecorder) GetGenesis(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGenesis", reflect.TypeOf((*MockNodeServer)(nil).GetGenesis), arg0, arg1)
}

// GetVersion mocks base method
func (m *MockNodeServer) GetVersion(arg0 context.Context, arg1 *types.Empty) (*v1alpha1.Version, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVersion", arg0, arg1)
	ret0, _ := ret[0].(*v1alpha1.Version)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVersion indicates an expected call of GetVersion
func (mr *MockNodeServerMockRecorder) GetVersion(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVersion", reflect.TypeOf((*MockNodeServer)(nil).GetVersion), arg0, arg1)
}

// ListImplementedServices mocks base method
func (m *MockNodeServer) ListImplementedServices(arg0 context.Context, arg1 *types.Empty) (*v1alpha1.ImplementedServices, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListImplementedServices", arg0, arg1)
	ret0, _ := ret[0].(*v1alpha1.ImplementedServices)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListImplementedServices indicates an expected call of ListImplementedServices
func (mr *MockNodeServerMockRecorder) ListImplementedServices(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListImplementedServices", reflect.TypeOf((*MockNodeServer)(nil).ListImplementedServices), arg0, arg1)
}

// ListPeers mocks base method
func (m *MockNodeServer) ListPeers(arg0 context.Context, arg1 *types.Empty) (*v1alpha1.Peers, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPeers", arg0, arg1)
	ret0, _ := ret[0].(*v1alpha1.Peers)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPeers indicates an expected call of ListPeers
func (mr *MockNodeServerMockRecorder) ListPeers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPeers", reflect.TypeOf((*MockNodeServer)(nil).ListPeers), arg0, arg1)
}