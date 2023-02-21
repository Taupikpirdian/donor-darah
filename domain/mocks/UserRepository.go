// Code generated by mockery v2.20.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/bxcodec/go-clean-arch/domain"
	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

// Register provides a mock function with given fields: ctx, us
func (_m *UserRepository) Register(ctx context.Context, us *domain.UserData) error {
	ret := _m.Called(ctx, us)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.UserData) error); ok {
		r0 = rf(ctx, us)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// StoreProfile provides a mock function with given fields: ctx, us
func (_m *UserRepository) StoreProfile(ctx context.Context, us *domain.UserData) error {
	ret := _m.Called(ctx, us)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.UserData) error); ok {
		r0 = rf(ctx, us)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewUserRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUserRepository(t mockConstructorTestingTNewUserRepository) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
