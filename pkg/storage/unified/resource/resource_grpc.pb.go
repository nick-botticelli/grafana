// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: resource.proto

package resource

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	ResourceStore_Read_FullMethodName   = "/resource.ResourceStore/Read"
	ResourceStore_Create_FullMethodName = "/resource.ResourceStore/Create"
	ResourceStore_Update_FullMethodName = "/resource.ResourceStore/Update"
	ResourceStore_Delete_FullMethodName = "/resource.ResourceStore/Delete"
	ResourceStore_List_FullMethodName   = "/resource.ResourceStore/List"
	ResourceStore_List2_FullMethodName  = "/resource.ResourceStore/List2"
	ResourceStore_Watch_FullMethodName  = "/resource.ResourceStore/Watch"
)

// ResourceStoreClient is the client API for ResourceStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// This provides the CRUD+List+Watch support needed for a k8s apiserver
// The semantics and behaviors of this service are constrained by kubernetes
// This does not understand the resource schemas, only deals with json bytes
// Clients should not use this interface directly; it is for use in API Servers
type ResourceStoreClient interface {
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error)
	Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error)
	Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error)
	Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error)
	// The results *may* include values that should not be returned to the user
	// This will perform best-effort filtering to increase performace.
	// NOTE: storage.Interface is ultimatly responsible for the final filtering
	List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	List2(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error)
	// The results *may* include values that should not be returned to the user
	// This will perform best-effort filtering to increase performace.
	// NOTE: storage.Interface is ultimatly responsible for the final filtering
	Watch(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (ResourceStore_WatchClient, error)
}

type resourceStoreClient struct {
	cc grpc.ClientConnInterface
}

func NewResourceStoreClient(cc grpc.ClientConnInterface) ResourceStoreClient {
	return &resourceStoreClient{cc}
}

