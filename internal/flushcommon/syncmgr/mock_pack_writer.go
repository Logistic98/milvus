// Code generated by mockery v2.46.0. DO NOT EDIT.

package syncmgr

import (
	context "context"

	datapb "github.com/milvus-io/milvus/pkg/proto/datapb"
	mock "github.com/stretchr/testify/mock"
)

// MockPackWriter is an autogenerated mock type for the PackWriter type
type MockPackWriter struct {
	mock.Mock
}

type MockPackWriter_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPackWriter) EXPECT() *MockPackWriter_Expecter {
	return &MockPackWriter_Expecter{mock: &_m.Mock}
}

// Write provides a mock function with given fields: ctx, pack
func (_m *MockPackWriter) Write(ctx context.Context, pack *SyncPack) ([]*datapb.Binlog, *datapb.Binlog, *datapb.Binlog, *datapb.Binlog, int64, error) {
	ret := _m.Called(ctx, pack)

	if len(ret) == 0 {
		panic("no return value specified for Write")
	}

	var r0 []*datapb.Binlog
	var r1 *datapb.Binlog
	var r2 *datapb.Binlog
	var r3 *datapb.Binlog
	var r4 int64
	var r5 error
	if rf, ok := ret.Get(0).(func(context.Context, *SyncPack) ([]*datapb.Binlog, *datapb.Binlog, *datapb.Binlog, *datapb.Binlog, int64, error)); ok {
		return rf(ctx, pack)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *SyncPack) []*datapb.Binlog); ok {
		r0 = rf(ctx, pack)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*datapb.Binlog)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *SyncPack) *datapb.Binlog); ok {
		r1 = rf(ctx, pack)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*datapb.Binlog)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, *SyncPack) *datapb.Binlog); ok {
		r2 = rf(ctx, pack)
	} else {
		if ret.Get(2) != nil {
			r2 = ret.Get(2).(*datapb.Binlog)
		}
	}

	if rf, ok := ret.Get(3).(func(context.Context, *SyncPack) *datapb.Binlog); ok {
		r3 = rf(ctx, pack)
	} else {
		if ret.Get(3) != nil {
			r3 = ret.Get(3).(*datapb.Binlog)
		}
	}

	if rf, ok := ret.Get(4).(func(context.Context, *SyncPack) int64); ok {
		r4 = rf(ctx, pack)
	} else {
		r4 = ret.Get(4).(int64)
	}

	if rf, ok := ret.Get(5).(func(context.Context, *SyncPack) error); ok {
		r5 = rf(ctx, pack)
	} else {
		r5 = ret.Error(5)
	}

	return r0, r1, r2, r3, r4, r5
}

// MockPackWriter_Write_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Write'
type MockPackWriter_Write_Call struct {
	*mock.Call
}

// Write is a helper method to define mock.On call
//   - ctx context.Context
//   - pack *SyncPack
func (_e *MockPackWriter_Expecter) Write(ctx interface{}, pack interface{}) *MockPackWriter_Write_Call {
	return &MockPackWriter_Write_Call{Call: _e.mock.On("Write", ctx, pack)}
}

func (_c *MockPackWriter_Write_Call) Run(run func(ctx context.Context, pack *SyncPack)) *MockPackWriter_Write_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*SyncPack))
	})
	return _c
}

func (_c *MockPackWriter_Write_Call) Return(inserts []*datapb.Binlog, deletes *datapb.Binlog, stats *datapb.Binlog, bm25Stats *datapb.Binlog, size int64, err error) *MockPackWriter_Write_Call {
	_c.Call.Return(inserts, deletes, stats, bm25Stats, size, err)
	return _c
}

func (_c *MockPackWriter_Write_Call) RunAndReturn(run func(context.Context, *SyncPack) ([]*datapb.Binlog, *datapb.Binlog, *datapb.Binlog, *datapb.Binlog, int64, error)) *MockPackWriter_Write_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPackWriter creates a new instance of MockPackWriter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPackWriter(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPackWriter {
	mock := &MockPackWriter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
