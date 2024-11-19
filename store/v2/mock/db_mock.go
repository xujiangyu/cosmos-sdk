// Code generated by MockGen. DO NOT EDIT.
// Source: ./types.go
//
// Generated by this command:
//
//	mockgen -package mock -destination ./db_mock.go -source ./types.go
//

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	store "cosmossdk.io/core/store"
	proof "cosmossdk.io/store/v2/proof"
	gomock "go.uber.org/mock/gomock"
)

// MockStateCommitter is a mock of StateCommitter interface.
type MockStateCommitter struct {
	ctrl     *gomock.Controller
	recorder *MockStateCommitterMockRecorder
	isgomock struct{}
}

// MockStateCommitterMockRecorder is the mock recorder for MockStateCommitter.
type MockStateCommitterMockRecorder struct {
	mock *MockStateCommitter
}

// NewMockStateCommitter creates a new mock instance.
func NewMockStateCommitter(ctrl *gomock.Controller) *MockStateCommitter {
	mock := &MockStateCommitter{ctrl: ctrl}
	mock.recorder = &MockStateCommitterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStateCommitter) EXPECT() *MockStateCommitterMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockStateCommitter) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockStateCommitterMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockStateCommitter)(nil).Close))
}

// Commit mocks base method.
func (m *MockStateCommitter) Commit(version uint64) (*proof.CommitInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit", version)
	ret0, _ := ret[0].(*proof.CommitInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Commit indicates an expected call of Commit.
func (mr *MockStateCommitterMockRecorder) Commit(version any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockStateCommitter)(nil).Commit), version)
}

// Get mocks base method.
func (m *MockStateCommitter) Get(storeKey []byte, version uint64, key []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", storeKey, version, key)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockStateCommitterMockRecorder) Get(storeKey, version, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockStateCommitter)(nil).Get), storeKey, version, key)
}

// GetCommitInfo mocks base method.
func (m *MockStateCommitter) GetCommitInfo(version uint64) (*proof.CommitInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommitInfo", version)
	ret0, _ := ret[0].(*proof.CommitInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommitInfo indicates an expected call of GetCommitInfo.
func (mr *MockStateCommitterMockRecorder) GetCommitInfo(version any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommitInfo", reflect.TypeOf((*MockStateCommitter)(nil).GetCommitInfo), version)
}

// GetLatestVersion mocks base method.
func (m *MockStateCommitter) GetLatestVersion() (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestVersion")
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestVersion indicates an expected call of GetLatestVersion.
func (mr *MockStateCommitterMockRecorder) GetLatestVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestVersion", reflect.TypeOf((*MockStateCommitter)(nil).GetLatestVersion))
}

// GetProof mocks base method.
func (m *MockStateCommitter) GetProof(storeKey []byte, version uint64, key []byte) ([]proof.CommitmentOp, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProof", storeKey, version, key)
	ret0, _ := ret[0].([]proof.CommitmentOp)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProof indicates an expected call of GetProof.
func (mr *MockStateCommitterMockRecorder) GetProof(storeKey, version, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProof", reflect.TypeOf((*MockStateCommitter)(nil).GetProof), storeKey, version, key)
}

// LoadVersion mocks base method.
func (m *MockStateCommitter) LoadVersion(targetVersion uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadVersion", targetVersion)
	ret0, _ := ret[0].(error)
	return ret0
}

// LoadVersion indicates an expected call of LoadVersion.
func (mr *MockStateCommitterMockRecorder) LoadVersion(targetVersion any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadVersion", reflect.TypeOf((*MockStateCommitter)(nil).LoadVersion), targetVersion)
}

// LoadVersionAndUpgrade mocks base method.
func (m *MockStateCommitter) LoadVersionAndUpgrade(version uint64, upgrades *store.StoreUpgrades) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadVersionAndUpgrade", version, upgrades)
	ret0, _ := ret[0].(error)
	return ret0
}

// LoadVersionAndUpgrade indicates an expected call of LoadVersionAndUpgrade.
func (mr *MockStateCommitterMockRecorder) LoadVersionAndUpgrade(version, upgrades any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadVersionAndUpgrade", reflect.TypeOf((*MockStateCommitter)(nil).LoadVersionAndUpgrade), version, upgrades)
}

// PausePruning mocks base method.
func (m *MockStateCommitter) PausePruning(pause bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PausePruning", pause)
}

// PausePruning indicates an expected call of PausePruning.
func (mr *MockStateCommitterMockRecorder) PausePruning(pause any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PausePruning", reflect.TypeOf((*MockStateCommitter)(nil).PausePruning), pause)
}

// Prune mocks base method.
func (m *MockStateCommitter) Prune(version uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Prune", version)
	ret0, _ := ret[0].(error)
	return ret0
}

// Prune indicates an expected call of Prune.
func (mr *MockStateCommitterMockRecorder) Prune(version any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prune", reflect.TypeOf((*MockStateCommitter)(nil).Prune), version)
}

// SetInitialVersion mocks base method.
func (m *MockStateCommitter) SetInitialVersion(version uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetInitialVersion", version)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetInitialVersion indicates an expected call of SetInitialVersion.
func (mr *MockStateCommitterMockRecorder) SetInitialVersion(version any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetInitialVersion", reflect.TypeOf((*MockStateCommitter)(nil).SetInitialVersion), version)
}

// WriteChangeset mocks base method.
func (m *MockStateCommitter) WriteChangeset(cs *store.Changeset) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteChangeset", cs)
	ret0, _ := ret[0].(error)
	return ret0
}

// WriteChangeset indicates an expected call of WriteChangeset.
func (mr *MockStateCommitterMockRecorder) WriteChangeset(cs any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteChangeset", reflect.TypeOf((*MockStateCommitter)(nil).WriteChangeset), cs)
}

// MockStateStorage is a mock of StateStorage interface.
type MockStateStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStateStorageMockRecorder
	isgomock struct{}
}

// MockStateStorageMockRecorder is the mock recorder for MockStateStorage.
type MockStateStorageMockRecorder struct {
	mock *MockStateStorage
}

// NewMockStateStorage creates a new mock instance.
func NewMockStateStorage(ctrl *gomock.Controller) *MockStateStorage {
	mock := &MockStateStorage{ctrl: ctrl}
	mock.recorder = &MockStateStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStateStorage) EXPECT() *MockStateStorageMockRecorder {
	return m.recorder
}

// ApplyChangeset mocks base method.
func (m *MockStateStorage) ApplyChangeset(cs *store.Changeset) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ApplyChangeset", cs)
	ret0, _ := ret[0].(error)
	return ret0
}

