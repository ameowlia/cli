// Code generated by counterfeiter. DO NOT EDIT.
package v7pushactionfakes

import (
	io "io"
	sync "sync"

	sharedaction "code.cloudfoundry.org/cli/actor/sharedaction"
	v7action "code.cloudfoundry.org/cli/actor/v7action"
	v7pushaction "code.cloudfoundry.org/cli/actor/v7pushaction"
)

type FakeV7Actor struct {
	CreateApplicationInSpaceStub        func(v7action.Application, string) (v7action.Application, v7action.Warnings, error)
	createApplicationInSpaceMutex       sync.RWMutex
	createApplicationInSpaceArgsForCall []struct {
		arg1 v7action.Application
		arg2 string
	}
	createApplicationInSpaceReturns struct {
		result1 v7action.Application
		result2 v7action.Warnings
		result3 error
	}
	createApplicationInSpaceReturnsOnCall map[int]struct {
		result1 v7action.Application
		result2 v7action.Warnings
		result3 error
	}
	CreateBitsPackageByApplicationStub        func(string) (v7action.Package, v7action.Warnings, error)
	createBitsPackageByApplicationMutex       sync.RWMutex
	createBitsPackageByApplicationArgsForCall []struct {
		arg1 string
	}
	createBitsPackageByApplicationReturns struct {
		result1 v7action.Package
		result2 v7action.Warnings
		result3 error
	}
	createBitsPackageByApplicationReturnsOnCall map[int]struct {
		result1 v7action.Package
		result2 v7action.Warnings
		result3 error
	}
	CreateDockerPackageByApplicationStub        func(string, v7action.DockerImageCredentials) (v7action.Package, v7action.Warnings, error)
	createDockerPackageByApplicationMutex       sync.RWMutex
	createDockerPackageByApplicationArgsForCall []struct {
		arg1 string
		arg2 v7action.DockerImageCredentials
	}
	createDockerPackageByApplicationReturns struct {
		result1 v7action.Package
		result2 v7action.Warnings
		result3 error
	}
	createDockerPackageByApplicationReturnsOnCall map[int]struct {
		result1 v7action.Package
		result2 v7action.Warnings
		result3 error
	}
	GetApplicationByNameAndSpaceStub        func(string, string) (v7action.Application, v7action.Warnings, error)
	getApplicationByNameAndSpaceMutex       sync.RWMutex
	getApplicationByNameAndSpaceArgsForCall []struct {
		arg1 string
		arg2 string
	}
	getApplicationByNameAndSpaceReturns struct {
		result1 v7action.Application
		result2 v7action.Warnings
		result3 error
	}
	getApplicationByNameAndSpaceReturnsOnCall map[int]struct {
		result1 v7action.Application
		result2 v7action.Warnings
		result3 error
	}
	PollBuildStub        func(string, string) (v7action.Droplet, v7action.Warnings, error)
	pollBuildMutex       sync.RWMutex
	pollBuildArgsForCall []struct {
		arg1 string
		arg2 string
	}
	pollBuildReturns struct {
		result1 v7action.Droplet
		result2 v7action.Warnings
		result3 error
	}
	pollBuildReturnsOnCall map[int]struct {
		result1 v7action.Droplet
		result2 v7action.Warnings
		result3 error
	}
	PollPackageStub        func(v7action.Package) (v7action.Package, v7action.Warnings, error)
	pollPackageMutex       sync.RWMutex
	pollPackageArgsForCall []struct {
		arg1 v7action.Package
	}
	pollPackageReturns struct {
		result1 v7action.Package
		result2 v7action.Warnings
		result3 error
	}
	pollPackageReturnsOnCall map[int]struct {
		result1 v7action.Package
		result2 v7action.Warnings
		result3 error
	}
	ScaleProcessByApplicationStub        func(string, v7action.Process) (v7action.Warnings, error)
	scaleProcessByApplicationMutex       sync.RWMutex
	scaleProcessByApplicationArgsForCall []struct {
		arg1 string
		arg2 v7action.Process
	}
	scaleProcessByApplicationReturns struct {
		result1 v7action.Warnings
		result2 error
	}
	scaleProcessByApplicationReturnsOnCall map[int]struct {
		result1 v7action.Warnings
		result2 error
	}
	SetApplicationDropletStub        func(string, string) (v7action.Warnings, error)
	setApplicationDropletMutex       sync.RWMutex
	setApplicationDropletArgsForCall []struct {
		arg1 string
		arg2 string
	}
	setApplicationDropletReturns struct {
		result1 v7action.Warnings
		result2 error
	}
	setApplicationDropletReturnsOnCall map[int]struct {
		result1 v7action.Warnings
		result2 error
	}
	StageApplicationPackageStub        func(string) (v7action.Build, v7action.Warnings, error)
	stageApplicationPackageMutex       sync.RWMutex
	stageApplicationPackageArgsForCall []struct {
		arg1 string
	}
	stageApplicationPackageReturns struct {
		result1 v7action.Build
		result2 v7action.Warnings
		result3 error
	}
	stageApplicationPackageReturnsOnCall map[int]struct {
		result1 v7action.Build
		result2 v7action.Warnings
		result3 error
	}
	UpdateApplicationStub        func(v7action.Application) (v7action.Application, v7action.Warnings, error)
	updateApplicationMutex       sync.RWMutex
	updateApplicationArgsForCall []struct {
		arg1 v7action.Application
	}
	updateApplicationReturns struct {
		result1 v7action.Application
		result2 v7action.Warnings
		result3 error
	}
	updateApplicationReturnsOnCall map[int]struct {
		result1 v7action.Application
		result2 v7action.Warnings
		result3 error
	}
	UpdateProcessByTypeAndApplicationStub        func(string, string, v7action.Process) (v7action.Warnings, error)
	updateProcessByTypeAndApplicationMutex       sync.RWMutex
	updateProcessByTypeAndApplicationArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 v7action.Process
	}
	updateProcessByTypeAndApplicationReturns struct {
		result1 v7action.Warnings
		result2 error
	}
	updateProcessByTypeAndApplicationReturnsOnCall map[int]struct {
		result1 v7action.Warnings
		result2 error
	}
	UploadBitsPackageStub        func(v7action.Package, []sharedaction.Resource, io.Reader, int64) (v7action.Package, v7action.Warnings, error)
	uploadBitsPackageMutex       sync.RWMutex
	uploadBitsPackageArgsForCall []struct {
		arg1 v7action.Package
		arg2 []sharedaction.Resource
		arg3 io.Reader
		arg4 int64
	}
	uploadBitsPackageReturns struct {
		result1 v7action.Package
		result2 v7action.Warnings
		result3 error
	}
	uploadBitsPackageReturnsOnCall map[int]struct {
		result1 v7action.Package
		result2 v7action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeV7Actor) CreateApplicationInSpace(arg1 v7action.Application, arg2 string) (v7action.Application, v7action.Warnings, error) {
	fake.createApplicationInSpaceMutex.Lock()
	ret, specificReturn := fake.createApplicationInSpaceReturnsOnCall[len(fake.createApplicationInSpaceArgsForCall)]
	fake.createApplicationInSpaceArgsForCall = append(fake.createApplicationInSpaceArgsForCall, struct {
		arg1 v7action.Application
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("CreateApplicationInSpace", []interface{}{arg1, arg2})
	fake.createApplicationInSpaceMutex.Unlock()
	if fake.CreateApplicationInSpaceStub != nil {
		return fake.CreateApplicationInSpaceStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.createApplicationInSpaceReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeV7Actor) CreateApplicationInSpaceCallCount() int {
	fake.createApplicationInSpaceMutex.RLock()
	defer fake.createApplicationInSpaceMutex.RUnlock()
	return len(fake.createApplicationInSpaceArgsForCall)
}

func (fake *FakeV7Actor) CreateApplicationInSpaceCalls(stub func(v7action.Application, string) (v7action.Application, v7action.Warnings, error)) {
	fake.createApplicationInSpaceMutex.Lock()
	defer fake.createApplicationInSpaceMutex.Unlock()
	fake.CreateApplicationInSpaceStub = stub
}

func (fake *FakeV7Actor) CreateApplicationInSpaceArgsForCall(i int) (v7action.Application, string) {
	fake.createApplicationInSpaceMutex.RLock()
	defer fake.createApplicationInSpaceMutex.RUnlock()
	argsForCall := fake.createApplicationInSpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeV7Actor) CreateApplicationInSpaceReturns(result1 v7action.Application, result2 v7action.Warnings, result3 error) {
	fake.createApplicationInSpaceMutex.Lock()
	defer fake.createApplicationInSpaceMutex.Unlock()
	fake.CreateApplicationInSpaceStub = nil
	fake.createApplicationInSpaceReturns = struct {
		result1 v7action.Application
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV7Actor) CreateApplicationInSpaceReturnsOnCall(i int, result1 v7action.Application, result2 v7action.Warnings, result3 error) {
	fake.createApplicationInSpaceMutex.Lock()
	defer fake.createApplicationInSpaceMutex.Unlock()
	fake.CreateApplicationInSpaceStub = nil
	if fake.createApplicationInSpaceReturnsOnCall == nil {
		fake.createApplicationInSpaceReturnsOnCall = make(map[int]struct {
			result1 v7action.Application
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.createApplicationInSpaceReturnsOnCall[i] = struct {
		result1 v7action.Application
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV7Actor) CreateBitsPackageByApplication(arg1 string) (v7action.Package, v7action.Warnings, error) {
	fake.createBitsPackageByApplicationMutex.Lock()
	ret, specificReturn := fake.createBitsPackageByApplicationReturnsOnCall[len(fake.createBitsPackageByApplicationArgsForCall)]
	fake.createBitsPackageByApplicationArgsForCall = append(fake.createBitsPackageByApplicationArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("CreateBitsPackageByApplication", []interface{}{arg1})
	fake.createBitsPackageByApplicationMutex.Unlock()
	if fake.CreateBitsPackageByApplicationStub != nil {
		return fake.CreateBitsPackageByApplicationStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.createBitsPackageByApplicationReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeV7Actor) CreateBitsPackageByApplicationCallCount() int {
	fake.createBitsPackageByApplicationMutex.RLock()
	defer fake.createBitsPackageByApplicationMutex.RUnlock()
	return len(fake.createBitsPackageByApplicationArgsForCall)
}

func (fake *FakeV7Actor) CreateBitsPackageByApplicationCalls(stub func(string) (v7action.Package, v7action.Warnings, error)) {
	fake.createBitsPackageByApplicationMutex.Lock()
	defer fake.createBitsPackageByApplicationMutex.Unlock()
	fake.CreateBitsPackageByApplicationStub = stub
}

func (fake *FakeV7Actor) CreateBitsPackageByApplicationArgsForCall(i int) string {
	fake.createBitsPackageByApplicationMutex.RLock()
	defer fake.createBitsPackageByApplicationMutex.RUnlock()
	argsForCall := fake.createBitsPackageByApplicationArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeV7Actor) CreateBitsPackageByApplicationReturns(result1 v7action.Package, result2 v7action.Warnings, result3 error) {
	fake.createBitsPackageByApplicationMutex.Lock()
	defer fake.createBitsPackageByApplicationMutex.Unlock()
	fake.CreateBitsPackageByApplicationStub = nil
	fake.createBitsPackageByApplicationReturns = struct {
		result1 v7action.Package
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV7Actor) CreateBitsPackageByApplicationReturnsOnCall(i int, result1 v7action.Package, result2 v7action.Warnings, result3 error) {
	fake.createBitsPackageByApplicationMutex.Lock()
	defer fake.createBitsPackageByApplicationMutex.Unlock()
	fake.CreateBitsPackageByApplicationStub = nil
	if fake.createBitsPackageByApplicationReturnsOnCall == nil {
		fake.createBitsPackageByApplicationReturnsOnCall = make(map[int]struct {
			result1 v7action.Package
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.createBitsPackageByApplicationReturnsOnCall[i] = struct {
		result1 v7action.Package
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV7Actor) CreateDockerPackageByApplication(arg1 string, arg2 v7action.DockerImageCredentials) (v7action.Package, v7action.Warnings, error) {
	fake.createDockerPackageByApplicationMutex.Lock()
	ret, specificReturn := fake.createDockerPackageByApplicationReturnsOnCall[len(fake.createDockerPackageByApplicationArgsForCall)]
	fake.createDockerPackageByApplicationArgsForCall = append(fake.createDockerPackageByApplicationArgsForCall, struct {
		arg1 string
		arg2 v7action.DockerImageCredentials
	}{arg1, arg2})
	fake.recordInvocation("CreateDockerPackageByApplication", []interface{}{arg1, arg2})
	fake.createDockerPackageByApplicationMutex.Unlock()
	if fake.CreateDockerPackageByApplicationStub != nil {
		return fake.CreateDockerPackageByApplicationStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.createDockerPackageByApplicationReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeV7Actor) CreateDockerPackageByApplicationCallCount() int {
	fake.createDockerPackageByApplicationMutex.RLock()
	defer fake.createDockerPackageByApplicationMutex.RUnlock()
	return len(fake.createDockerPackageByApplicationArgsForCall)
}

func (fake *FakeV7Actor) CreateDockerPackageByApplicationCalls(stub func(string, v7action.DockerImageCredentials) (v7action.Package, v7action.Warnings, error)) {
	fake.createDockerPackageByApplicationMutex.Lock()
	defer fake.createDockerPackageByApplicationMutex.Unlock()
	fake.CreateDockerPackageByApplicationStub = stub
}

func (fake *FakeV7Actor) CreateDockerPackageByApplicationArgsForCall(i int) (string, v7action.DockerImageCredentials) {
	fake.createDockerPackageByApplicationMutex.RLock()
	defer fake.createDockerPackageByApplicationMutex.RUnlock()
	argsForCall := fake.createDockerPackageByApplicationArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeV7Actor) CreateDockerPackageByApplicationReturns(result1 v7action.Package, result2 v7action.Warnings, result3 error) {
	fake.createDockerPackageByApplicationMutex.Lock()
	defer fake.createDockerPackageByApplicationMutex.Unlock()
	fake.CreateDockerPackageByApplicationStub = nil
	fake.createDockerPackageByApplicationReturns = struct {
		result1 v7action.Package
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV7Actor) CreateDockerPackageByApplicationReturnsOnCall(i int, result1 v7action.Package, result2 v7action.Warnings, result3 error) {
	fake.createDockerPackageByApplicationMutex.Lock()
	defer fake.createDockerPackageByApplicationMutex.Unlock()
	fake.CreateDockerPackageByApplicationStub = nil
	if fake.createDockerPackageByApplicationReturnsOnCall == nil {
		fake.createDockerPackageByApplicationReturnsOnCall = make(map[int]struct {
			result1 v7action.Package
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.createDockerPackageByApplicationReturnsOnCall[i] = struct {
		result1 v7action.Package
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV7Actor) GetApplicationByNameAndSpace(arg1 string, arg2 string) (v7action.Application, v7action.Warnings, error) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	ret, specificReturn := fake.getApplicationByNameAndSpaceReturnsOnCall[len(fake.getApplicationByNameAndSpaceArgsForCall)]
	fake.getApplicationByNameAndSpaceArgsForCall = append(fake.getApplicationByNameAndSpaceArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("GetApplicationByNameAndSpace", []interface{}{arg1, arg2})
	fake.getApplicationByNameAndSpaceMutex.Unlock()
	if fake.GetApplicationByNameAndSpaceStub != nil {
		return fake.GetApplicationByNameAndSpaceStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getApplicationByNameAndSpaceReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeV7Actor) GetApplicationByNameAndSpaceCallCount() int {
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	return len(fake.getApplicationByNameAndSpaceArgsForCall)
}

func (fake *FakeV7Actor) GetApplicationByNameAndSpaceCalls(stub func(string, string) (v7action.Application, v7action.Warnings, error)) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	defer fake.getApplicationByNameAndSpaceMutex.Unlock()
	fake.GetApplicationByNameAndSpaceStub = stub
}

func (fake *FakeV7Actor) GetApplicationByNameAndSpaceArgsForCall(i int) (string, string) {
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	argsForCall := fake.getApplicationByNameAndSpaceArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeV7Actor) GetApplicationByNameAndSpaceReturns(result1 v7action.Application, result2 v7action.Warnings, result3 error) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	defer fake.getApplicationByNameAndSpaceMutex.Unlock()
	fake.GetApplicationByNameAndSpaceStub = nil
	fake.getApplicationByNameAndSpaceReturns = struct {
		result1 v7action.Application
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV7Actor) GetApplicationByNameAndSpaceReturnsOnCall(i int, result1 v7action.Application, result2 v7action.Warnings, result3 error) {
	fake.getApplicationByNameAndSpaceMutex.Lock()
	defer fake.getApplicationByNameAndSpaceMutex.Unlock()
	fake.GetApplicationByNameAndSpaceStub = nil
	if fake.getApplicationByNameAndSpaceReturnsOnCall == nil {
		fake.getApplicationByNameAndSpaceReturnsOnCall = make(map[int]struct {
			result1 v7action.Application
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.getApplicationByNameAndSpaceReturnsOnCall[i] = struct {
		result1 v7action.Application
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV7Actor) PollBuild(arg1 string, arg2 string) (v7action.Droplet, v7action.Warnings, error) {
	fake.pollBuildMutex.Lock()
	ret, specificReturn := fake.pollBuildReturnsOnCall[len(fake.pollBuildArgsForCall)]
	fake.pollBuildArgsForCall = append(fake.pollBuildArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("PollBuild", []interface{}{arg1, arg2})
	fake.pollBuildMutex.Unlock()
	if fake.PollBuildStub != nil {
		return fake.PollBuildStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.pollBuildReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeV7Actor) PollBuildCallCount() int {
	fake.pollBuildMutex.RLock()
	defer fake.pollBuildMutex.RUnlock()
	return len(fake.pollBuildArgsForCall)
}

func (fake *FakeV7Actor) PollBuildCalls(stub func(string, string) (v7action.Droplet, v7action.Warnings, error)) {
	fake.pollBuildMutex.Lock()
	defer fake.pollBuildMutex.Unlock()
	fake.PollBuildStub = stub
}

func (fake *FakeV7Actor) PollBuildArgsForCall(i int) (string, string) {
	fake.pollBuildMutex.RLock()
	defer fake.pollBuildMutex.RUnlock()
	argsForCall := fake.pollBuildArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeV7Actor) PollBuildReturns(result1 v7action.Droplet, result2 v7action.Warnings, result3 error) {
	fake.pollBuildMutex.Lock()
	defer fake.pollBuildMutex.Unlock()
	fake.PollBuildStub = nil
	fake.pollBuildReturns = struct {
		result1 v7action.Droplet
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV7Actor) PollBuildReturnsOnCall(i int, result1 v7action.Droplet, result2 v7action.Warnings, result3 error) {
	fake.pollBuildMutex.Lock()
	defer fake.pollBuildMutex.Unlock()
	fake.PollBuildStub = nil
	if fake.pollBuildReturnsOnCall == nil {
		fake.pollBuildReturnsOnCall = make(map[int]struct {
			result1 v7action.Droplet
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.pollBuildReturnsOnCall[i] = struct {
		result1 v7action.Droplet
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV7Actor) PollPackage(arg1 v7action.Package) (v7action.Package, v7action.Warnings, error) {
	fake.pollPackageMutex.Lock()
	ret, specificReturn := fake.pollPackageReturnsOnCall[len(fake.pollPackageArgsForCall)]
	fake.pollPackageArgsForCall = append(fake.pollPackageArgsForCall, struct {
		arg1 v7action.Package
	}{arg1})
	fake.recordInvocation("PollPackage", []interface{}{arg1})
	fake.pollPackageMutex.Unlock()
	if fake.PollPackageStub != nil {
		return fake.PollPackageStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.pollPackageReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeV7Actor) PollPackageCallCount() int {
	fake.pollPackageMutex.RLock()
	defer fake.pollPackageMutex.RUnlock()
	return len(fake.pollPackageArgsForCall)
}

func (fake *FakeV7Actor) PollPackageCalls(stub func(v7action.Package) (v7action.Package, v7action.Warnings, error)) {
	fake.pollPackageMutex.Lock()
	defer fake.pollPackageMutex.Unlock()
	fake.PollPackageStub = stub
}

func (fake *FakeV7Actor) PollPackageArgsForCall(i int) v7action.Package {
	fake.pollPackageMutex.RLock()
	defer fake.pollPackageMutex.RUnlock()
	argsForCall := fake.pollPackageArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeV7Actor) PollPackageReturns(result1 v7action.Package, result2 v7action.Warnings, result3 error) {
	fake.pollPackageMutex.Lock()
	defer fake.pollPackageMutex.Unlock()
	fake.PollPackageStub = nil
	fake.pollPackageReturns = struct {
		result1 v7action.Package
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV7Actor) PollPackageReturnsOnCall(i int, result1 v7action.Package, result2 v7action.Warnings, result3 error) {
	fake.pollPackageMutex.Lock()
	defer fake.pollPackageMutex.Unlock()
	fake.PollPackageStub = nil
	if fake.pollPackageReturnsOnCall == nil {
		fake.pollPackageReturnsOnCall = make(map[int]struct {
			result1 v7action.Package
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.pollPackageReturnsOnCall[i] = struct {
		result1 v7action.Package
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV7Actor) ScaleProcessByApplication(arg1 string, arg2 v7action.Process) (v7action.Warnings, error) {
	fake.scaleProcessByApplicationMutex.Lock()
	ret, specificReturn := fake.scaleProcessByApplicationReturnsOnCall[len(fake.scaleProcessByApplicationArgsForCall)]
	fake.scaleProcessByApplicationArgsForCall = append(fake.scaleProcessByApplicationArgsForCall, struct {
		arg1 string
		arg2 v7action.Process
	}{arg1, arg2})
	fake.recordInvocation("ScaleProcessByApplication", []interface{}{arg1, arg2})
	fake.scaleProcessByApplicationMutex.Unlock()
	if fake.ScaleProcessByApplicationStub != nil {
		return fake.ScaleProcessByApplicationStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.scaleProcessByApplicationReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeV7Actor) ScaleProcessByApplicationCallCount() int {
	fake.scaleProcessByApplicationMutex.RLock()
	defer fake.scaleProcessByApplicationMutex.RUnlock()
	return len(fake.scaleProcessByApplicationArgsForCall)
}

func (fake *FakeV7Actor) ScaleProcessByApplicationCalls(stub func(string, v7action.Process) (v7action.Warnings, error)) {
	fake.scaleProcessByApplicationMutex.Lock()
	defer fake.scaleProcessByApplicationMutex.Unlock()
	fake.ScaleProcessByApplicationStub = stub
}

func (fake *FakeV7Actor) ScaleProcessByApplicationArgsForCall(i int) (string, v7action.Process) {
	fake.scaleProcessByApplicationMutex.RLock()
	defer fake.scaleProcessByApplicationMutex.RUnlock()
	argsForCall := fake.scaleProcessByApplicationArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeV7Actor) ScaleProcessByApplicationReturns(result1 v7action.Warnings, result2 error) {
	fake.scaleProcessByApplicationMutex.Lock()
	defer fake.scaleProcessByApplicationMutex.Unlock()
	fake.ScaleProcessByApplicationStub = nil
	fake.scaleProcessByApplicationReturns = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeV7Actor) ScaleProcessByApplicationReturnsOnCall(i int, result1 v7action.Warnings, result2 error) {
	fake.scaleProcessByApplicationMutex.Lock()
	defer fake.scaleProcessByApplicationMutex.Unlock()
	fake.ScaleProcessByApplicationStub = nil
	if fake.scaleProcessByApplicationReturnsOnCall == nil {
		fake.scaleProcessByApplicationReturnsOnCall = make(map[int]struct {
			result1 v7action.Warnings
			result2 error
		})
	}
	fake.scaleProcessByApplicationReturnsOnCall[i] = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeV7Actor) SetApplicationDroplet(arg1 string, arg2 string) (v7action.Warnings, error) {
	fake.setApplicationDropletMutex.Lock()
	ret, specificReturn := fake.setApplicationDropletReturnsOnCall[len(fake.setApplicationDropletArgsForCall)]
	fake.setApplicationDropletArgsForCall = append(fake.setApplicationDropletArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("SetApplicationDroplet", []interface{}{arg1, arg2})
	fake.setApplicationDropletMutex.Unlock()
	if fake.SetApplicationDropletStub != nil {
		return fake.SetApplicationDropletStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.setApplicationDropletReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeV7Actor) SetApplicationDropletCallCount() int {
	fake.setApplicationDropletMutex.RLock()
	defer fake.setApplicationDropletMutex.RUnlock()
	return len(fake.setApplicationDropletArgsForCall)
}

func (fake *FakeV7Actor) SetApplicationDropletCalls(stub func(string, string) (v7action.Warnings, error)) {
	fake.setApplicationDropletMutex.Lock()
	defer fake.setApplicationDropletMutex.Unlock()
	fake.SetApplicationDropletStub = stub
}

func (fake *FakeV7Actor) SetApplicationDropletArgsForCall(i int) (string, string) {
	fake.setApplicationDropletMutex.RLock()
	defer fake.setApplicationDropletMutex.RUnlock()
	argsForCall := fake.setApplicationDropletArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeV7Actor) SetApplicationDropletReturns(result1 v7action.Warnings, result2 error) {
	fake.setApplicationDropletMutex.Lock()
	defer fake.setApplicationDropletMutex.Unlock()
	fake.SetApplicationDropletStub = nil
	fake.setApplicationDropletReturns = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeV7Actor) SetApplicationDropletReturnsOnCall(i int, result1 v7action.Warnings, result2 error) {
	fake.setApplicationDropletMutex.Lock()
	defer fake.setApplicationDropletMutex.Unlock()
	fake.SetApplicationDropletStub = nil
	if fake.setApplicationDropletReturnsOnCall == nil {
		fake.setApplicationDropletReturnsOnCall = make(map[int]struct {
			result1 v7action.Warnings
			result2 error
		})
	}
	fake.setApplicationDropletReturnsOnCall[i] = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeV7Actor) StageApplicationPackage(arg1 string) (v7action.Build, v7action.Warnings, error) {
	fake.stageApplicationPackageMutex.Lock()
	ret, specificReturn := fake.stageApplicationPackageReturnsOnCall[len(fake.stageApplicationPackageArgsForCall)]
	fake.stageApplicationPackageArgsForCall = append(fake.stageApplicationPackageArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("StageApplicationPackage", []interface{}{arg1})
	fake.stageApplicationPackageMutex.Unlock()
	if fake.StageApplicationPackageStub != nil {
		return fake.StageApplicationPackageStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.stageApplicationPackageReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeV7Actor) StageApplicationPackageCallCount() int {
	fake.stageApplicationPackageMutex.RLock()
	defer fake.stageApplicationPackageMutex.RUnlock()
	return len(fake.stageApplicationPackageArgsForCall)
}

func (fake *FakeV7Actor) StageApplicationPackageCalls(stub func(string) (v7action.Build, v7action.Warnings, error)) {
	fake.stageApplicationPackageMutex.Lock()
	defer fake.stageApplicationPackageMutex.Unlock()
	fake.StageApplicationPackageStub = stub
}

func (fake *FakeV7Actor) StageApplicationPackageArgsForCall(i int) string {
	fake.stageApplicationPackageMutex.RLock()
	defer fake.stageApplicationPackageMutex.RUnlock()
	argsForCall := fake.stageApplicationPackageArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeV7Actor) StageApplicationPackageReturns(result1 v7action.Build, result2 v7action.Warnings, result3 error) {
	fake.stageApplicationPackageMutex.Lock()
	defer fake.stageApplicationPackageMutex.Unlock()
	fake.StageApplicationPackageStub = nil
	fake.stageApplicationPackageReturns = struct {
		result1 v7action.Build
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV7Actor) StageApplicationPackageReturnsOnCall(i int, result1 v7action.Build, result2 v7action.Warnings, result3 error) {
	fake.stageApplicationPackageMutex.Lock()
	defer fake.stageApplicationPackageMutex.Unlock()
	fake.StageApplicationPackageStub = nil
	if fake.stageApplicationPackageReturnsOnCall == nil {
		fake.stageApplicationPackageReturnsOnCall = make(map[int]struct {
			result1 v7action.Build
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.stageApplicationPackageReturnsOnCall[i] = struct {
		result1 v7action.Build
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV7Actor) UpdateApplication(arg1 v7action.Application) (v7action.Application, v7action.Warnings, error) {
	fake.updateApplicationMutex.Lock()
	ret, specificReturn := fake.updateApplicationReturnsOnCall[len(fake.updateApplicationArgsForCall)]
	fake.updateApplicationArgsForCall = append(fake.updateApplicationArgsForCall, struct {
		arg1 v7action.Application
	}{arg1})
	fake.recordInvocation("UpdateApplication", []interface{}{arg1})
	fake.updateApplicationMutex.Unlock()
	if fake.UpdateApplicationStub != nil {
		return fake.UpdateApplicationStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.updateApplicationReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeV7Actor) UpdateApplicationCallCount() int {
	fake.updateApplicationMutex.RLock()
	defer fake.updateApplicationMutex.RUnlock()
	return len(fake.updateApplicationArgsForCall)
}

func (fake *FakeV7Actor) UpdateApplicationCalls(stub func(v7action.Application) (v7action.Application, v7action.Warnings, error)) {
	fake.updateApplicationMutex.Lock()
	defer fake.updateApplicationMutex.Unlock()
	fake.UpdateApplicationStub = stub
}

func (fake *FakeV7Actor) UpdateApplicationArgsForCall(i int) v7action.Application {
	fake.updateApplicationMutex.RLock()
	defer fake.updateApplicationMutex.RUnlock()
	argsForCall := fake.updateApplicationArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeV7Actor) UpdateApplicationReturns(result1 v7action.Application, result2 v7action.Warnings, result3 error) {
	fake.updateApplicationMutex.Lock()
	defer fake.updateApplicationMutex.Unlock()
	fake.UpdateApplicationStub = nil
	fake.updateApplicationReturns = struct {
		result1 v7action.Application
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV7Actor) UpdateApplicationReturnsOnCall(i int, result1 v7action.Application, result2 v7action.Warnings, result3 error) {
	fake.updateApplicationMutex.Lock()
	defer fake.updateApplicationMutex.Unlock()
	fake.UpdateApplicationStub = nil
	if fake.updateApplicationReturnsOnCall == nil {
		fake.updateApplicationReturnsOnCall = make(map[int]struct {
			result1 v7action.Application
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.updateApplicationReturnsOnCall[i] = struct {
		result1 v7action.Application
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV7Actor) UpdateProcessByTypeAndApplication(arg1 string, arg2 string, arg3 v7action.Process) (v7action.Warnings, error) {
	fake.updateProcessByTypeAndApplicationMutex.Lock()
	ret, specificReturn := fake.updateProcessByTypeAndApplicationReturnsOnCall[len(fake.updateProcessByTypeAndApplicationArgsForCall)]
	fake.updateProcessByTypeAndApplicationArgsForCall = append(fake.updateProcessByTypeAndApplicationArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 v7action.Process
	}{arg1, arg2, arg3})
	fake.recordInvocation("UpdateProcessByTypeAndApplication", []interface{}{arg1, arg2, arg3})
	fake.updateProcessByTypeAndApplicationMutex.Unlock()
	if fake.UpdateProcessByTypeAndApplicationStub != nil {
		return fake.UpdateProcessByTypeAndApplicationStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.updateProcessByTypeAndApplicationReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeV7Actor) UpdateProcessByTypeAndApplicationCallCount() int {
	fake.updateProcessByTypeAndApplicationMutex.RLock()
	defer fake.updateProcessByTypeAndApplicationMutex.RUnlock()
	return len(fake.updateProcessByTypeAndApplicationArgsForCall)
}

func (fake *FakeV7Actor) UpdateProcessByTypeAndApplicationCalls(stub func(string, string, v7action.Process) (v7action.Warnings, error)) {
	fake.updateProcessByTypeAndApplicationMutex.Lock()
	defer fake.updateProcessByTypeAndApplicationMutex.Unlock()
	fake.UpdateProcessByTypeAndApplicationStub = stub
}

func (fake *FakeV7Actor) UpdateProcessByTypeAndApplicationArgsForCall(i int) (string, string, v7action.Process) {
	fake.updateProcessByTypeAndApplicationMutex.RLock()
	defer fake.updateProcessByTypeAndApplicationMutex.RUnlock()
	argsForCall := fake.updateProcessByTypeAndApplicationArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeV7Actor) UpdateProcessByTypeAndApplicationReturns(result1 v7action.Warnings, result2 error) {
	fake.updateProcessByTypeAndApplicationMutex.Lock()
	defer fake.updateProcessByTypeAndApplicationMutex.Unlock()
	fake.UpdateProcessByTypeAndApplicationStub = nil
	fake.updateProcessByTypeAndApplicationReturns = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeV7Actor) UpdateProcessByTypeAndApplicationReturnsOnCall(i int, result1 v7action.Warnings, result2 error) {
	fake.updateProcessByTypeAndApplicationMutex.Lock()
	defer fake.updateProcessByTypeAndApplicationMutex.Unlock()
	fake.UpdateProcessByTypeAndApplicationStub = nil
	if fake.updateProcessByTypeAndApplicationReturnsOnCall == nil {
		fake.updateProcessByTypeAndApplicationReturnsOnCall = make(map[int]struct {
			result1 v7action.Warnings
			result2 error
		})
	}
	fake.updateProcessByTypeAndApplicationReturnsOnCall[i] = struct {
		result1 v7action.Warnings
		result2 error
	}{result1, result2}
}

func (fake *FakeV7Actor) UploadBitsPackage(arg1 v7action.Package, arg2 []sharedaction.Resource, arg3 io.Reader, arg4 int64) (v7action.Package, v7action.Warnings, error) {
	var arg2Copy []sharedaction.Resource
	if arg2 != nil {
		arg2Copy = make([]sharedaction.Resource, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.uploadBitsPackageMutex.Lock()
	ret, specificReturn := fake.uploadBitsPackageReturnsOnCall[len(fake.uploadBitsPackageArgsForCall)]
	fake.uploadBitsPackageArgsForCall = append(fake.uploadBitsPackageArgsForCall, struct {
		arg1 v7action.Package
		arg2 []sharedaction.Resource
		arg3 io.Reader
		arg4 int64
	}{arg1, arg2Copy, arg3, arg4})
	fake.recordInvocation("UploadBitsPackage", []interface{}{arg1, arg2Copy, arg3, arg4})
	fake.uploadBitsPackageMutex.Unlock()
	if fake.UploadBitsPackageStub != nil {
		return fake.UploadBitsPackageStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.uploadBitsPackageReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeV7Actor) UploadBitsPackageCallCount() int {
	fake.uploadBitsPackageMutex.RLock()
	defer fake.uploadBitsPackageMutex.RUnlock()
	return len(fake.uploadBitsPackageArgsForCall)
}

func (fake *FakeV7Actor) UploadBitsPackageCalls(stub func(v7action.Package, []sharedaction.Resource, io.Reader, int64) (v7action.Package, v7action.Warnings, error)) {
	fake.uploadBitsPackageMutex.Lock()
	defer fake.uploadBitsPackageMutex.Unlock()
	fake.UploadBitsPackageStub = stub
}

func (fake *FakeV7Actor) UploadBitsPackageArgsForCall(i int) (v7action.Package, []sharedaction.Resource, io.Reader, int64) {
	fake.uploadBitsPackageMutex.RLock()
	defer fake.uploadBitsPackageMutex.RUnlock()
	argsForCall := fake.uploadBitsPackageArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeV7Actor) UploadBitsPackageReturns(result1 v7action.Package, result2 v7action.Warnings, result3 error) {
	fake.uploadBitsPackageMutex.Lock()
	defer fake.uploadBitsPackageMutex.Unlock()
	fake.UploadBitsPackageStub = nil
	fake.uploadBitsPackageReturns = struct {
		result1 v7action.Package
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV7Actor) UploadBitsPackageReturnsOnCall(i int, result1 v7action.Package, result2 v7action.Warnings, result3 error) {
	fake.uploadBitsPackageMutex.Lock()
	defer fake.uploadBitsPackageMutex.Unlock()
	fake.UploadBitsPackageStub = nil
	if fake.uploadBitsPackageReturnsOnCall == nil {
		fake.uploadBitsPackageReturnsOnCall = make(map[int]struct {
			result1 v7action.Package
			result2 v7action.Warnings
			result3 error
		})
	}
	fake.uploadBitsPackageReturnsOnCall[i] = struct {
		result1 v7action.Package
		result2 v7action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV7Actor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createApplicationInSpaceMutex.RLock()
	defer fake.createApplicationInSpaceMutex.RUnlock()
	fake.createBitsPackageByApplicationMutex.RLock()
	defer fake.createBitsPackageByApplicationMutex.RUnlock()
	fake.createDockerPackageByApplicationMutex.RLock()
	defer fake.createDockerPackageByApplicationMutex.RUnlock()
	fake.getApplicationByNameAndSpaceMutex.RLock()
	defer fake.getApplicationByNameAndSpaceMutex.RUnlock()
	fake.pollBuildMutex.RLock()
	defer fake.pollBuildMutex.RUnlock()
	fake.pollPackageMutex.RLock()
	defer fake.pollPackageMutex.RUnlock()
	fake.scaleProcessByApplicationMutex.RLock()
	defer fake.scaleProcessByApplicationMutex.RUnlock()
	fake.setApplicationDropletMutex.RLock()
	defer fake.setApplicationDropletMutex.RUnlock()
	fake.stageApplicationPackageMutex.RLock()
	defer fake.stageApplicationPackageMutex.RUnlock()
	fake.updateApplicationMutex.RLock()
	defer fake.updateApplicationMutex.RUnlock()
	fake.updateProcessByTypeAndApplicationMutex.RLock()
	defer fake.updateProcessByTypeAndApplicationMutex.RUnlock()
	fake.uploadBitsPackageMutex.RLock()
	defer fake.uploadBitsPackageMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeV7Actor) recordInvocation(key string, args []interface{}) {
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

var _ v7pushaction.V7Actor = new(FakeV7Actor)
