// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             (unknown)
// source: authzed/api/materialize/v0/watchpermissionsets.proto

package v0

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
	WatchPermissionSetsService_WatchPermissionSets_FullMethodName  = "/authzed.api.materialize.v0.WatchPermissionSetsService/WatchPermissionSets"
	WatchPermissionSetsService_LookupPermissionSets_FullMethodName = "/authzed.api.materialize.v0.WatchPermissionSetsService/LookupPermissionSets"
)

// WatchPermissionSetsServiceClient is the client API for WatchPermissionSetsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type WatchPermissionSetsServiceClient interface {
	// WatchPermissionSets returns a stream of changes to the sets which can be used to compute the watched permissions.
	//
	// WatchPermissionSets lets consumers achieve the same thing as WatchPermissions, but trades off a simpler usage model with
	// significantly lower computational requirements. Unlike WatchPermissions, this method returns changes to the sets of permissions,
	// rather than the individual permissions. Permission sets are a normalized form of the computed permissions, which
	// means that the consumer must perform an extra computation over this representation to obtain the final computed
	// permissions, typically by intersecting the provided sets.
	//
	// For example, this would look like a JOIN between the
	// materialize permission sets table in a target relation database, the table with the resources to authorize access
	// to, and the table with the subject (e.g. a user).
	//
	// In exchange, the number of changes issued by WatchPermissionSets will be several orders of magnitude less than those
	// emitted by WatchPermissions, which has several implications:
	// - significantly less resources to compute the sets
	// - significantly less messages to stream over the network
	// - significantly less events to ingest on the consumer side
	// - less ingestion lag from the origin SpiceDB mutation
	//
	// The type of scenarios WatchPermissionSets is particularly well suited is when a single change
	// in the origin SpiceDB can yield millions of changes. For example, in the GitHub authorization model, assigning a role
	// to a top-level team of an organization with hundreds of thousands of employees can lead to an explosion of
	// permission change events that would require a lot of computational resources to process, both on Materialize and
	// the consumer side.
	//
	// WatchPermissionSets is thus recommended for any larger scale use case where the fan-out in permission changes that
	// emerges from a specific schema and data shape is too large to handle effectively.
	//
	// The API does not offer a sharding mechanism and thus there should only be one consumer per target system.
	// Implementing an active-active HA consumer setup over the same target system will require coordinating which
	// revisions have been consumed in order to prevent transitioning to an inconsistent state.
	WatchPermissionSets(ctx context.Context, in *WatchPermissionSetsRequest, opts ...grpc.CallOption) (WatchPermissionSetsService_WatchPermissionSetsClient, error)
	// LookupPermissionSets returns the current state of the permission sets which can be used to derive the computed permissions.
	// It's typically used to backfill the state of the permission sets in the consumer side.
	//
	// It's a cursored API and the consumer is responsible to keep track of the cursor and use it on each subsequent call.
	// Each stream will return <N> permission sets defined by the specified request limit. The server will keep streaming until
	// the sets per stream is hit, or the current state of the sets is reached,
	// whatever happens first, and then close the stream. The server will indicate there are no more changes to stream
	// through the `completed_members` in the cursor.
	//
	// There may be many elements to stream, and so the consumer should be prepared to resume the stream from the last
	// cursor received. Once completed, the consumer may start streaming permission set changes using WatchPermissionSets
	// and the revision token from the last LookupPermissionSets response.
	LookupPermissionSets(ctx context.Context, in *LookupPermissionSetsRequest, opts ...grpc.CallOption) (WatchPermissionSetsService_LookupPermissionSetsClient, error)
}

type watchPermissionSetsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewWatchPermissionSetsServiceClient(cc grpc.ClientConnInterface) WatchPermissionSetsServiceClient {
	return &watchPermissionSetsServiceClient{cc}
}

