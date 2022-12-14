// Code generated by counterfeiter. DO NOT EDIT.
package servicefakes

import (
	"example-project/model"
	"example-project/service"
	"sync"
)

type FakeDatabaseInterface struct {
	GetByIDStub        func(string) model.Employee
	getByIDMutex       sync.RWMutex
	getByIDArgsForCall []struct {
		arg1 string
	}
	getByIDReturns struct {
		result1 model.Employee
	}
	getByIDReturnsOnCall map[int]struct {
		result1 model.Employee
	}
	UpdateManyStub        func([]interface{}) interface{}
	updateManyMutex       sync.RWMutex
	updateManyArgsForCall []struct {
		arg1 []interface{}
	}
	updateManyReturns struct {
		result1 interface{}
	}
	updateManyReturnsOnCall map[int]struct {
		result1 interface{}
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDatabaseInterface) GetByID(arg1 string) model.Employee {
	fake.getByIDMutex.Lock()
	ret, specificReturn := fake.getByIDReturnsOnCall[len(fake.getByIDArgsForCall)]
	fake.getByIDArgsForCall = append(fake.getByIDArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetByIDStub
	fakeReturns := fake.getByIDReturns
	fake.recordInvocation("GetByID", []interface{}{arg1})
	fake.getByIDMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeDatabaseInterface) GetByIDCallCount() int {
	fake.getByIDMutex.RLock()
	defer fake.getByIDMutex.RUnlock()
	return len(fake.getByIDArgsForCall)
}

func (fake *FakeDatabaseInterface) GetByIDCalls(stub func(string) model.Employee) {
	fake.getByIDMutex.Lock()
	defer fake.getByIDMutex.Unlock()
	fake.GetByIDStub = stub
}

func (fake *FakeDatabaseInterface) GetByIDArgsForCall(i int) string {
	fake.getByIDMutex.RLock()
	defer fake.getByIDMutex.RUnlock()
	argsForCall := fake.getByIDArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDatabaseInterface) GetByIDReturns(result1 model.Employee) {
	fake.getByIDMutex.Lock()
	defer fake.getByIDMutex.Unlock()
	fake.GetByIDStub = nil
	fake.getByIDReturns = struct {
		result1 model.Employee
	}{result1}
}

func (fake *FakeDatabaseInterface) GetByIDReturnsOnCall(i int, result1 model.Employee) {
	fake.getByIDMutex.Lock()
	defer fake.getByIDMutex.Unlock()
	fake.GetByIDStub = nil
	if fake.getByIDReturnsOnCall == nil {
		fake.getByIDReturnsOnCall = make(map[int]struct {
			result1 model.Employee
		})
	}
	fake.getByIDReturnsOnCall[i] = struct {
		result1 model.Employee
	}{result1}
}

func (fake *FakeDatabaseInterface) UpdateMany(arg1 []interface{}) interface{} {
	var arg1Copy []interface{}
	if arg1 != nil {
		arg1Copy = make([]interface{}, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.updateManyMutex.Lock()
	ret, specificReturn := fake.updateManyReturnsOnCall[len(fake.updateManyArgsForCall)]
	fake.updateManyArgsForCall = append(fake.updateManyArgsForCall, struct {
		arg1 []interface{}
	}{arg1Copy})
	stub := fake.UpdateManyStub
	fakeReturns := fake.updateManyReturns
	fake.recordInvocation("UpdateMany", []interface{}{arg1Copy})
	fake.updateManyMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeDatabaseInterface) UpdateManyCallCount() int {
	fake.updateManyMutex.RLock()
	defer fake.updateManyMutex.RUnlock()
	return len(fake.updateManyArgsForCall)
}

func (fake *FakeDatabaseInterface) UpdateManyCalls(stub func([]interface{}) interface{}) {
	fake.updateManyMutex.Lock()
	defer fake.updateManyMutex.Unlock()
	fake.UpdateManyStub = stub
}

func (fake *FakeDatabaseInterface) UpdateManyArgsForCall(i int) []interface{} {
	fake.updateManyMutex.RLock()
	defer fake.updateManyMutex.RUnlock()
	argsForCall := fake.updateManyArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDatabaseInterface) UpdateManyReturns(result1 interface{}) {
	fake.updateManyMutex.Lock()
	defer fake.updateManyMutex.Unlock()
	fake.UpdateManyStub = nil
	fake.updateManyReturns = struct {
		result1 interface{}
	}{result1}
}

func (fake *FakeDatabaseInterface) UpdateManyReturnsOnCall(i int, result1 interface{}) {
	fake.updateManyMutex.Lock()
	defer fake.updateManyMutex.Unlock()
	fake.UpdateManyStub = nil
	if fake.updateManyReturnsOnCall == nil {
		fake.updateManyReturnsOnCall = make(map[int]struct {
			result1 interface{}
		})
	}
	fake.updateManyReturnsOnCall[i] = struct {
		result1 interface{}
	}{result1}
}

func (fake *FakeDatabaseInterface) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getByIDMutex.RLock()
	defer fake.getByIDMutex.RUnlock()
	fake.updateManyMutex.RLock()
	defer fake.updateManyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDatabaseInterface) recordInvocation(key string, args []interface{}) {
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

var _ service.DatabaseInterface = new(FakeDatabaseInterface)
