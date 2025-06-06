// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.27.0--rc1
// source: panel/panel.proto

package panel

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	PanelService_Login_FullMethodName                = "/panel.v1.PanelService/Login"
	PanelService_ListServers_FullMethodName          = "/panel.v1.PanelService/ListServers"
	PanelService_CreateServer_FullMethodName         = "/panel.v1.PanelService/CreateServer"
	PanelService_GetServerDetails_FullMethodName     = "/panel.v1.PanelService/GetServerDetails"
	PanelService_ListDatabases_FullMethodName        = "/panel.v1.PanelService/ListDatabases"
	PanelService_ActionServer_FullMethodName         = "/panel.v1.PanelService/ActionServer"
	PanelService_CreateDatabaseServer_FullMethodName = "/panel.v1.PanelService/CreateDatabaseServer"
	PanelService_UpdateServer_FullMethodName         = "/panel.v1.PanelService/UpdateServer"
	PanelService_DeleteServer_FullMethodName         = "/panel.v1.PanelService/DeleteServer"
	PanelService_DeleteDatabase_FullMethodName       = "/panel.v1.PanelService/DeleteDatabase"
	PanelService_GetServerLogs_FullMethodName        = "/panel.v1.PanelService/GetServerLogs"
	PanelService_ExecuteCommand_FullMethodName       = "/panel.v1.PanelService/ExecuteCommand"
	PanelService_UploadArchive_FullMethodName        = "/panel.v1.PanelService/UploadArchive"
)

// PanelServiceClient is the client API for PanelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PanelServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
	ListServers(ctx context.Context, in *ListServersRequest, opts ...grpc.CallOption) (*ListServersResponse, error)
	CreateServer(ctx context.Context, in *CreateServerRequest, opts ...grpc.CallOption) (*CreateServerResponse, error)
	GetServerDetails(ctx context.Context, in *ServerDetailsRequest, opts ...grpc.CallOption) (*ServerDetailsResponse, error)
	ListDatabases(ctx context.Context, in *GetDatabasesRequest, opts ...grpc.CallOption) (*GetDatabasesResponse, error)
	ActionServer(ctx context.Context, in *ActionServerRequest, opts ...grpc.CallOption) (*ActionServerResponse, error)
	CreateDatabaseServer(ctx context.Context, in *CreateDatabaseRequest, opts ...grpc.CallOption) (*CreateDatabaseResponse, error)
	UpdateServer(ctx context.Context, in *UpdateServerRequest, opts ...grpc.CallOption) (*UpdateServerResponse, error)
	DeleteServer(ctx context.Context, in *DeleteServerRequest, opts ...grpc.CallOption) (*DeleteServerResponse, error)
	DeleteDatabase(ctx context.Context, in *DeleteDatabaseRequest, opts ...grpc.CallOption) (*DeleteDatabaseResponse, error)
	// Получение логов сервера
	GetServerLogs(ctx context.Context, in *ServerLogRequest, opts ...grpc.CallOption) (*ServerLogResponse, error)
	// Выполнение команды на сервере
	ExecuteCommand(ctx context.Context, in *ExecuteCommandRequest, opts ...grpc.CallOption) (*ExecuteCommandResponse, error)
	// Загрузка архива на сервер
	UploadArchive(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[ArchiveChunk, UploadStatus], error)
}

type panelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPanelServiceClient(cc grpc.ClientConnInterface) PanelServiceClient {
	return &panelServiceClient{cc}
}

