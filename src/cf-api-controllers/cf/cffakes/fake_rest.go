// Code generated by counterfeiter. DO NOT EDIT.
package cffakes

import (
	"io"
	"net/http"
	"sync"

	"code.cloudfoundry.org/capi-k8s-release/src/cf-api-controllers/cf"
)

type FakeRest struct {
	PatchStub        func(string, string, io.Reader) (*http.Response, error)
	patchMutex       sync.RWMutex
	patchArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 io.Reader
	}
	patchReturns struct {
		result1 *http.Response
		result2 error
	}
	patchReturnsOnCall map[int]struct {
		result1 *http.Response
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRest) Patch(arg1 string, arg2 string, arg3 io.Reader) (*http.Response, error) {
	fake.patchMutex.Lock()
	ret, specificReturn := fake.patchReturnsOnCall[len(fake.patchArgsForCall)]
	fake.patchArgsForCall = append(fake.patchArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 io.Reader
	}{arg1, arg2, arg3})
	fake.recordInvocation("Patch", []interface{}{arg1, arg2, arg3})
	fake.patchMutex.Unlock()
	if fake.PatchStub != nil {
		return fake.PatchStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.patchReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRest) PatchCallCount() int {
	fake.patchMutex.RLock()
	defer fake.patchMutex.RUnlock()
	return len(fake.patchArgsForCall)
}

func (fake *FakeRest) PatchCalls(stub func(string, string, io.Reader) (*http.Response, error)) {
	fake.patchMutex.Lock()
	defer fake.patchMutex.Unlock()
	fake.PatchStub = stub
}

func (fake *FakeRest) PatchArgsForCall(i int) (string, string, io.Reader) {
	fake.patchMutex.RLock()
	defer fake.patchMutex.RUnlock()
	argsForCall := fake.patchArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeRest) PatchReturns(result1 *http.Response, result2 error) {
	fake.patchMutex.Lock()
	defer fake.patchMutex.Unlock()
	fake.PatchStub = nil
	fake.patchReturns = struct {
		result1 *http.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeRest) PatchReturnsOnCall(i int, result1 *http.Response, result2 error) {
	fake.patchMutex.Lock()
	defer fake.patchMutex.Unlock()
	fake.PatchStub = nil
	if fake.patchReturnsOnCall == nil {
		fake.patchReturnsOnCall = make(map[int]struct {
			result1 *http.Response
			result2 error
		})
	}
	fake.patchReturnsOnCall[i] = struct {
		result1 *http.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeRest) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.patchMutex.RLock()
	defer fake.patchMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRest) recordInvocation(key string, args []interface{}) {
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

var _ cf.Rest = new(FakeRest)
