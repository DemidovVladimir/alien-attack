// Code generated by MockGen. DO NOT EDIT.
// Source: usecase/usecase.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockWorldUseCase is a mock of WorldUseCase interface.
type MockWorldUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockWorldUseCaseMockRecorder
}

// MockWorldUseCaseMockRecorder is the mock recorder for MockWorldUseCase.
type MockWorldUseCaseMockRecorder struct {
	mock *MockWorldUseCase
}

// NewMockWorldUseCase creates a new mock instance.
func NewMockWorldUseCase(ctrl *gomock.Controller) *MockWorldUseCase {
	mock := &MockWorldUseCase{ctrl: ctrl}
	mock.recorder = &MockWorldUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWorldUseCase) EXPECT() *MockWorldUseCaseMockRecorder {
	return m.recorder
}

// GetCityByName mocks base method.
func (m *MockWorldUseCase) GetCityByName(arg0 string) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCityByName", arg0)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCityByName indicates an expected call of GetCityByName.
func (mr *MockWorldUseCaseMockRecorder) GetCityByName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCityByName", reflect.TypeOf((*MockWorldUseCase)(nil).GetCityByName), arg0)
}

// GetRandomNeighbor mocks base method.
func (m *MockWorldUseCase) GetRandomNeighbor(arg0 string, arg1 int64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRandomNeighbor", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRandomNeighbor indicates an expected call of GetRandomNeighbor.
func (mr *MockWorldUseCaseMockRecorder) GetRandomNeighbor(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRandomNeighbor", reflect.TypeOf((*MockWorldUseCase)(nil).GetRandomNeighbor), arg0, arg1)
}

// ProvideRandomCity mocks base method.
func (m *MockWorldUseCase) ProvideRandomCity(arg0 int64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProvideRandomCity", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ProvideRandomCity indicates an expected call of ProvideRandomCity.
func (mr *MockWorldUseCaseMockRecorder) ProvideRandomCity(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProvideRandomCity", reflect.TypeOf((*MockWorldUseCase)(nil).ProvideRandomCity), arg0)
}