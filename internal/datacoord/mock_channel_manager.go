// Code generated by mockery v2.46.0. DO NOT EDIT.

package datacoord

import (
	context "context"

	datapb "github.com/milvus-io/milvus/pkg/proto/datapb"
	mock "github.com/stretchr/testify/mock"
)

// MockChannelManager is an autogenerated mock type for the ChannelManager type
type MockChannelManager struct {
	mock.Mock
}

type MockChannelManager_Expecter struct {
	mock *mock.Mock
}

func (_m *MockChannelManager) EXPECT() *MockChannelManager_Expecter {
	return &MockChannelManager_Expecter{mock: &_m.Mock}
}

// AddNode provides a mock function with given fields: nodeID
func (_m *MockChannelManager) AddNode(nodeID int64) error {
	ret := _m.Called(nodeID)

	if len(ret) == 0 {
		panic("no return value specified for AddNode")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(nodeID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockChannelManager_AddNode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddNode'
type MockChannelManager_AddNode_Call struct {
	*mock.Call
}

// AddNode is a helper method to define mock.On call
//   - nodeID int64
func (_e *MockChannelManager_Expecter) AddNode(nodeID interface{}) *MockChannelManager_AddNode_Call {
	return &MockChannelManager_AddNode_Call{Call: _e.mock.On("AddNode", nodeID)}
}

func (_c *MockChannelManager_AddNode_Call) Run(run func(nodeID int64)) *MockChannelManager_AddNode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *MockChannelManager_AddNode_Call) Return(_a0 error) *MockChannelManager_AddNode_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockChannelManager_AddNode_Call) RunAndReturn(run func(int64) error) *MockChannelManager_AddNode_Call {
	_c.Call.Return(run)
	return _c
}

// Close provides a mock function with given fields:
func (_m *MockChannelManager) Close() {
	_m.Called()
}

// MockChannelManager_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockChannelManager_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MockChannelManager_Expecter) Close() *MockChannelManager_Close_Call {
	return &MockChannelManager_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *MockChannelManager_Close_Call) Run(run func()) *MockChannelManager_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockChannelManager_Close_Call) Return() *MockChannelManager_Close_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockChannelManager_Close_Call) RunAndReturn(run func()) *MockChannelManager_Close_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteNode provides a mock function with given fields: nodeID
func (_m *MockChannelManager) DeleteNode(nodeID int64) error {
	ret := _m.Called(nodeID)

	if len(ret) == 0 {
		panic("no return value specified for DeleteNode")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(nodeID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockChannelManager_DeleteNode_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteNode'
type MockChannelManager_DeleteNode_Call struct {
	*mock.Call
}

// DeleteNode is a helper method to define mock.On call
//   - nodeID int64
func (_e *MockChannelManager_Expecter) DeleteNode(nodeID interface{}) *MockChannelManager_DeleteNode_Call {
	return &MockChannelManager_DeleteNode_Call{Call: _e.mock.On("DeleteNode", nodeID)}
}

func (_c *MockChannelManager_DeleteNode_Call) Run(run func(nodeID int64)) *MockChannelManager_DeleteNode_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *MockChannelManager_DeleteNode_Call) Return(_a0 error) *MockChannelManager_DeleteNode_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockChannelManager_DeleteNode_Call) RunAndReturn(run func(int64) error) *MockChannelManager_DeleteNode_Call {
	_c.Call.Return(run)
	return _c
}

// FindWatcher provides a mock function with given fields: channel
func (_m *MockChannelManager) FindWatcher(channel string) (int64, error) {
	ret := _m.Called(channel)

	if len(ret) == 0 {
		panic("no return value specified for FindWatcher")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (int64, error)); ok {
		return rf(channel)
	}
	if rf, ok := ret.Get(0).(func(string) int64); ok {
		r0 = rf(channel)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(channel)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockChannelManager_FindWatcher_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindWatcher'
type MockChannelManager_FindWatcher_Call struct {
	*mock.Call
}

// FindWatcher is a helper method to define mock.On call
//   - channel string
func (_e *MockChannelManager_Expecter) FindWatcher(channel interface{}) *MockChannelManager_FindWatcher_Call {
	return &MockChannelManager_FindWatcher_Call{Call: _e.mock.On("FindWatcher", channel)}
}

func (_c *MockChannelManager_FindWatcher_Call) Run(run func(channel string)) *MockChannelManager_FindWatcher_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *MockChannelManager_FindWatcher_Call) Return(_a0 int64, _a1 error) *MockChannelManager_FindWatcher_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockChannelManager_FindWatcher_Call) RunAndReturn(run func(string) (int64, error)) *MockChannelManager_FindWatcher_Call {
	_c.Call.Return(run)
	return _c
}

// GetChannel provides a mock function with given fields: nodeID, channel
func (_m *MockChannelManager) GetChannel(nodeID int64, channel string) (RWChannel, bool) {
	ret := _m.Called(nodeID, channel)

	if len(ret) == 0 {
		panic("no return value specified for GetChannel")
	}

	var r0 RWChannel
	var r1 bool
	if rf, ok := ret.Get(0).(func(int64, string) (RWChannel, bool)); ok {
		return rf(nodeID, channel)
	}
	if rf, ok := ret.Get(0).(func(int64, string) RWChannel); ok {
		r0 = rf(nodeID, channel)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(RWChannel)
		}
	}

	if rf, ok := ret.Get(1).(func(int64, string) bool); ok {
		r1 = rf(nodeID, channel)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// MockChannelManager_GetChannel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetChannel'
type MockChannelManager_GetChannel_Call struct {
	*mock.Call
}

// GetChannel is a helper method to define mock.On call
//   - nodeID int64
//   - channel string
func (_e *MockChannelManager_Expecter) GetChannel(nodeID interface{}, channel interface{}) *MockChannelManager_GetChannel_Call {
	return &MockChannelManager_GetChannel_Call{Call: _e.mock.On("GetChannel", nodeID, channel)}
}

func (_c *MockChannelManager_GetChannel_Call) Run(run func(nodeID int64, channel string)) *MockChannelManager_GetChannel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64), args[1].(string))
	})
	return _c
}

