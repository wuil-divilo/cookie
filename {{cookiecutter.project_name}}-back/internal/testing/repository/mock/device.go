// Code generated by MockGen. DO NOT EDIT.
// Source: ./{{cookiecutter.model_name}}.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	model "github.com/divilo/{{cookiecutter.project_name}}-back/internal/model"
	gomock "github.com/golang/mock/gomock"
)

// Mock{{cookiecutter.model_name.capitalize()}}Repository is a mock of {{cookiecutter.model_name.capitalize()}}Repository interface.
type Mock{{cookiecutter.model_name.capitalize()}}Repository struct {
	ctrl     *gomock.Controller
	recorder *Mock{{cookiecutter.model_name.capitalize()}}RepositoryMockRecorder
}

// Mock{{cookiecutter.model_name.capitalize()}}RepositoryMockRecorder is the mock recorder for Mock{{cookiecutter.model_name.capitalize()}}Repository.
type Mock{{cookiecutter.model_name.capitalize()}}RepositoryMockRecorder struct {
	mock *Mock{{cookiecutter.model_name.capitalize()}}Repository
}

// NewMock{{cookiecutter.model_name.capitalize()}}Repository creates a new mock instance.
func NewMock{{cookiecutter.model_name.capitalize()}}Repository(ctrl *gomock.Controller) *Mock{{cookiecutter.model_name.capitalize()}}Repository {
	mock := &Mock{{cookiecutter.model_name.capitalize()}}Repository{ctrl: ctrl}
	mock.recorder = &Mock{{cookiecutter.model_name.capitalize()}}RepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mock{{cookiecutter.model_name.capitalize()}}Repository) EXPECT() *Mock{{cookiecutter.model_name.capitalize()}}RepositoryMockRecorder {
	return m.recorder
}

// FilterByID mocks base method.
func (m *Mock{{cookiecutter.model_name.capitalize()}}Repository) FilterByID(ctx context.Context, dvcID string) (model.{{cookiecutter.model_name.capitalize()}}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilterByID", ctx, dvcID)
	ret0, _ := ret[0].(model.{{cookiecutter.model_name.capitalize()}})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilterByID indicates an expected call of FilterByID.
func (mr *Mock{{cookiecutter.model_name.capitalize()}}RepositoryMockRecorder) FilterByID(ctx, dvcID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterByID", reflect.TypeOf((*Mock{{cookiecutter.model_name.capitalize()}}Repository)(nil).FilterByID), ctx, dvcID)
}

// Upsert mocks base method.
func (m *Mock{{cookiecutter.model_name.capitalize()}}Repository) Upsert(ctx context.Context, {{cookiecutter.model_name}} model.{{cookiecutter.model_name.capitalize()}}) (model.{{cookiecutter.model_name.capitalize()}}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", ctx, {{cookiecutter.model_name}})
	ret0, _ := ret[0].(model.{{cookiecutter.model_name.capitalize()}})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Upsert indicates an expected call of Upsert.
func (mr *Mock{{cookiecutter.model_name.capitalize()}}RepositoryMockRecorder) Upsert(ctx, {{cookiecutter.model_name}} interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*Mock{{cookiecutter.model_name.capitalize()}}Repository)(nil).Upsert), ctx, {{cookiecutter.model_name}})
}
