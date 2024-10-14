// Code generated by mockery v2.41.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "github.com/babylonlabs-io/babylon-staking-indexer/internal/db/model"
)

// DbInterface is an autogenerated mock type for the DbInterface type
type DbInterface struct {
	mock.Mock
}

// GetFinalityProviderByBtcPk provides a mock function with given fields: ctx, btcPk
func (_m *DbInterface) GetFinalityProviderByBtcPk(ctx context.Context, btcPk string) (model.FinalityProviderDetails, error) {
	ret := _m.Called(ctx, btcPk)

	if len(ret) == 0 {
		panic("no return value specified for GetFinalityProviderByBtcPk")
	}

	var r0 model.FinalityProviderDetails
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (model.FinalityProviderDetails, error)); ok {
		return rf(ctx, btcPk)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) model.FinalityProviderDetails); ok {
		r0 = rf(ctx, btcPk)
	} else {
		r0 = ret.Get(0).(model.FinalityProviderDetails)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, btcPk)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Ping provides a mock function with given fields: ctx
func (_m *DbInterface) Ping(ctx context.Context) error {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for Ping")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SaveNewFinalityProvider provides a mock function with given fields: ctx, fpDoc
func (_m *DbInterface) SaveNewFinalityProvider(ctx context.Context, fpDoc *model.FinalityProviderDetails) error {
	ret := _m.Called(ctx, fpDoc)

	if len(ret) == 0 {
		panic("no return value specified for SaveNewFinalityProvider")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.FinalityProviderDetails) error); ok {
		r0 = rf(ctx, fpDoc)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateFinalityProviderDetailsFromEvent provides a mock function with given fields: ctx, detailsToUpdate
func (_m *DbInterface) UpdateFinalityProviderDetailsFromEvent(ctx context.Context, detailsToUpdate *model.FinalityProviderDetails) error {
	ret := _m.Called(ctx, detailsToUpdate)

	if len(ret) == 0 {
		panic("no return value specified for UpdateFinalityProviderDetailsFromEvent")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.FinalityProviderDetails) error); ok {
		r0 = rf(ctx, detailsToUpdate)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateFinalityProviderState provides a mock function with given fields: ctx, btcPk, newState
func (_m *DbInterface) UpdateFinalityProviderState(ctx context.Context, btcPk string, newState string) error {
	ret := _m.Called(ctx, btcPk, newState)

	if len(ret) == 0 {
		panic("no return value specified for UpdateFinalityProviderState")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, btcPk, newState)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewDbInterface creates a new instance of DbInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewDbInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *DbInterface {
	mock := &DbInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
