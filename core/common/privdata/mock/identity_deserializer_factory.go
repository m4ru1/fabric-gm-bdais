// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/m4ru1/fabric-gm-bdais/msp"
)

type IdentityDeserializerFactory struct {
	GetIdentityDeserializerStub        func(string) msp.IdentityDeserializer
	getIdentityDeserializerMutex       sync.RWMutex
	getIdentityDeserializerArgsForCall []struct {
		arg1 string
	}
	getIdentityDeserializerReturns struct {
		result1 msp.IdentityDeserializer
	}
	getIdentityDeserializerReturnsOnCall map[int]struct {
		result1 msp.IdentityDeserializer
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *IdentityDeserializerFactory) GetIdentityDeserializer(arg1 string) msp.IdentityDeserializer {
	fake.getIdentityDeserializerMutex.Lock()
	ret, specificReturn := fake.getIdentityDeserializerReturnsOnCall[len(fake.getIdentityDeserializerArgsForCall)]
	fake.getIdentityDeserializerArgsForCall = append(fake.getIdentityDeserializerArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetIdentityDeserializer", []interface{}{arg1})
	fake.getIdentityDeserializerMutex.Unlock()
	if fake.GetIdentityDeserializerStub != nil {
		return fake.GetIdentityDeserializerStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.getIdentityDeserializerReturns
	return fakeReturns.result1
}

func (fake *IdentityDeserializerFactory) GetIdentityDeserializerCallCount() int {
	fake.getIdentityDeserializerMutex.RLock()
	defer fake.getIdentityDeserializerMutex.RUnlock()
	return len(fake.getIdentityDeserializerArgsForCall)
}

func (fake *IdentityDeserializerFactory) GetIdentityDeserializerCalls(stub func(string) msp.IdentityDeserializer) {
	fake.getIdentityDeserializerMutex.Lock()
	defer fake.getIdentityDeserializerMutex.Unlock()
	fake.GetIdentityDeserializerStub = stub
}

func (fake *IdentityDeserializerFactory) GetIdentityDeserializerArgsForCall(i int) string {
	fake.getIdentityDeserializerMutex.RLock()
	defer fake.getIdentityDeserializerMutex.RUnlock()
	argsForCall := fake.getIdentityDeserializerArgsForCall[i]
	return argsForCall.arg1
}

func (fake *IdentityDeserializerFactory) GetIdentityDeserializerReturns(result1 msp.IdentityDeserializer) {
	fake.getIdentityDeserializerMutex.Lock()
	defer fake.getIdentityDeserializerMutex.Unlock()
	fake.GetIdentityDeserializerStub = nil
	fake.getIdentityDeserializerReturns = struct {
		result1 msp.IdentityDeserializer
	}{result1}
}

func (fake *IdentityDeserializerFactory) GetIdentityDeserializerReturnsOnCall(i int, result1 msp.IdentityDeserializer) {
	fake.getIdentityDeserializerMutex.Lock()
	defer fake.getIdentityDeserializerMutex.Unlock()
	fake.GetIdentityDeserializerStub = nil
	if fake.getIdentityDeserializerReturnsOnCall == nil {
		fake.getIdentityDeserializerReturnsOnCall = make(map[int]struct {
			result1 msp.IdentityDeserializer
		})
	}
	fake.getIdentityDeserializerReturnsOnCall[i] = struct {
		result1 msp.IdentityDeserializer
	}{result1}
}

func (fake *IdentityDeserializerFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getIdentityDeserializerMutex.RLock()
	defer fake.getIdentityDeserializerMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *IdentityDeserializerFactory) recordInvocation(key string, args []interface{}) {
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
