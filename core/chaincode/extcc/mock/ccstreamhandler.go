// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/m4ru1/fabric-gm-bdais/core/chaincode/extcc"
	"github.com/m4ru1/fabric-gm-bdais/core/container/ccintf"
)

type StreamHandler struct {
	HandleChaincodeStreamStub        func(ccintf.ChaincodeStream) error
	handleChaincodeStreamMutex       sync.RWMutex
	handleChaincodeStreamArgsForCall []struct {
		arg1 ccintf.ChaincodeStream
	}
	handleChaincodeStreamReturns struct {
		result1 error
	}
	handleChaincodeStreamReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *StreamHandler) HandleChaincodeStream(arg1 ccintf.ChaincodeStream) error {
	fake.handleChaincodeStreamMutex.Lock()
	ret, specificReturn := fake.handleChaincodeStreamReturnsOnCall[len(fake.handleChaincodeStreamArgsForCall)]
	fake.handleChaincodeStreamArgsForCall = append(fake.handleChaincodeStreamArgsForCall, struct {
		arg1 ccintf.ChaincodeStream
	}{arg1})
	fake.recordInvocation("HandleChaincodeStream", []interface{}{arg1})
	fake.handleChaincodeStreamMutex.Unlock()
	if fake.HandleChaincodeStreamStub != nil {
		return fake.HandleChaincodeStreamStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.handleChaincodeStreamReturns
	return fakeReturns.result1
}

func (fake *StreamHandler) HandleChaincodeStreamCallCount() int {
	fake.handleChaincodeStreamMutex.RLock()
	defer fake.handleChaincodeStreamMutex.RUnlock()
	return len(fake.handleChaincodeStreamArgsForCall)
}

func (fake *StreamHandler) HandleChaincodeStreamCalls(stub func(ccintf.ChaincodeStream) error) {
	fake.handleChaincodeStreamMutex.Lock()
	defer fake.handleChaincodeStreamMutex.Unlock()
	fake.HandleChaincodeStreamStub = stub
}

func (fake *StreamHandler) HandleChaincodeStreamArgsForCall(i int) ccintf.ChaincodeStream {
	fake.handleChaincodeStreamMutex.RLock()
	defer fake.handleChaincodeStreamMutex.RUnlock()
	argsForCall := fake.handleChaincodeStreamArgsForCall[i]
	return argsForCall.arg1
}

func (fake *StreamHandler) HandleChaincodeStreamReturns(result1 error) {
	fake.handleChaincodeStreamMutex.Lock()
	defer fake.handleChaincodeStreamMutex.Unlock()
	fake.HandleChaincodeStreamStub = nil
	fake.handleChaincodeStreamReturns = struct {
		result1 error
	}{result1}
}

func (fake *StreamHandler) HandleChaincodeStreamReturnsOnCall(i int, result1 error) {
	fake.handleChaincodeStreamMutex.Lock()
	defer fake.handleChaincodeStreamMutex.Unlock()
	fake.HandleChaincodeStreamStub = nil
	if fake.handleChaincodeStreamReturnsOnCall == nil {
		fake.handleChaincodeStreamReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.handleChaincodeStreamReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *StreamHandler) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.handleChaincodeStreamMutex.RLock()
	defer fake.handleChaincodeStreamMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *StreamHandler) recordInvocation(key string, args []interface{}) {
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

var _ extcc.StreamHandler = new(StreamHandler)
