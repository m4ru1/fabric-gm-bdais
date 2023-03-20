// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import msp "github.com/m4ru1/fabric-gm-bdais/msp"

import protoutil "github.com/m4ru1/fabric-gm-bdais/protoutil"

// Policy is an autogenerated mock type for the Policy type
type Policy struct {
	mock.Mock
}

// EvaluateIdentities provides a mock function with given fields: identities
func (_m *Policy) EvaluateIdentities(identities []msp.Identity) error {
	ret := _m.Called(identities)

	var r0 error
	if rf, ok := ret.Get(0).(func([]msp.Identity) error); ok {
		r0 = rf(identities)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// EvaluateSignedData provides a mock function with given fields: signatureSet
func (_m *Policy) EvaluateSignedData(signatureSet []*protoutil.SignedData) error {
	ret := _m.Called(signatureSet)

	var r0 error
	if rf, ok := ret.Get(0).(func([]*protoutil.SignedData) error); ok {
		r0 = rf(signatureSet)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