func (c *watchPermissionSetsServiceClient) WatchPermissionSets(ctx context.Context, in *WatchPermissionSetsRequest, opts ...grpc.CallOption) (WatchPermissionSetsService_WatchPermissionSetsClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &WatchPermissionSetsService_ServiceDesc.Streams[0], WatchPermissionSetsService_WatchPermissionSets_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &watchPermissionSetsServiceWatchPermissionSetsClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchPermissionSetsService_WatchPermissionSetsClient interface {
	Recv() (*WatchPermissionSetsResponse, error)
	grpc.ClientStream
}

type watchPermissionSetsServiceWatchPermissionSetsClient struct {
	grpc.ClientStream
}

func (x *watchPermissionSetsServiceWatchPermissionSetsClient) Recv() (*WatchPermissionSetsResponse, error) {
	m := new(WatchPermissionSetsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *watchPermissionSetsServiceClient) LookupPermissionSets(ctx context.Context, in *LookupPermissionSetsRequest, opts ...grpc.CallOption) (WatchPermissionSetsService_LookupPermissionSetsClient, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &WatchPermissionSetsService_ServiceDesc.Streams[1], WatchPermissionSetsService_LookupPermissionSets_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &watchPermissionSetsServiceLookupPermissionSetsClient{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type WatchPermissionSetsService_LookupPermissionSetsClient interface {
	Recv() (*LookupPermissionSetsResponse, error)
	grpc.ClientStream
}

type watchPermissionSetsServiceLookupPermissionSetsClient struct {
	grpc.ClientStream
}

func (x *watchPermissionSetsServiceLookupPermissionSetsClient) Recv() (*LookupPermissionSetsResponse, error) {
	m := new(LookupPermissionSetsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// WatchPermissionSetsServiceServer is the server API for WatchPermissionSetsService service.
// All implementations must embed UnimplementedWatchPermissionSetsServiceServer
// for forward compatibility
type WatchPermissionSetsServiceServer interface {
	// WatchPermissionSets returns a stream of changes to the sets which can be used to compute the watched permissions.
	//
	// WatchPermissionSets lets consumers achieve the same thing as WatchPermissions, but trades off a simpler usage model with
	// significantly lower computational requirements. Unlike WatchPermissions, this method returns changes to the sets of permissions,
	// rather than the individual permissions. Permission sets are a normalized form of the computed permissions, which
	// means that the consumer must perform an extra computation over this representation to obtain the final computed
	// permissions, typically by intersecting the provided sets.
	//
	// For example, this would look like a JOIN between the
	// materialize permission sets table in a target relation database, the table with the resources to authorize access
	// to, and the table with the subject (e.g. a user).
	//
	// In exchange, the number of changes issued by WatchPermissionSets will be several orders of magnitude less than those
	// emitted by WatchPermissions, which has several implications:
	// - significantly less resources to compute the sets
	// - significantly less messages to stream over the network
	// - significantly less events to ingest on the consumer side
	// - less ingestion lag from the origin SpiceDB mutation
	//
	// The type of scenarios WatchPermissionSets is particularly well suited is when a single change
	// in the origin SpiceDB can yield millions of changes. For example, in the GitHub authorization model, assigning a role
	// to a top-level team of an organization with hundreds of thousands of employees can lead to an explosion of
	// permission change events that would require a lot of computational resources to process, both on Materialize and
	// the consumer side.
	//
	// WatchPermissionSets is thus recommended for any larger scale use case where the fan-out in permission changes that
	// emerges from a specific schema and data shape is too large to handle effectively.
	//
	// The API does not offer a sharding mechanism and thus there should only be one consumer per target system.
	// Implementing an active-active HA consumer setup over the same target system will require coordinating which
	// revisions have been consumed in order to prevent transitioning to an inconsistent state.
	WatchPermissionSets(*WatchPermissionSetsRequest, WatchPermissionSetsService_WatchPermissionSetsServer) error
	// LookupPermissionSets returns the current state of the permission sets which can be used to derive the computed permissions.
	// It's typically used to backfill the state of the permission sets in the consumer side.
	//
	// It's a cursored API and the consumer is responsible to keep track of the cursor and use it on each subsequent call.
	// Each stream will return <N> permission sets defined by the specified request limit. The server will keep streaming until
	// the sets per stream is hit, or the current state of the sets is reached,
	// whatever happens first, and then close the stream. The server will indicate there are no more changes to stream
	// through the `completed_members` in the cursor.
	//
	// There may be many elements to stream, and so the consumer should be prepared to resume the stream from the last
	// cursor received. Once completed, the consumer may start streaming permission set changes using WatchPermissionSets
	// and the revision token from the last LookupPermissionSets response.
	LookupPermissionSets(*LookupPermissionSetsRequest, WatchPermissionSetsService_LookupPermissionSetsServer) error
	mustEmbedUnimplementedWatchPermissionSetsServiceServer()
}

// UnimplementedWatchPermissionSetsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedWatchPermissionSetsServiceServer struct {
}

func (UnimplementedWatchPermissionSetsServiceServer) WatchPermissionSets(*WatchPermissionSetsRequest, WatchPermissionSetsService_WatchPermissionSetsServer) error {
	return status.Errorf(codes.Unimplemented, "method WatchPermissionSets not implemented")
}
func (UnimplementedWatchPermissionSetsServiceServer) LookupPermissionSets(*LookupPermissionSetsRequest, WatchPermissionSetsService_LookupPermissionSetsServer) error {
	return status.Errorf(codes.Unimplemented, "method LookupPermissionSets not implemented")
}
func (UnimplementedWatchPermissionSetsServiceServer) mustEmbedUnimplementedWatchPermissionSetsServiceServer() {
}

// UnsafeWatchPermissionSetsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to WatchPermissionSetsServiceServer will
// result in compilation errors.
type UnsafeWatchPermissionSetsServiceServer interface {
	mustEmbedUnimplementedWatchPermissionSetsServiceServer()
}

func RegisterWatchPermissionSetsServiceServer(s grpc.ServiceRegistrar, srv WatchPermissionSetsServiceServer) {
	s.RegisterService(&WatchPermissionSetsService_ServiceDesc, srv)
}

func _WatchPermissionSetsService_WatchPermissionSets_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(WatchPermissionSetsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WatchPermissionSetsServiceServer).WatchPermissionSets(m, &watchPermissionSetsServiceWatchPermissionSetsServer{ServerStream: stream})
}

type WatchPermissionSetsService_WatchPermissionSetsServer interface {
	Send(*WatchPermissionSetsResponse) error
	grpc.ServerStream
}

type watchPermissionSetsServiceWatchPermissionSetsServer struct {
	grpc.ServerStream
}

func (x *watchPermissionSetsServiceWatchPermissionSetsServer) Send(m *WatchPermissionSetsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _WatchPermissionSetsService_LookupPermissionSets_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LookupPermissionSetsRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(WatchPermissionSetsServiceServer).LookupPermissionSets(m, &watchPermissionSetsServiceLookupPermissionSetsServer{ServerStream: stream})
}

type WatchPermissionSetsService_LookupPermissionSetsServer interface {
	Send(*LookupPermissionSetsResponse) error
	grpc.ServerStream
}

type watchPermissionSetsServiceLookupPermissionSetsServer struct {
	grpc.ServerStream
}

func (x *watchPermissionSetsServiceLookupPermissionSetsServer) Send(m *LookupPermissionSetsResponse) error {
	return x.ServerStream.SendMsg(m)
}

// WatchPermissionSetsService_ServiceDesc is the grpc.ServiceDesc for WatchPermissionSetsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var WatchPermissionSetsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "authzed.api.materialize.v0.WatchPermissionSetsService",
	HandlerType: (*WatchPermissionSetsServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "WatchPermissionSets",
			Handler:       _WatchPermissionSetsService_WatchPermissionSets_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "LookupPermissionSets",
			Handler:       _WatchPermissionSetsService_LookupPermissionSets_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "authzed/api/materialize/v0/watchpermissionsets.proto",
}
