// This file was generated by counterfeiter
package fakes

import "sync"

type CCClient struct {
	GetLiveAppGUIDsStub        func(token string, appGUIDs []string) (map[string]struct{}, error)
	getLiveAppGUIDsMutex       sync.RWMutex
	getLiveAppGUIDsArgsForCall []struct {
		token    string
		appGUIDs []string
	}
	getLiveAppGUIDsReturns struct {
		result1 map[string]struct{}
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *CCClient) GetLiveAppGUIDs(token string, appGUIDs []string) (map[string]struct{}, error) {
	var appGUIDsCopy []string
	if appGUIDs != nil {
		appGUIDsCopy = make([]string, len(appGUIDs))
		copy(appGUIDsCopy, appGUIDs)
	}
	fake.getLiveAppGUIDsMutex.Lock()
	fake.getLiveAppGUIDsArgsForCall = append(fake.getLiveAppGUIDsArgsForCall, struct {
		token    string
		appGUIDs []string
	}{token, appGUIDsCopy})
	fake.recordInvocation("GetLiveAppGUIDs", []interface{}{token, appGUIDsCopy})
	fake.getLiveAppGUIDsMutex.Unlock()
	if fake.GetLiveAppGUIDsStub != nil {
		return fake.GetLiveAppGUIDsStub(token, appGUIDs)
	} else {
		return fake.getLiveAppGUIDsReturns.result1, fake.getLiveAppGUIDsReturns.result2
	}
}

func (fake *CCClient) GetLiveAppGUIDsCallCount() int {
	fake.getLiveAppGUIDsMutex.RLock()
	defer fake.getLiveAppGUIDsMutex.RUnlock()
	return len(fake.getLiveAppGUIDsArgsForCall)
}

func (fake *CCClient) GetLiveAppGUIDsArgsForCall(i int) (string, []string) {
	fake.getLiveAppGUIDsMutex.RLock()
	defer fake.getLiveAppGUIDsMutex.RUnlock()
	return fake.getLiveAppGUIDsArgsForCall[i].token, fake.getLiveAppGUIDsArgsForCall[i].appGUIDs
}

func (fake *CCClient) GetLiveAppGUIDsReturns(result1 map[string]struct{}, result2 error) {
	fake.GetLiveAppGUIDsStub = nil
	fake.getLiveAppGUIDsReturns = struct {
		result1 map[string]struct{}
		result2 error
	}{result1, result2}
}

func (fake *CCClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getLiveAppGUIDsMutex.RLock()
	defer fake.getLiveAppGUIDsMutex.RUnlock()
	return fake.invocations
}

func (fake *CCClient) recordInvocation(key string, args []interface{}) {
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
