// Code generated by MockGen. DO NOT EDIT.
// Source: repos.go

package repos

import (
	v1alpha1 "github.com/fluxcd/source-controller/api/v1alpha1"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRepos is a mock of Repos interface
type MockRepos struct {
	ctrl     *gomock.Controller
	recorder *MockReposMockRecorder
}

// MockReposMockRecorder is the mock recorder for MockRepos
type MockReposMockRecorder struct {
	mock *MockRepos
}

// NewMockRepos creates a new mock instance
func NewMockRepos(ctrl *gomock.Controller) *MockRepos {
	mock := &MockRepos{ctrl: ctrl}
	mock.recorder = &MockReposMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockRepos) EXPECT() *MockReposMockRecorder {
	return _m.recorder
}

// Add mocks base method
func (_m *MockRepos) Add(srcRepo *v1alpha1.GitRepository) Repo {
	ret := _m.ctrl.Call(_m, "Add", srcRepo)
	ret0, _ := ret[0].(Repo)
	return ret0
}

// Add indicates an expected call of Add
func (_mr *MockReposMockRecorder) Add(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Add", reflect.TypeOf((*MockRepos)(nil).Add), arg0)
}

// Get mocks base method
func (_m *MockRepos) Get(name string) Repo {
	ret := _m.ctrl.Call(_m, "Get", name)
	ret0, _ := ret[0].(Repo)
	return ret0
}

// Get indicates an expected call of Get
func (_mr *MockReposMockRecorder) Get(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Get", reflect.TypeOf((*MockRepos)(nil).Get), arg0)
}

// Delete mocks base method
func (_m *MockRepos) Delete(name string) {
	_m.ctrl.Call(_m, "Delete", name)
}

// Delete indicates an expected call of Delete
func (_mr *MockReposMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Delete", reflect.TypeOf((*MockRepos)(nil).Delete), arg0)
}

// List mocks base method
func (_m *MockRepos) List() map[string]Repo {
	ret := _m.ctrl.Call(_m, "List")
	ret0, _ := ret[0].(map[string]Repo)
	return ret0
}

// List indicates an expected call of List
func (_mr *MockReposMockRecorder) List() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "List", reflect.TypeOf((*MockRepos)(nil).List))
}

// MockRepo is a mock of Repo interface
type MockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRepoMockRecorder
}

// MockRepoMockRecorder is the mock recorder for MockRepo
type MockRepoMockRecorder struct {
	mock *MockRepo
}

// NewMockRepo creates a new mock instance
func NewMockRepo(ctrl *gomock.Controller) *MockRepo {
	mock := &MockRepo{ctrl: ctrl}
	mock.recorder = &MockRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockRepo) EXPECT() *MockRepoMockRecorder {
	return _m.recorder
}

// GetName mocks base method
func (_m *MockRepo) GetName() string {
	ret := _m.ctrl.Call(_m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName
func (_mr *MockRepoMockRecorder) GetName() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetName", reflect.TypeOf((*MockRepo)(nil).GetName))
}

// GetSourceName mocks base method
func (_m *MockRepo) GetSourceName() string {
	ret := _m.ctrl.Call(_m, "GetSourceName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetSourceName indicates an expected call of GetSourceName
func (_mr *MockRepoMockRecorder) GetSourceName() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetSourceName", reflect.TypeOf((*MockRepo)(nil).GetSourceName))
}

// GetSourceNameSpace mocks base method
func (_m *MockRepo) GetSourceNameSpace() string {
	ret := _m.ctrl.Call(_m, "GetSourceNameSpace")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetSourceNameSpace indicates an expected call of GetSourceNameSpace
func (_mr *MockRepoMockRecorder) GetSourceNameSpace() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetSourceNameSpace", reflect.TypeOf((*MockRepo)(nil).GetSourceNameSpace))
}

// SyncRepo mocks base method
func (_m *MockRepo) SyncRepo() error {
	ret := _m.ctrl.Call(_m, "SyncRepo")
	ret0, _ := ret[0].(error)
	return ret0
}

// SyncRepo indicates an expected call of SyncRepo
func (_mr *MockRepoMockRecorder) SyncRepo() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "SyncRepo", reflect.TypeOf((*MockRepo)(nil).SyncRepo))
}

// GetDataPath mocks base method
func (_m *MockRepo) GetDataPath() string {
	ret := _m.ctrl.Call(_m, "GetDataPath")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetDataPath indicates an expected call of GetDataPath
func (_mr *MockRepoMockRecorder) GetDataPath() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "GetDataPath", reflect.TypeOf((*MockRepo)(nil).GetDataPath))
}

// LinkData mocks base method
func (_m *MockRepo) LinkData(layerPath string, sourcePath string) error {
	ret := _m.ctrl.Call(_m, "LinkData", layerPath, sourcePath)
	ret0, _ := ret[0].(error)
	return ret0
}

// LinkData indicates an expected call of LinkData
func (_mr *MockRepoMockRecorder) LinkData(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "LinkData", reflect.TypeOf((*MockRepo)(nil).LinkData), arg0, arg1)
}
