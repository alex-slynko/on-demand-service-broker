// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/pivotal-cf/brokerapi"
	"github.com/pivotal-cf/on-demand-service-broker/broker"
	"github.com/pivotal-cf/on-demand-service-broker/brokerclient/broker_response"
	"github.com/pivotal-cf/on-demand-service-broker/upgrader"
)

type FakeBrokerServices struct {
	InstancesStub        func() ([]string, error)
	instancesMutex       sync.RWMutex
	instancesArgsForCall []struct{}
	instancesReturns     struct {
		result1 []string
		result2 error
	}
	instancesReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	UpgradeInstanceStub        func(instance string) (broker_response.UpgradeOperation, error)
	upgradeInstanceMutex       sync.RWMutex
	upgradeInstanceArgsForCall []struct {
		instance string
	}
	upgradeInstanceReturns struct {
		result1 broker_response.UpgradeOperation
		result2 error
	}
	upgradeInstanceReturnsOnCall map[int]struct {
		result1 broker_response.UpgradeOperation
		result2 error
	}
	LastOperationStub        func(instance string, operationData broker.OperationData) (brokerapi.LastOperation, error)
	lastOperationMutex       sync.RWMutex
	lastOperationArgsForCall []struct {
		instance      string
		operationData broker.OperationData
	}
	lastOperationReturns struct {
		result1 brokerapi.LastOperation
		result2 error
	}
	lastOperationReturnsOnCall map[int]struct {
		result1 brokerapi.LastOperation
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBrokerServices) Instances() ([]string, error) {
	fake.instancesMutex.Lock()
	ret, specificReturn := fake.instancesReturnsOnCall[len(fake.instancesArgsForCall)]
	fake.instancesArgsForCall = append(fake.instancesArgsForCall, struct{}{})
	fake.recordInvocation("Instances", []interface{}{})
	fake.instancesMutex.Unlock()
	if fake.InstancesStub != nil {
		return fake.InstancesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.instancesReturns.result1, fake.instancesReturns.result2
}

func (fake *FakeBrokerServices) InstancesCallCount() int {
	fake.instancesMutex.RLock()
	defer fake.instancesMutex.RUnlock()
	return len(fake.instancesArgsForCall)
}

func (fake *FakeBrokerServices) InstancesReturns(result1 []string, result2 error) {
	fake.InstancesStub = nil
	fake.instancesReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeBrokerServices) InstancesReturnsOnCall(i int, result1 []string, result2 error) {
	fake.InstancesStub = nil
	if fake.instancesReturnsOnCall == nil {
		fake.instancesReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.instancesReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeBrokerServices) UpgradeInstance(instance string) (broker_response.UpgradeOperation, error) {
	fake.upgradeInstanceMutex.Lock()
	ret, specificReturn := fake.upgradeInstanceReturnsOnCall[len(fake.upgradeInstanceArgsForCall)]
	fake.upgradeInstanceArgsForCall = append(fake.upgradeInstanceArgsForCall, struct {
		instance string
	}{instance})
	fake.recordInvocation("UpgradeInstance", []interface{}{instance})
	fake.upgradeInstanceMutex.Unlock()
	if fake.UpgradeInstanceStub != nil {
		return fake.UpgradeInstanceStub(instance)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.upgradeInstanceReturns.result1, fake.upgradeInstanceReturns.result2
}

func (fake *FakeBrokerServices) UpgradeInstanceCallCount() int {
	fake.upgradeInstanceMutex.RLock()
	defer fake.upgradeInstanceMutex.RUnlock()
	return len(fake.upgradeInstanceArgsForCall)
}

func (fake *FakeBrokerServices) UpgradeInstanceArgsForCall(i int) string {
	fake.upgradeInstanceMutex.RLock()
	defer fake.upgradeInstanceMutex.RUnlock()
	return fake.upgradeInstanceArgsForCall[i].instance
}

func (fake *FakeBrokerServices) UpgradeInstanceReturns(result1 broker_response.UpgradeOperation, result2 error) {
	fake.UpgradeInstanceStub = nil
	fake.upgradeInstanceReturns = struct {
		result1 broker_response.UpgradeOperation
		result2 error
	}{result1, result2}
}

func (fake *FakeBrokerServices) UpgradeInstanceReturnsOnCall(i int, result1 broker_response.UpgradeOperation, result2 error) {
	fake.UpgradeInstanceStub = nil
	if fake.upgradeInstanceReturnsOnCall == nil {
		fake.upgradeInstanceReturnsOnCall = make(map[int]struct {
			result1 broker_response.UpgradeOperation
			result2 error
		})
	}
	fake.upgradeInstanceReturnsOnCall[i] = struct {
		result1 broker_response.UpgradeOperation
		result2 error
	}{result1, result2}
}

func (fake *FakeBrokerServices) LastOperation(instance string, operationData broker.OperationData) (brokerapi.LastOperation, error) {
	fake.lastOperationMutex.Lock()
	ret, specificReturn := fake.lastOperationReturnsOnCall[len(fake.lastOperationArgsForCall)]
	fake.lastOperationArgsForCall = append(fake.lastOperationArgsForCall, struct {
		instance      string
		operationData broker.OperationData
	}{instance, operationData})
	fake.recordInvocation("LastOperation", []interface{}{instance, operationData})
	fake.lastOperationMutex.Unlock()
	if fake.LastOperationStub != nil {
		return fake.LastOperationStub(instance, operationData)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.lastOperationReturns.result1, fake.lastOperationReturns.result2
}

func (fake *FakeBrokerServices) LastOperationCallCount() int {
	fake.lastOperationMutex.RLock()
	defer fake.lastOperationMutex.RUnlock()
	return len(fake.lastOperationArgsForCall)
}

func (fake *FakeBrokerServices) LastOperationArgsForCall(i int) (string, broker.OperationData) {
	fake.lastOperationMutex.RLock()
	defer fake.lastOperationMutex.RUnlock()
	return fake.lastOperationArgsForCall[i].instance, fake.lastOperationArgsForCall[i].operationData
}

func (fake *FakeBrokerServices) LastOperationReturns(result1 brokerapi.LastOperation, result2 error) {
	fake.LastOperationStub = nil
	fake.lastOperationReturns = struct {
		result1 brokerapi.LastOperation
		result2 error
	}{result1, result2}
}

func (fake *FakeBrokerServices) LastOperationReturnsOnCall(i int, result1 brokerapi.LastOperation, result2 error) {
	fake.LastOperationStub = nil
	if fake.lastOperationReturnsOnCall == nil {
		fake.lastOperationReturnsOnCall = make(map[int]struct {
			result1 brokerapi.LastOperation
			result2 error
		})
	}
	fake.lastOperationReturnsOnCall[i] = struct {
		result1 brokerapi.LastOperation
		result2 error
	}{result1, result2}
}

func (fake *FakeBrokerServices) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.instancesMutex.RLock()
	defer fake.instancesMutex.RUnlock()
	fake.upgradeInstanceMutex.RLock()
	defer fake.upgradeInstanceMutex.RUnlock()
	fake.lastOperationMutex.RLock()
	defer fake.lastOperationMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeBrokerServices) recordInvocation(key string, args []interface{}) {
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

var _ upgrader.BrokerServices = new(FakeBrokerServices)
