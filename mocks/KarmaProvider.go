// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	discordgo "github.com/bwmarrin/discordgo"

	mock "github.com/stretchr/testify/mock"
)

// KarmaProvider is an autogenerated mock type for the Provider type
type KarmaProvider struct {
	mock.Mock
}

// ApplyPenalty provides a mock function with given fields: guildID, userID
func (_m *KarmaProvider) ApplyPenalty(guildID string, userID string) error {
	ret := _m.Called(guildID, userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(guildID, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CheckAndUpdate provides a mock function with given fields: guildID, executorID, object, value
func (_m *KarmaProvider) CheckAndUpdate(guildID string, executorID string, object *discordgo.User, value int) (bool, error) {
	ret := _m.Called(guildID, executorID, object, value)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string, *discordgo.User, int) bool); ok {
		r0 = rf(guildID, executorID, object, value)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, *discordgo.User, int) error); ok {
		r1 = rf(guildID, executorID, object, value)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetState provides a mock function with given fields: guildID
func (_m *KarmaProvider) GetState(guildID string) (bool, error) {
	ret := _m.Called(guildID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(guildID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(guildID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsBlockListed provides a mock function with given fields: guildID, userID
func (_m *KarmaProvider) IsBlockListed(guildID string, userID string) (bool, error) {
	ret := _m.Called(guildID, userID)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(guildID, userID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(guildID, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: guildID, userID, executorID, value
func (_m *KarmaProvider) Update(guildID string, userID string, executorID string, value int) error {
	ret := _m.Called(guildID, userID, executorID, value)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, int) error); ok {
		r0 = rf(guildID, userID, executorID, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewKarmaProvider interface {
	mock.TestingT
	Cleanup(func())
}

// NewKarmaProvider creates a new instance of KarmaProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewKarmaProvider(t mockConstructorTestingTNewKarmaProvider) *KarmaProvider {
	mock := &KarmaProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
