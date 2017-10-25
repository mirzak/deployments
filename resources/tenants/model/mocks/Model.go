// Copyright 2017 Northern.tech AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.
package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"

// Model is an autogenerated mock type for the Model type
type Model struct {
	mock.Mock
}

// ProvisionTenant provides a mock function with given fields: ctx, tenant_id
func (_m *Model) ProvisionTenant(ctx context.Context, tenant_id string) error {
	ret := _m.Called(ctx, tenant_id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, tenant_id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
