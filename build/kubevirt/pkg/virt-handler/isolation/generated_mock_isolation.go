// Automatically generated by MockGen. DO NOT EDIT!
// Source: isolation.go

package isolation

import (
	gomock "github.com/golang/mock/gomock"
	mountinfo "github.com/moby/sys/mountinfo"
)

// Mock of IsolationResult interface
type MockIsolationResult struct {
	ctrl     *gomock.Controller
	recorder *_MockIsolationResultRecorder
}

// Recorder for MockIsolationResult (not exported)
type _MockIsolationResultRecorder struct {
	mock *MockIsolationResult
}

func NewMockIsolationResult(ctrl *gomock.Controller) *MockIsolationResult {
	mock := &MockIsolationResult{ctrl: ctrl}
	mock.recorder = &_MockIsolationResultRecorder{mock}
	return mock
}

func (_m *MockIsolationResult) EXPECT() *_MockIsolationResultRecorder {
	return _m.recorder
}

func (_m *MockIsolationResult) Pid() int {
	ret := _m.ctrl.Call(_m, "Pid")
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockIsolationResultRecorder) Pid() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Pid")
}

func (_m *MockIsolationResult) PPid() int {
	ret := _m.ctrl.Call(_m, "PPid")
	ret0, _ := ret[0].(int)
	return ret0
}

func (_mr *_MockIsolationResultRecorder) PPid() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PPid")
}

func (_m *MockIsolationResult) PIDNamespace() string {
	ret := _m.ctrl.Call(_m, "PIDNamespace")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockIsolationResultRecorder) PIDNamespace() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PIDNamespace")
}

func (_m *MockIsolationResult) MountRoot() string {
	ret := _m.ctrl.Call(_m, "MountRoot")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockIsolationResultRecorder) MountRoot() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MountRoot")
}

func (_m *MockIsolationResult) MountNamespace() string {
	ret := _m.ctrl.Call(_m, "MountNamespace")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockIsolationResultRecorder) MountNamespace() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "MountNamespace")
}

func (_m *MockIsolationResult) Mounts(_param0 mountinfo.FilterFunc) ([]*mountinfo.Info, error) {
	ret := _m.ctrl.Call(_m, "Mounts", _param0)
	ret0, _ := ret[0].([]*mountinfo.Info)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockIsolationResultRecorder) Mounts(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Mounts", arg0)
}
