// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cluster-api-provider-hcloud/cluster-api-provider-hcloud/pkg/scope (interfaces: HcloudClient,Manifests,Packer)

// Package mock_scope is a generated GoMock package.
package mock_scope

import (
	context "context"
	v1alpha4 "github.com/cluster-api-provider-hcloud/cluster-api-provider-hcloud/api/v1alpha4"
	api "github.com/cluster-api-provider-hcloud/cluster-api-provider-hcloud/pkg/packer/api"
	logr "github.com/go-logr/logr"
	gomock "github.com/golang/mock/gomock"
	hcloud "github.com/hetznercloud/hcloud-go/hcloud"
	clientcmd "k8s.io/client-go/tools/clientcmd"
	reflect "reflect"
)

// MockHcloudClient is a mock of HcloudClient interface
type MockHcloudClient struct {
	ctrl     *gomock.Controller
	recorder *MockHcloudClientMockRecorder
}

// MockHcloudClientMockRecorder is the mock recorder for MockHcloudClient
type MockHcloudClientMockRecorder struct {
	mock *MockHcloudClient
}

// NewMockHcloudClient creates a new mock instance
func NewMockHcloudClient(ctrl *gomock.Controller) *MockHcloudClient {
	mock := &MockHcloudClient{ctrl: ctrl}
	mock.recorder = &MockHcloudClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockHcloudClient) EXPECT() *MockHcloudClientMockRecorder {
	return m.recorder
}

// CreateLoadBalancer mocks base method
func (m *MockHcloudClient) CreateLoadBalancer(arg0 context.Context, arg1 hcloud.LoadBalancerCreateOpts) (hcloud.LoadBalancerCreateResult, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateLoadBalancer", arg0, arg1)
	ret0, _ := ret[0].(hcloud.LoadBalancerCreateResult)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateLoadBalancer indicates an expected call of CreateLoadBalancer
func (mr *MockHcloudClientMockRecorder) CreateLoadBalancer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateLoadBalancer", reflect.TypeOf((*MockHcloudClient)(nil).CreateLoadBalancer), arg0, arg1)
}

