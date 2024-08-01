/*
Copyright the Velero contributors.

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
// Code generated by mockery v2.43.2. DO NOT EDIT.

package v1

import (
	mock "github.com/stretchr/testify/mock"
	runtime "k8s.io/apimachinery/pkg/runtime"

	velero "github.com/vmware-tanzu/velero/pkg/plugin/velero"

	velerov1 "github.com/vmware-tanzu/velero/pkg/apis/velero/v1"
)

// ItemBlockAction is an autogenerated mock type for the ItemBlockAction type
type ItemBlockAction struct {
	mock.Mock
}

// AppliesTo provides a mock function with given fields:
func (_m *ItemBlockAction) AppliesTo() (velero.ResourceSelector, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for AppliesTo")
	}

	var r0 velero.ResourceSelector
	var r1 error
	if rf, ok := ret.Get(0).(func() (velero.ResourceSelector, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() velero.ResourceSelector); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(velero.ResourceSelector)
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetRelatedItems provides a mock function with given fields: item, backup
func (_m *ItemBlockAction) GetRelatedItems(item runtime.Unstructured, backup *velerov1.Backup) ([]velero.ResourceIdentifier, error) {
	ret := _m.Called(item, backup)

	if len(ret) == 0 {
		panic("no return value specified for GetRelatedItems")
	}

	var r0 []velero.ResourceIdentifier
	var r1 error
	if rf, ok := ret.Get(0).(func(runtime.Unstructured, *velerov1.Backup) ([]velero.ResourceIdentifier, error)); ok {
		return rf(item, backup)
	}
	if rf, ok := ret.Get(0).(func(runtime.Unstructured, *velerov1.Backup) []velero.ResourceIdentifier); ok {
		r0 = rf(item, backup)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]velero.ResourceIdentifier)
		}
	}

	if rf, ok := ret.Get(1).(func(runtime.Unstructured, *velerov1.Backup) error); ok {
		r1 = rf(item, backup)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Name provides a mock function with given fields:
func (_m *ItemBlockAction) Name() string {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Name")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NewItemBlockAction creates a new instance of ItemBlockAction. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewItemBlockAction(t interface {
	mock.TestingT
	Cleanup(func())
}) *ItemBlockAction {
	mock := &ItemBlockAction{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