func (c *panelServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(LoginResponse)
	err := c.cc.Invoke(ctx, PanelService_Login_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *panelServiceClient) ListServers(ctx context.Context, in *ListServersRequest, opts ...grpc.CallOption) (*ListServersResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListServersResponse)
	err := c.cc.Invoke(ctx, PanelService_ListServers_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *panelServiceClient) CreateServer(ctx context.Context, in *CreateServerRequest, opts ...grpc.CallOption) (*CreateServerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateServerResponse)
	err := c.cc.Invoke(ctx, PanelService_CreateServer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *panelServiceClient) GetServerDetails(ctx context.Context, in *ServerDetailsRequest, opts ...grpc.CallOption) (*ServerDetailsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ServerDetailsResponse)
	err := c.cc.Invoke(ctx, PanelService_GetServerDetails_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *panelServiceClient) ListDatabases(ctx context.Context, in *GetDatabasesRequest, opts ...grpc.CallOption) (*GetDatabasesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetDatabasesResponse)
	err := c.cc.Invoke(ctx, PanelService_ListDatabases_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *panelServiceClient) ActionServer(ctx context.Context, in *ActionServerRequest, opts ...grpc.CallOption) (*ActionServerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ActionServerResponse)
	err := c.cc.Invoke(ctx, PanelService_ActionServer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *panelServiceClient) CreateDatabaseServer(ctx context.Context, in *CreateDatabaseRequest, opts ...grpc.CallOption) (*CreateDatabaseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateDatabaseResponse)
	err := c.cc.Invoke(ctx, PanelService_CreateDatabaseServer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *panelServiceClient) UpdateServer(ctx context.Context, in *UpdateServerRequest, opts ...grpc.CallOption) (*UpdateServerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateServerResponse)
	err := c.cc.Invoke(ctx, PanelService_UpdateServer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *panelServiceClient) DeleteServer(ctx context.Context, in *DeleteServerRequest, opts ...grpc.CallOption) (*DeleteServerResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteServerResponse)
	err := c.cc.Invoke(ctx, PanelService_DeleteServer_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *panelServiceClient) DeleteDatabase(ctx context.Context, in *DeleteDatabaseRequest, opts ...grpc.CallOption) (*DeleteDatabaseResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteDatabaseResponse)
	err := c.cc.Invoke(ctx, PanelService_DeleteDatabase_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *panelServiceClient) GetServerLogs(ctx context.Context, in *ServerLogRequest, opts ...grpc.CallOption) (*ServerLogResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ServerLogResponse)
	err := c.cc.Invoke(ctx, PanelService_GetServerLogs_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *panelServiceClient) ExecuteCommand(ctx context.Context, in *ExecuteCommandRequest, opts ...grpc.CallOption) (*ExecuteCommandResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ExecuteCommandResponse)
	err := c.cc.Invoke(ctx, PanelService_ExecuteCommand_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *panelServiceClient) UploadArchive(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[ArchiveChunk, UploadStatus], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &PanelService_ServiceDesc.Streams[0], PanelService_UploadArchive_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ArchiveChunk, UploadStatus]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type PanelService_UploadArchiveClient = grpc.ClientStreamingClient[ArchiveChunk, UploadStatus]

// PanelServiceServer is the server API for PanelService service.
// All implementations must embed UnimplementedPanelServiceServer
// for forward compatibility.
type PanelServiceServer interface {
	Login(context.Context, *LoginRequest) (*LoginResponse, error)
	ListServers(context.Context, *ListServersRequest) (*ListServersResponse, error)
	CreateServer(context.Context, *CreateServerRequest) (*CreateServerResponse, error)
	GetServerDetails(context.Context, *ServerDetailsRequest) (*ServerDetailsResponse, error)
	ListDatabases(context.Context, *GetDatabasesRequest) (*GetDatabasesResponse, error)
	ActionServer(context.Context, *ActionServerRequest) (*ActionServerResponse, error)
	CreateDatabaseServer(context.Context, *CreateDatabaseRequest) (*CreateDatabaseResponse, error)
	UpdateServer(context.Context, *UpdateServerRequest) (*UpdateServerResponse, error)
	DeleteServer(context.Context, *DeleteServerRequest) (*DeleteServerResponse, error)
	DeleteDatabase(context.Context, *DeleteDatabaseRequest) (*DeleteDatabaseResponse, error)
	// Получение логов сервера
	GetServerLogs(context.Context, *ServerLogRequest) (*ServerLogResponse, error)
	// Выполнение команды на сервере
	ExecuteCommand(context.Context, *ExecuteCommandRequest) (*ExecuteCommandResponse, error)
	// Загрузка архива на сервер
	UploadArchive(grpc.ClientStreamingServer[ArchiveChunk, UploadStatus]) error
	mustEmbedUnimplementedPanelServiceServer()
}

// UnimplementedPanelServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedPanelServiceServer struct{}

func (UnimplementedPanelServiceServer) Login(context.Context, *LoginRequest) (*LoginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedPanelServiceServer) ListServers(context.Context, *ListServersRequest) (*ListServersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListServers not implemented")
}
func (UnimplementedPanelServiceServer) CreateServer(context.Context, *CreateServerRequest) (*CreateServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateServer not implemented")
}
func (UnimplementedPanelServiceServer) GetServerDetails(context.Context, *ServerDetailsRequest) (*ServerDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerDetails not implemented")
}
func (UnimplementedPanelServiceServer) ListDatabases(context.Context, *GetDatabasesRequest) (*GetDatabasesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListDatabases not implemented")
}
func (UnimplementedPanelServiceServer) ActionServer(context.Context, *ActionServerRequest) (*ActionServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ActionServer not implemented")
}
func (UnimplementedPanelServiceServer) CreateDatabaseServer(context.Context, *CreateDatabaseRequest) (*CreateDatabaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateDatabaseServer not implemented")
}
func (UnimplementedPanelServiceServer) UpdateServer(context.Context, *UpdateServerRequest) (*UpdateServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateServer not implemented")
}
func (UnimplementedPanelServiceServer) DeleteServer(context.Context, *DeleteServerRequest) (*DeleteServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteServer not implemented")
}
func (UnimplementedPanelServiceServer) DeleteDatabase(context.Context, *DeleteDatabaseRequest) (*DeleteDatabaseResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteDatabase not implemented")
}
func (UnimplementedPanelServiceServer) GetServerLogs(context.Context, *ServerLogRequest) (*ServerLogResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetServerLogs not implemented")
}
func (UnimplementedPanelServiceServer) ExecuteCommand(context.Context, *ExecuteCommandRequest) (*ExecuteCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ExecuteCommand not implemented")
}
func (UnimplementedPanelServiceServer) UploadArchive(grpc.ClientStreamingServer[ArchiveChunk, UploadStatus]) error {
	return status.Errorf(codes.Unimplemented, "method UploadArchive not implemented")
}
func (UnimplementedPanelServiceServer) mustEmbedUnimplementedPanelServiceServer() {}
func (UnimplementedPanelServiceServer) testEmbeddedByValue()                      {}

// UnsafePanelServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PanelServiceServer will
// result in compilation errors.
type UnsafePanelServiceServer interface {
	mustEmbedUnimplementedPanelServiceServer()
}

func RegisterPanelServiceServer(s grpc.ServiceRegistrar, srv PanelServiceServer) {
	// If the following call pancis, it indicates UnimplementedPanelServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&PanelService_ServiceDesc, srv)
}

func _PanelService_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PanelServiceServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PanelService_Login_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PanelServiceServer).Login(ctx, req.(*LoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PanelService_ListServers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListServersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PanelServiceServer).ListServers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PanelService_ListServers_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PanelServiceServer).ListServers(ctx, req.(*ListServersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PanelService_CreateServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PanelServiceServer).CreateServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PanelService_CreateServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PanelServiceServer).CreateServer(ctx, req.(*CreateServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PanelService_GetServerDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PanelServiceServer).GetServerDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PanelService_GetServerDetails_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PanelServiceServer).GetServerDetails(ctx, req.(*ServerDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PanelService_ListDatabases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDatabasesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PanelServiceServer).ListDatabases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PanelService_ListDatabases_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PanelServiceServer).ListDatabases(ctx, req.(*GetDatabasesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PanelService_ActionServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ActionServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PanelServiceServer).ActionServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PanelService_ActionServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PanelServiceServer).ActionServer(ctx, req.(*ActionServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PanelService_CreateDatabaseServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateDatabaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PanelServiceServer).CreateDatabaseServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PanelService_CreateDatabaseServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PanelServiceServer).CreateDatabaseServer(ctx, req.(*CreateDatabaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PanelService_UpdateServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PanelServiceServer).UpdateServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PanelService_UpdateServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PanelServiceServer).UpdateServer(ctx, req.(*UpdateServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PanelService_DeleteServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PanelServiceServer).DeleteServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PanelService_DeleteServer_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PanelServiceServer).DeleteServer(ctx, req.(*DeleteServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PanelService_DeleteDatabase_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDatabaseRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PanelServiceServer).DeleteDatabase(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PanelService_DeleteDatabase_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PanelServiceServer).DeleteDatabase(ctx, req.(*DeleteDatabaseRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PanelService_GetServerLogs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ServerLogRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PanelServiceServer).GetServerLogs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PanelService_GetServerLogs_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PanelServiceServer).GetServerLogs(ctx, req.(*ServerLogRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PanelService_ExecuteCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ExecuteCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PanelServiceServer).ExecuteCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PanelService_ExecuteCommand_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PanelServiceServer).ExecuteCommand(ctx, req.(*ExecuteCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PanelService_UploadArchive_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(PanelServiceServer).UploadArchive(&grpc.GenericServerStream[ArchiveChunk, UploadStatus]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type PanelService_UploadArchiveServer = grpc.ClientStreamingServer[ArchiveChunk, UploadStatus]

// PanelService_ServiceDesc is the grpc.ServiceDesc for PanelService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PanelService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "panel.v1.PanelService",
	HandlerType: (*PanelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _PanelService_Login_Handler,
		},
		{
			MethodName: "ListServers",
			Handler:    _PanelService_ListServers_Handler,
		},
		{
			MethodName: "CreateServer",
			Handler:    _PanelService_CreateServer_Handler,
		},
		{
			MethodName: "GetServerDetails",
			Handler:    _PanelService_GetServerDetails_Handler,
		},
		{
			MethodName: "ListDatabases",
			Handler:    _PanelService_ListDatabases_Handler,
		},
		{
			MethodName: "ActionServer",
			Handler:    _PanelService_ActionServer_Handler,
		},
		{
			MethodName: "CreateDatabaseServer",
			Handler:    _PanelService_CreateDatabaseServer_Handler,
		},
		{
			MethodName: "UpdateServer",
			Handler:    _PanelService_UpdateServer_Handler,
		},
		{
			MethodName: "DeleteServer",
			Handler:    _PanelService_DeleteServer_Handler,
		},
		{
			MethodName: "DeleteDatabase",
			Handler:    _PanelService_DeleteDatabase_Handler,
		},
		{
			MethodName: "GetServerLogs",
			Handler:    _PanelService_GetServerLogs_Handler,
		},
		{
			MethodName: "ExecuteCommand",
			Handler:    _PanelService_ExecuteCommand_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "UploadArchive",
			Handler:       _PanelService_UploadArchive_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "panel/panel.proto",
}
