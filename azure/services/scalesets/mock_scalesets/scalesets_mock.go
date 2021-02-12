/*
Copyright The Kubernetes Authors.

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
// Source: ../scalesets.go

// Package mock_scalesets is a generated GoMock package.
package mock_scalesets

import (
	context "context"
	autorest "github.com/Azure/go-autorest/autorest"
	logr "github.com/go-logr/logr"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	v1alpha3 "sigs.k8s.io/cluster-api-provider-azure/api/v1alpha3"
	azure "sigs.k8s.io/cluster-api-provider-azure/azure"
	v1alpha30 "sigs.k8s.io/cluster-api-provider-azure/exp/api/v1alpha3"
)

// MockScaleSetScope is a mock of ScaleSetScope interface.
type MockScaleSetScope struct {
	ctrl     *gomock.Controller
	recorder *MockScaleSetScopeMockRecorder
}

// MockScaleSetScopeMockRecorder is the mock recorder for MockScaleSetScope.
type MockScaleSetScopeMockRecorder struct {
	mock *MockScaleSetScope
}

// NewMockScaleSetScope creates a new mock instance.
func NewMockScaleSetScope(ctrl *gomock.Controller) *MockScaleSetScope {
	mock := &MockScaleSetScope{ctrl: ctrl}
	mock.recorder = &MockScaleSetScopeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScaleSetScope) EXPECT() *MockScaleSetScopeMockRecorder {
	return m.recorder
}

// Info mocks base method.
func (m *MockScaleSetScope) Info(msg string, keysAndValues ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{msg}
	for _, a := range keysAndValues {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Info", varargs...)
}

// Info indicates an expected call of Info.
func (mr *MockScaleSetScopeMockRecorder) Info(msg interface{}, keysAndValues ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{msg}, keysAndValues...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockScaleSetScope)(nil).Info), varargs...)
}

// Enabled mocks base method.
func (m *MockScaleSetScope) Enabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Enabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Enabled indicates an expected call of Enabled.
func (mr *MockScaleSetScopeMockRecorder) Enabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enabled", reflect.TypeOf((*MockScaleSetScope)(nil).Enabled))
}

// Error mocks base method.
func (m *MockScaleSetScope) Error(err error, msg string, keysAndValues ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{err, msg}
	for _, a := range keysAndValues {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Error", varargs...)
}

// Error indicates an expected call of Error.
func (mr *MockScaleSetScopeMockRecorder) Error(err, msg interface{}, keysAndValues ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{err, msg}, keysAndValues...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Error", reflect.TypeOf((*MockScaleSetScope)(nil).Error), varargs...)
}

// V mocks base method.
func (m *MockScaleSetScope) V(level int) logr.InfoLogger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "V", level)
	ret0, _ := ret[0].(logr.InfoLogger)
	return ret0
}

// V indicates an expected call of V.
func (mr *MockScaleSetScopeMockRecorder) V(level interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "V", reflect.TypeOf((*MockScaleSetScope)(nil).V), level)
}

// WithValues mocks base method.
func (m *MockScaleSetScope) WithValues(keysAndValues ...interface{}) logr.Logger {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range keysAndValues {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "WithValues", varargs...)
	ret0, _ := ret[0].(logr.Logger)
	return ret0
}

// WithValues indicates an expected call of WithValues.
func (mr *MockScaleSetScopeMockRecorder) WithValues(keysAndValues ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithValues", reflect.TypeOf((*MockScaleSetScope)(nil).WithValues), keysAndValues...)
}

// WithName mocks base method.
func (m *MockScaleSetScope) WithName(name string) logr.Logger {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithName", name)
	ret0, _ := ret[0].(logr.Logger)
	return ret0
}

// WithName indicates an expected call of WithName.
func (mr *MockScaleSetScopeMockRecorder) WithName(name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithName", reflect.TypeOf((*MockScaleSetScope)(nil).WithName), name)
}

// SubscriptionID mocks base method.
func (m *MockScaleSetScope) SubscriptionID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscriptionID")
	ret0, _ := ret[0].(string)
	return ret0
}

// SubscriptionID indicates an expected call of SubscriptionID.
func (mr *MockScaleSetScopeMockRecorder) SubscriptionID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscriptionID", reflect.TypeOf((*MockScaleSetScope)(nil).SubscriptionID))
}

// ClientID mocks base method.
func (m *MockScaleSetScope) ClientID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClientID indicates an expected call of ClientID.
func (mr *MockScaleSetScopeMockRecorder) ClientID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientID", reflect.TypeOf((*MockScaleSetScope)(nil).ClientID))
}

// ClientSecret mocks base method.
func (m *MockScaleSetScope) ClientSecret() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientSecret")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClientSecret indicates an expected call of ClientSecret.
func (mr *MockScaleSetScopeMockRecorder) ClientSecret() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientSecret", reflect.TypeOf((*MockScaleSetScope)(nil).ClientSecret))
}

// CloudEnvironment mocks base method.
func (m *MockScaleSetScope) CloudEnvironment() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudEnvironment")
	ret0, _ := ret[0].(string)
	return ret0
}

// CloudEnvironment indicates an expected call of CloudEnvironment.
func (mr *MockScaleSetScopeMockRecorder) CloudEnvironment() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudEnvironment", reflect.TypeOf((*MockScaleSetScope)(nil).CloudEnvironment))
}

// TenantID mocks base method.
func (m *MockScaleSetScope) TenantID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TenantID")
	ret0, _ := ret[0].(string)
	return ret0
}

// TenantID indicates an expected call of TenantID.
func (mr *MockScaleSetScopeMockRecorder) TenantID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TenantID", reflect.TypeOf((*MockScaleSetScope)(nil).TenantID))
}

// BaseURI mocks base method.
func (m *MockScaleSetScope) BaseURI() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BaseURI")
	ret0, _ := ret[0].(string)
	return ret0
}

// BaseURI indicates an expected call of BaseURI.
func (mr *MockScaleSetScopeMockRecorder) BaseURI() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BaseURI", reflect.TypeOf((*MockScaleSetScope)(nil).BaseURI))
}

// Authorizer mocks base method.
func (m *MockScaleSetScope) Authorizer() autorest.Authorizer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authorizer")
	ret0, _ := ret[0].(autorest.Authorizer)
	return ret0
}

// Authorizer indicates an expected call of Authorizer.
func (mr *MockScaleSetScopeMockRecorder) Authorizer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authorizer", reflect.TypeOf((*MockScaleSetScope)(nil).Authorizer))
}

// HashKey mocks base method.
func (m *MockScaleSetScope) HashKey() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HashKey")
	ret0, _ := ret[0].(string)
	return ret0
}

// HashKey indicates an expected call of HashKey.
func (mr *MockScaleSetScopeMockRecorder) HashKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashKey", reflect.TypeOf((*MockScaleSetScope)(nil).HashKey))
}

// ResourceGroup mocks base method.
func (m *MockScaleSetScope) ResourceGroup() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResourceGroup")
	ret0, _ := ret[0].(string)
	return ret0
}

// ResourceGroup indicates an expected call of ResourceGroup.
func (mr *MockScaleSetScopeMockRecorder) ResourceGroup() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResourceGroup", reflect.TypeOf((*MockScaleSetScope)(nil).ResourceGroup))
}

// ClusterName mocks base method.
func (m *MockScaleSetScope) ClusterName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterName")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClusterName indicates an expected call of ClusterName.
func (mr *MockScaleSetScopeMockRecorder) ClusterName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterName", reflect.TypeOf((*MockScaleSetScope)(nil).ClusterName))
}

// Location mocks base method.
func (m *MockScaleSetScope) Location() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Location")
	ret0, _ := ret[0].(string)
	return ret0
}

// Location indicates an expected call of Location.
func (mr *MockScaleSetScopeMockRecorder) Location() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Location", reflect.TypeOf((*MockScaleSetScope)(nil).Location))
}

// AdditionalTags mocks base method.
func (m *MockScaleSetScope) AdditionalTags() v1alpha3.Tags {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdditionalTags")
	ret0, _ := ret[0].(v1alpha3.Tags)
	return ret0
}

// AdditionalTags indicates an expected call of AdditionalTags.
func (mr *MockScaleSetScopeMockRecorder) AdditionalTags() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdditionalTags", reflect.TypeOf((*MockScaleSetScope)(nil).AdditionalTags))
}

// AvailabilitySetEnabled mocks base method.
func (m *MockScaleSetScope) AvailabilitySetEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AvailabilitySetEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// AvailabilitySetEnabled indicates an expected call of AvailabilitySetEnabled.
func (mr *MockScaleSetScopeMockRecorder) AvailabilitySetEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AvailabilitySetEnabled", reflect.TypeOf((*MockScaleSetScope)(nil).AvailabilitySetEnabled))
}

// GetBootstrapData mocks base method.
func (m *MockScaleSetScope) GetBootstrapData(ctx context.Context) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBootstrapData", ctx)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBootstrapData indicates an expected call of GetBootstrapData.
func (mr *MockScaleSetScopeMockRecorder) GetBootstrapData(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBootstrapData", reflect.TypeOf((*MockScaleSetScope)(nil).GetBootstrapData), ctx)
}

// GetLongRunningOperationState mocks base method.
func (m *MockScaleSetScope) GetLongRunningOperationState() *v1alpha3.Future {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLongRunningOperationState")
	ret0, _ := ret[0].(*v1alpha3.Future)
	return ret0
}

// GetLongRunningOperationState indicates an expected call of GetLongRunningOperationState.
func (mr *MockScaleSetScopeMockRecorder) GetLongRunningOperationState() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLongRunningOperationState", reflect.TypeOf((*MockScaleSetScope)(nil).GetLongRunningOperationState))
}

// GetVMImage mocks base method.
func (m *MockScaleSetScope) GetVMImage() (*v1alpha3.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVMImage")
	ret0, _ := ret[0].(*v1alpha3.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetVMImage indicates an expected call of GetVMImage.
func (mr *MockScaleSetScopeMockRecorder) GetVMImage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVMImage", reflect.TypeOf((*MockScaleSetScope)(nil).GetVMImage))
}

// MaxSurge mocks base method.
func (m *MockScaleSetScope) MaxSurge() int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaxSurge")
	ret0, _ := ret[0].(int32)
	return ret0
}

// MaxSurge indicates an expected call of MaxSurge.
func (mr *MockScaleSetScopeMockRecorder) MaxSurge() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaxSurge", reflect.TypeOf((*MockScaleSetScope)(nil).MaxSurge))
}

// MaxUnavailable mocks base method.
func (m *MockScaleSetScope) MaxUnavailable() int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MaxUnavailable")
	ret0, _ := ret[0].(int32)
	return ret0
}

// MaxUnavailable indicates an expected call of MaxUnavailable.
func (mr *MockScaleSetScopeMockRecorder) MaxUnavailable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MaxUnavailable", reflect.TypeOf((*MockScaleSetScope)(nil).MaxUnavailable))
}

// ScaleSetSpec mocks base method.
func (m *MockScaleSetScope) ScaleSetSpec() azure.ScaleSetSpec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScaleSetSpec")
	ret0, _ := ret[0].(azure.ScaleSetSpec)
	return ret0
}

// ScaleSetSpec indicates an expected call of ScaleSetSpec.
func (mr *MockScaleSetScopeMockRecorder) ScaleSetSpec() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScaleSetSpec", reflect.TypeOf((*MockScaleSetScope)(nil).ScaleSetSpec))
}

// VMSSExtensionSpecs mocks base method.
func (m *MockScaleSetScope) VMSSExtensionSpecs() []azure.VMSSExtensionSpec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VMSSExtensionSpecs")
	ret0, _ := ret[0].([]azure.VMSSExtensionSpec)
	return ret0
}

// VMSSExtensionSpecs indicates an expected call of VMSSExtensionSpecs.
func (mr *MockScaleSetScopeMockRecorder) VMSSExtensionSpecs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VMSSExtensionSpecs", reflect.TypeOf((*MockScaleSetScope)(nil).VMSSExtensionSpecs))
}

// SetAnnotation mocks base method.
func (m *MockScaleSetScope) SetAnnotation(arg0, arg1 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetAnnotation", arg0, arg1)
}

// SetAnnotation indicates an expected call of SetAnnotation.
func (mr *MockScaleSetScopeMockRecorder) SetAnnotation(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetAnnotation", reflect.TypeOf((*MockScaleSetScope)(nil).SetAnnotation), arg0, arg1)
}

// SetLongRunningOperationState mocks base method.
func (m *MockScaleSetScope) SetLongRunningOperationState(arg0 *v1alpha3.Future) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLongRunningOperationState", arg0)
}

// SetLongRunningOperationState indicates an expected call of SetLongRunningOperationState.
func (mr *MockScaleSetScopeMockRecorder) SetLongRunningOperationState(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLongRunningOperationState", reflect.TypeOf((*MockScaleSetScope)(nil).SetLongRunningOperationState), arg0)
}

// SetProviderID mocks base method.
func (m *MockScaleSetScope) SetProviderID(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetProviderID", arg0)
}

// SetProviderID indicates an expected call of SetProviderID.
func (mr *MockScaleSetScopeMockRecorder) SetProviderID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetProviderID", reflect.TypeOf((*MockScaleSetScope)(nil).SetProviderID), arg0)
}

// SetVMSSState mocks base method.
func (m *MockScaleSetScope) SetVMSSState(arg0 *v1alpha30.VMSS) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetVMSSState", arg0)
}

// SetVMSSState indicates an expected call of SetVMSSState.
func (mr *MockScaleSetScopeMockRecorder) SetVMSSState(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetVMSSState", reflect.TypeOf((*MockScaleSetScope)(nil).SetVMSSState), arg0)
}
