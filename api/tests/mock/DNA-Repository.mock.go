package mock

import (
	"reflect"

	"github.com/elissonalvesilva/interview-meli/api/domain/protocols"
	"github.com/golang/mock/gomock"
)

// MockDNARepository is a mock of DNARepository interface.
type MockDNARepository struct {
	ctrl     *gomock.Controller
	recorder *MockDNARepositoryMockRecorder
}

// MockDNARepositoryMockRecorder is the mock recorder for MockDNARepository.
type MockDNARepositoryMockRecorder struct {
	mock *MockDNARepository
}

// NewMockDNARepository creates a new mock instance.
func NewMockDNARepository(ctrl *gomock.Controller) *MockDNARepository {
	mock := &MockDNARepository{ctrl: ctrl}
	mock.recorder = &MockDNARepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDNARepository) EXPECT() *MockDNARepositoryMockRecorder {
	return m.recorder
}

// StatsHumanSimian mocks base method.
func (m *MockDNARepository) StatsHumanSimian() (protocols.StatsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StatsHumanSimian")
	ret0, _ := ret[0].(protocols.StatsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// StatsHumanSimian indicates an expected call of StatsHumanSimian.
func (mr *MockDNARepositoryMockRecorder) StatsHumanSimian() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StatsHumanSimian", reflect.TypeOf((*MockDNARepository)(nil).StatsHumanSimian))
}

