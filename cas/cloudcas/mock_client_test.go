// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/adtsign/certificates/cas/cloudcas (interfaces: CertificateAuthorityClient)

// Package cloudcas is a generated GoMock package.
package cloudcas

import (
	context "context"
	reflect "reflect"

	privateca "cloud.google.com/go/security/privateca/apiv1"
	privatecapb "cloud.google.com/go/security/privateca/apiv1/privatecapb"
	gomock "github.com/golang/mock/gomock"
	gax "github.com/googleapis/gax-go/v2"
)

// MockCertificateAuthorityClient is a mock of CertificateAuthorityClient interface.
type MockCertificateAuthorityClient struct {
	ctrl     *gomock.Controller
	recorder *MockCertificateAuthorityClientMockRecorder
}

// MockCertificateAuthorityClientMockRecorder is the mock recorder for MockCertificateAuthorityClient.
type MockCertificateAuthorityClientMockRecorder struct {
	mock *MockCertificateAuthorityClient
}

// NewMockCertificateAuthorityClient creates a new mock instance.
func NewMockCertificateAuthorityClient(ctrl *gomock.Controller) *MockCertificateAuthorityClient {
	mock := &MockCertificateAuthorityClient{ctrl: ctrl}
	mock.recorder = &MockCertificateAuthorityClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCertificateAuthorityClient) EXPECT() *MockCertificateAuthorityClientMockRecorder {
	return m.recorder
}

