// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/hyperledger/fabric-protos-go/common"
	"github.com/m4ru1/fabric-gm-bdais/common/ledger"
)

type FileLedgerBlockStore struct {
	AddBlockStub        func(*common.Block) error
	addBlockMutex       sync.RWMutex
	addBlockArgsForCall []struct {
		arg1 *common.Block
	}
	addBlockReturns struct {
		result1 error
	}
	addBlockReturnsOnCall map[int]struct {
		result1 error
	}
	GetBlockchainInfoStub        func() (*common.BlockchainInfo, error)
	getBlockchainInfoMutex       sync.RWMutex
	getBlockchainInfoArgsForCall []struct {
	}
	getBlockchainInfoReturns struct {
		result1 *common.BlockchainInfo
		result2 error
	}
	getBlockchainInfoReturnsOnCall map[int]struct {
		result1 *common.BlockchainInfo
		result2 error
	}
	RetrieveBlockByNumberStub        func(uint64) (*common.Block, error)
	retrieveBlockByNumberMutex       sync.RWMutex
	retrieveBlockByNumberArgsForCall []struct {
		arg1 uint64
	}
	retrieveBlockByNumberReturns struct {
		result1 *common.Block
		result2 error
	}
	retrieveBlockByNumberReturnsOnCall map[int]struct {
		result1 *common.Block
		result2 error
	}
	RetrieveBlocksStub        func(uint64) (ledger.ResultsIterator, error)
	retrieveBlocksMutex       sync.RWMutex
	retrieveBlocksArgsForCall []struct {
		arg1 uint64
	}
	retrieveBlocksReturns struct {
		result1 ledger.ResultsIterator
		result2 error
	}
	retrieveBlocksReturnsOnCall map[int]struct {
		result1 ledger.ResultsIterator
		result2 error
	}
	ShutdownStub        func()
	shutdownMutex       sync.RWMutex
	shutdownArgsForCall []struct {
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FileLedgerBlockStore) AddBlock(arg1 *common.Block) error {
	fake.addBlockMutex.Lock()
	ret, specificReturn := fake.addBlockReturnsOnCall[len(fake.addBlockArgsForCall)]
	fake.addBlockArgsForCall = append(fake.addBlockArgsForCall, struct {
		arg1 *common.Block
	}{arg1})
	stub := fake.AddBlockStub
	fakeReturns := fake.addBlockReturns
	fake.recordInvocation("AddBlock", []interface{}{arg1})
	fake.addBlockMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FileLedgerBlockStore) AddBlockCallCount() int {
	fake.addBlockMutex.RLock()
	defer fake.addBlockMutex.RUnlock()
	return len(fake.addBlockArgsForCall)
}

func (fake *FileLedgerBlockStore) AddBlockCalls(stub func(*common.Block) error) {
	fake.addBlockMutex.Lock()
	defer fake.addBlockMutex.Unlock()
	fake.AddBlockStub = stub
}

func (fake *FileLedgerBlockStore) AddBlockArgsForCall(i int) *common.Block {
	fake.addBlockMutex.RLock()
	defer fake.addBlockMutex.RUnlock()
	argsForCall := fake.addBlockArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FileLedgerBlockStore) AddBlockReturns(result1 error) {
	fake.addBlockMutex.Lock()
	defer fake.addBlockMutex.Unlock()
	fake.AddBlockStub = nil
	fake.addBlockReturns = struct {
		result1 error
	}{result1}
}

func (fake *FileLedgerBlockStore) AddBlockReturnsOnCall(i int, result1 error) {
	fake.addBlockMutex.Lock()
	defer fake.addBlockMutex.Unlock()
	fake.AddBlockStub = nil
	if fake.addBlockReturnsOnCall == nil {
		fake.addBlockReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addBlockReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FileLedgerBlockStore) GetBlockchainInfo() (*common.BlockchainInfo, error) {
	fake.getBlockchainInfoMutex.Lock()
	ret, specificReturn := fake.getBlockchainInfoReturnsOnCall[len(fake.getBlockchainInfoArgsForCall)]
	fake.getBlockchainInfoArgsForCall = append(fake.getBlockchainInfoArgsForCall, struct {
	}{})
	stub := fake.GetBlockchainInfoStub
	fakeReturns := fake.getBlockchainInfoReturns
	fake.recordInvocation("GetBlockchainInfo", []interface{}{})
	fake.getBlockchainInfoMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FileLedgerBlockStore) GetBlockchainInfoCallCount() int {
	fake.getBlockchainInfoMutex.RLock()
	defer fake.getBlockchainInfoMutex.RUnlock()
	return len(fake.getBlockchainInfoArgsForCall)
}

func (fake *FileLedgerBlockStore) GetBlockchainInfoCalls(stub func() (*common.BlockchainInfo, error)) {
	fake.getBlockchainInfoMutex.Lock()
	defer fake.getBlockchainInfoMutex.Unlock()
	fake.GetBlockchainInfoStub = stub
}

func (fake *FileLedgerBlockStore) GetBlockchainInfoReturns(result1 *common.BlockchainInfo, result2 error) {
	fake.getBlockchainInfoMutex.Lock()
	defer fake.getBlockchainInfoMutex.Unlock()
	fake.GetBlockchainInfoStub = nil
	fake.getBlockchainInfoReturns = struct {
		result1 *common.BlockchainInfo
		result2 error
	}{result1, result2}
}

func (fake *FileLedgerBlockStore) GetBlockchainInfoReturnsOnCall(i int, result1 *common.BlockchainInfo, result2 error) {
	fake.getBlockchainInfoMutex.Lock()
	defer fake.getBlockchainInfoMutex.Unlock()
	fake.GetBlockchainInfoStub = nil
	if fake.getBlockchainInfoReturnsOnCall == nil {
		fake.getBlockchainInfoReturnsOnCall = make(map[int]struct {
			result1 *common.BlockchainInfo
			result2 error
		})
	}
	fake.getBlockchainInfoReturnsOnCall[i] = struct {
		result1 *common.BlockchainInfo
		result2 error
	}{result1, result2}
}

func (fake *FileLedgerBlockStore) RetrieveBlockByNumber(arg1 uint64) (*common.Block, error) {
	fake.retrieveBlockByNumberMutex.Lock()
	ret, specificReturn := fake.retrieveBlockByNumberReturnsOnCall[len(fake.retrieveBlockByNumberArgsForCall)]
	fake.retrieveBlockByNumberArgsForCall = append(fake.retrieveBlockByNumberArgsForCall, struct {
		arg1 uint64
	}{arg1})
	stub := fake.RetrieveBlockByNumberStub
	fakeReturns := fake.retrieveBlockByNumberReturns
	fake.recordInvocation("RetrieveBlockByNumber", []interface{}{arg1})
	fake.retrieveBlockByNumberMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FileLedgerBlockStore) RetrieveBlockByNumberCallCount() int {
	fake.retrieveBlockByNumberMutex.RLock()
	defer fake.retrieveBlockByNumberMutex.RUnlock()
	return len(fake.retrieveBlockByNumberArgsForCall)
}

func (fake *FileLedgerBlockStore) RetrieveBlockByNumberCalls(stub func(uint64) (*common.Block, error)) {
	fake.retrieveBlockByNumberMutex.Lock()
	defer fake.retrieveBlockByNumberMutex.Unlock()
	fake.RetrieveBlockByNumberStub = stub
}

func (fake *FileLedgerBlockStore) RetrieveBlockByNumberArgsForCall(i int) uint64 {
	fake.retrieveBlockByNumberMutex.RLock()
	defer fake.retrieveBlockByNumberMutex.RUnlock()
	argsForCall := fake.retrieveBlockByNumberArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FileLedgerBlockStore) RetrieveBlockByNumberReturns(result1 *common.Block, result2 error) {
	fake.retrieveBlockByNumberMutex.Lock()
	defer fake.retrieveBlockByNumberMutex.Unlock()
	fake.RetrieveBlockByNumberStub = nil
	fake.retrieveBlockByNumberReturns = struct {
		result1 *common.Block
		result2 error
	}{result1, result2}
}

func (fake *FileLedgerBlockStore) RetrieveBlockByNumberReturnsOnCall(i int, result1 *common.Block, result2 error) {
	fake.retrieveBlockByNumberMutex.Lock()
	defer fake.retrieveBlockByNumberMutex.Unlock()
	fake.RetrieveBlockByNumberStub = nil
	if fake.retrieveBlockByNumberReturnsOnCall == nil {
		fake.retrieveBlockByNumberReturnsOnCall = make(map[int]struct {
			result1 *common.Block
			result2 error
		})
	}
	fake.retrieveBlockByNumberReturnsOnCall[i] = struct {
		result1 *common.Block
		result2 error
	}{result1, result2}
}

func (fake *FileLedgerBlockStore) RetrieveBlocks(arg1 uint64) (ledger.ResultsIterator, error) {
	fake.retrieveBlocksMutex.Lock()
	ret, specificReturn := fake.retrieveBlocksReturnsOnCall[len(fake.retrieveBlocksArgsForCall)]
	fake.retrieveBlocksArgsForCall = append(fake.retrieveBlocksArgsForCall, struct {
		arg1 uint64
	}{arg1})
	stub := fake.RetrieveBlocksStub
	fakeReturns := fake.retrieveBlocksReturns
	fake.recordInvocation("RetrieveBlocks", []interface{}{arg1})
	fake.retrieveBlocksMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FileLedgerBlockStore) RetrieveBlocksCallCount() int {
	fake.retrieveBlocksMutex.RLock()
	defer fake.retrieveBlocksMutex.RUnlock()
	return len(fake.retrieveBlocksArgsForCall)
}

func (fake *FileLedgerBlockStore) RetrieveBlocksCalls(stub func(uint64) (ledger.ResultsIterator, error)) {
	fake.retrieveBlocksMutex.Lock()
	defer fake.retrieveBlocksMutex.Unlock()
	fake.RetrieveBlocksStub = stub
}

func (fake *FileLedgerBlockStore) RetrieveBlocksArgsForCall(i int) uint64 {
	fake.retrieveBlocksMutex.RLock()
	defer fake.retrieveBlocksMutex.RUnlock()
	argsForCall := fake.retrieveBlocksArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FileLedgerBlockStore) RetrieveBlocksReturns(result1 ledger.ResultsIterator, result2 error) {
	fake.retrieveBlocksMutex.Lock()
	defer fake.retrieveBlocksMutex.Unlock()
	fake.RetrieveBlocksStub = nil
	fake.retrieveBlocksReturns = struct {
		result1 ledger.ResultsIterator
		result2 error
	}{result1, result2}
}

func (fake *FileLedgerBlockStore) RetrieveBlocksReturnsOnCall(i int, result1 ledger.ResultsIterator, result2 error) {
	fake.retrieveBlocksMutex.Lock()
	defer fake.retrieveBlocksMutex.Unlock()
	fake.RetrieveBlocksStub = nil
	if fake.retrieveBlocksReturnsOnCall == nil {
		fake.retrieveBlocksReturnsOnCall = make(map[int]struct {
			result1 ledger.ResultsIterator
			result2 error
		})
	}
	fake.retrieveBlocksReturnsOnCall[i] = struct {
		result1 ledger.ResultsIterator
		result2 error
	}{result1, result2}
}

func (fake *FileLedgerBlockStore) Shutdown() {
	fake.shutdownMutex.Lock()
	fake.shutdownArgsForCall = append(fake.shutdownArgsForCall, struct {
	}{})
	stub := fake.ShutdownStub
	fake.recordInvocation("Shutdown", []interface{}{})
	fake.shutdownMutex.Unlock()
	if stub != nil {
		fake.ShutdownStub()
	}
}

func (fake *FileLedgerBlockStore) ShutdownCallCount() int {
	fake.shutdownMutex.RLock()
	defer fake.shutdownMutex.RUnlock()
	return len(fake.shutdownArgsForCall)
}

func (fake *FileLedgerBlockStore) ShutdownCalls(stub func()) {
	fake.shutdownMutex.Lock()
	defer fake.shutdownMutex.Unlock()
	fake.ShutdownStub = stub
}

func (fake *FileLedgerBlockStore) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addBlockMutex.RLock()
	defer fake.addBlockMutex.RUnlock()
	fake.getBlockchainInfoMutex.RLock()
	defer fake.getBlockchainInfoMutex.RUnlock()
	fake.retrieveBlockByNumberMutex.RLock()
	defer fake.retrieveBlockByNumberMutex.RUnlock()
	fake.retrieveBlocksMutex.RLock()
	defer fake.retrieveBlocksMutex.RUnlock()
	fake.shutdownMutex.RLock()
	defer fake.shutdownMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FileLedgerBlockStore) recordInvocation(key string, args []interface{}) {
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