// CreateNetwork mocks base method
func (m *MockHcloudClient) CreateNetwork(arg0 context.Context, arg1 hcloud.NetworkCreateOpts) (*hcloud.Network, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNetwork", arg0, arg1)
	ret0, _ := ret[0].(*hcloud.Network)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateNetwork indicates an expected call of CreateNetwork
func (mr *MockHcloudClientMockRecorder) CreateNetwork(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNetwork", reflect.TypeOf((*MockHcloudClient)(nil).CreateNetwork), arg0, arg1)
}

// CreateServer mocks base method
func (m *MockHcloudClient) CreateServer(arg0 context.Context, arg1 hcloud.ServerCreateOpts) (hcloud.ServerCreateResult, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateServer", arg0, arg1)
	ret0, _ := ret[0].(hcloud.ServerCreateResult)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateServer indicates an expected call of CreateServer
func (mr *MockHcloudClientMockRecorder) CreateServer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateServer", reflect.TypeOf((*MockHcloudClient)(nil).CreateServer), arg0, arg1)
}

// CreateVolume mocks base method
func (m *MockHcloudClient) CreateVolume(arg0 context.Context, arg1 hcloud.VolumeCreateOpts) (hcloud.VolumeCreateResult, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateVolume", arg0, arg1)
	ret0, _ := ret[0].(hcloud.VolumeCreateResult)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// CreateVolume indicates an expected call of CreateVolume
func (mr *MockHcloudClientMockRecorder) CreateVolume(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateVolume", reflect.TypeOf((*MockHcloudClient)(nil).CreateVolume), arg0, arg1)
}

// DeleteLoadBalancer mocks base method
func (m *MockHcloudClient) DeleteLoadBalancer(arg0 context.Context, arg1 *hcloud.LoadBalancer) (*hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteLoadBalancer", arg0, arg1)
	ret0, _ := ret[0].(*hcloud.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteLoadBalancer indicates an expected call of DeleteLoadBalancer
func (mr *MockHcloudClientMockRecorder) DeleteLoadBalancer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLoadBalancer", reflect.TypeOf((*MockHcloudClient)(nil).DeleteLoadBalancer), arg0, arg1)
}

// DeleteNetwork mocks base method
func (m *MockHcloudClient) DeleteNetwork(arg0 context.Context, arg1 *hcloud.Network) (*hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNetwork", arg0, arg1)
	ret0, _ := ret[0].(*hcloud.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteNetwork indicates an expected call of DeleteNetwork
func (mr *MockHcloudClientMockRecorder) DeleteNetwork(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNetwork", reflect.TypeOf((*MockHcloudClient)(nil).DeleteNetwork), arg0, arg1)
}

// DeleteServer mocks base method
func (m *MockHcloudClient) DeleteServer(arg0 context.Context, arg1 *hcloud.Server) (*hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteServer", arg0, arg1)
	ret0, _ := ret[0].(*hcloud.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteServer indicates an expected call of DeleteServer
func (mr *MockHcloudClientMockRecorder) DeleteServer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteServer", reflect.TypeOf((*MockHcloudClient)(nil).DeleteServer), arg0, arg1)
}

// DeleteVolume mocks base method
func (m *MockHcloudClient) DeleteVolume(arg0 context.Context, arg1 *hcloud.Volume) (*hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteVolume", arg0, arg1)
	ret0, _ := ret[0].(*hcloud.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteVolume indicates an expected call of DeleteVolume
func (mr *MockHcloudClientMockRecorder) DeleteVolume(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteVolume", reflect.TypeOf((*MockHcloudClient)(nil).DeleteVolume), arg0, arg1)
}

// ListImages mocks base method
func (m *MockHcloudClient) ListImages(arg0 context.Context, arg1 hcloud.ImageListOpts) ([]*hcloud.Image, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListImages", arg0, arg1)
	ret0, _ := ret[0].([]*hcloud.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListImages indicates an expected call of ListImages
func (mr *MockHcloudClientMockRecorder) ListImages(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListImages", reflect.TypeOf((*MockHcloudClient)(nil).ListImages), arg0, arg1)
}

// ListLoadBalancers mocks base method
func (m *MockHcloudClient) ListLoadBalancers(arg0 context.Context, arg1 hcloud.LoadBalancerListOpts) ([]*hcloud.LoadBalancer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListLoadBalancers", arg0, arg1)
	ret0, _ := ret[0].([]*hcloud.LoadBalancer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListLoadBalancers indicates an expected call of ListLoadBalancers
func (mr *MockHcloudClientMockRecorder) ListLoadBalancers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLoadBalancers", reflect.TypeOf((*MockHcloudClient)(nil).ListLoadBalancers), arg0, arg1)
}

// ListLocation mocks base method
func (m *MockHcloudClient) ListLocation(arg0 context.Context) ([]*hcloud.Location, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListLocation", arg0)
	ret0, _ := ret[0].([]*hcloud.Location)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListLocation indicates an expected call of ListLocation
func (mr *MockHcloudClientMockRecorder) ListLocation(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLocation", reflect.TypeOf((*MockHcloudClient)(nil).ListLocation), arg0)
}

// ListNetworks mocks base method
func (m *MockHcloudClient) ListNetworks(arg0 context.Context, arg1 hcloud.NetworkListOpts) ([]*hcloud.Network, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNetworks", arg0, arg1)
	ret0, _ := ret[0].([]*hcloud.Network)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListNetworks indicates an expected call of ListNetworks
func (mr *MockHcloudClientMockRecorder) ListNetworks(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNetworks", reflect.TypeOf((*MockHcloudClient)(nil).ListNetworks), arg0, arg1)
}

// ListSSHKeys mocks base method
func (m *MockHcloudClient) ListSSHKeys(arg0 context.Context, arg1 hcloud.SSHKeyListOpts) ([]*hcloud.SSHKey, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSSHKeys", arg0, arg1)
	ret0, _ := ret[0].([]*hcloud.SSHKey)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListSSHKeys indicates an expected call of ListSSHKeys
func (mr *MockHcloudClientMockRecorder) ListSSHKeys(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSSHKeys", reflect.TypeOf((*MockHcloudClient)(nil).ListSSHKeys), arg0, arg1)
}

// ListServers mocks base method
func (m *MockHcloudClient) ListServers(arg0 context.Context, arg1 hcloud.ServerListOpts) ([]*hcloud.Server, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListServers", arg0, arg1)
	ret0, _ := ret[0].([]*hcloud.Server)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListServers indicates an expected call of ListServers
func (mr *MockHcloudClientMockRecorder) ListServers(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListServers", reflect.TypeOf((*MockHcloudClient)(nil).ListServers), arg0, arg1)
}

// ListVolumes mocks base method
func (m *MockHcloudClient) ListVolumes(arg0 context.Context, arg1 hcloud.VolumeListOpts) ([]*hcloud.Volume, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListVolumes", arg0, arg1)
	ret0, _ := ret[0].([]*hcloud.Volume)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListVolumes indicates an expected call of ListVolumes
func (mr *MockHcloudClientMockRecorder) ListVolumes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListVolumes", reflect.TypeOf((*MockHcloudClient)(nil).ListVolumes), arg0, arg1)
}

// ShutdownServer mocks base method
func (m *MockHcloudClient) ShutdownServer(arg0 context.Context, arg1 *hcloud.Server) (*hcloud.Action, *hcloud.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShutdownServer", arg0, arg1)
	ret0, _ := ret[0].(*hcloud.Action)
	ret1, _ := ret[1].(*hcloud.Response)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ShutdownServer indicates an expected call of ShutdownServer
func (mr *MockHcloudClientMockRecorder) ShutdownServer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShutdownServer", reflect.TypeOf((*MockHcloudClient)(nil).ShutdownServer), arg0, arg1)
}

// Token mocks base method
func (m *MockHcloudClient) Token() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Token")
	ret0, _ := ret[0].(string)
	return ret0
}

// Token indicates an expected call of Token
func (mr *MockHcloudClientMockRecorder) Token() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Token", reflect.TypeOf((*MockHcloudClient)(nil).Token))
}

// MockManifests is a mock of Manifests interface
type MockManifests struct {
	ctrl     *gomock.Controller
	recorder *MockManifestsMockRecorder
}

// MockManifestsMockRecorder is the mock recorder for MockManifests
type MockManifestsMockRecorder struct {
	mock *MockManifests
}

// NewMockManifests creates a new mock instance
func NewMockManifests(ctrl *gomock.Controller) *MockManifests {
	mock := &MockManifests{ctrl: ctrl}
	mock.recorder = &MockManifestsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockManifests) EXPECT() *MockManifestsMockRecorder {
	return m.recorder
}

// Apply mocks base method
func (m *MockManifests) Apply(arg0 context.Context, arg1 clientcmd.ClientConfig, arg2 map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Apply", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Apply indicates an expected call of Apply
func (mr *MockManifestsMockRecorder) Apply(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockManifests)(nil).Apply), arg0, arg1, arg2)
}

// Hash mocks base method
func (m *MockManifests) Hash(arg0 map[string]string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hash", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Hash indicates an expected call of Hash
func (mr *MockManifestsMockRecorder) Hash(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hash", reflect.TypeOf((*MockManifests)(nil).Hash), arg0)
}

// MockPacker is a mock of Packer interface
type MockPacker struct {
	ctrl     *gomock.Controller
	recorder *MockPackerMockRecorder
}

// MockPackerMockRecorder is the mock recorder for MockPacker
type MockPackerMockRecorder struct {
	mock *MockPacker
}

// NewMockPacker creates a new mock instance
func NewMockPacker(ctrl *gomock.Controller) *MockPacker {
	mock := &MockPacker{ctrl: ctrl}
	mock.recorder = &MockPackerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPacker) EXPECT() *MockPackerMockRecorder {
	return m.recorder
}

// EnsureImage mocks base method
func (m *MockPacker) EnsureImage(arg0 context.Context, arg1 logr.Logger, arg2 api.HcloudClient, arg3 *api.PackerParameters) (*v1alpha4.HcloudImageID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureImage", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*v1alpha4.HcloudImageID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnsureImage indicates an expected call of EnsureImage
func (mr *MockPackerMockRecorder) EnsureImage(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureImage", reflect.TypeOf((*MockPacker)(nil).EnsureImage), arg0, arg1, arg2, arg3)
}
