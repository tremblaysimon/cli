// Code generated by counterfeiter. DO NOT EDIT.
package apifakes

import (
	"sync"

	"code.cloudfoundry.org/cli/cf/api"
	"code.cloudfoundry.org/cli/cf/models"
)

type FakeServicePlanVisibilityRepository struct {
	CreateStub        func(string, string) error
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 string
		arg2 string
	}
	createReturns struct {
		result1 error
	}
	createReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteStub        func(string) error
	deleteMutex       sync.RWMutex
	deleteArgsForCall []struct {
		arg1 string
	}
	deleteReturns struct {
		result1 error
	}
	deleteReturnsOnCall map[int]struct {
		result1 error
	}
	ListStub        func() ([]models.ServicePlanVisibilityFields, error)
	listMutex       sync.RWMutex
	listArgsForCall []struct {
	}
	listReturns struct {
		result1 []models.ServicePlanVisibilityFields
		result2 error
	}
	listReturnsOnCall map[int]struct {
		result1 []models.ServicePlanVisibilityFields
		result2 error
	}
	SearchStub        func(map[string]string) ([]models.ServicePlanVisibilityFields, error)
	searchMutex       sync.RWMutex
	searchArgsForCall []struct {
		arg1 map[string]string
	}
	searchReturns struct {
		result1 []models.ServicePlanVisibilityFields
		result2 error
	}
	searchReturnsOnCall map[int]struct {
		result1 []models.ServicePlanVisibilityFields
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeServicePlanVisibilityRepository) Create(arg1 string, arg2 string) error {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("Create", []interface{}{arg1, arg2})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.createReturns
	return fakeReturns.result1
}

func (fake *FakeServicePlanVisibilityRepository) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeServicePlanVisibilityRepository) CreateCalls(stub func(string, string) error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *FakeServicePlanVisibilityRepository) CreateArgsForCall(i int) (string, string) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeServicePlanVisibilityRepository) CreateReturns(result1 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeServicePlanVisibilityRepository) CreateReturnsOnCall(i int, result1 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeServicePlanVisibilityRepository) Delete(arg1 string) error {
	fake.deleteMutex.Lock()
	ret, specificReturn := fake.deleteReturnsOnCall[len(fake.deleteArgsForCall)]
	fake.deleteArgsForCall = append(fake.deleteArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("Delete", []interface{}{arg1})
	fake.deleteMutex.Unlock()
	if fake.DeleteStub != nil {
		return fake.DeleteStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteReturns
	return fakeReturns.result1
}

func (fake *FakeServicePlanVisibilityRepository) DeleteCallCount() int {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	return len(fake.deleteArgsForCall)
}

func (fake *FakeServicePlanVisibilityRepository) DeleteCalls(stub func(string) error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = stub
}

func (fake *FakeServicePlanVisibilityRepository) DeleteArgsForCall(i int) string {
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	argsForCall := fake.deleteArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeServicePlanVisibilityRepository) DeleteReturns(result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	fake.deleteReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeServicePlanVisibilityRepository) DeleteReturnsOnCall(i int, result1 error) {
	fake.deleteMutex.Lock()
	defer fake.deleteMutex.Unlock()
	fake.DeleteStub = nil
	if fake.deleteReturnsOnCall == nil {
		fake.deleteReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeServicePlanVisibilityRepository) List() ([]models.ServicePlanVisibilityFields, error) {
	fake.listMutex.Lock()
	ret, specificReturn := fake.listReturnsOnCall[len(fake.listArgsForCall)]
	fake.listArgsForCall = append(fake.listArgsForCall, struct {
	}{})
	fake.recordInvocation("List", []interface{}{})
	fake.listMutex.Unlock()
	if fake.ListStub != nil {
		return fake.ListStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeServicePlanVisibilityRepository) ListCallCount() int {
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	return len(fake.listArgsForCall)
}

func (fake *FakeServicePlanVisibilityRepository) ListCalls(stub func() ([]models.ServicePlanVisibilityFields, error)) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = stub
}

func (fake *FakeServicePlanVisibilityRepository) ListReturns(result1 []models.ServicePlanVisibilityFields, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	fake.listReturns = struct {
		result1 []models.ServicePlanVisibilityFields
		result2 error
	}{result1, result2}
}

func (fake *FakeServicePlanVisibilityRepository) ListReturnsOnCall(i int, result1 []models.ServicePlanVisibilityFields, result2 error) {
	fake.listMutex.Lock()
	defer fake.listMutex.Unlock()
	fake.ListStub = nil
	if fake.listReturnsOnCall == nil {
		fake.listReturnsOnCall = make(map[int]struct {
			result1 []models.ServicePlanVisibilityFields
			result2 error
		})
	}
	fake.listReturnsOnCall[i] = struct {
		result1 []models.ServicePlanVisibilityFields
		result2 error
	}{result1, result2}
}

func (fake *FakeServicePlanVisibilityRepository) Search(arg1 map[string]string) ([]models.ServicePlanVisibilityFields, error) {
	fake.searchMutex.Lock()
	ret, specificReturn := fake.searchReturnsOnCall[len(fake.searchArgsForCall)]
	fake.searchArgsForCall = append(fake.searchArgsForCall, struct {
		arg1 map[string]string
	}{arg1})
	fake.recordInvocation("Search", []interface{}{arg1})
	fake.searchMutex.Unlock()
	if fake.SearchStub != nil {
		return fake.SearchStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.searchReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeServicePlanVisibilityRepository) SearchCallCount() int {
	fake.searchMutex.RLock()
	defer fake.searchMutex.RUnlock()
	return len(fake.searchArgsForCall)
}

func (fake *FakeServicePlanVisibilityRepository) SearchCalls(stub func(map[string]string) ([]models.ServicePlanVisibilityFields, error)) {
	fake.searchMutex.Lock()
	defer fake.searchMutex.Unlock()
	fake.SearchStub = stub
}

func (fake *FakeServicePlanVisibilityRepository) SearchArgsForCall(i int) map[string]string {
	fake.searchMutex.RLock()
	defer fake.searchMutex.RUnlock()
	argsForCall := fake.searchArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeServicePlanVisibilityRepository) SearchReturns(result1 []models.ServicePlanVisibilityFields, result2 error) {
	fake.searchMutex.Lock()
	defer fake.searchMutex.Unlock()
	fake.SearchStub = nil
	fake.searchReturns = struct {
		result1 []models.ServicePlanVisibilityFields
		result2 error
	}{result1, result2}
}

func (fake *FakeServicePlanVisibilityRepository) SearchReturnsOnCall(i int, result1 []models.ServicePlanVisibilityFields, result2 error) {
	fake.searchMutex.Lock()
	defer fake.searchMutex.Unlock()
	fake.SearchStub = nil
	if fake.searchReturnsOnCall == nil {
		fake.searchReturnsOnCall = make(map[int]struct {
			result1 []models.ServicePlanVisibilityFields
			result2 error
		})
	}
	fake.searchReturnsOnCall[i] = struct {
		result1 []models.ServicePlanVisibilityFields
		result2 error
	}{result1, result2}
}

func (fake *FakeServicePlanVisibilityRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.deleteMutex.RLock()
	defer fake.deleteMutex.RUnlock()
	fake.listMutex.RLock()
	defer fake.listMutex.RUnlock()
	fake.searchMutex.RLock()
	defer fake.searchMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeServicePlanVisibilityRepository) recordInvocation(key string, args []interface{}) {
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

var _ api.ServicePlanVisibilityRepository = new(FakeServicePlanVisibilityRepository)