func (_c *MockChannelManager_GetChannel_Call) Return(_a0 RWChannel, _a1 bool) *MockChannelManager_GetChannel_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockChannelManager_GetChannel_Call) RunAndReturn(run func(int64, string) (RWChannel, bool)) *MockChannelManager_GetChannel_Call {
	_c.Call.Return(run)
	return _c
}

// GetChannelNamesByCollectionID provides a mock function with given fields: collectionID
func (_m *MockChannelManager) GetChannelNamesByCollectionID(collectionID int64) []string {
	ret := _m.Called(collectionID)

	if len(ret) == 0 {
		panic("no return value specified for GetChannelNamesByCollectionID")
	}

	var r0 []string
	if rf, ok := ret.Get(0).(func(int64) []string); ok {
		r0 = rf(collectionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	return r0
}

// MockChannelManager_GetChannelNamesByCollectionID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetChannelNamesByCollectionID'
type MockChannelManager_GetChannelNamesByCollectionID_Call struct {
	*mock.Call
}

// GetChannelNamesByCollectionID is a helper method to define mock.On call
//   - collectionID int64
func (_e *MockChannelManager_Expecter) GetChannelNamesByCollectionID(collectionID interface{}) *MockChannelManager_GetChannelNamesByCollectionID_Call {
	return &MockChannelManager_GetChannelNamesByCollectionID_Call{Call: _e.mock.On("GetChannelNamesByCollectionID", collectionID)}
}

func (_c *MockChannelManager_GetChannelNamesByCollectionID_Call) Run(run func(collectionID int64)) *MockChannelManager_GetChannelNamesByCollectionID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *MockChannelManager_GetChannelNamesByCollectionID_Call) Return(_a0 []string) *MockChannelManager_GetChannelNamesByCollectionID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockChannelManager_GetChannelNamesByCollectionID_Call) RunAndReturn(run func(int64) []string) *MockChannelManager_GetChannelNamesByCollectionID_Call {
	_c.Call.Return(run)
	return _c
}

// GetChannelWatchInfos provides a mock function with given fields:
func (_m *MockChannelManager) GetChannelWatchInfos() map[int64]map[string]*datapb.ChannelWatchInfo {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetChannelWatchInfos")
	}

	var r0 map[int64]map[string]*datapb.ChannelWatchInfo
	if rf, ok := ret.Get(0).(func() map[int64]map[string]*datapb.ChannelWatchInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[int64]map[string]*datapb.ChannelWatchInfo)
		}
	}

	return r0
}

// MockChannelManager_GetChannelWatchInfos_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetChannelWatchInfos'
type MockChannelManager_GetChannelWatchInfos_Call struct {
	*mock.Call
}

// GetChannelWatchInfos is a helper method to define mock.On call
func (_e *MockChannelManager_Expecter) GetChannelWatchInfos() *MockChannelManager_GetChannelWatchInfos_Call {
	return &MockChannelManager_GetChannelWatchInfos_Call{Call: _e.mock.On("GetChannelWatchInfos")}
}

