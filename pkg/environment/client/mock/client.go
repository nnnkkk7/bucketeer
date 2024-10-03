// Code generated by MockGen. DO NOT EDIT.
// Source: client.go
//
// Generated by this command:
//
//	mockgen -source=client.go -package=mock -destination=./mock/client.go
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"

	environment "github.com/bucketeer-io/bucketeer/proto/environment"
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

// ArchiveEnvironmentV2 mocks base method.
func (m *MockClient) ArchiveEnvironmentV2(ctx context.Context, in *environment.ArchiveEnvironmentV2Request, opts ...grpc.CallOption) (*environment.ArchiveEnvironmentV2Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ArchiveEnvironmentV2", varargs...)
	ret0, _ := ret[0].(*environment.ArchiveEnvironmentV2Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ArchiveEnvironmentV2 indicates an expected call of ArchiveEnvironmentV2.
func (mr *MockClientMockRecorder) ArchiveEnvironmentV2(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ArchiveEnvironmentV2", reflect.TypeOf((*MockClient)(nil).ArchiveEnvironmentV2), varargs...)
}

// ArchiveOrganization mocks base method.
func (m *MockClient) ArchiveOrganization(ctx context.Context, in *environment.ArchiveOrganizationRequest, opts ...grpc.CallOption) (*environment.ArchiveOrganizationResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ArchiveOrganization", varargs...)
	ret0, _ := ret[0].(*environment.ArchiveOrganizationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ArchiveOrganization indicates an expected call of ArchiveOrganization.
func (mr *MockClientMockRecorder) ArchiveOrganization(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ArchiveOrganization", reflect.TypeOf((*MockClient)(nil).ArchiveOrganization), varargs...)
}

// Close mocks base method.
func (m *MockClient) Close() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Close")
}

// Close indicates an expected call of Close.
func (mr *MockClientMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockClient)(nil).Close))
}

