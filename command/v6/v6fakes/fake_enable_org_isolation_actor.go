// Code generated by counterfeiter. DO NOT EDIT.
package v6fakes

import (
	sync "sync"

	v3action "code.cloudfoundry.org/cli/actor/v3action"
	v6 "code.cloudfoundry.org/cli/command/v6"
)

type FakeEnableOrgIsolationActor struct {
	EntitleIsolationSegmentToOrganizationByNameStub        func(string, string) (v3action.Warnings, error)
	entitleIsolationSegmentToOrganizationByNameMutex       sync.RWMutex
	entitleIsolationSegmentToOrganizationByNameArgsForCall []struct {
		arg1 string
		arg2 string
	}
	entitleIsolationSegmentToOrganizationByNameReturns struct {
		result1 v3action.Warnings
		result2 error
	}
	entitleIsolationSegmentToOrganizationByNameReturnsOnCall map[int]struct {
		result1 v3action.Warnings
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEnableOrgIsolationActor) EntitleIsolationSegmentToOrganizationByName(arg1 string, arg2 string) (v3action.Warnings, error) {
	fake.entitleIsolationSegmentToOrganizationByNameMutex.Lock()
	ret, specificReturn := fake.entitleIsolationSegmentToOrganizationByNameReturnsOnCall[len(fake.entitleIsolationSegmentToOrganizationByNameArgsForCall)]
	fake.entitleIsolationSegmentToOrganizationByNameArgsForCall = append(fake.entitleIsolationSegmentToOrganizationByNameArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("EntitleIsolationSegmentToOrganizationByName", []interface{}{arg1, arg2})
	fake.entitleIsolationSegmentToOrganizationByNameMutex.Unlock()
	if fake.EntitleIsolationSegmentToOrganizationByNameStub != nil {
		return fake.EntitleIsolationSegmentToOrganizationByNameStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.entitleIsolationSegmentToOrganizationByNameReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeEnableOrgIsolationActor) EntitleIsolationSegmentToOrganizationByNameCallCount() int {
	fake.entitleIsolationSegmentToOrganizationByNameMutex.RLock()
	defer fake.entitleIsolationSegmentToOrganizationByNameMutex.RUnlock()
	return len(fake.entitleIsolationSegmentToOrganizationByNameArgsForCall)
}

func (fake *FakeEnableOrgIsolationActor) EntitleIsolationSegmentToOrganizationByNameCalls(stub func(string, string) (v3action.Warnings, error)) {
	fake.entitleIsolationSegmentToOrganizationByNameMutex.Lock()
	defer fake.entitleIsolationSegmentToOrganizationByNameMutex.Unlock()
	fake.EntitleIsolationSegmentToOrganizationByNameStub = stub
}

func (fake *FakeEnableOrgIsolationActor) EntitleIsolationSegmentToOrganizationByNameArgsForCall(i int) (string, string) {
	fake.entitleIsolationSegmentToOrganizationByNameMutex.RLock()
	defer fake.entitleIsolationSegmentToOrganizationByNameMutex.RUnlock()
	argsForCall := fake.entitleIsolationSegmentToOrganizationByNameArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeEnableOrgIsolationActor) EntitleIsolationSegmentToOrganizationByNameReturns(result1 v3action.Warnings, result2 error) {
	fake.entitleIsolationSegmentToOrganizationByNameMutex.Lock()
	defer fake.entitleIsolationSegmentToOrganizationByNameMutex.Unlock()
	fake.EntitleIsolationSegmentToOrganizationByNameStub = nil
	fake.entitleIsolationSegmentToOrganizationByNameReturns = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeEnableOrgIsolationActor) EntitleIsolationSegmentToOrganizationByNameReturnsOnCall(i int, result1 v3action.Warnings, result2 error) {
	fake.entitleIsolationSegmentToOrganizationByNameMutex.Lock()
	defer fake.entitleIsolationSegmentToOrganizationByNameMutex.Unlock()
	fake.EntitleIsolationSegmentToOrganizationByNameStub = nil
	if fake.entitleIsolationSegmentToOrganizationByNameReturnsOnCall == nil {
		fake.entitleIsolationSegmentToOrganizationByNameReturnsOnCall = make(map[int]struct {
			result1 v3action.Warnings
			result2 error
		})
	}
	fake.entitleIsolationSegmentToOrganizationByNameReturnsOnCall[i] = struct {
		result1 v3action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeEnableOrgIsolationActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.entitleIsolationSegmentToOrganizationByNameMutex.RLock()
	defer fake.entitleIsolationSegmentToOrganizationByNameMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeEnableOrgIsolationActor) recordInvocation(key string, args []interface{}) {
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

var _ v6.EnableOrgIsolationActor = new(FakeEnableOrgIsolationActor)