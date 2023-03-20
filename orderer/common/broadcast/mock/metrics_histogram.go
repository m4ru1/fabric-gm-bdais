// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/m4ru1/fabric-gm-bdais/common/metrics"
)

type MetricsHistogram struct {
	ObserveStub        func(float64)
	observeMutex       sync.RWMutex
	observeArgsForCall []struct {
		arg1 float64
	}
	WithStub        func(...string) metrics.Histogram
	withMutex       sync.RWMutex
	withArgsForCall []struct {
		arg1 []string
	}
	withReturns struct {
		result1 metrics.Histogram
	}
	withReturnsOnCall map[int]struct {
		result1 metrics.Histogram
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *MetricsHistogram) Observe(arg1 float64) {
	fake.observeMutex.Lock()
	fake.observeArgsForCall = append(fake.observeArgsForCall, struct {
		arg1 float64
	}{arg1})
	fake.recordInvocation("Observe", []interface{}{arg1})
	fake.observeMutex.Unlock()
	if fake.ObserveStub != nil {
		fake.ObserveStub(arg1)
	}
}

func (fake *MetricsHistogram) ObserveCallCount() int {
	fake.observeMutex.RLock()
	defer fake.observeMutex.RUnlock()
	return len(fake.observeArgsForCall)
}

func (fake *MetricsHistogram) ObserveCalls(stub func(float64)) {
	fake.observeMutex.Lock()
	defer fake.observeMutex.Unlock()
	fake.ObserveStub = stub
}

func (fake *MetricsHistogram) ObserveArgsForCall(i int) float64 {
	fake.observeMutex.RLock()
	defer fake.observeMutex.RUnlock()
	argsForCall := fake.observeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *MetricsHistogram) With(arg1 ...string) metrics.Histogram {
	fake.withMutex.Lock()
	ret, specificReturn := fake.withReturnsOnCall[len(fake.withArgsForCall)]
	fake.withArgsForCall = append(fake.withArgsForCall, struct {
		arg1 []string
	}{arg1})
	fake.recordInvocation("With", []interface{}{arg1})
	fake.withMutex.Unlock()
	if fake.WithStub != nil {
		return fake.WithStub(arg1...)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.withReturns
	return fakeReturns.result1
}

func (fake *MetricsHistogram) WithCallCount() int {
	fake.withMutex.RLock()
	defer fake.withMutex.RUnlock()
	return len(fake.withArgsForCall)
}

func (fake *MetricsHistogram) WithCalls(stub func(...string) metrics.Histogram) {
	fake.withMutex.Lock()
	defer fake.withMutex.Unlock()
	fake.WithStub = stub
}

func (fake *MetricsHistogram) WithArgsForCall(i int) []string {
	fake.withMutex.RLock()
	defer fake.withMutex.RUnlock()
	argsForCall := fake.withArgsForCall[i]
	return argsForCall.arg1
}

func (fake *MetricsHistogram) WithReturns(result1 metrics.Histogram) {
	fake.withMutex.Lock()
	defer fake.withMutex.Unlock()
	fake.WithStub = nil
	fake.withReturns = struct {
		result1 metrics.Histogram
	}{result1}
}

func (fake *MetricsHistogram) WithReturnsOnCall(i int, result1 metrics.Histogram) {
	fake.withMutex.Lock()
	defer fake.withMutex.Unlock()
	fake.WithStub = nil
	if fake.withReturnsOnCall == nil {
		fake.withReturnsOnCall = make(map[int]struct {
			result1 metrics.Histogram
		})
	}
	fake.withReturnsOnCall[i] = struct {
		result1 metrics.Histogram
	}{result1}
}

func (fake *MetricsHistogram) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.observeMutex.RLock()
	defer fake.observeMutex.RUnlock()
	fake.withMutex.RLock()
	defer fake.withMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *MetricsHistogram) recordInvocation(key string, args []interface{}) {
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
