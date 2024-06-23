// Code generated by MockGen. DO NOT EDIT.
// Source: ./core/repo/account.go

// Package repo is a generated GoMock package.
package repo

import (
	reflect "reflect"

	model "github.com/asepnur/sharing/example/testing/combine/account/core/model"
	gomock "github.com/golang/mock/gomock"
)

// MockAccountRepo is a mock of AccountRepo interface.
type MockAccountRepo struct {
	ctrl     *gomock.Controller
	recorder *MockAccountRepoMockRecorder
}

// MockAccountRepoMockRecorder is the mock recorder for MockAccountRepo.
type MockAccountRepoMockRecorder struct {
	mock *MockAccountRepo
}

// NewMockAccountRepo creates a new mock instance.
func NewMockAccountRepo(ctrl *gomock.Controller) *MockAccountRepo {
	mock := &MockAccountRepo{ctrl: ctrl}
	mock.recorder = &MockAccountRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountRepo) EXPECT() *MockAccountRepoMockRecorder {
	return m.recorder
}

// FindByUsername mocks base method.
func (m *MockAccountRepo) FindByUsername(username string) (*model.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUsername", username)
	ret0, _ := ret[0].(*model.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUsername indicates an expected call of FindByUsername.
func (mr *MockAccountRepoMockRecorder) FindByUsername(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUsername", reflect.TypeOf((*MockAccountRepo)(nil).FindByUsername), username)
}

// Insert mocks base method.
func (m *MockAccountRepo) Insert(account *model.Account) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Insert", account)
	ret0, _ := ret[0].(error)
	return ret0
}

// Insert indicates an expected call of Insert.
func (mr *MockAccountRepoMockRecorder) Insert(account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Insert", reflect.TypeOf((*MockAccountRepo)(nil).Insert), account)
}