func (_c *MockChannelManager_GetChannelWatchInfos_Call) Run(run func()) *MockChannelManager_GetChannelWatchInfos_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockChannelManager_GetChannelWatchInfos_Call) Return(_a0 map[int64]map[string]*datapb.ChannelWatchInfo) *MockChannelManager_GetChannelWatchInfos_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockChannelManager_GetChannelWatchInfos_Call) RunAndReturn(run func() map[int64]map[string]*datapb.ChannelWatchInfo) *MockChannelManager_GetChannelWatchInfos_Call {
	_c.Call.Return(run)
	return _c
}

// GetChannelsByCollectionID provides a mock function with given fields: collectionID
func (_m *MockChannelManager) GetChannelsByCollectionID(collectionID int64) []RWChannel {
	ret := _m.Called(collectionID)

	if len(ret) == 0 {
		panic("no return value specified for GetChannelsByCollectionID")
	}

	var r0 []RWChannel
	if rf, ok := ret.Get(0).(func(int64) []RWChannel); ok {
		r0 = rf(collectionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]RWChannel)
		}
	}

	return r0
}

// MockChannelManager_GetChannelsByCollectionID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetChannelsByCollectionID'
type MockChannelManager_GetChannelsByCollectionID_Call struct {
	*mock.Call
}

// GetChannelsByCollectionID is a helper method to define mock.On call
//   - collectionID int64
func (_e *MockChannelManager_Expecter) GetChannelsByCollectionID(collectionID interface{}) *MockChannelManager_GetChannelsByCollectionID_Call {
	return &MockChannelManager_GetChannelsByCollectionID_Call{Call: _e.mock.On("GetChannelsByCollectionID", collectionID)}
}

func (_c *MockChannelManager_GetChannelsByCollectionID_Call) Run(run func(collectionID int64)) *MockChannelManager_GetChannelsByCollectionID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *MockChannelManager_GetChannelsByCollectionID_Call) Return(_a0 []RWChannel) *MockChannelManager_GetChannelsByCollectionID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockChannelManager_GetChannelsByCollectionID_Call) RunAndReturn(run func(int64) []RWChannel) *MockChannelManager_GetChannelsByCollectionID_Call {
	_c.Call.Return(run)
	return _c
}

// GetNodeChannelsByCollectionID provides a mock function with given fields: collectionID
func (_m *MockChannelManager) GetNodeChannelsByCollectionID(collectionID int64) map[int64][]string {
	ret := _m.Called(collectionID)

	if len(ret) == 0 {
		panic("no return value specified for GetNodeChannelsByCollectionID")
	}

	var r0 map[int64][]string
	if rf, ok := ret.Get(0).(func(int64) map[int64][]string); ok {
		r0 = rf(collectionID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[int64][]string)
		}
	}

	return r0
}

// MockChannelManager_GetNodeChannelsByCollectionID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetNodeChannelsByCollectionID'
type MockChannelManager_GetNodeChannelsByCollectionID_Call struct {
	*mock.Call
}

// GetNodeChannelsByCollectionID is a helper method to define mock.On call
//   - collectionID int64
func (_e *MockChannelManager_Expecter) GetNodeChannelsByCollectionID(collectionID interface{}) *MockChannelManager_GetNodeChannelsByCollectionID_Call {
	return &MockChannelManager_GetNodeChannelsByCollectionID_Call{Call: _e.mock.On("GetNodeChannelsByCollectionID", collectionID)}
}

func (_c *MockChannelManager_GetNodeChannelsByCollectionID_Call) Run(run func(collectionID int64)) *MockChannelManager_GetNodeChannelsByCollectionID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *MockChannelManager_GetNodeChannelsByCollectionID_Call) Return(_a0 map[int64][]string) *MockChannelManager_GetNodeChannelsByCollectionID_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockChannelManager_GetNodeChannelsByCollectionID_Call) RunAndReturn(run func(int64) map[int64][]string) *MockChannelManager_GetNodeChannelsByCollectionID_Call {
	_c.Call.Return(run)
	return _c
}

// Match provides a mock function with given fields: nodeID, channel
func (_m *MockChannelManager) Match(nodeID int64, channel string) bool {
	ret := _m.Called(nodeID, channel)

	if len(ret) == 0 {
		panic("no return value specified for Match")
	}

	var r0 bool
	if rf, ok := ret.Get(0).(func(int64, string) bool); ok {
		r0 = rf(nodeID, channel)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}

// MockChannelManager_Match_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Match'
type MockChannelManager_Match_Call struct {
	*mock.Call
}

// Match is a helper method to define mock.On call
//   - nodeID int64
//   - channel string
func (_e *MockChannelManager_Expecter) Match(nodeID interface{}, channel interface{}) *MockChannelManager_Match_Call {
	return &MockChannelManager_Match_Call{Call: _e.mock.On("Match", nodeID, channel)}
}

func (_c *MockChannelManager_Match_Call) Run(run func(nodeID int64, channel string)) *MockChannelManager_Match_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64), args[1].(string))
	})
	return _c
}