// ApplyChangeset indicates an expected call of ApplyChangeset.
func (mr *MockStateStorageMockRecorder) ApplyChangeset(cs any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ApplyChangeset", reflect.TypeOf((*MockStateStorage)(nil).ApplyChangeset), cs)
}

// Close mocks base method.
func (m *MockStateStorage) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockStateStorageMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockStateStorage)(nil).Close))
}

// Get mocks base method.
func (m *MockStateStorage) Get(storeKey []byte, version uint64, key []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", storeKey, version, key)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockStateStorageMockRecorder) Get(storeKey, version, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockStateStorage)(nil).Get), storeKey, version, key)
}

// GetLatestVersion mocks base method.
func (m *MockStateStorage) GetLatestVersion() (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLatestVersion")
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLatestVersion indicates an expected call of GetLatestVersion.
func (mr *MockStateStorageMockRecorder) GetLatestVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLatestVersion", reflect.TypeOf((*MockStateStorage)(nil).GetLatestVersion))
}

// Has mocks base method.
func (m *MockStateStorage) Has(storeKey []byte, version uint64, key []byte) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", storeKey, version, key)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Has indicates an expected call of Has.
func (mr *MockStateStorageMockRecorder) Has(storeKey, version, key any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockStateStorage)(nil).Has), storeKey, version, key)
}

// Iterator mocks base method.
func (m *MockStateStorage) Iterator(storeKey []byte, version uint64, start, end []byte) (store.Iterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Iterator", storeKey, version, start, end)
	ret0, _ := ret[0].(store.Iterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Iterator indicates an expected call of Iterator.
func (mr *MockStateStorageMockRecorder) Iterator(storeKey, version, start, end any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Iterator", reflect.TypeOf((*MockStateStorage)(nil).Iterator), storeKey, version, start, end)
}

// PausePruning mocks base method.
func (m *MockStateStorage) PausePruning(pause bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "PausePruning", pause)
}

// PausePruning indicates an expected call of PausePruning.
func (mr *MockStateStorageMockRecorder) PausePruning(pause any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PausePruning", reflect.TypeOf((*MockStateStorage)(nil).PausePruning), pause)
}

// Prune mocks base method.
func (m *MockStateStorage) Prune(version uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Prune", version)
	ret0, _ := ret[0].(error)
	return ret0
}

// Prune indicates an expected call of Prune.
func (mr *MockStateStorageMockRecorder) Prune(version any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prune", reflect.TypeOf((*MockStateStorage)(nil).Prune), version)
}

// PruneStoreKeys mocks base method.
func (m *MockStateStorage) PruneStoreKeys(storeKeys []string, version uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PruneStoreKeys", storeKeys, version)
	ret0, _ := ret[0].(error)
	return ret0
}

// PruneStoreKeys indicates an expected call of PruneStoreKeys.
func (mr *MockStateStorageMockRecorder) PruneStoreKeys(storeKeys, version any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PruneStoreKeys", reflect.TypeOf((*MockStateStorage)(nil).PruneStoreKeys), storeKeys, version)
}

// ReverseIterator mocks base method.
func (m *MockStateStorage) ReverseIterator(storeKey []byte, version uint64, start, end []byte) (store.Iterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReverseIterator", storeKey, version, start, end)
	ret0, _ := ret[0].(store.Iterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReverseIterator indicates an expected call of ReverseIterator.
func (mr *MockStateStorageMockRecorder) ReverseIterator(storeKey, version, start, end any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReverseIterator", reflect.TypeOf((*MockStateStorage)(nil).ReverseIterator), storeKey, version, start, end)
}

// SetLatestVersion mocks base method.
func (m *MockStateStorage) SetLatestVersion(version uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetLatestVersion", version)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetLatestVersion indicates an expected call of SetLatestVersion.
func (mr *MockStateStorageMockRecorder) SetLatestVersion(version any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLatestVersion", reflect.TypeOf((*MockStateStorage)(nil).SetLatestVersion), version)
}

// VersionExists mocks base method.
func (m *MockStateStorage) VersionExists(v uint64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VersionExists", v)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VersionExists indicates an expected call of VersionExists.
func (mr *MockStateStorageMockRecorder) VersionExists(v any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VersionExists", reflect.TypeOf((*MockStateStorage)(nil).VersionExists), v)
}
