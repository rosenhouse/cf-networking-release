// This file was generated by counterfeiter
package fakes

import "sync"

type UAAClient struct {
	GetTokenStub        func() (string, error)
	getTokenMutex       sync.RWMutex
	getTokenArgsForCall []struct{}
	getTokenReturns     struct {
		result1 string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *UAAClient) GetToken() (string, error) {
	fake.getTokenMutex.Lock()
	fake.getTokenArgsForCall = append(fake.getTokenArgsForCall, struct{}{})
	fake.recordInvocation("GetToken", []interface{}{})
	fake.getTokenMutex.Unlock()
	if fake.GetTokenStub != nil {
		return fake.GetTokenStub()
	} else {
		return fake.getTokenReturns.result1, fake.getTokenReturns.result2
	}
}

func (fake *UAAClient) GetTokenCallCount() int {
	fake.getTokenMutex.RLock()
	defer fake.getTokenMutex.RUnlock()
	return len(fake.getTokenArgsForCall)
}

func (fake *UAAClient) GetTokenReturns(result1 string, result2 error) {
	fake.GetTokenStub = nil
	fake.getTokenReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *UAAClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getTokenMutex.RLock()
	defer fake.getTokenMutex.RUnlock()
	return fake.invocations
}

func (fake *UAAClient) recordInvocation(key string, args []interface{}) {
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
