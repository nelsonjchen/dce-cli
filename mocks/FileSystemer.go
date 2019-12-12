// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	os "os"

	mock "github.com/stretchr/testify/mock"
)

// FileSystemer is an autogenerated mock type for the FileSystemer type
type FileSystemer struct {
	mock.Mock
}

// Chdir provides a mock function with given fields: path
func (_m *FileSystemer) Chdir(path string) {
	_m.Called(path)
}

// GetConfigFile provides a mock function with given fields:
func (_m *FileSystemer) GetConfigFile() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// GetHomeDir provides a mock function with given fields:
func (_m *FileSystemer) GetHomeDir() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// IsExistingFile provides a mock function with given fields: path
func (_m *FileSystemer) IsExistingFile(path string) bool {
	ret := _m.Called(path)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MvToTempDir provides a mock function with given fields: prefix
func (_m *FileSystemer) MvToTempDir(prefix string) (string, string) {
	ret := _m.Called(prefix)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(prefix)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 string
	if rf, ok := ret.Get(1).(func(string) string); ok {
		r1 = rf(prefix)
	} else {
		r1 = ret.Get(1).(string)
	}

	return r0, r1
}

// ReadDir provides a mock function with given fields: path
func (_m *FileSystemer) ReadDir(path string) []os.FileInfo {
	ret := _m.Called(path)

	var r0 []os.FileInfo
	if rf, ok := ret.Get(0).(func(string) []os.FileInfo); ok {
		r0 = rf(path)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]os.FileInfo)
		}
	}

	return r0
}

// ReadFromFile provides a mock function with given fields: path
func (_m *FileSystemer) ReadFromFile(path string) string {
	ret := _m.Called(path)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(path)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// ReadInConfig provides a mock function with given fields:
func (_m *FileSystemer) ReadInConfig() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RemoveAll provides a mock function with given fields: path
func (_m *FileSystemer) RemoveAll(path string) {
	_m.Called(path)
}

// Unarchive provides a mock function with given fields: source, destination
func (_m *FileSystemer) Unarchive(source string, destination string) {
	_m.Called(source, destination)
}

// WriteConfig provides a mock function with given fields:
func (_m *FileSystemer) WriteConfig() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WriteFile provides a mock function with given fields: fileName, data
func (_m *FileSystemer) WriteFile(fileName string, data string) {
	_m.Called(fileName, data)
}
