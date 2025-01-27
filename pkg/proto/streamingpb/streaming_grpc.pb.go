// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v3.21.4
// source: streaming.proto

package streamingpb

import (
	context "context"
	milvuspb "github.com/milvus-io/milvus-proto/go-api/v2/milvuspb"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	StreamingNodeStateService_GetComponentStates_FullMethodName = "/milvus.proto.streaming.StreamingNodeStateService/GetComponentStates"
)

// StreamingNodeStateServiceClient is the client API for StreamingNodeStateService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamingNodeStateServiceClient interface {
	GetComponentStates(ctx context.Context, in *milvuspb.GetComponentStatesRequest, opts ...grpc.CallOption) (*milvuspb.ComponentStates, error)
}

type streamingNodeStateServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamingNodeStateServiceClient(cc grpc.ClientConnInterface) StreamingNodeStateServiceClient {
	return &streamingNodeStateServiceClient{cc}
}

func (c *streamingNodeStateServiceClient) GetComponentStates(ctx context.Context, in *milvuspb.GetComponentStatesRequest, opts ...grpc.CallOption) (*milvuspb.ComponentStates, error) {
	out := new(milvuspb.ComponentStates)
	err := c.cc.Invoke(ctx, StreamingNodeStateService_GetComponentStates_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StreamingNodeStateServiceServer is the server API for StreamingNodeStateService service.
// All implementations should embed UnimplementedStreamingNodeStateServiceServer
// for forward compatibility
type StreamingNodeStateServiceServer interface {
	GetComponentStates(context.Context, *milvuspb.GetComponentStatesRequest) (*milvuspb.ComponentStates, error)
}

// UnimplementedStreamingNodeStateServiceServer should be embedded to have forward compatible implementations.
type UnimplementedStreamingNodeStateServiceServer struct {
}

func (UnimplementedStreamingNodeStateServiceServer) GetComponentStates(context.Context, *milvuspb.GetComponentStatesRequest) (*milvuspb.ComponentStates, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetComponentStates not implemented")
}

// UnsafeStreamingNodeStateServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamingNodeStateServiceServer will
// result in compilation errors.
type UnsafeStreamingNodeStateServiceServer interface {
	mustEmbedUnimplementedStreamingNodeStateServiceServer()
}

func RegisterStreamingNodeStateServiceServer(s grpc.ServiceRegistrar, srv StreamingNodeStateServiceServer) {
	s.RegisterService(&StreamingNodeStateService_ServiceDesc, srv)
}

func _StreamingNodeStateService_GetComponentStates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(milvuspb.GetComponentStatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamingNodeStateServiceServer).GetComponentStates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StreamingNodeStateService_GetComponentStates_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamingNodeStateServiceServer).GetComponentStates(ctx, req.(*milvuspb.GetComponentStatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StreamingNodeStateService_ServiceDesc is the grpc.ServiceDesc for StreamingNodeStateService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StreamingNodeStateService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "milvus.proto.streaming.StreamingNodeStateService",
	HandlerType: (*StreamingNodeStateServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetComponentStates",
			Handler:    _StreamingNodeStateService_GetComponentStates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "streaming.proto",
}

const (
	StreamingCoordBroadcastService_Broadcast_FullMethodName = "/milvus.proto.streaming.StreamingCoordBroadcastService/Broadcast"
)

// StreamingCoordBroadcastServiceClient is the client API for StreamingCoordBroadcastService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamingCoordBroadcastServiceClient interface {
	// Broadcast receives broadcast messages from other component and make sure that the message is broadcast to all wal.
	// It performs an atomic broadcast to all wal, achieve eventual consistency.
	Broadcast(ctx context.Context, in *BroadcastRequest, opts ...grpc.CallOption) (*BroadcastResponse, error)
}

type streamingCoordBroadcastServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamingCoordBroadcastServiceClient(cc grpc.ClientConnInterface) StreamingCoordBroadcastServiceClient {
	return &streamingCoordBroadcastServiceClient{cc}
}

func (c *streamingCoordBroadcastServiceClient) Broadcast(ctx context.Context, in *BroadcastRequest, opts ...grpc.CallOption) (*BroadcastResponse, error) {
	out := new(BroadcastResponse)
	err := c.cc.Invoke(ctx, StreamingCoordBroadcastService_Broadcast_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StreamingCoordBroadcastServiceServer is the server API for StreamingCoordBroadcastService service.
// All implementations should embed UnimplementedStreamingCoordBroadcastServiceServer
// for forward compatibility
type StreamingCoordBroadcastServiceServer interface {
	// Broadcast receives broadcast messages from other component and make sure that the message is broadcast to all wal.
	// It performs an atomic broadcast to all wal, achieve eventual consistency.
	Broadcast(context.Context, *BroadcastRequest) (*BroadcastResponse, error)
}

// UnimplementedStreamingCoordBroadcastServiceServer should be embedded to have forward compatible implementations.
type UnimplementedStreamingCoordBroadcastServiceServer struct {
}

func (UnimplementedStreamingCoordBroadcastServiceServer) Broadcast(context.Context, *BroadcastRequest) (*BroadcastResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Broadcast not implemented")
}

// UnsafeStreamingCoordBroadcastServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamingCoordBroadcastServiceServer will
// result in compilation errors.
type UnsafeStreamingCoordBroadcastServiceServer interface {
	mustEmbedUnimplementedStreamingCoordBroadcastServiceServer()
}

func RegisterStreamingCoordBroadcastServiceServer(s grpc.ServiceRegistrar, srv StreamingCoordBroadcastServiceServer) {
	s.RegisterService(&StreamingCoordBroadcastService_ServiceDesc, srv)
}

func _StreamingCoordBroadcastService_Broadcast_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BroadcastRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamingCoordBroadcastServiceServer).Broadcast(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StreamingCoordBroadcastService_Broadcast_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamingCoordBroadcastServiceServer).Broadcast(ctx, req.(*BroadcastRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StreamingCoordBroadcastService_ServiceDesc is the grpc.ServiceDesc for StreamingCoordBroadcastService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StreamingCoordBroadcastService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "milvus.proto.streaming.StreamingCoordBroadcastService",
	HandlerType: (*StreamingCoordBroadcastServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Broadcast",
			Handler:    _StreamingCoordBroadcastService_Broadcast_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "streaming.proto",
}

const (
	StreamingCoordAssignmentService_AssignmentDiscover_FullMethodName = "/milvus.proto.streaming.StreamingCoordAssignmentService/AssignmentDiscover"
)

// StreamingCoordAssignmentServiceClient is the client API for StreamingCoordAssignmentService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamingCoordAssignmentServiceClient interface {
	// AssignmentDiscover is used to discover all log nodes managed by the
	// streamingcoord. Channel assignment information will be pushed to client
	// by stream.
	AssignmentDiscover(ctx context.Context, opts ...grpc.CallOption) (StreamingCoordAssignmentService_AssignmentDiscoverClient, error)
}

type streamingCoordAssignmentServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamingCoordAssignmentServiceClient(cc grpc.ClientConnInterface) StreamingCoordAssignmentServiceClient {
	return &streamingCoordAssignmentServiceClient{cc}
}

func (c *streamingCoordAssignmentServiceClient) AssignmentDiscover(ctx context.Context, opts ...grpc.CallOption) (StreamingCoordAssignmentService_AssignmentDiscoverClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamingCoordAssignmentService_ServiceDesc.Streams[0], StreamingCoordAssignmentService_AssignmentDiscover_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &streamingCoordAssignmentServiceAssignmentDiscoverClient{stream}
	return x, nil
}

type StreamingCoordAssignmentService_AssignmentDiscoverClient interface {
	Send(*AssignmentDiscoverRequest) error
	Recv() (*AssignmentDiscoverResponse, error)
	grpc.ClientStream
}

type streamingCoordAssignmentServiceAssignmentDiscoverClient struct {
	grpc.ClientStream
}

func (x *streamingCoordAssignmentServiceAssignmentDiscoverClient) Send(m *AssignmentDiscoverRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamingCoordAssignmentServiceAssignmentDiscoverClient) Recv() (*AssignmentDiscoverResponse, error) {
	m := new(AssignmentDiscoverResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamingCoordAssignmentServiceServer is the server API for StreamingCoordAssignmentService service.
// All implementations should embed UnimplementedStreamingCoordAssignmentServiceServer
// for forward compatibility
type StreamingCoordAssignmentServiceServer interface {
	// AssignmentDiscover is used to discover all log nodes managed by the
	// streamingcoord. Channel assignment information will be pushed to client
	// by stream.
	AssignmentDiscover(StreamingCoordAssignmentService_AssignmentDiscoverServer) error
}

// UnimplementedStreamingCoordAssignmentServiceServer should be embedded to have forward compatible implementations.
type UnimplementedStreamingCoordAssignmentServiceServer struct {
}

func (UnimplementedStreamingCoordAssignmentServiceServer) AssignmentDiscover(StreamingCoordAssignmentService_AssignmentDiscoverServer) error {
	return status.Errorf(codes.Unimplemented, "method AssignmentDiscover not implemented")
}

// UnsafeStreamingCoordAssignmentServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamingCoordAssignmentServiceServer will
// result in compilation errors.
type UnsafeStreamingCoordAssignmentServiceServer interface {
	mustEmbedUnimplementedStreamingCoordAssignmentServiceServer()
}

func RegisterStreamingCoordAssignmentServiceServer(s grpc.ServiceRegistrar, srv StreamingCoordAssignmentServiceServer) {
	s.RegisterService(&StreamingCoordAssignmentService_ServiceDesc, srv)
}

func _StreamingCoordAssignmentService_AssignmentDiscover_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamingCoordAssignmentServiceServer).AssignmentDiscover(&streamingCoordAssignmentServiceAssignmentDiscoverServer{stream})
}

type StreamingCoordAssignmentService_AssignmentDiscoverServer interface {
	Send(*AssignmentDiscoverResponse) error
	Recv() (*AssignmentDiscoverRequest, error)
	grpc.ServerStream
}

type streamingCoordAssignmentServiceAssignmentDiscoverServer struct {
	grpc.ServerStream
}

func (x *streamingCoordAssignmentServiceAssignmentDiscoverServer) Send(m *AssignmentDiscoverResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamingCoordAssignmentServiceAssignmentDiscoverServer) Recv() (*AssignmentDiscoverRequest, error) {
	m := new(AssignmentDiscoverRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamingCoordAssignmentService_ServiceDesc is the grpc.ServiceDesc for StreamingCoordAssignmentService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StreamingCoordAssignmentService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "milvus.proto.streaming.StreamingCoordAssignmentService",
	HandlerType: (*StreamingCoordAssignmentServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "AssignmentDiscover",
			Handler:       _StreamingCoordAssignmentService_AssignmentDiscover_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "streaming.proto",
}

const (
	StreamingNodeHandlerService_Produce_FullMethodName = "/milvus.proto.streaming.StreamingNodeHandlerService/Produce"
	StreamingNodeHandlerService_Consume_FullMethodName = "/milvus.proto.streaming.StreamingNodeHandlerService/Consume"
)

// StreamingNodeHandlerServiceClient is the client API for StreamingNodeHandlerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamingNodeHandlerServiceClient interface {
	// Produce is a bi-directional streaming RPC to send messages to a channel.
	// All messages sent to a channel will be assigned a unique messageID.
	// The messageID is used to identify the message in the channel.
	// The messageID isn't promised to be monotonous increasing with the
	// sequence of responsing. Error: If channel isn't assign to this log node,
	// the RPC will return error CHANNEL_NOT_EXIST. If channel is moving away to
	// other log node, the RPC will return error CHANNEL_FENCED.
	Produce(ctx context.Context, opts ...grpc.CallOption) (StreamingNodeHandlerService_ProduceClient, error)
	// Consume is a server streaming RPC to receive messages from a channel.
	// All message after given startMessageID and excluding will be sent to the
	// client by stream. If no more message in the channel, the stream will be
	// blocked until new message coming. Error: If channel isn't assign to this
	// log node, the RPC will return error CHANNEL_NOT_EXIST. If channel is
	// moving away to other log node, the RPC will return error CHANNEL_FENCED.
	Consume(ctx context.Context, opts ...grpc.CallOption) (StreamingNodeHandlerService_ConsumeClient, error)
}

type streamingNodeHandlerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamingNodeHandlerServiceClient(cc grpc.ClientConnInterface) StreamingNodeHandlerServiceClient {
	return &streamingNodeHandlerServiceClient{cc}
}

func (c *streamingNodeHandlerServiceClient) Produce(ctx context.Context, opts ...grpc.CallOption) (StreamingNodeHandlerService_ProduceClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamingNodeHandlerService_ServiceDesc.Streams[0], StreamingNodeHandlerService_Produce_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &streamingNodeHandlerServiceProduceClient{stream}
	return x, nil
}

type StreamingNodeHandlerService_ProduceClient interface {
	Send(*ProduceRequest) error
	Recv() (*ProduceResponse, error)
	grpc.ClientStream
}

type streamingNodeHandlerServiceProduceClient struct {
	grpc.ClientStream
}

func (x *streamingNodeHandlerServiceProduceClient) Send(m *ProduceRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamingNodeHandlerServiceProduceClient) Recv() (*ProduceResponse, error) {
	m := new(ProduceResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *streamingNodeHandlerServiceClient) Consume(ctx context.Context, opts ...grpc.CallOption) (StreamingNodeHandlerService_ConsumeClient, error) {
	stream, err := c.cc.NewStream(ctx, &StreamingNodeHandlerService_ServiceDesc.Streams[1], StreamingNodeHandlerService_Consume_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &streamingNodeHandlerServiceConsumeClient{stream}
	return x, nil
}

type StreamingNodeHandlerService_ConsumeClient interface {
	Send(*ConsumeRequest) error
	Recv() (*ConsumeResponse, error)
	grpc.ClientStream
}

type streamingNodeHandlerServiceConsumeClient struct {
	grpc.ClientStream
}

func (x *streamingNodeHandlerServiceConsumeClient) Send(m *ConsumeRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *streamingNodeHandlerServiceConsumeClient) Recv() (*ConsumeResponse, error) {
	m := new(ConsumeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamingNodeHandlerServiceServer is the server API for StreamingNodeHandlerService service.
// All implementations should embed UnimplementedStreamingNodeHandlerServiceServer
// for forward compatibility
type StreamingNodeHandlerServiceServer interface {
	// Produce is a bi-directional streaming RPC to send messages to a channel.
	// All messages sent to a channel will be assigned a unique messageID.
	// The messageID is used to identify the message in the channel.
	// The messageID isn't promised to be monotonous increasing with the
	// sequence of responsing. Error: If channel isn't assign to this log node,
	// the RPC will return error CHANNEL_NOT_EXIST. If channel is moving away to
	// other log node, the RPC will return error CHANNEL_FENCED.
	Produce(StreamingNodeHandlerService_ProduceServer) error
	// Consume is a server streaming RPC to receive messages from a channel.
	// All message after given startMessageID and excluding will be sent to the
	// client by stream. If no more message in the channel, the stream will be
	// blocked until new message coming. Error: If channel isn't assign to this
	// log node, the RPC will return error CHANNEL_NOT_EXIST. If channel is
	// moving away to other log node, the RPC will return error CHANNEL_FENCED.
	Consume(StreamingNodeHandlerService_ConsumeServer) error
}

// UnimplementedStreamingNodeHandlerServiceServer should be embedded to have forward compatible implementations.
type UnimplementedStreamingNodeHandlerServiceServer struct {
}

func (UnimplementedStreamingNodeHandlerServiceServer) Produce(StreamingNodeHandlerService_ProduceServer) error {
	return status.Errorf(codes.Unimplemented, "method Produce not implemented")
}
func (UnimplementedStreamingNodeHandlerServiceServer) Consume(StreamingNodeHandlerService_ConsumeServer) error {
	return status.Errorf(codes.Unimplemented, "method Consume not implemented")
}

// UnsafeStreamingNodeHandlerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamingNodeHandlerServiceServer will
// result in compilation errors.
type UnsafeStreamingNodeHandlerServiceServer interface {
	mustEmbedUnimplementedStreamingNodeHandlerServiceServer()
}

func RegisterStreamingNodeHandlerServiceServer(s grpc.ServiceRegistrar, srv StreamingNodeHandlerServiceServer) {
	s.RegisterService(&StreamingNodeHandlerService_ServiceDesc, srv)
}

func _StreamingNodeHandlerService_Produce_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamingNodeHandlerServiceServer).Produce(&streamingNodeHandlerServiceProduceServer{stream})
}

type StreamingNodeHandlerService_ProduceServer interface {
	Send(*ProduceResponse) error
	Recv() (*ProduceRequest, error)
	grpc.ServerStream
}

type streamingNodeHandlerServiceProduceServer struct {
	grpc.ServerStream
}

func (x *streamingNodeHandlerServiceProduceServer) Send(m *ProduceResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamingNodeHandlerServiceProduceServer) Recv() (*ProduceRequest, error) {
	m := new(ProduceRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _StreamingNodeHandlerService_Consume_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(StreamingNodeHandlerServiceServer).Consume(&streamingNodeHandlerServiceConsumeServer{stream})
}

type StreamingNodeHandlerService_ConsumeServer interface {
	Send(*ConsumeResponse) error
	Recv() (*ConsumeRequest, error)
	grpc.ServerStream
}

type streamingNodeHandlerServiceConsumeServer struct {
	grpc.ServerStream
}

func (x *streamingNodeHandlerServiceConsumeServer) Send(m *ConsumeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *streamingNodeHandlerServiceConsumeServer) Recv() (*ConsumeRequest, error) {
	m := new(ConsumeRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// StreamingNodeHandlerService_ServiceDesc is the grpc.ServiceDesc for StreamingNodeHandlerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StreamingNodeHandlerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "milvus.proto.streaming.StreamingNodeHandlerService",
	HandlerType: (*StreamingNodeHandlerServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Produce",
			Handler:       _StreamingNodeHandlerService_Produce_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "Consume",
			Handler:       _StreamingNodeHandlerService_Consume_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "streaming.proto",
}

const (
	StreamingNodeManagerService_Assign_FullMethodName        = "/milvus.proto.streaming.StreamingNodeManagerService/Assign"
	StreamingNodeManagerService_Remove_FullMethodName        = "/milvus.proto.streaming.StreamingNodeManagerService/Remove"
	StreamingNodeManagerService_CollectStatus_FullMethodName = "/milvus.proto.streaming.StreamingNodeManagerService/CollectStatus"
)

// StreamingNodeManagerServiceClient is the client API for StreamingNodeManagerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type StreamingNodeManagerServiceClient interface {
	// Assign is a unary RPC to assign a channel on a log node.
	// Block until the channel assignd is ready to read or write on the log
	// node. Error: If the channel already exists, return error with code
	// CHANNEL_EXIST.
	Assign(ctx context.Context, in *StreamingNodeManagerAssignRequest, opts ...grpc.CallOption) (*StreamingNodeManagerAssignResponse, error)
	// Remove is unary RPC to remove a channel on a log node.
	// Data of the channel on flying would be sent or flused as much as
	// possible. Block until the resource of channel is released on the log
	// node. New incoming request of handler of this channel will be rejected
	// with special error. Error: If the channel does not exist, return error
	// with code CHANNEL_NOT_EXIST.
	Remove(ctx context.Context, in *StreamingNodeManagerRemoveRequest, opts ...grpc.CallOption) (*StreamingNodeManagerRemoveResponse, error)
	// rpc CollectStatus() ...
	// CollectStatus is unary RPC to collect all avaliable channel info and load
	// balance info on a log node. Used to recover channel info on log coord,
	// collect balance info and health check.
	CollectStatus(ctx context.Context, in *StreamingNodeManagerCollectStatusRequest, opts ...grpc.CallOption) (*StreamingNodeManagerCollectStatusResponse, error)
}

type streamingNodeManagerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewStreamingNodeManagerServiceClient(cc grpc.ClientConnInterface) StreamingNodeManagerServiceClient {
	return &streamingNodeManagerServiceClient{cc}
}

func (c *streamingNodeManagerServiceClient) Assign(ctx context.Context, in *StreamingNodeManagerAssignRequest, opts ...grpc.CallOption) (*StreamingNodeManagerAssignResponse, error) {
	out := new(StreamingNodeManagerAssignResponse)
	err := c.cc.Invoke(ctx, StreamingNodeManagerService_Assign_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamingNodeManagerServiceClient) Remove(ctx context.Context, in *StreamingNodeManagerRemoveRequest, opts ...grpc.CallOption) (*StreamingNodeManagerRemoveResponse, error) {
	out := new(StreamingNodeManagerRemoveResponse)
	err := c.cc.Invoke(ctx, StreamingNodeManagerService_Remove_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *streamingNodeManagerServiceClient) CollectStatus(ctx context.Context, in *StreamingNodeManagerCollectStatusRequest, opts ...grpc.CallOption) (*StreamingNodeManagerCollectStatusResponse, error) {
	out := new(StreamingNodeManagerCollectStatusResponse)
	err := c.cc.Invoke(ctx, StreamingNodeManagerService_CollectStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// StreamingNodeManagerServiceServer is the server API for StreamingNodeManagerService service.
// All implementations should embed UnimplementedStreamingNodeManagerServiceServer
// for forward compatibility
type StreamingNodeManagerServiceServer interface {
	// Assign is a unary RPC to assign a channel on a log node.
	// Block until the channel assignd is ready to read or write on the log
	// node. Error: If the channel already exists, return error with code
	// CHANNEL_EXIST.
	Assign(context.Context, *StreamingNodeManagerAssignRequest) (*StreamingNodeManagerAssignResponse, error)
	// Remove is unary RPC to remove a channel on a log node.
	// Data of the channel on flying would be sent or flused as much as
	// possible. Block until the resource of channel is released on the log
	// node. New incoming request of handler of this channel will be rejected
	// with special error. Error: If the channel does not exist, return error
	// with code CHANNEL_NOT_EXIST.
	Remove(context.Context, *StreamingNodeManagerRemoveRequest) (*StreamingNodeManagerRemoveResponse, error)
	// rpc CollectStatus() ...
	// CollectStatus is unary RPC to collect all avaliable channel info and load
	// balance info on a log node. Used to recover channel info on log coord,
	// collect balance info and health check.
	CollectStatus(context.Context, *StreamingNodeManagerCollectStatusRequest) (*StreamingNodeManagerCollectStatusResponse, error)
}

// UnimplementedStreamingNodeManagerServiceServer should be embedded to have forward compatible implementations.
type UnimplementedStreamingNodeManagerServiceServer struct {
}

func (UnimplementedStreamingNodeManagerServiceServer) Assign(context.Context, *StreamingNodeManagerAssignRequest) (*StreamingNodeManagerAssignResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Assign not implemented")
}
func (UnimplementedStreamingNodeManagerServiceServer) Remove(context.Context, *StreamingNodeManagerRemoveRequest) (*StreamingNodeManagerRemoveResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Remove not implemented")
}
func (UnimplementedStreamingNodeManagerServiceServer) CollectStatus(context.Context, *StreamingNodeManagerCollectStatusRequest) (*StreamingNodeManagerCollectStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CollectStatus not implemented")
}

// UnsafeStreamingNodeManagerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to StreamingNodeManagerServiceServer will
// result in compilation errors.
type UnsafeStreamingNodeManagerServiceServer interface {
	mustEmbedUnimplementedStreamingNodeManagerServiceServer()
}

func RegisterStreamingNodeManagerServiceServer(s grpc.ServiceRegistrar, srv StreamingNodeManagerServiceServer) {
	s.RegisterService(&StreamingNodeManagerService_ServiceDesc, srv)
}

func _StreamingNodeManagerService_Assign_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StreamingNodeManagerAssignRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamingNodeManagerServiceServer).Assign(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StreamingNodeManagerService_Assign_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamingNodeManagerServiceServer).Assign(ctx, req.(*StreamingNodeManagerAssignRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreamingNodeManagerService_Remove_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StreamingNodeManagerRemoveRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamingNodeManagerServiceServer).Remove(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StreamingNodeManagerService_Remove_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamingNodeManagerServiceServer).Remove(ctx, req.(*StreamingNodeManagerRemoveRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _StreamingNodeManagerService_CollectStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StreamingNodeManagerCollectStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StreamingNodeManagerServiceServer).CollectStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: StreamingNodeManagerService_CollectStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StreamingNodeManagerServiceServer).CollectStatus(ctx, req.(*StreamingNodeManagerCollectStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// StreamingNodeManagerService_ServiceDesc is the grpc.ServiceDesc for StreamingNodeManagerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var StreamingNodeManagerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "milvus.proto.streaming.StreamingNodeManagerService",
	HandlerType: (*StreamingNodeManagerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Assign",
			Handler:    _StreamingNodeManagerService_Assign_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _StreamingNodeManagerService_Remove_Handler,
		},
		{
			MethodName: "CollectStatus",
			Handler:    _StreamingNodeManagerService_CollectStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "streaming.proto",
}