// ActivateCertificateAuthority mocks base method.
func (m *MockCertificateAuthorityClient) ActivateCertificateAuthority(arg0 context.Context, arg1 *privatecapb.ActivateCertificateAuthorityRequest, arg2 ...gax.CallOption) (*privateca.ActivateCertificateAuthorityOperation, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ActivateCertificateAuthority", varargs...)
	ret0, _ := ret[0].(*privateca.ActivateCertificateAuthorityOperation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ActivateCertificateAuthority indicates an expected call of ActivateCertificateAuthority.
func (mr *MockCertificateAuthorityClientMockRecorder) ActivateCertificateAuthority(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ActivateCertificateAuthority", reflect.TypeOf((*MockCertificateAuthorityClient)(nil).ActivateCertificateAuthority), varargs...)
}

// CreateCaPool mocks base method.
func (m *MockCertificateAuthorityClient) CreateCaPool(arg0 context.Context, arg1 *privatecapb.CreateCaPoolRequest, arg2 ...gax.CallOption) (*privateca.CreateCaPoolOperation, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateCaPool", varargs...)
	ret0, _ := ret[0].(*privateca.CreateCaPoolOperation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCaPool indicates an expected call of CreateCaPool.
func (mr *MockCertificateAuthorityClientMockRecorder) CreateCaPool(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCaPool", reflect.TypeOf((*MockCertificateAuthorityClient)(nil).CreateCaPool), varargs...)
}

// CreateCertificate mocks base method.
func (m *MockCertificateAuthorityClient) CreateCertificate(arg0 context.Context, arg1 *privatecapb.CreateCertificateRequest, arg2 ...gax.CallOption) (*privatecapb.Certificate, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateCertificate", varargs...)
	ret0, _ := ret[0].(*privatecapb.Certificate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCertificate indicates an expected call of CreateCertificate.
func (mr *MockCertificateAuthorityClientMockRecorder) CreateCertificate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCertificate", reflect.TypeOf((*MockCertificateAuthorityClient)(nil).CreateCertificate), varargs...)
}

// CreateCertificateAuthority mocks base method.
func (m *MockCertificateAuthorityClient) CreateCertificateAuthority(arg0 context.Context, arg1 *privatecapb.CreateCertificateAuthorityRequest, arg2 ...gax.CallOption) (*privateca.CreateCertificateAuthorityOperation, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateCertificateAuthority", varargs...)
	ret0, _ := ret[0].(*privateca.CreateCertificateAuthorityOperation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCertificateAuthority indicates an expected call of CreateCertificateAuthority.
func (mr *MockCertificateAuthorityClientMockRecorder) CreateCertificateAuthority(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCertificateAuthority", reflect.TypeOf((*MockCertificateAuthorityClient)(nil).CreateCertificateAuthority), varargs...)
}

// EnableCertificateAuthority mocks base method.
func (m *MockCertificateAuthorityClient) EnableCertificateAuthority(arg0 context.Context, arg1 *privatecapb.EnableCertificateAuthorityRequest, arg2 ...gax.CallOption) (*privateca.EnableCertificateAuthorityOperation, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EnableCertificateAuthority", varargs...)
	ret0, _ := ret[0].(*privateca.EnableCertificateAuthorityOperation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnableCertificateAuthority indicates an expected call of EnableCertificateAuthority.
func (mr *MockCertificateAuthorityClientMockRecorder) EnableCertificateAuthority(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableCertificateAuthority", reflect.TypeOf((*MockCertificateAuthorityClient)(nil).EnableCertificateAuthority), varargs...)
}

// FetchCertificateAuthorityCsr mocks base method.
func (m *MockCertificateAuthorityClient) FetchCertificateAuthorityCsr(arg0 context.Context, arg1 *privatecapb.FetchCertificateAuthorityCsrRequest, arg2 ...gax.CallOption) (*privatecapb.FetchCertificateAuthorityCsrResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FetchCertificateAuthorityCsr", varargs...)
	ret0, _ := ret[0].(*privatecapb.FetchCertificateAuthorityCsrResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchCertificateAuthorityCsr indicates an expected call of FetchCertificateAuthorityCsr.
func (mr *MockCertificateAuthorityClientMockRecorder) FetchCertificateAuthorityCsr(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchCertificateAuthorityCsr", reflect.TypeOf((*MockCertificateAuthorityClient)(nil).FetchCertificateAuthorityCsr), varargs...)
}

// GetCaPool mocks base method.
func (m *MockCertificateAuthorityClient) GetCaPool(arg0 context.Context, arg1 *privatecapb.GetCaPoolRequest, arg2 ...gax.CallOption) (*privatecapb.CaPool, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCaPool", varargs...)
	ret0, _ := ret[0].(*privatecapb.CaPool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCaPool indicates an expected call of GetCaPool.
func (mr *MockCertificateAuthorityClientMockRecorder) GetCaPool(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCaPool", reflect.TypeOf((*MockCertificateAuthorityClient)(nil).GetCaPool), varargs...)
}

// GetCertificateAuthority mocks base method.
func (m *MockCertificateAuthorityClient) GetCertificateAuthority(arg0 context.Context, arg1 *privatecapb.GetCertificateAuthorityRequest, arg2 ...gax.CallOption) (*privatecapb.CertificateAuthority, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetCertificateAuthority", varargs...)
	ret0, _ := ret[0].(*privatecapb.CertificateAuthority)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCertificateAuthority indicates an expected call of GetCertificateAuthority.
func (mr *MockCertificateAuthorityClientMockRecorder) GetCertificateAuthority(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCertificateAuthority", reflect.TypeOf((*MockCertificateAuthorityClient)(nil).GetCertificateAuthority), varargs...)
}

// RevokeCertificate mocks base method.
func (m *MockCertificateAuthorityClient) RevokeCertificate(arg0 context.Context, arg1 *privatecapb.RevokeCertificateRequest, arg2 ...gax.CallOption) (*privatecapb.Certificate, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "RevokeCertificate", varargs...)
	ret0, _ := ret[0].(*privatecapb.Certificate)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RevokeCertificate indicates an expected call of RevokeCertificate.
func (mr *MockCertificateAuthorityClientMockRecorder) RevokeCertificate(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeCertificate", reflect.TypeOf((*MockCertificateAuthorityClient)(nil).RevokeCertificate), varargs...)
}
