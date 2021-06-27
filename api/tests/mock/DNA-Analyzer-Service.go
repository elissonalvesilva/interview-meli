package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAnalyzeGRPC is a mock of AnalyzeGRPC interface.
type MockAnalyzeGRPC struct {
	ctrl     *gomock.Controller
	recorder *MockAnalyzeGRPCMockRecorder
}

// MockAnalyzeGRPCMockRecorder is the mock recorder for MockAnalyzeGRPC.
type MockAnalyzeGRPCMockRecorder struct {
	mock *MockAnalyzeGRPC
}

// NewMockAnalyzeGRPC creates a new mock instance.
func NewMockAnalyzeGRPC(ctrl *gomock.Controller) *MockAnalyzeGRPC {
	mock := &MockAnalyzeGRPC{ctrl: ctrl}
	mock.recorder = &MockAnalyzeGRPCMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAnalyzeGRPC) EXPECT() *MockAnalyzeGRPCMockRecorder {
	return m.recorder
}

// AnalyzeDNA mocks base method.
func (m *MockAnalyzeGRPC) AnalyzeDNA(arg0 []string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AnalyzeDNA", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AnalyzeDNA indicates an expected call of AnalyzeDNA.
func (mr *MockAnalyzeGRPCMockRecorder) AnalyzeDNA(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AnalyzeDNA", reflect.TypeOf((*MockAnalyzeGRPC)(nil).AnalyzeDNA), arg0)
}
