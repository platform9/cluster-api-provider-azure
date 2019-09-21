/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-06-01/network/networkapi (interfaces: LoadBalancersClientAPI)

// Package mock_internalloadbalancers is a generated GoMock package.
package mock_internalloadbalancers

import (
	context "context"
	reflect "reflect"

	network "github.com/Azure/azure-sdk-for-go/services/network/mgmt/2019-06-01/network"
	gomock "github.com/golang/mock/gomock"
)

// MockLoadBalancersClientAPI is a mock of LoadBalancersClientAPI interface
type MockLoadBalancersClientAPI struct {
	ctrl     *gomock.Controller
	recorder *MockLoadBalancersClientAPIMockRecorder
}

// MockLoadBalancersClientAPIMockRecorder is the mock recorder for MockLoadBalancersClientAPI
type MockLoadBalancersClientAPIMockRecorder struct {
	mock *MockLoadBalancersClientAPI
}

// NewMockLoadBalancersClientAPI creates a new mock instance
func NewMockLoadBalancersClientAPI(ctrl *gomock.Controller) *MockLoadBalancersClientAPI {
	mock := &MockLoadBalancersClientAPI{ctrl: ctrl}
	mock.recorder = &MockLoadBalancersClientAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLoadBalancersClientAPI) EXPECT() *MockLoadBalancersClientAPIMockRecorder {
	return m.recorder
}

// CreateOrUpdate mocks base method
func (m *MockLoadBalancersClientAPI) CreateOrUpdate(arg0 context.Context, arg1, arg2 string, arg3 network.LoadBalancer) (network.LoadBalancersCreateOrUpdateFuture, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrUpdate", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(network.LoadBalancersCreateOrUpdateFuture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrUpdate indicates an expected call of CreateOrUpdate
func (mr *MockLoadBalancersClientAPIMockRecorder) CreateOrUpdate(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrUpdate", reflect.TypeOf((*MockLoadBalancersClientAPI)(nil).CreateOrUpdate), arg0, arg1, arg2, arg3)
}

// Delete mocks base method
func (m *MockLoadBalancersClientAPI) Delete(arg0 context.Context, arg1, arg2 string) (network.LoadBalancersDeleteFuture, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0, arg1, arg2)
	ret0, _ := ret[0].(network.LoadBalancersDeleteFuture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Delete indicates an expected call of Delete
func (mr *MockLoadBalancersClientAPIMockRecorder) Delete(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockLoadBalancersClientAPI)(nil).Delete), arg0, arg1, arg2)
}

// Get mocks base method
func (m *MockLoadBalancersClientAPI) Get(arg0 context.Context, arg1, arg2, arg3 string) (network.LoadBalancer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(network.LoadBalancer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockLoadBalancersClientAPIMockRecorder) Get(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockLoadBalancersClientAPI)(nil).Get), arg0, arg1, arg2, arg3)
}

// List mocks base method
func (m *MockLoadBalancersClientAPI) List(arg0 context.Context, arg1 string) (network.LoadBalancerListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(network.LoadBalancerListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockLoadBalancersClientAPIMockRecorder) List(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockLoadBalancersClientAPI)(nil).List), arg0, arg1)
}

// ListAll mocks base method
func (m *MockLoadBalancersClientAPI) ListAll(arg0 context.Context) (network.LoadBalancerListResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAll", arg0)
	ret0, _ := ret[0].(network.LoadBalancerListResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAll indicates an expected call of ListAll
func (mr *MockLoadBalancersClientAPIMockRecorder) ListAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAll", reflect.TypeOf((*MockLoadBalancersClientAPI)(nil).ListAll), arg0)
}

// UpdateTags mocks base method
func (m *MockLoadBalancersClientAPI) UpdateTags(arg0 context.Context, arg1, arg2 string, arg3 network.TagsObject) (network.LoadBalancersUpdateTagsFuture, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTags", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(network.LoadBalancersUpdateTagsFuture)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTags indicates an expected call of UpdateTags
func (mr *MockLoadBalancersClientAPIMockRecorder) UpdateTags(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTags", reflect.TypeOf((*MockLoadBalancersClientAPI)(nil).UpdateTags), arg0, arg1, arg2, arg3)
}