func (_c *MockChannelManager_Match_Call) Return(_a0 bool) *MockChannelManager_Match_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockChannelManager_Match_Call) RunAndReturn(run func(int64, string) bool) *MockChannelManager_Match_Call {
	_c.Call.Return(run)
	return _c
}

// Release provides a mock function with given fields: nodeID, channelName
func (_m *MockChannelManager) Release(nodeID int64, channelName string) error {
	ret := _m.Called(nodeID, channelName)

	if len(ret) == 0 {
		panic("no return value specified for Release")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, string) error); ok {
		r0 = rf(nodeID, channelName)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockChannelManager_Release_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Release'
type MockChannelManager_Release_Call struct {
	*mock.Call
}

// Release is a helper method to define mock.On call
//   - nodeID int64
//   - channelName string
func (_e *MockChannelManager_Expecter) Release(nodeID interface{}, channelName interface{}) *MockChannelManager_Release_Call {
	return &MockChannelManager_Release_Call{Call: _e.mock.On("Release", nodeID, channelName)}
}

func (_c *MockChannelManager_Release_Call) Run(run func(nodeID int64, channelName string)) *MockChannelManager_Release_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64), args[1].(string))
	})
	return _c
}

func (_c *MockChannelManager_Release_Call) Return(_a0 error) *MockChannelManager_Release_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockChannelManager_Release_Call) RunAndReturn(run func(int64, string) error) *MockChannelManager_Release_Call {
	_c.Call.Return(run)
	return _c
}

// Startup provides a mock function with given fields: ctx, legacyNodes, allNodes
func (_m *MockChannelManager) Startup(ctx context.Context, legacyNodes []int64, allNodes []int64) error {
	ret := _m.Called(ctx, legacyNodes, allNodes)

	if len(ret) == 0 {
		panic("no return value specified for Startup")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, []int64, []int64) error); ok {
		r0 = rf(ctx, legacyNodes, allNodes)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockChannelManager_Startup_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Startup'
type MockChannelManager_Startup_Call struct {
	*mock.Call
}

// Startup is a helper method to define mock.On call
//   - ctx context.Context
//   - legacyNodes []int64
//   - allNodes []int64
func (_e *MockChannelManager_Expecter) Startup(ctx interface{}, legacyNodes interface{}, allNodes interface{}) *MockChannelManager_Startup_Call {
	return &MockChannelManager_Startup_Call{Call: _e.mock.On("Startup", ctx, legacyNodes, allNodes)}
}

func (_c *MockChannelManager_Startup_Call) Run(run func(ctx context.Context, legacyNodes []int64, allNodes []int64)) *MockChannelManager_Startup_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].([]int64), args[2].([]int64))
	})
	return _c
}

func (_c *MockChannelManager_Startup_Call) Return(_a0 error) *MockChannelManager_Startup_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockChannelManager_Startup_Call) RunAndReturn(run func(context.Context, []int64, []int64) error) *MockChannelManager_Startup_Call {
	_c.Call.Return(run)
	return _c
}

// Watch provides a mock function with given fields: ctx, ch
func (_m *MockChannelManager) Watch(ctx context.Context, ch RWChannel) error {
	ret := _m.Called(ctx, ch)

	if len(ret) == 0 {
		panic("no return value specified for Watch")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, RWChannel) error); ok {
		r0 = rf(ctx, ch)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockChannelManager_Watch_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Watch'
type MockChannelManager_Watch_Call struct {
	*mock.Call
}

// Watch is a helper method to define mock.On call
//   - ctx context.Context
//   - ch RWChannel
func (_e *MockChannelManager_Expecter) Watch(ctx interface{}, ch interface{}) *MockChannelManager_Watch_Call {
	return &MockChannelManager_Watch_Call{Call: _e.mock.On("Watch", ctx, ch)}
}

func (_c *MockChannelManager_Watch_Call) Run(run func(ctx context.Context, ch RWChannel)) *MockChannelManager_Watch_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(RWChannel))
	})
	return _c
}

func (_c *MockChannelManager_Watch_Call) Return(_a0 error) *MockChannelManager_Watch_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockChannelManager_Watch_Call) RunAndReturn(run func(context.Context, RWChannel) error) *MockChannelManager_Watch_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockChannelManager creates a new instance of MockChannelManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockChannelManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockChannelManager {
	mock := &MockChannelManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
