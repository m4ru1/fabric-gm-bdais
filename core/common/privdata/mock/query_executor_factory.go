// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/m4ru1/fabric-gm-bdais/core/ledger"
)

type QueryExecutorFactory struct {
	NewQueryExecutorStub        func() (ledger.QueryExecutor, error)
	newQueryExecutorMutex       sync.RWMutex
	newQueryExecutorArgsForCall []struct {
	}
	newQueryExecutorReturns struct {
		result1 ledger.QueryExecutor
		result2 error
	}
	newQueryExecutorReturnsOnCall map[int]struct {
		result1 ledger.QueryExecutor
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *QueryExecutorFactory) NewQueryExecutor() (ledger.QueryExecutor, error) {
	fake.newQueryExecutorMutex.Lock()
	ret, specificReturn := fake.newQueryExecutorReturnsOnCall[len(fake.newQueryExecutorArgsForCall)]
	fake.newQueryExecutorArgsForCall = append(fake.newQueryExecutorArgsForCall, struct {
	}{})
	fake.recordInvocation("NewQueryExecutor", []interface{}{})
	fake.newQueryExecutorMutex.Unlock()
	if fake.NewQueryExecutorStub != nil {
		return fake.NewQueryExecutorStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.newQueryExecutorReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *QueryExecutorFactory) NewQueryExecutorCallCount() int {
	fake.newQueryExecutorMutex.RLock()
	defer fake.newQueryExecutorMutex.RUnlock()
	return len(fake.newQueryExecutorArgsForCall)
}

func (fake *QueryExecutorFactory) NewQueryExecutorCalls(stub func() (ledger.QueryExecutor, error)) {
	fake.newQueryExecutorMutex.Lock()
	defer fake.newQueryExecutorMutex.Unlock()
	fake.NewQueryExecutorStub = stub
}

func (fake *QueryExecutorFactory) NewQueryExecutorReturns(result1 ledger.QueryExecutor, result2 error) {
	fake.newQueryExecutorMutex.Lock()
	defer fake.newQueryExecutorMutex.Unlock()
	fake.NewQueryExecutorStub = nil
	fake.newQueryExecutorReturns = struct {
		result1 ledger.QueryExecutor
		result2 error
	}{result1, result2}
}

func (fake *QueryExecutorFactory) NewQueryExecutorReturnsOnCall(i int, result1 ledger.QueryExecutor, result2 error) {
	fake.newQueryExecutorMutex.Lock()
	defer fake.newQueryExecutorMutex.Unlock()
	fake.NewQueryExecutorStub = nil
	if fake.newQueryExecutorReturnsOnCall == nil {
		fake.newQueryExecutorReturnsOnCall = make(map[int]struct {
			result1 ledger.QueryExecutor
			result2 error
		})
	}
	fake.newQueryExecutorReturnsOnCall[i] = struct {
		result1 ledger.QueryExecutor
		result2 error
	}{result1, result2}
}

func (fake *QueryExecutorFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.newQueryExecutorMutex.RLock()
	defer fake.newQueryExecutorMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *QueryExecutorFactory) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}
