// Code generated by MockGen. DO NOT EDIT.
// Source: filesystem.go
//
// Generated by this command:
//
//	mockgen -source=filesystem.go -destination=../../pkg/mocks/fs_mock.go -package=mocks os DirEntry
//

// Package mocks is a generated GoMock package.
package mocks

import (
	fs "crawler/internal/fs"
	os "os"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockFileSystem is a mock of FileSystem interface.
type MockFileSystem struct {
	ctrl     *gomock.Controller
	recorder *MockFileSystemMockRecorder
	isgomock struct{}
}

// MockFileSystemMockRecorder is the mock recorder for MockFileSystem.
type MockFileSystemMockRecorder struct {
	mock *MockFileSystem
}

// NewMockFileSystem creates a new mock instance.
func NewMockFileSystem(ctrl *gomock.Controller) *MockFileSystem {
	mock := &MockFileSystem{ctrl: ctrl}
	mock.recorder = &MockFileSystemMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileSystem) EXPECT() *MockFileSystemMockRecorder {
	return m.recorder
}

// Join mocks base method.
func (m *MockFileSystem) Join(elem ...string) string {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range elem {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Join", varargs...)
	ret0, _ := ret[0].(string)
	return ret0
}

// Join indicates an expected call of Join.
func (mr *MockFileSystemMockRecorder) Join(elem ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Join", reflect.TypeOf((*MockFileSystem)(nil).Join), elem...)
}

// Open mocks base method.
func (m *MockFileSystem) Open(name string) (fs.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Open", name)
	ret0, _ := ret[0].(fs.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Open indicates an expected call of Open.
func (mr *MockFileSystemMockRecorder) Open(name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockFileSystem)(nil).Open), name)
}

// ReadDir mocks base method.
func (m *MockFileSystem) ReadDir(name string) ([]os.DirEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReadDir", name)
	ret0, _ := ret[0].([]os.DirEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReadDir indicates an expected call of ReadDir.
func (mr *MockFileSystemMockRecorder) ReadDir(name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReadDir", reflect.TypeOf((*MockFileSystem)(nil).ReadDir), name)
}

// MockFile is a mock of File interface.
type MockFile struct {
	ctrl     *gomock.Controller
	recorder *MockFileMockRecorder
	isgomock struct{}
}

// MockFileMockRecorder is the mock recorder for MockFile.
type MockFileMockRecorder struct {
	mock *MockFile
}

// NewMockFile creates a new mock instance.
func NewMockFile(ctrl *gomock.Controller) *MockFile {
	mock := &MockFile{ctrl: ctrl}
	mock.recorder = &MockFileMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFile) EXPECT() *MockFileMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockFile) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockFileMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockFile)(nil).Close))
}

// Read mocks base method.
func (m *MockFile) Read(p []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", p)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockFileMockRecorder) Read(p any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockFile)(nil).Read), p)
}