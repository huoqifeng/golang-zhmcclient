// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.ibm.com/zhmcclient/golang-zhmcclient/pkg/zhmcclient"
)

type JobAPI struct {
	CancelJobStub        func(string) error
	cancelJobMutex       sync.RWMutex
	cancelJobArgsForCall []struct {
		arg1 string
	}
	cancelJobReturns struct {
		result1 error
	}
	cancelJobReturnsOnCall map[int]struct {
		result1 error
	}
	DeleteJobStub        func(string) error
	deleteJobMutex       sync.RWMutex
	deleteJobArgsForCall []struct {
		arg1 string
	}
	deleteJobReturns struct {
		result1 error
	}
	deleteJobReturnsOnCall map[int]struct {
		result1 error
	}
	QueryJobStub        func(string) (*zhmcclient.Job, error)
	queryJobMutex       sync.RWMutex
	queryJobArgsForCall []struct {
		arg1 string
	}
	queryJobReturns struct {
		result1 *zhmcclient.Job
		result2 error
	}
	queryJobReturnsOnCall map[int]struct {
		result1 *zhmcclient.Job
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *JobAPI) CancelJob(arg1 string) error {
	fake.cancelJobMutex.Lock()
	ret, specificReturn := fake.cancelJobReturnsOnCall[len(fake.cancelJobArgsForCall)]
	fake.cancelJobArgsForCall = append(fake.cancelJobArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.CancelJobStub
	fakeReturns := fake.cancelJobReturns
	fake.recordInvocation("CancelJob", []interface{}{arg1})
	fake.cancelJobMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *JobAPI) CancelJobCallCount() int {
	fake.cancelJobMutex.RLock()
	defer fake.cancelJobMutex.RUnlock()
	return len(fake.cancelJobArgsForCall)
}

func (fake *JobAPI) CancelJobCalls(stub func(string) error) {
	fake.cancelJobMutex.Lock()
	defer fake.cancelJobMutex.Unlock()
	fake.CancelJobStub = stub
}

func (fake *JobAPI) CancelJobArgsForCall(i int) string {
	fake.cancelJobMutex.RLock()
	defer fake.cancelJobMutex.RUnlock()
	argsForCall := fake.cancelJobArgsForCall[i]
	return argsForCall.arg1
}

func (fake *JobAPI) CancelJobReturns(result1 error) {
	fake.cancelJobMutex.Lock()
	defer fake.cancelJobMutex.Unlock()
	fake.CancelJobStub = nil
	fake.cancelJobReturns = struct {
		result1 error
	}{result1}
}

func (fake *JobAPI) CancelJobReturnsOnCall(i int, result1 error) {
	fake.cancelJobMutex.Lock()
	defer fake.cancelJobMutex.Unlock()
	fake.CancelJobStub = nil
	if fake.cancelJobReturnsOnCall == nil {
		fake.cancelJobReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.cancelJobReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *JobAPI) DeleteJob(arg1 string) error {
	fake.deleteJobMutex.Lock()
	ret, specificReturn := fake.deleteJobReturnsOnCall[len(fake.deleteJobArgsForCall)]
	fake.deleteJobArgsForCall = append(fake.deleteJobArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.DeleteJobStub
	fakeReturns := fake.deleteJobReturns
	fake.recordInvocation("DeleteJob", []interface{}{arg1})
	fake.deleteJobMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *JobAPI) DeleteJobCallCount() int {
	fake.deleteJobMutex.RLock()
	defer fake.deleteJobMutex.RUnlock()
	return len(fake.deleteJobArgsForCall)
}

func (fake *JobAPI) DeleteJobCalls(stub func(string) error) {
	fake.deleteJobMutex.Lock()
	defer fake.deleteJobMutex.Unlock()
	fake.DeleteJobStub = stub
}

func (fake *JobAPI) DeleteJobArgsForCall(i int) string {
	fake.deleteJobMutex.RLock()
	defer fake.deleteJobMutex.RUnlock()
	argsForCall := fake.deleteJobArgsForCall[i]
	return argsForCall.arg1
}

func (fake *JobAPI) DeleteJobReturns(result1 error) {
	fake.deleteJobMutex.Lock()
	defer fake.deleteJobMutex.Unlock()
	fake.DeleteJobStub = nil
	fake.deleteJobReturns = struct {
		result1 error
	}{result1}
}

func (fake *JobAPI) DeleteJobReturnsOnCall(i int, result1 error) {
	fake.deleteJobMutex.Lock()
	defer fake.deleteJobMutex.Unlock()
	fake.DeleteJobStub = nil
	if fake.deleteJobReturnsOnCall == nil {
		fake.deleteJobReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteJobReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *JobAPI) QueryJob(arg1 string) (*zhmcclient.Job, error) {
	fake.queryJobMutex.Lock()
	ret, specificReturn := fake.queryJobReturnsOnCall[len(fake.queryJobArgsForCall)]
	fake.queryJobArgsForCall = append(fake.queryJobArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.QueryJobStub
	fakeReturns := fake.queryJobReturns
	fake.recordInvocation("QueryJob", []interface{}{arg1})
	fake.queryJobMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *JobAPI) QueryJobCallCount() int {
	fake.queryJobMutex.RLock()
	defer fake.queryJobMutex.RUnlock()
	return len(fake.queryJobArgsForCall)
}

func (fake *JobAPI) QueryJobCalls(stub func(string) (*zhmcclient.Job, error)) {
	fake.queryJobMutex.Lock()
	defer fake.queryJobMutex.Unlock()
	fake.QueryJobStub = stub
}

func (fake *JobAPI) QueryJobArgsForCall(i int) string {
	fake.queryJobMutex.RLock()
	defer fake.queryJobMutex.RUnlock()
	argsForCall := fake.queryJobArgsForCall[i]
	return argsForCall.arg1
}

func (fake *JobAPI) QueryJobReturns(result1 *zhmcclient.Job, result2 error) {
	fake.queryJobMutex.Lock()
	defer fake.queryJobMutex.Unlock()
	fake.QueryJobStub = nil
	fake.queryJobReturns = struct {
		result1 *zhmcclient.Job
		result2 error
	}{result1, result2}
}

func (fake *JobAPI) QueryJobReturnsOnCall(i int, result1 *zhmcclient.Job, result2 error) {
	fake.queryJobMutex.Lock()
	defer fake.queryJobMutex.Unlock()
	fake.QueryJobStub = nil
	if fake.queryJobReturnsOnCall == nil {
		fake.queryJobReturnsOnCall = make(map[int]struct {
			result1 *zhmcclient.Job
			result2 error
		})
	}
	fake.queryJobReturnsOnCall[i] = struct {
		result1 *zhmcclient.Job
		result2 error
	}{result1, result2}
}

func (fake *JobAPI) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.cancelJobMutex.RLock()
	defer fake.cancelJobMutex.RUnlock()
	fake.deleteJobMutex.RLock()
	defer fake.deleteJobMutex.RUnlock()
	fake.queryJobMutex.RLock()
	defer fake.queryJobMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *JobAPI) recordInvocation(key string, args []interface{}) {
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

var _ zhmcclient.JobAPI = new(JobAPI)