// ConvertTrialOrganization mocks base method.
func (m *MockClient) ConvertTrialOrganization(ctx context.Context, in *environment.ConvertTrialOrganizationRequest, opts ...grpc.CallOption) (*environment.ConvertTrialOrganizationResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ConvertTrialOrganization", varargs...)
	ret0, _ := ret[0].(*environment.ConvertTrialOrganizationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConvertTrialOrganization indicates an expected call of ConvertTrialOrganization.
func (mr *MockClientMockRecorder) ConvertTrialOrganization(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvertTrialOrganization", reflect.TypeOf((*MockClient)(nil).ConvertTrialOrganization), varargs...)
}

// ConvertTrialProject mocks base method.
func (m *MockClient) ConvertTrialProject(ctx context.Context, in *environment.ConvertTrialProjectRequest, opts ...grpc.CallOption) (*environment.ConvertTrialProjectResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ConvertTrialProject", varargs...)
	ret0, _ := ret[0].(*environment.ConvertTrialProjectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ConvertTrialProject indicates an expected call of ConvertTrialProject.
func (mr *MockClientMockRecorder) ConvertTrialProject(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConvertTrialProject", reflect.TypeOf((*MockClient)(nil).ConvertTrialProject), varargs...)
}

// CreateEnvironmentV2 mocks base method.
func (m *MockClient) CreateEnvironmentV2(ctx context.Context, in *environment.CreateEnvironmentV2Request, opts ...grpc.CallOption) (*environment.CreateEnvironmentV2Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateEnvironmentV2", varargs...)
	ret0, _ := ret[0].(*environment.CreateEnvironmentV2Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateEnvironmentV2 indicates an expected call of CreateEnvironmentV2.
func (mr *MockClientMockRecorder) CreateEnvironmentV2(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateEnvironmentV2", reflect.TypeOf((*MockClient)(nil).CreateEnvironmentV2), varargs...)
}

// CreateOrganization mocks base method.
func (m *MockClient) CreateOrganization(ctx context.Context, in *environment.CreateOrganizationRequest, opts ...grpc.CallOption) (*environment.CreateOrganizationResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateOrganization", varargs...)
	ret0, _ := ret[0].(*environment.CreateOrganizationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrganization indicates an expected call of CreateOrganization.
func (mr *MockClientMockRecorder) CreateOrganization(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrganization", reflect.TypeOf((*MockClient)(nil).CreateOrganization), varargs...)
}

// CreateProject mocks base method.
func (m *MockClient) CreateProject(ctx context.Context, in *environment.CreateProjectRequest, opts ...grpc.CallOption) (*environment.CreateProjectResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateProject", varargs...)
	ret0, _ := ret[0].(*environment.CreateProjectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProject indicates an expected call of CreateProject.
func (mr *MockClientMockRecorder) CreateProject(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProject", reflect.TypeOf((*MockClient)(nil).CreateProject), varargs...)
}

// CreateTrialProject mocks base method.
func (m *MockClient) CreateTrialProject(ctx context.Context, in *environment.CreateTrialProjectRequest, opts ...grpc.CallOption) (*environment.CreateTrialProjectResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CreateTrialProject", varargs...)
	ret0, _ := ret[0].(*environment.CreateTrialProjectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTrialProject indicates an expected call of CreateTrialProject.
func (mr *MockClientMockRecorder) CreateTrialProject(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTrialProject", reflect.TypeOf((*MockClient)(nil).CreateTrialProject), varargs...)
}

// DisableOrganization mocks base method.
func (m *MockClient) DisableOrganization(ctx context.Context, in *environment.DisableOrganizationRequest, opts ...grpc.CallOption) (*environment.DisableOrganizationResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DisableOrganization", varargs...)
	ret0, _ := ret[0].(*environment.DisableOrganizationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DisableOrganization indicates an expected call of DisableOrganization.
func (mr *MockClientMockRecorder) DisableOrganization(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableOrganization", reflect.TypeOf((*MockClient)(nil).DisableOrganization), varargs...)
}

// DisableProject mocks base method.
func (m *MockClient) DisableProject(ctx context.Context, in *environment.DisableProjectRequest, opts ...grpc.CallOption) (*environment.DisableProjectResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "DisableProject", varargs...)
	ret0, _ := ret[0].(*environment.DisableProjectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DisableProject indicates an expected call of DisableProject.
func (mr *MockClientMockRecorder) DisableProject(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DisableProject", reflect.TypeOf((*MockClient)(nil).DisableProject), varargs...)
}

// EnableOrganization mocks base method.
func (m *MockClient) EnableOrganization(ctx context.Context, in *environment.EnableOrganizationRequest, opts ...grpc.CallOption) (*environment.EnableOrganizationResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EnableOrganization", varargs...)
	ret0, _ := ret[0].(*environment.EnableOrganizationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnableOrganization indicates an expected call of EnableOrganization.
func (mr *MockClientMockRecorder) EnableOrganization(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableOrganization", reflect.TypeOf((*MockClient)(nil).EnableOrganization), varargs...)
}

// EnableProject mocks base method.
func (m *MockClient) EnableProject(ctx context.Context, in *environment.EnableProjectRequest, opts ...grpc.CallOption) (*environment.EnableProjectResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "EnableProject", varargs...)
	ret0, _ := ret[0].(*environment.EnableProjectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnableProject indicates an expected call of EnableProject.
func (mr *MockClientMockRecorder) EnableProject(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableProject", reflect.TypeOf((*MockClient)(nil).EnableProject), varargs...)
}

// GetEnvironmentV2 mocks base method.
func (m *MockClient) GetEnvironmentV2(ctx context.Context, in *environment.GetEnvironmentV2Request, opts ...grpc.CallOption) (*environment.GetEnvironmentV2Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetEnvironmentV2", varargs...)
	ret0, _ := ret[0].(*environment.GetEnvironmentV2Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetEnvironmentV2 indicates an expected call of GetEnvironmentV2.
func (mr *MockClientMockRecorder) GetEnvironmentV2(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEnvironmentV2", reflect.TypeOf((*MockClient)(nil).GetEnvironmentV2), varargs...)
}

// GetOrganization mocks base method.
func (m *MockClient) GetOrganization(ctx context.Context, in *environment.GetOrganizationRequest, opts ...grpc.CallOption) (*environment.GetOrganizationResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetOrganization", varargs...)
	ret0, _ := ret[0].(*environment.GetOrganizationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOrganization indicates an expected call of GetOrganization.
func (mr *MockClientMockRecorder) GetOrganization(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOrganization", reflect.TypeOf((*MockClient)(nil).GetOrganization), varargs...)
}

// GetProject mocks base method.
func (m *MockClient) GetProject(ctx context.Context, in *environment.GetProjectRequest, opts ...grpc.CallOption) (*environment.GetProjectResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetProject", varargs...)
	ret0, _ := ret[0].(*environment.GetProjectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProject indicates an expected call of GetProject.
func (mr *MockClientMockRecorder) GetProject(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProject", reflect.TypeOf((*MockClient)(nil).GetProject), varargs...)
}

// ListEnvironmentsV2 mocks base method.
func (m *MockClient) ListEnvironmentsV2(ctx context.Context, in *environment.ListEnvironmentsV2Request, opts ...grpc.CallOption) (*environment.ListEnvironmentsV2Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListEnvironmentsV2", varargs...)
	ret0, _ := ret[0].(*environment.ListEnvironmentsV2Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEnvironmentsV2 indicates an expected call of ListEnvironmentsV2.
func (mr *MockClientMockRecorder) ListEnvironmentsV2(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEnvironmentsV2", reflect.TypeOf((*MockClient)(nil).ListEnvironmentsV2), varargs...)
}

// ListOrganizations mocks base method.
func (m *MockClient) ListOrganizations(ctx context.Context, in *environment.ListOrganizationsRequest, opts ...grpc.CallOption) (*environment.ListOrganizationsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListOrganizations", varargs...)
	ret0, _ := ret[0].(*environment.ListOrganizationsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListOrganizations indicates an expected call of ListOrganizations.
func (mr *MockClientMockRecorder) ListOrganizations(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListOrganizations", reflect.TypeOf((*MockClient)(nil).ListOrganizations), varargs...)
}

// ListProjects mocks base method.
func (m *MockClient) ListProjects(ctx context.Context, in *environment.ListProjectsRequest, opts ...grpc.CallOption) (*environment.ListProjectsResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListProjects", varargs...)
	ret0, _ := ret[0].(*environment.ListProjectsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProjects indicates an expected call of ListProjects.
func (mr *MockClientMockRecorder) ListProjects(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjects", reflect.TypeOf((*MockClient)(nil).ListProjects), varargs...)
}

// ListProjectsV2 mocks base method.
func (m *MockClient) ListProjectsV2(ctx context.Context, in *environment.ListProjectsV2Request, opts ...grpc.CallOption) (*environment.ListProjectsV2Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListProjectsV2", varargs...)
	ret0, _ := ret[0].(*environment.ListProjectsV2Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListProjectsV2 indicates an expected call of ListProjectsV2.
func (mr *MockClientMockRecorder) ListProjectsV2(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListProjectsV2", reflect.TypeOf((*MockClient)(nil).ListProjectsV2), varargs...)
}

// UnarchiveEnvironmentV2 mocks base method.
func (m *MockClient) UnarchiveEnvironmentV2(ctx context.Context, in *environment.UnarchiveEnvironmentV2Request, opts ...grpc.CallOption) (*environment.UnarchiveEnvironmentV2Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UnarchiveEnvironmentV2", varargs...)
	ret0, _ := ret[0].(*environment.UnarchiveEnvironmentV2Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnarchiveEnvironmentV2 indicates an expected call of UnarchiveEnvironmentV2.
func (mr *MockClientMockRecorder) UnarchiveEnvironmentV2(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnarchiveEnvironmentV2", reflect.TypeOf((*MockClient)(nil).UnarchiveEnvironmentV2), varargs...)
}

// UnarchiveOrganization mocks base method.
func (m *MockClient) UnarchiveOrganization(ctx context.Context, in *environment.UnarchiveOrganizationRequest, opts ...grpc.CallOption) (*environment.UnarchiveOrganizationResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UnarchiveOrganization", varargs...)
	ret0, _ := ret[0].(*environment.UnarchiveOrganizationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UnarchiveOrganization indicates an expected call of UnarchiveOrganization.
func (mr *MockClientMockRecorder) UnarchiveOrganization(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnarchiveOrganization", reflect.TypeOf((*MockClient)(nil).UnarchiveOrganization), varargs...)
}

// UpdateEnvironmentV2 mocks base method.
func (m *MockClient) UpdateEnvironmentV2(ctx context.Context, in *environment.UpdateEnvironmentV2Request, opts ...grpc.CallOption) (*environment.UpdateEnvironmentV2Response, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateEnvironmentV2", varargs...)
	ret0, _ := ret[0].(*environment.UpdateEnvironmentV2Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateEnvironmentV2 indicates an expected call of UpdateEnvironmentV2.
func (mr *MockClientMockRecorder) UpdateEnvironmentV2(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateEnvironmentV2", reflect.TypeOf((*MockClient)(nil).UpdateEnvironmentV2), varargs...)
}

// UpdateOrganization mocks base method.
func (m *MockClient) UpdateOrganization(ctx context.Context, in *environment.UpdateOrganizationRequest, opts ...grpc.CallOption) (*environment.UpdateOrganizationResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateOrganization", varargs...)
	ret0, _ := ret[0].(*environment.UpdateOrganizationResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOrganization indicates an expected call of UpdateOrganization.
func (mr *MockClientMockRecorder) UpdateOrganization(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrganization", reflect.TypeOf((*MockClient)(nil).UpdateOrganization), varargs...)
}

// UpdateProject mocks base method.
func (m *MockClient) UpdateProject(ctx context.Context, in *environment.UpdateProjectRequest, opts ...grpc.CallOption) (*environment.UpdateProjectResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateProject", varargs...)
	ret0, _ := ret[0].(*environment.UpdateProjectResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateProject indicates an expected call of UpdateProject.
func (mr *MockClientMockRecorder) UpdateProject(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProject", reflect.TypeOf((*MockClient)(nil).UpdateProject), varargs...)
}
