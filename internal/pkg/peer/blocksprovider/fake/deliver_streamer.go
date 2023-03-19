// Code generated by counterfeiter. DO NOT EDIT.
package fake

import (
	"context"
	"sync"

	"github.com/hyperledger/fabric-protos-go/orderer"
	"github.com/m4ru1/fabric-gm-bdais/internal/pkg/peer/blocksprovider"
	"google.golang.org/grpc"
)

type DeliverStreamer struct {
	DeliverStub        func(context.Context, *grpc.ClientConn) (orderer.AtomicBroadcast_DeliverClient, error)
	deliverMutex       sync.RWMutex
	deliverArgsForCall []struct {
		arg1 context.Context
		arg2 *grpc.ClientConn
	}
	deliverReturns struct {
		result1 orderer.AtomicBroadcast_DeliverClient
		result2 error
	}
	deliverReturnsOnCall map[int]struct {
		result1 orderer.AtomicBroadcast_DeliverClient
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *DeliverStreamer) Deliver(arg1 context.Context, arg2 *grpc.ClientConn) (orderer.AtomicBroadcast_DeliverClient, error) {
	fake.deliverMutex.Lock()
	ret, specificReturn := fake.deliverReturnsOnCall[len(fake.deliverArgsForCall)]
	fake.deliverArgsForCall = append(fake.deliverArgsForCall, struct {
		arg1 context.Context
		arg2 *grpc.ClientConn
	}{arg1, arg2})
	fake.recordInvocation("Deliver", []interface{}{arg1, arg2})
	fake.deliverMutex.Unlock()
	if fake.DeliverStub != nil {
		return fake.DeliverStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.deliverReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *DeliverStreamer) DeliverCallCount() int {
	fake.deliverMutex.RLock()
	defer fake.deliverMutex.RUnlock()
	return len(fake.deliverArgsForCall)
}

func (fake *DeliverStreamer) DeliverCalls(stub func(context.Context, *grpc.ClientConn) (orderer.AtomicBroadcast_DeliverClient, error)) {
	fake.deliverMutex.Lock()
	defer fake.deliverMutex.Unlock()
	fake.DeliverStub = stub
}

func (fake *DeliverStreamer) DeliverArgsForCall(i int) (context.Context, *grpc.ClientConn) {
	fake.deliverMutex.RLock()
	defer fake.deliverMutex.RUnlock()
	argsForCall := fake.deliverArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *DeliverStreamer) DeliverReturns(result1 orderer.AtomicBroadcast_DeliverClient, result2 error) {
	fake.deliverMutex.Lock()
	defer fake.deliverMutex.Unlock()
	fake.DeliverStub = nil
	fake.deliverReturns = struct {
		result1 orderer.AtomicBroadcast_DeliverClient
		result2 error
	}{result1, result2}
}

func (fake *DeliverStreamer) DeliverReturnsOnCall(i int, result1 orderer.AtomicBroadcast_DeliverClient, result2 error) {
	fake.deliverMutex.Lock()
	defer fake.deliverMutex.Unlock()
	fake.DeliverStub = nil
	if fake.deliverReturnsOnCall == nil {
		fake.deliverReturnsOnCall = make(map[int]struct {
			result1 orderer.AtomicBroadcast_DeliverClient
			result2 error
		})
	}
	fake.deliverReturnsOnCall[i] = struct {
		result1 orderer.AtomicBroadcast_DeliverClient
		result2 error
	}{result1, result2}
}

func (fake *DeliverStreamer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deliverMutex.RLock()
	defer fake.deliverMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *DeliverStreamer) recordInvocation(key string, args []interface{}) {
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

var _ blocksprovider.DeliverStreamer = new(DeliverStreamer)