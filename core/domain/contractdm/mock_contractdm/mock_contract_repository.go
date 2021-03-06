// Code generated by MockGen. DO NOT EDIT.
// Source: contract_repository.go

// Package mock_contractdm is a generated GoMock package.
package mock_contractdm

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	contractdm "github.com/naoyakurokawa/ddd_menta/core/domain/contractdm"
)

// MockContractRepository is a mock of ContractRepository interface.
type MockContractRepository struct {
	ctrl     *gomock.Controller
	recorder *MockContractRepositoryMockRecorder
}

// MockContractRepositoryMockRecorder is the mock recorder for MockContractRepository.
type MockContractRepositoryMockRecorder struct {
	mock *MockContractRepository
}

// NewMockContractRepository creates a new mock instance.
func NewMockContractRepository(ctrl *gomock.Controller) *MockContractRepository {
	mock := &MockContractRepository{ctrl: ctrl}
	mock.recorder = &MockContractRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockContractRepository) EXPECT() *MockContractRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockContractRepository) Create(contract *contractdm.Contract) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", contract)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockContractRepositoryMockRecorder) Create(contract interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockContractRepository)(nil).Create), contract)
}

// FindByID mocks base method.
func (m *MockContractRepository) FindByID(contractID contractdm.ContractID) (*contractdm.Contract, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", contractID)
	ret0, _ := ret[0].(*contractdm.Contract)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockContractRepositoryMockRecorder) FindByID(contractID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockContractRepository)(nil).FindByID), contractID)
}

// UpdateContractStatus mocks base method.
func (m *MockContractRepository) UpdateContractStatus(contractID contractdm.ContractID, contractStatus contractdm.ContractStatus) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateContractStatus", contractID, contractStatus)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateContractStatus indicates an expected call of UpdateContractStatus.
func (mr *MockContractRepositoryMockRecorder) UpdateContractStatus(contractID, contractStatus interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateContractStatus", reflect.TypeOf((*MockContractRepository)(nil).UpdateContractStatus), contractID, contractStatus)
}
