// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	ledger "github.com/m4ru1/fabric-gm-bdais/core/ledger"
	mock "github.com/stretchr/testify/mock"
)

// ConfigHistoryRetriever is an autogenerated mock type for the ConfigHistoryRetriever type
type ConfigHistoryRetriever struct {
	mock.Mock
}

// MostRecentCollectionConfigBelow provides a mock function with given fields: blockNum, chaincodeName
func (_m *ConfigHistoryRetriever) MostRecentCollectionConfigBelow(blockNum uint64, chaincodeName string) (*ledger.CollectionConfigInfo, error) {
	ret := _m.Called(blockNum, chaincodeName)

	var r0 *ledger.CollectionConfigInfo
	if rf, ok := ret.Get(0).(func(uint64, string) *ledger.CollectionConfigInfo); ok {
		r0 = rf(blockNum, chaincodeName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*ledger.CollectionConfigInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(uint64, string) error); ok {
		r1 = rf(blockNum, chaincodeName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
