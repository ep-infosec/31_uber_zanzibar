// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/uber/zanzibar/examples/example-gateway/build/clients/contacts (interfaces: Client)

// Package clientmock is a generated GoMock package.
package clientmock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	contacts "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients-idl/clients/contacts/contacts"
	runtime "github.com/uber/zanzibar/runtime"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// HTTPClient mocks base method.
func (m *MockClient) HTTPClient() *runtime.HTTPClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HTTPClient")
	ret0, _ := ret[0].(*runtime.HTTPClient)
	return ret0
}

// HTTPClient indicates an expected call of HTTPClient.
func (mr *MockClientMockRecorder) HTTPClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HTTPClient", reflect.TypeOf((*MockClient)(nil).HTTPClient))
}

// SaveContacts mocks base method.
func (m *MockClient) SaveContacts(arg0 context.Context, arg1 map[string]string, arg2 *contacts.Contacts_SaveContacts_Args, arg3 *runtime.TimeoutAndRetryOptions) (context.Context, *contacts.SaveContactsResponse, map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveContacts", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].(*contacts.SaveContactsResponse)
	ret2, _ := ret[2].(map[string]string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// SaveContacts indicates an expected call of SaveContacts.
func (mr *MockClientMockRecorder) SaveContacts(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveContacts", reflect.TypeOf((*MockClient)(nil).SaveContacts), arg0, arg1, arg2, arg3)
}

// TestURLURL mocks base method.
func (m *MockClient) TestURLURL(arg0 context.Context, arg1 map[string]string, arg2 *runtime.TimeoutAndRetryOptions) (context.Context, string, map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TestURLURL", arg0, arg1, arg2)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(map[string]string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// TestURLURL indicates an expected call of TestURLURL.
func (mr *MockClientMockRecorder) TestURLURL(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TestURLURL", reflect.TypeOf((*MockClient)(nil).TestURLURL), arg0, arg1, arg2)
}
