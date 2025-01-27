// Code generated by mockery v2.46.0. DO NOT EDIT.

package mock_manager

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	types "github.com/milvus-io/milvus/pkg/streaming/util/types"
)

// MockManagerClient is an autogenerated mock type for the ManagerClient type
type MockManagerClient struct {
	mock.Mock
}

type MockManagerClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockManagerClient) EXPECT() *MockManagerClient_Expecter {
	return &MockManagerClient_Expecter{mock: &_m.Mock}
}

// Assign provides a mock function with given fields: ctx, pchannel
func (_m *MockManagerClient) Assign(ctx context.Context, pchannel types.PChannelInfoAssigned) error {
	ret := _m.Called(ctx, pchannel)

	if len(ret) == 0 {
		panic("no return value specified for Assign")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, types.PChannelInfoAssigned) error); ok {
		r0 = rf(ctx, pchannel)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockManagerClient_Assign_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Assign'
type MockManagerClient_Assign_Call struct {
	*mock.Call
}

// Assign is a helper method to define mock.On call
//   - ctx context.Context
//   - pchannel types.PChannelInfoAssigned
func (_e *MockManagerClient_Expecter) Assign(ctx interface{}, pchannel interface{}) *MockManagerClient_Assign_Call {
	return &MockManagerClient_Assign_Call{Call: _e.mock.On("Assign", ctx, pchannel)}
}

func (_c *MockManagerClient_Assign_Call) Run(run func(ctx context.Context, pchannel types.PChannelInfoAssigned)) *MockManagerClient_Assign_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(types.PChannelInfoAssigned))
	})
	return _c
}

func (_c *MockManagerClient_Assign_Call) Return(_a0 error) *MockManagerClient_Assign_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockManagerClient_Assign_Call) RunAndReturn(run func(context.Context, types.PChannelInfoAssigned) error) *MockManagerClient_Assign_Call {
	_c.Call.Return(run)
	return _c
}

// Close provides a mock function with given fields:
func (_m *MockManagerClient) Close() {
	_m.Called()
}

// MockManagerClient_Close_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Close'
type MockManagerClient_Close_Call struct {
	*mock.Call
}

// Close is a helper method to define mock.On call
func (_e *MockManagerClient_Expecter) Close() *MockManagerClient_Close_Call {
	return &MockManagerClient_Close_Call{Call: _e.mock.On("Close")}
}

func (_c *MockManagerClient_Close_Call) Run(run func()) *MockManagerClient_Close_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockManagerClient_Close_Call) Return() *MockManagerClient_Close_Call {
	_c.Call.Return()
	return _c
}

func (_c *MockManagerClient_Close_Call) RunAndReturn(run func()) *MockManagerClient_Close_Call {
	_c.Call.Return(run)
	return _c
}

// CollectAllStatus provides a mock function with given fields: ctx
func (_m *MockManagerClient) CollectAllStatus(ctx context.Context) (map[int64]*types.StreamingNodeStatus, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for CollectAllStatus")
	}

	var r0 map[int64]*types.StreamingNodeStatus
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (map[int64]*types.StreamingNodeStatus, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) map[int64]*types.StreamingNodeStatus); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[int64]*types.StreamingNodeStatus)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockManagerClient_CollectAllStatus_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CollectAllStatus'
type MockManagerClient_CollectAllStatus_Call struct {
	*mock.Call
}

// CollectAllStatus is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockManagerClient_Expecter) CollectAllStatus(ctx interface{}) *MockManagerClient_CollectAllStatus_Call {
	return &MockManagerClient_CollectAllStatus_Call{Call: _e.mock.On("CollectAllStatus", ctx)}
}

func (_c *MockManagerClient_CollectAllStatus_Call) Run(run func(ctx context.Context)) *MockManagerClient_CollectAllStatus_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockManagerClient_CollectAllStatus_Call) Return(_a0 map[int64]*types.StreamingNodeStatus, _a1 error) *MockManagerClient_CollectAllStatus_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockManagerClient_CollectAllStatus_Call) RunAndReturn(run func(context.Context) (map[int64]*types.StreamingNodeStatus, error)) *MockManagerClient_CollectAllStatus_Call {
	_c.Call.Return(run)
	return _c
}

// Remove provides a mock function with given fields: ctx, pchannel
func (_m *MockManagerClient) Remove(ctx context.Context, pchannel types.PChannelInfoAssigned) error {
	ret := _m.Called(ctx, pchannel)

	if len(ret) == 0 {
		panic("no return value specified for Remove")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, types.PChannelInfoAssigned) error); ok {
		r0 = rf(ctx, pchannel)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockManagerClient_Remove_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Remove'
type MockManagerClient_Remove_Call struct {
	*mock.Call
}

// Remove is a helper method to define mock.On call
//   - ctx context.Context
//   - pchannel types.PChannelInfoAssigned
func (_e *MockManagerClient_Expecter) Remove(ctx interface{}, pchannel interface{}) *MockManagerClient_Remove_Call {
	return &MockManagerClient_Remove_Call{Call: _e.mock.On("Remove", ctx, pchannel)}
}

func (_c *MockManagerClient_Remove_Call) Run(run func(ctx context.Context, pchannel types.PChannelInfoAssigned)) *MockManagerClient_Remove_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(types.PChannelInfoAssigned))
	})
	return _c
}

func (_c *MockManagerClient_Remove_Call) Return(_a0 error) *MockManagerClient_Remove_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockManagerClient_Remove_Call) RunAndReturn(run func(context.Context, types.PChannelInfoAssigned) error) *MockManagerClient_Remove_Call {
	_c.Call.Return(run)
	return _c
}

// WatchNodeChanged provides a mock function with given fields: ctx
func (_m *MockManagerClient) WatchNodeChanged(ctx context.Context) (<-chan struct{}, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for WatchNodeChanged")
	}

	var r0 <-chan struct{}
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (<-chan struct{}, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) <-chan struct{}); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan struct{})
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockManagerClient_WatchNodeChanged_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WatchNodeChanged'
type MockManagerClient_WatchNodeChanged_Call struct {
	*mock.Call
}

// WatchNodeChanged is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockManagerClient_Expecter) WatchNodeChanged(ctx interface{}) *MockManagerClient_WatchNodeChanged_Call {
	return &MockManagerClient_WatchNodeChanged_Call{Call: _e.mock.On("WatchNodeChanged", ctx)}
}

func (_c *MockManagerClient_WatchNodeChanged_Call) Run(run func(ctx context.Context)) *MockManagerClient_WatchNodeChanged_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockManagerClient_WatchNodeChanged_Call) Return(_a0 <-chan struct{}, _a1 error) *MockManagerClient_WatchNodeChanged_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockManagerClient_WatchNodeChanged_Call) RunAndReturn(run func(context.Context) (<-chan struct{}, error)) *MockManagerClient_WatchNodeChanged_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockManagerClient creates a new instance of MockManagerClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockManagerClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockManagerClient {
	mock := &MockManagerClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
