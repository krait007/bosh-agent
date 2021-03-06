// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-agent/platform/openiscsi"
)

type FakeOpenIscsi struct {
	SetupStub        func(iqn, username, password string) (err error)
	setupMutex       sync.RWMutex
	setupArgsForCall []struct {
		iqn      string
		username string
		password string
	}
	setupReturns struct {
		result1 error
	}
	setupReturnsOnCall map[int]struct {
		result1 error
	}
	StartStub        func() (err error)
	startMutex       sync.RWMutex
	startArgsForCall []struct{}
	startReturns     struct {
		result1 error
	}
	startReturnsOnCall map[int]struct {
		result1 error
	}
	StopStub        func() (err error)
	stopMutex       sync.RWMutex
	stopArgsForCall []struct{}
	stopReturns     struct {
		result1 error
	}
	stopReturnsOnCall map[int]struct {
		result1 error
	}
	RestartStub        func() (err error)
	restartMutex       sync.RWMutex
	restartArgsForCall []struct{}
	restartReturns     struct {
		result1 error
	}
	restartReturnsOnCall map[int]struct {
		result1 error
	}
	DiscoveryStub        func(ipAddress string) (err error)
	discoveryMutex       sync.RWMutex
	discoveryArgsForCall []struct {
		ipAddress string
	}
	discoveryReturns struct {
		result1 error
	}
	discoveryReturnsOnCall map[int]struct {
		result1 error
	}
	LoginStub        func() (err error)
	loginMutex       sync.RWMutex
	loginArgsForCall []struct{}
	loginReturns     struct {
		result1 error
	}
	loginReturnsOnCall map[int]struct {
		result1 error
	}
	LogoutStub        func() (err error)
	logoutMutex       sync.RWMutex
	logoutArgsForCall []struct{}
	logoutReturns     struct {
		result1 error
	}
	logoutReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeOpenIscsi) Setup(iqn string, username string, password string) (err error) {
	fake.setupMutex.Lock()
	ret, specificReturn := fake.setupReturnsOnCall[len(fake.setupArgsForCall)]
	fake.setupArgsForCall = append(fake.setupArgsForCall, struct {
		iqn      string
		username string
		password string
	}{iqn, username, password})
	fake.recordInvocation("Setup", []interface{}{iqn, username, password})
	fake.setupMutex.Unlock()
	if fake.SetupStub != nil {
		return fake.SetupStub(iqn, username, password)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.setupReturns.result1
}

func (fake *FakeOpenIscsi) SetupCallCount() int {
	fake.setupMutex.RLock()
	defer fake.setupMutex.RUnlock()
	return len(fake.setupArgsForCall)
}

func (fake *FakeOpenIscsi) SetupArgsForCall(i int) (string, string, string) {
	fake.setupMutex.RLock()
	defer fake.setupMutex.RUnlock()
	return fake.setupArgsForCall[i].iqn, fake.setupArgsForCall[i].username, fake.setupArgsForCall[i].password
}

func (fake *FakeOpenIscsi) SetupReturns(result1 error) {
	fake.SetupStub = nil
	fake.setupReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOpenIscsi) SetupReturnsOnCall(i int, result1 error) {
	fake.SetupStub = nil
	if fake.setupReturnsOnCall == nil {
		fake.setupReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.setupReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeOpenIscsi) Start() (err error) {
	fake.startMutex.Lock()
	ret, specificReturn := fake.startReturnsOnCall[len(fake.startArgsForCall)]
	fake.startArgsForCall = append(fake.startArgsForCall, struct{}{})
	fake.recordInvocation("Start", []interface{}{})
	fake.startMutex.Unlock()
	if fake.StartStub != nil {
		return fake.StartStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.startReturns.result1
}

func (fake *FakeOpenIscsi) StartCallCount() int {
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	return len(fake.startArgsForCall)
}

func (fake *FakeOpenIscsi) StartReturns(result1 error) {
	fake.StartStub = nil
	fake.startReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOpenIscsi) StartReturnsOnCall(i int, result1 error) {
	fake.StartStub = nil
	if fake.startReturnsOnCall == nil {
		fake.startReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.startReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeOpenIscsi) Stop() (err error) {
	fake.stopMutex.Lock()
	ret, specificReturn := fake.stopReturnsOnCall[len(fake.stopArgsForCall)]
	fake.stopArgsForCall = append(fake.stopArgsForCall, struct{}{})
	fake.recordInvocation("Stop", []interface{}{})
	fake.stopMutex.Unlock()
	if fake.StopStub != nil {
		return fake.StopStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.stopReturns.result1
}

func (fake *FakeOpenIscsi) StopCallCount() int {
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	return len(fake.stopArgsForCall)
}

func (fake *FakeOpenIscsi) StopReturns(result1 error) {
	fake.StopStub = nil
	fake.stopReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOpenIscsi) StopReturnsOnCall(i int, result1 error) {
	fake.StopStub = nil
	if fake.stopReturnsOnCall == nil {
		fake.stopReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.stopReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeOpenIscsi) Restart() (err error) {
	fake.restartMutex.Lock()
	ret, specificReturn := fake.restartReturnsOnCall[len(fake.restartArgsForCall)]
	fake.restartArgsForCall = append(fake.restartArgsForCall, struct{}{})
	fake.recordInvocation("Restart", []interface{}{})
	fake.restartMutex.Unlock()
	if fake.RestartStub != nil {
		return fake.RestartStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.restartReturns.result1
}

func (fake *FakeOpenIscsi) RestartCallCount() int {
	fake.restartMutex.RLock()
	defer fake.restartMutex.RUnlock()
	return len(fake.restartArgsForCall)
}

func (fake *FakeOpenIscsi) RestartReturns(result1 error) {
	fake.RestartStub = nil
	fake.restartReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOpenIscsi) RestartReturnsOnCall(i int, result1 error) {
	fake.RestartStub = nil
	if fake.restartReturnsOnCall == nil {
		fake.restartReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.restartReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeOpenIscsi) Discovery(ipAddress string) (err error) {
	fake.discoveryMutex.Lock()
	ret, specificReturn := fake.discoveryReturnsOnCall[len(fake.discoveryArgsForCall)]
	fake.discoveryArgsForCall = append(fake.discoveryArgsForCall, struct {
		ipAddress string
	}{ipAddress})
	fake.recordInvocation("Discovery", []interface{}{ipAddress})
	fake.discoveryMutex.Unlock()
	if fake.DiscoveryStub != nil {
		return fake.DiscoveryStub(ipAddress)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.discoveryReturns.result1
}

func (fake *FakeOpenIscsi) DiscoveryCallCount() int {
	fake.discoveryMutex.RLock()
	defer fake.discoveryMutex.RUnlock()
	return len(fake.discoveryArgsForCall)
}

func (fake *FakeOpenIscsi) DiscoveryArgsForCall(i int) string {
	fake.discoveryMutex.RLock()
	defer fake.discoveryMutex.RUnlock()
	return fake.discoveryArgsForCall[i].ipAddress
}

func (fake *FakeOpenIscsi) DiscoveryReturns(result1 error) {
	fake.DiscoveryStub = nil
	fake.discoveryReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOpenIscsi) DiscoveryReturnsOnCall(i int, result1 error) {
	fake.DiscoveryStub = nil
	if fake.discoveryReturnsOnCall == nil {
		fake.discoveryReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.discoveryReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeOpenIscsi) Login() (err error) {
	fake.loginMutex.Lock()
	ret, specificReturn := fake.loginReturnsOnCall[len(fake.loginArgsForCall)]
	fake.loginArgsForCall = append(fake.loginArgsForCall, struct{}{})
	fake.recordInvocation("Login", []interface{}{})
	fake.loginMutex.Unlock()
	if fake.LoginStub != nil {
		return fake.LoginStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.loginReturns.result1
}

func (fake *FakeOpenIscsi) LoginCallCount() int {
	fake.loginMutex.RLock()
	defer fake.loginMutex.RUnlock()
	return len(fake.loginArgsForCall)
}

func (fake *FakeOpenIscsi) LoginReturns(result1 error) {
	fake.LoginStub = nil
	fake.loginReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOpenIscsi) LoginReturnsOnCall(i int, result1 error) {
	fake.LoginStub = nil
	if fake.loginReturnsOnCall == nil {
		fake.loginReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.loginReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeOpenIscsi) Logout() (err error) {
	fake.logoutMutex.Lock()
	ret, specificReturn := fake.logoutReturnsOnCall[len(fake.logoutArgsForCall)]
	fake.logoutArgsForCall = append(fake.logoutArgsForCall, struct{}{})
	fake.recordInvocation("Logout", []interface{}{})
	fake.logoutMutex.Unlock()
	if fake.LogoutStub != nil {
		return fake.LogoutStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.logoutReturns.result1
}

func (fake *FakeOpenIscsi) LogoutCallCount() int {
	fake.logoutMutex.RLock()
	defer fake.logoutMutex.RUnlock()
	return len(fake.logoutArgsForCall)
}

func (fake *FakeOpenIscsi) LogoutReturns(result1 error) {
	fake.LogoutStub = nil
	fake.logoutReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeOpenIscsi) LogoutReturnsOnCall(i int, result1 error) {
	fake.LogoutStub = nil
	if fake.logoutReturnsOnCall == nil {
		fake.logoutReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.logoutReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeOpenIscsi) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.setupMutex.RLock()
	defer fake.setupMutex.RUnlock()
	fake.startMutex.RLock()
	defer fake.startMutex.RUnlock()
	fake.stopMutex.RLock()
	defer fake.stopMutex.RUnlock()
	fake.restartMutex.RLock()
	defer fake.restartMutex.RUnlock()
	fake.discoveryMutex.RLock()
	defer fake.discoveryMutex.RUnlock()
	fake.loginMutex.RLock()
	defer fake.loginMutex.RUnlock()
	fake.logoutMutex.RLock()
	defer fake.logoutMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeOpenIscsi) recordInvocation(key string, args []interface{}) {
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

var _ openiscsi.OpenIscsi = new(FakeOpenIscsi)