func (c *resourceStoreClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ReadResponse)
	err := c.cc.Invoke(ctx, ResourceStore_Read_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceStoreClient) Create(ctx context.Context, in *CreateRequest, opts ...grpc.CallOption) (*CreateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateResponse)
	err := c.cc.Invoke(ctx, ResourceStore_Create_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceStoreClient) Update(ctx context.Context, in *UpdateRequest, opts ...grpc.CallOption) (*UpdateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(UpdateResponse)
	err := c.cc.Invoke(ctx, ResourceStore_Update_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceStoreClient) Delete(ctx context.Context, in *DeleteRequest, opts ...grpc.CallOption) (*DeleteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DeleteResponse)
	err := c.cc.Invoke(ctx, ResourceStore_Delete_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceStoreClient) List(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, ResourceStore_List_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceStoreClient) List2(ctx context.Context, in *ListRequest, opts ...grpc.CallOption) (*ListResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ListResponse)
	err := c.cc.Invoke(ctx, ResourceStore_List2_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceStoreClient) Watch(ctx context.Context, in *WatchRequest, opts ...grpc.CallOption) (ResourceStore_WatchClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ResourceStore_ServiceDesc.Streams[0], ResourceStore_Watch_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &resourceStoreWatchClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ResourceStore_WatchClient interface {
	Recv() (*WatchEvent, error)
	grpc.ClientStream
}

type resourceStoreWatchClient struct {
	grpc.ClientStream
}

func (x *resourceStoreWatchClient) Recv() (*WatchEvent, error) {
	m := new(WatchEvent)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ResourceStoreServer is the server API for ResourceStore service.
// All implementations should embed UnimplementedResourceStoreServer
// for forward compatibility
//
// This provides the CRUD+List+Watch support needed for a k8s apiserver
// The semantics and behaviors of this service are constrained by kubernetes
// This does not understand the resource schemas, only deals with json bytes
// Clients should not use this interface directly; it is for use in API Servers
type ResourceStoreServer interface {
	Read(context.Context, *ReadRequest) (*ReadResponse, error)
	Create(context.Context, *CreateRequest) (*CreateResponse, error)
	Update(context.Context, *UpdateRequest) (*UpdateResponse, error)
	Delete(context.Context, *DeleteRequest) (*DeleteResponse, error)
	// The results *may* include values that should not be returned to the user
	// This will perform best-effort filtering to increase performace.
	// NOTE: storage.Interface is ultimatly responsible for the final filtering
	List(context.Context, *ListRequest) (*ListResponse, error)
	List2(context.Context, *ListRequest) (*ListResponse, error)
	// The results *may* include values that should not be returned to the user
	// This will perform best-effort filtering to increase performace.
	// NOTE: storage.Interface is ultimatly responsible for the final filtering
	Watch(*WatchRequest, ResourceStore_WatchServer) error
}

// UnimplementedResourceStoreServer should be embedded to have forward compatible implementations.
type UnimplementedResourceStoreServer struct {
}

func (UnimplementedResourceStoreServer) Read(context.Context, *ReadRequest) (*ReadResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Read not implemented")
}
func (UnimplementedResourceStoreServer) Create(context.Context, *CreateRequest) (*CreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Create not implemented")
}
func (UnimplementedResourceStoreServer) Update(context.Context, *UpdateRequest) (*UpdateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Update not implemented")
}
func (UnimplementedResourceStoreServer) Delete(context.Context, *DeleteRequest) (*DeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Delete not implemented")
}
func (UnimplementedResourceStoreServer) List(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List not implemented")
}
func (UnimplementedResourceStoreServer) List2(context.Context, *ListRequest) (*ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method List2 not implemented")
}
func (UnimplementedResourceStoreServer) Watch(*WatchRequest, ResourceStore_WatchServer) error {
	return status.Errorf(codes.Unimplemented, "method Watch not implemented")
}

// UnsafeResourceStoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResourceStoreServer will
// result in compilation errors.
type UnsafeResourceStoreServer interface {
	mustEmbedUnimplementedResourceStoreServer()
}

func RegisterResourceStoreServer(s grpc.ServiceRegistrar, srv ResourceStoreServer) {
	s.RegisterService(&ResourceStore_ServiceDesc, srv)
}

func _ResourceStore_Read_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceStoreServer).Read(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceStore_Read_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceStoreServer).Read(ctx, req.(*ReadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceStore_Create_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceStoreServer).Create(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceStore_Create_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceStoreServer).Create(ctx, req.(*CreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceStore_Update_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceStoreServer).Update(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceStore_Update_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceStoreServer).Update(ctx, req.(*UpdateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceStore_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceStoreServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceStore_Delete_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceStoreServer).Delete(ctx, req.(*DeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceStore_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceStoreServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceStore_List_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceStoreServer).List(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceStore_List2_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceStoreServer).List2(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceStore_List2_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceStoreServer).List2(ctx, req.(*ListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceStore_Watch_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ResourceStoreServer).Watch(m, &resourceStoreWatchServer{ServerStream: stream})
}

type ResourceStore_WatchServer interface {
	Send(*WatchEvent) error
	grpc.ServerStream
}

type resourceStoreWatchServer struct {
	grpc.ServerStream
}

func (x *resourceStoreWatchServer) Send(m *WatchEvent) error {
	return x.ServerStream.SendMsg(m)
}

// ResourceStore_ServiceDesc is the grpc.ServiceDesc for ResourceStore service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ResourceStore_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "resource.ResourceStore",
	HandlerType: (*ResourceStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Read",
			Handler:    _ResourceStore_Read_Handler,
		},
		{
			MethodName: "Create",
			Handler:    _ResourceStore_Create_Handler,
		},
		{
			MethodName: "Update",
			Handler:    _ResourceStore_Update_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _ResourceStore_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _ResourceStore_List_Handler,
		},
		{
			MethodName: "List2",
			Handler:    _ResourceStore_List2_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Watch",
			Handler:       _ResourceStore_Watch_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "resource.proto",
}

const (
	ResourceIndex_Search_FullMethodName  = "/resource.ResourceIndex/Search"
	ResourceIndex_History_FullMethodName = "/resource.ResourceIndex/History"
	ResourceIndex_Origin_FullMethodName  = "/resource.ResourceIndex/Origin"
)

// ResourceIndexClient is the client API for ResourceIndex service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Unlike the ResourceStore, this service can be exposed to clients directly
// It should be implemented with efficient indexes and does not need read-after-write semantics
type ResourceIndexClient interface {
	Search(ctx context.Context, in *ResourceSearchRequest, opts ...grpc.CallOption) (*ResourceSearchResponse, error)
	// Show resource history (and trash)
	History(ctx context.Context, in *HistoryRequest, opts ...grpc.CallOption) (*HistoryResponse, error)
	// Used for efficient provisioning
	Origin(ctx context.Context, in *OriginRequest, opts ...grpc.CallOption) (*OriginResponse, error)
}

type resourceIndexClient struct {
	cc grpc.ClientConnInterface
}

func NewResourceIndexClient(cc grpc.ClientConnInterface) ResourceIndexClient {
	return &resourceIndexClient{cc}
}

func (c *resourceIndexClient) Search(ctx context.Context, in *ResourceSearchRequest, opts ...grpc.CallOption) (*ResourceSearchResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ResourceSearchResponse)
	err := c.cc.Invoke(ctx, ResourceIndex_Search_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceIndexClient) History(ctx context.Context, in *HistoryRequest, opts ...grpc.CallOption) (*HistoryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HistoryResponse)
	err := c.cc.Invoke(ctx, ResourceIndex_History_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *resourceIndexClient) Origin(ctx context.Context, in *OriginRequest, opts ...grpc.CallOption) (*OriginResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(OriginResponse)
	err := c.cc.Invoke(ctx, ResourceIndex_Origin_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResourceIndexServer is the server API for ResourceIndex service.
// All implementations should embed UnimplementedResourceIndexServer
// for forward compatibility
//
// Unlike the ResourceStore, this service can be exposed to clients directly
// It should be implemented with efficient indexes and does not need read-after-write semantics
type ResourceIndexServer interface {
	Search(context.Context, *ResourceSearchRequest) (*ResourceSearchResponse, error)
	// Show resource history (and trash)
	History(context.Context, *HistoryRequest) (*HistoryResponse, error)
	// Used for efficient provisioning
	Origin(context.Context, *OriginRequest) (*OriginResponse, error)
}

// UnimplementedResourceIndexServer should be embedded to have forward compatible implementations.
type UnimplementedResourceIndexServer struct {
}

func (UnimplementedResourceIndexServer) Search(context.Context, *ResourceSearchRequest) (*ResourceSearchResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}
func (UnimplementedResourceIndexServer) History(context.Context, *HistoryRequest) (*HistoryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method History not implemented")
}
func (UnimplementedResourceIndexServer) Origin(context.Context, *OriginRequest) (*OriginResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Origin not implemented")
}

// UnsafeResourceIndexServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ResourceIndexServer will
// result in compilation errors.
type UnsafeResourceIndexServer interface {
	mustEmbedUnimplementedResourceIndexServer()
}

func RegisterResourceIndexServer(s grpc.ServiceRegistrar, srv ResourceIndexServer) {
	s.RegisterService(&ResourceIndex_ServiceDesc, srv)
}

func _ResourceIndex_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ResourceSearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceIndexServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceIndex_Search_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceIndexServer).Search(ctx, req.(*ResourceSearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceIndex_History_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceIndexServer).History(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceIndex_History_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceIndexServer).History(ctx, req.(*HistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ResourceIndex_Origin_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OriginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResourceIndexServer).Origin(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ResourceIndex_Origin_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResourceIndexServer).Origin(ctx, req.(*OriginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// ResourceIndex_ServiceDesc is the grpc.ServiceDesc for ResourceIndex service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ResourceIndex_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "resource.ResourceIndex",
	HandlerType: (*ResourceIndexServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _ResourceIndex_Search_Handler,
		},
		{
			MethodName: "History",
			Handler:    _ResourceIndex_History_Handler,
		},
		{
			MethodName: "Origin",
			Handler:    _ResourceIndex_Origin_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "resource.proto",
}

const (
	BlobStore_PutBlob_FullMethodName = "/resource.BlobStore/PutBlob"
	BlobStore_GetBlob_FullMethodName = "/resource.BlobStore/GetBlob"
)

// BlobStoreClient is the client API for BlobStore service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlobStoreClient interface {
	// Upload a blob that will be saved in a resource
	PutBlob(ctx context.Context, in *PutBlobRequest, opts ...grpc.CallOption) (*PutBlobResponse, error)
	// Get blob contents.  When possible, this will return a signed URL
	// For large payloads, signed URLs are required to avoid protobuf message size limits
	GetBlob(ctx context.Context, in *GetBlobRequest, opts ...grpc.CallOption) (*GetBlobResponse, error)
}

type blobStoreClient struct {
	cc grpc.ClientConnInterface
}

func NewBlobStoreClient(cc grpc.ClientConnInterface) BlobStoreClient {
	return &blobStoreClient{cc}
}

func (c *blobStoreClient) PutBlob(ctx context.Context, in *PutBlobRequest, opts ...grpc.CallOption) (*PutBlobResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PutBlobResponse)
	err := c.cc.Invoke(ctx, BlobStore_PutBlob_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blobStoreClient) GetBlob(ctx context.Context, in *GetBlobRequest, opts ...grpc.CallOption) (*GetBlobResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetBlobResponse)
	err := c.cc.Invoke(ctx, BlobStore_GetBlob_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlobStoreServer is the server API for BlobStore service.
// All implementations should embed UnimplementedBlobStoreServer
// for forward compatibility
type BlobStoreServer interface {
	// Upload a blob that will be saved in a resource
	PutBlob(context.Context, *PutBlobRequest) (*PutBlobResponse, error)
	// Get blob contents.  When possible, this will return a signed URL
	// For large payloads, signed URLs are required to avoid protobuf message size limits
	GetBlob(context.Context, *GetBlobRequest) (*GetBlobResponse, error)
}

// UnimplementedBlobStoreServer should be embedded to have forward compatible implementations.
type UnimplementedBlobStoreServer struct {
}

func (UnimplementedBlobStoreServer) PutBlob(context.Context, *PutBlobRequest) (*PutBlobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PutBlob not implemented")
}
func (UnimplementedBlobStoreServer) GetBlob(context.Context, *GetBlobRequest) (*GetBlobResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetBlob not implemented")
}

// UnsafeBlobStoreServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlobStoreServer will
// result in compilation errors.
type UnsafeBlobStoreServer interface {
	mustEmbedUnimplementedBlobStoreServer()
}

func RegisterBlobStoreServer(s grpc.ServiceRegistrar, srv BlobStoreServer) {
	s.RegisterService(&BlobStore_ServiceDesc, srv)
}

func _BlobStore_PutBlob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PutBlobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlobStoreServer).PutBlob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlobStore_PutBlob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlobStoreServer).PutBlob(ctx, req.(*PutBlobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlobStore_GetBlob_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetBlobRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlobStoreServer).GetBlob(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BlobStore_GetBlob_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlobStoreServer).GetBlob(ctx, req.(*GetBlobRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BlobStore_ServiceDesc is the grpc.ServiceDesc for BlobStore service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BlobStore_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "resource.BlobStore",
	HandlerType: (*BlobStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PutBlob",
			Handler:    _BlobStore_PutBlob_Handler,
		},
		{
			MethodName: "GetBlob",
			Handler:    _BlobStore_GetBlob_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "resource.proto",
}

const (
	Diagnostics_IsHealthy_FullMethodName = "/resource.Diagnostics/IsHealthy"
)

// DiagnosticsClient is the client API for Diagnostics service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Clients can use this service directly
// NOTE: This is read only, and no read afer write guarantees
type DiagnosticsClient interface {
	// Check if the service is healthy
	IsHealthy(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error)
}

type diagnosticsClient struct {
	cc grpc.ClientConnInterface
}

func NewDiagnosticsClient(cc grpc.ClientConnInterface) DiagnosticsClient {
	return &diagnosticsClient{cc}
}

func (c *diagnosticsClient) IsHealthy(ctx context.Context, in *HealthCheckRequest, opts ...grpc.CallOption) (*HealthCheckResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(HealthCheckResponse)
	err := c.cc.Invoke(ctx, Diagnostics_IsHealthy_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DiagnosticsServer is the server API for Diagnostics service.
// All implementations should embed UnimplementedDiagnosticsServer
// for forward compatibility
//
// Clients can use this service directly
// NOTE: This is read only, and no read afer write guarantees
type DiagnosticsServer interface {
	// Check if the service is healthy
	IsHealthy(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error)
}

// UnimplementedDiagnosticsServer should be embedded to have forward compatible implementations.
type UnimplementedDiagnosticsServer struct {
}

func (UnimplementedDiagnosticsServer) IsHealthy(context.Context, *HealthCheckRequest) (*HealthCheckResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IsHealthy not implemented")
}

// UnsafeDiagnosticsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DiagnosticsServer will
// result in compilation errors.
type UnsafeDiagnosticsServer interface {
	mustEmbedUnimplementedDiagnosticsServer()
}

func RegisterDiagnosticsServer(s grpc.ServiceRegistrar, srv DiagnosticsServer) {
	s.RegisterService(&Diagnostics_ServiceDesc, srv)
}

func _Diagnostics_IsHealthy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HealthCheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DiagnosticsServer).IsHealthy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Diagnostics_IsHealthy_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DiagnosticsServer).IsHealthy(ctx, req.(*HealthCheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Diagnostics_ServiceDesc is the grpc.ServiceDesc for Diagnostics service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Diagnostics_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "resource.Diagnostics",
	HandlerType: (*DiagnosticsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IsHealthy",
			Handler:    _Diagnostics_IsHealthy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "resource.proto",
}
