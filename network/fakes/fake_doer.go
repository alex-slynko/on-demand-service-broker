// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"net/http"
	"sync"

	"github.com/pivotal-cf/on-demand-service-broker/network"
)

type FakeDoer struct {
	DoStub        func(request *http.Request) (*http.Response, error)
	doMutex       sync.RWMutex
	doArgsForCall []struct {
		request *http.Request
	}
	doReturns struct {
		result1 *http.Response
		result2 error
	}
	doReturnsOnCall map[int]struct {
		result1 *http.Response
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDoer) Do(request *http.Request) (*http.Response, error) {
	fake.doMutex.Lock()
	ret, specificReturn := fake.doReturnsOnCall[len(fake.doArgsForCall)]
	fake.doArgsForCall = append(fake.doArgsForCall, struct {
		request *http.Request
	}{request})
	fake.recordInvocation("Do", []interface{}{request})
	fake.doMutex.Unlock()
	if fake.DoStub != nil {
		return fake.DoStub(request)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.doReturns.result1, fake.doReturns.result2
}

func (fake *FakeDoer) DoCallCount() int {
	fake.doMutex.RLock()
	defer fake.doMutex.RUnlock()
	return len(fake.doArgsForCall)
}

func (fake *FakeDoer) DoArgsForCall(i int) *http.Request {
	fake.doMutex.RLock()
	defer fake.doMutex.RUnlock()
	return fake.doArgsForCall[i].request
}

func (fake *FakeDoer) DoReturns(result1 *http.Response, result2 error) {
	fake.DoStub = nil
	fake.doReturns = struct {
		result1 *http.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeDoer) DoReturnsOnCall(i int, result1 *http.Response, result2 error) {
	fake.DoStub = nil
	if fake.doReturnsOnCall == nil {
		fake.doReturnsOnCall = make(map[int]struct {
			result1 *http.Response
			result2 error
		})
	}
	fake.doReturnsOnCall[i] = struct {
		result1 *http.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeDoer) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.doMutex.RLock()
	defer fake.doMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDoer) recordInvocation(key string, args []interface{}) {
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

var _ network.Doer = new(FakeDoer)
