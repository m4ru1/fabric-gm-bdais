// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/m4ru1/fabric-gm-bdais/core/ledger/internal/version"
)

type SnapshotPvtdataHashesConsumer struct {
	ConsumeSnapshotDataStub        func(string, string, []byte, []byte, *version.Height) error
	consumeSnapshotDataMutex       sync.RWMutex
	consumeSnapshotDataArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 []byte
		arg4 []byte
		arg5 *version.Height
	}
	consumeSnapshotDataReturns struct {
		result1 error
	}
	consumeSnapshotDataReturnsOnCall map[int]struct {
		result1 error
	}
	DoneStub        func() error
	doneMutex       sync.RWMutex
	doneArgsForCall []struct {
	}
	doneReturns struct {
		result1 error
	}
	doneReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *SnapshotPvtdataHashesConsumer) ConsumeSnapshotData(arg1 string, arg2 string, arg3 []byte, arg4 []byte, arg5 *version.Height) error {
	var arg3Copy []byte
	if arg3 != nil {
		arg3Copy = make([]byte, len(arg3))
		copy(arg3Copy, arg3)
	}
	var arg4Copy []byte
	if arg4 != nil {
		arg4Copy = make([]byte, len(arg4))
		copy(arg4Copy, arg4)
	}
	fake.consumeSnapshotDataMutex.Lock()
	ret, specificReturn := fake.consumeSnapshotDataReturnsOnCall[len(fake.consumeSnapshotDataArgsForCall)]
	fake.consumeSnapshotDataArgsForCall = append(fake.consumeSnapshotDataArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 []byte
		arg4 []byte
		arg5 *version.Height
	}{arg1, arg2, arg3Copy, arg4Copy, arg5})
	fake.recordInvocation("ConsumeSnapshotData", []interface{}{arg1, arg2, arg3Copy, arg4Copy, arg5})
	fake.consumeSnapshotDataMutex.Unlock()
	if fake.ConsumeSnapshotDataStub != nil {
		return fake.ConsumeSnapshotDataStub(arg1, arg2, arg3, arg4, arg5)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.consumeSnapshotDataReturns
	return fakeReturns.result1
}

func (fake *SnapshotPvtdataHashesConsumer) ConsumeSnapshotDataCallCount() int {
	fake.consumeSnapshotDataMutex.RLock()
	defer fake.consumeSnapshotDataMutex.RUnlock()
	return len(fake.consumeSnapshotDataArgsForCall)
}

func (fake *SnapshotPvtdataHashesConsumer) ConsumeSnapshotDataCalls(stub func(string, string, []byte, []byte, *version.Height) error) {
	fake.consumeSnapshotDataMutex.Lock()
	defer fake.consumeSnapshotDataMutex.Unlock()
	fake.ConsumeSnapshotDataStub = stub
}

func (fake *SnapshotPvtdataHashesConsumer) ConsumeSnapshotDataArgsForCall(i int) (string, string, []byte, []byte, *version.Height) {
	fake.consumeSnapshotDataMutex.RLock()
	defer fake.consumeSnapshotDataMutex.RUnlock()
	argsForCall := fake.consumeSnapshotDataArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4, argsForCall.arg5
}

func (fake *SnapshotPvtdataHashesConsumer) ConsumeSnapshotDataReturns(result1 error) {
	fake.consumeSnapshotDataMutex.Lock()
	defer fake.consumeSnapshotDataMutex.Unlock()
	fake.ConsumeSnapshotDataStub = nil
	fake.consumeSnapshotDataReturns = struct {
		result1 error
	}{result1}
}

func (fake *SnapshotPvtdataHashesConsumer) ConsumeSnapshotDataReturnsOnCall(i int, result1 error) {
	fake.consumeSnapshotDataMutex.Lock()
	defer fake.consumeSnapshotDataMutex.Unlock()
	fake.ConsumeSnapshotDataStub = nil
	if fake.consumeSnapshotDataReturnsOnCall == nil {
		fake.consumeSnapshotDataReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.consumeSnapshotDataReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *SnapshotPvtdataHashesConsumer) Done() error {
	fake.doneMutex.Lock()
	ret, specificReturn := fake.doneReturnsOnCall[len(fake.doneArgsForCall)]
	fake.doneArgsForCall = append(fake.doneArgsForCall, struct {
	}{})
	fake.recordInvocation("Done", []interface{}{})
	fake.doneMutex.Unlock()
	if fake.DoneStub != nil {
		return fake.DoneStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.doneReturns
	return fakeReturns.result1
}

func (fake *SnapshotPvtdataHashesConsumer) DoneCallCount() int {
	fake.doneMutex.RLock()
	defer fake.doneMutex.RUnlock()
	return len(fake.doneArgsForCall)
}

func (fake *SnapshotPvtdataHashesConsumer) DoneCalls(stub func() error) {
	fake.doneMutex.Lock()
	defer fake.doneMutex.Unlock()
	fake.DoneStub = stub
}

func (fake *SnapshotPvtdataHashesConsumer) DoneReturns(result1 error) {
	fake.doneMutex.Lock()
	defer fake.doneMutex.Unlock()
	fake.DoneStub = nil
	fake.doneReturns = struct {
		result1 error
	}{result1}
}

func (fake *SnapshotPvtdataHashesConsumer) DoneReturnsOnCall(i int, result1 error) {
	fake.doneMutex.Lock()
	defer fake.doneMutex.Unlock()
	fake.DoneStub = nil
	if fake.doneReturnsOnCall == nil {
		fake.doneReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.doneReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *SnapshotPvtdataHashesConsumer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.consumeSnapshotDataMutex.RLock()
	defer fake.consumeSnapshotDataMutex.RUnlock()
	fake.doneMutex.RLock()
	defer fake.doneMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *SnapshotPvtdataHashesConsumer) recordInvocation(key string, args []interface{}) {
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
