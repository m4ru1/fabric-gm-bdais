// Code generated by counterfeiter. DO NOT EDIT.
package fake

import (
	"sync"

	"github.com/m4ru1/fabric-gm-bdais/core/chaincode"
)

type LaunchRegistry struct {
	DeregisterStub        func(string) error
	deregisterMutex       sync.RWMutex
	deregisterArgsForCall []struct {
		arg1 string
	}
	deregisterReturns struct {
		result1 error
	}
	deregisterReturnsOnCall map[int]struct {
		result1 error
	}
	LaunchingStub        func(string) (*chaincode.LaunchState, bool)
	launchingMutex       sync.RWMutex
	launchingArgsForCall []struct {
		arg1 string
	}
	launchingReturns struct {
		result1 *chaincode.LaunchState
		result2 bool
	}
	launchingReturnsOnCall map[int]struct {
		result1 *chaincode.LaunchState
		result2 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *LaunchRegistry) Deregister(arg1 string) error {
	fake.deregisterMutex.Lock()
	ret, specificReturn := fake.deregisterReturnsOnCall[len(fake.deregisterArgsForCall)]
	fake.deregisterArgsForCall = append(fake.deregisterArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Deregister", []interface{}{arg1})
	fake.deregisterMutex.Unlock()
	if fake.DeregisterStub != nil {
		return fake.DeregisterStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deregisterReturns
	return fakeReturns.result1
}

func (fake *LaunchRegistry) DeregisterCallCount() int {
	fake.deregisterMutex.RLock()
	defer fake.deregisterMutex.RUnlock()
	return len(fake.deregisterArgsForCall)
}

func (fake *LaunchRegistry) DeregisterCalls(stub func(string) error) {
	fake.deregisterMutex.Lock()
	defer fake.deregisterMutex.Unlock()
	fake.DeregisterStub = stub
}

func (fake *LaunchRegistry) DeregisterArgsForCall(i int) string {
	fake.deregisterMutex.RLock()
	defer fake.deregisterMutex.RUnlock()
	argsForCall := fake.deregisterArgsForCall[i]
	return argsForCall.arg1
}

func (fake *LaunchRegistry) DeregisterReturns(result1 error) {
	fake.deregisterMutex.Lock()
	defer fake.deregisterMutex.Unlock()
	fake.DeregisterStub = nil
	fake.deregisterReturns = struct {
		result1 error
	}{result1}
}

func (fake *LaunchRegistry) DeregisterReturnsOnCall(i int, result1 error) {
	fake.deregisterMutex.Lock()
	defer fake.deregisterMutex.Unlock()
	fake.DeregisterStub = nil
	if fake.deregisterReturnsOnCall == nil {
		fake.deregisterReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deregisterReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *LaunchRegistry) Launching(arg1 string) (*chaincode.LaunchState, bool) {
	fake.launchingMutex.Lock()
	ret, specificReturn := fake.launchingReturnsOnCall[len(fake.launchingArgsForCall)]
	fake.launchingArgsForCall = append(fake.launchingArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Launching", []interface{}{arg1})
	fake.launchingMutex.Unlock()
	if fake.LaunchingStub != nil {
		return fake.LaunchingStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.launchingReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *LaunchRegistry) LaunchingCallCount() int {
	fake.launchingMutex.RLock()
	defer fake.launchingMutex.RUnlock()
	return len(fake.launchingArgsForCall)
}

func (fake *LaunchRegistry) LaunchingCalls(stub func(string) (*chaincode.LaunchState, bool)) {
	fake.launchingMutex.Lock()
	defer fake.launchingMutex.Unlock()
	fake.LaunchingStub = stub
}

func (fake *LaunchRegistry) LaunchingArgsForCall(i int) string {
	fake.launchingMutex.RLock()
	defer fake.launchingMutex.RUnlock()
	argsForCall := fake.launchingArgsForCall[i]
	return argsForCall.arg1
}

func (fake *LaunchRegistry) LaunchingReturns(result1 *chaincode.LaunchState, result2 bool) {
	fake.launchingMutex.Lock()
	defer fake.launchingMutex.Unlock()
	fake.LaunchingStub = nil
	fake.launchingReturns = struct {
		result1 *chaincode.LaunchState
		result2 bool
	}{result1, result2}
}

func (fake *LaunchRegistry) LaunchingReturnsOnCall(i int, result1 *chaincode.LaunchState, result2 bool) {
	fake.launchingMutex.Lock()
	defer fake.launchingMutex.Unlock()
	fake.LaunchingStub = nil
	if fake.launchingReturnsOnCall == nil {
		fake.launchingReturnsOnCall = make(map[int]struct {
			result1 *chaincode.LaunchState
			result2 bool
		})
	}
	fake.launchingReturnsOnCall[i] = struct {
		result1 *chaincode.LaunchState
		result2 bool
	}{result1, result2}
}

func (fake *LaunchRegistry) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deregisterMutex.RLock()
	defer fake.deregisterMutex.RUnlock()
	fake.launchingMutex.RLock()
	defer fake.launchingMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *LaunchRegistry) recordInvocation(key string, args []interface{}) {
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