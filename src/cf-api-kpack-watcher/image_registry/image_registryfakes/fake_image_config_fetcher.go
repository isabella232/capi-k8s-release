// Code generated by counterfeiter. DO NOT EDIT.
package image_registryfakes

import (
	"sync"

	"code.cloudfoundry.org/capi-k8s-release/src/cf-api-kpack-watcher/image_registry"
	v1 "github.com/google/go-containerregistry/pkg/v1"
)

type FakeImageConfigFetcher struct {
	FetchImageConfigStub        func(string, string, string) (*v1.Config, error)
	fetchImageConfigMutex       sync.RWMutex
	fetchImageConfigArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	fetchImageConfigReturns struct {
		result1 *v1.Config
		result2 error
	}
	fetchImageConfigReturnsOnCall map[int]struct {
		result1 *v1.Config
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeImageConfigFetcher) FetchImageConfig(arg1 string, arg2 string, arg3 string) (*v1.Config, error) {
	fake.fetchImageConfigMutex.Lock()
	ret, specificReturn := fake.fetchImageConfigReturnsOnCall[len(fake.fetchImageConfigArgsForCall)]
	fake.fetchImageConfigArgsForCall = append(fake.fetchImageConfigArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("FetchImageConfig", []interface{}{arg1, arg2, arg3})
	fake.fetchImageConfigMutex.Unlock()
	if fake.FetchImageConfigStub != nil {
		return fake.FetchImageConfigStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.fetchImageConfigReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImageConfigFetcher) FetchImageConfigCallCount() int {
	fake.fetchImageConfigMutex.RLock()
	defer fake.fetchImageConfigMutex.RUnlock()
	return len(fake.fetchImageConfigArgsForCall)
}

func (fake *FakeImageConfigFetcher) FetchImageConfigCalls(stub func(string, string, string) (*v1.Config, error)) {
	fake.fetchImageConfigMutex.Lock()
	defer fake.fetchImageConfigMutex.Unlock()
	fake.FetchImageConfigStub = stub
}

func (fake *FakeImageConfigFetcher) FetchImageConfigArgsForCall(i int) (string, string, string) {
	fake.fetchImageConfigMutex.RLock()
	defer fake.fetchImageConfigMutex.RUnlock()
	argsForCall := fake.fetchImageConfigArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeImageConfigFetcher) FetchImageConfigReturns(result1 *v1.Config, result2 error) {
	fake.fetchImageConfigMutex.Lock()
	defer fake.fetchImageConfigMutex.Unlock()
	fake.FetchImageConfigStub = nil
	fake.fetchImageConfigReturns = struct {
		result1 *v1.Config
		result2 error
	}{result1, result2}
}

func (fake *FakeImageConfigFetcher) FetchImageConfigReturnsOnCall(i int, result1 *v1.Config, result2 error) {
	fake.fetchImageConfigMutex.Lock()
	defer fake.fetchImageConfigMutex.Unlock()
	fake.FetchImageConfigStub = nil
	if fake.fetchImageConfigReturnsOnCall == nil {
		fake.fetchImageConfigReturnsOnCall = make(map[int]struct {
			result1 *v1.Config
			result2 error
		})
	}
	fake.fetchImageConfigReturnsOnCall[i] = struct {
		result1 *v1.Config
		result2 error
	}{result1, result2}
}

func (fake *FakeImageConfigFetcher) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.fetchImageConfigMutex.RLock()
	defer fake.fetchImageConfigMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeImageConfigFetcher) recordInvocation(key string, args []interface{}) {
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

var _ image_registry.ImageConfigFetcher = new(FakeImageConfigFetcher)
