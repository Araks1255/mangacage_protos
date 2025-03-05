// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.29.3
// source: notifications.proto

package mangacage_protos

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
	Notifications_NotifyAboutNewTitleOnModeration_FullMethodName       = "/notifications.Notifications/NotifyAboutNewTitleOnModeration"
	Notifications_NotifyAboutNewChapterOnModeration_FullMethodName     = "/notifications.Notifications/NotifyAboutNewChapterOnModeration"
	Notifications_NotifyAboutReleaseOfNewChapterInTitle_FullMethodName = "/notifications.Notifications/NotifyAboutReleaseOfNewChapterInTitle"
	Notifications_SendPromocode_FullMethodName                         = "/notifications.Notifications/SendPromocode"
)

// NotificationsClient is the client API for Notifications service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NotificationsClient interface {
	NotifyAboutNewTitleOnModeration(ctx context.Context, in *TitleOnModeration, opts ...grpc.CallOption) (*Empty, error)
	NotifyAboutNewChapterOnModeration(ctx context.Context, in *ChapterOnModeration, opts ...grpc.CallOption) (*Empty, error)
	NotifyAboutReleaseOfNewChapterInTitle(ctx context.Context, in *ReleasedChapter, opts ...grpc.CallOption) (*Empty, error)
	SendPromocode(ctx context.Context, in *PromocodeRequest, opts ...grpc.CallOption) (*Empty, error)
}

type notificationsClient struct {
	cc grpc.ClientConnInterface
}

func NewNotificationsClient(cc grpc.ClientConnInterface) NotificationsClient {
	return &notificationsClient{cc}
}

func (c *notificationsClient) NotifyAboutNewTitleOnModeration(ctx context.Context, in *TitleOnModeration, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Notifications_NotifyAboutNewTitleOnModeration_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsClient) NotifyAboutNewChapterOnModeration(ctx context.Context, in *ChapterOnModeration, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Notifications_NotifyAboutNewChapterOnModeration_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsClient) NotifyAboutReleaseOfNewChapterInTitle(ctx context.Context, in *ReleasedChapter, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Notifications_NotifyAboutReleaseOfNewChapterInTitle_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *notificationsClient) SendPromocode(ctx context.Context, in *PromocodeRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, Notifications_SendPromocode_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NotificationsServer is the server API for Notifications service.
// All implementations must embed UnimplementedNotificationsServer
// for forward compatibility.
type NotificationsServer interface {
	NotifyAboutNewTitleOnModeration(context.Context, *TitleOnModeration) (*Empty, error)
	NotifyAboutNewChapterOnModeration(context.Context, *ChapterOnModeration) (*Empty, error)
	NotifyAboutReleaseOfNewChapterInTitle(context.Context, *ReleasedChapter) (*Empty, error)
	SendPromocode(context.Context, *PromocodeRequest) (*Empty, error)
	mustEmbedUnimplementedNotificationsServer()
}

// UnimplementedNotificationsServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedNotificationsServer struct{}

func (UnimplementedNotificationsServer) NotifyAboutNewTitleOnModeration(context.Context, *TitleOnModeration) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyAboutNewTitleOnModeration not implemented")
}
func (UnimplementedNotificationsServer) NotifyAboutNewChapterOnModeration(context.Context, *ChapterOnModeration) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyAboutNewChapterOnModeration not implemented")
}
func (UnimplementedNotificationsServer) NotifyAboutReleaseOfNewChapterInTitle(context.Context, *ReleasedChapter) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NotifyAboutReleaseOfNewChapterInTitle not implemented")
}
func (UnimplementedNotificationsServer) SendPromocode(context.Context, *PromocodeRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendPromocode not implemented")
}
func (UnimplementedNotificationsServer) mustEmbedUnimplementedNotificationsServer() {}
func (UnimplementedNotificationsServer) testEmbeddedByValue()                       {}

// UnsafeNotificationsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NotificationsServer will
// result in compilation errors.
type UnsafeNotificationsServer interface {
	mustEmbedUnimplementedNotificationsServer()
}

func RegisterNotificationsServer(s grpc.ServiceRegistrar, srv NotificationsServer) {
	// If the following call pancis, it indicates UnimplementedNotificationsServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Notifications_ServiceDesc, srv)
}

func _Notifications_NotifyAboutNewTitleOnModeration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TitleOnModeration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServer).NotifyAboutNewTitleOnModeration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Notifications_NotifyAboutNewTitleOnModeration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServer).NotifyAboutNewTitleOnModeration(ctx, req.(*TitleOnModeration))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifications_NotifyAboutNewChapterOnModeration_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ChapterOnModeration)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServer).NotifyAboutNewChapterOnModeration(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Notifications_NotifyAboutNewChapterOnModeration_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServer).NotifyAboutNewChapterOnModeration(ctx, req.(*ChapterOnModeration))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifications_NotifyAboutReleaseOfNewChapterInTitle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleasedChapter)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServer).NotifyAboutReleaseOfNewChapterInTitle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Notifications_NotifyAboutReleaseOfNewChapterInTitle_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServer).NotifyAboutReleaseOfNewChapterInTitle(ctx, req.(*ReleasedChapter))
	}
	return interceptor(ctx, in, info, handler)
}

func _Notifications_SendPromocode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PromocodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NotificationsServer).SendPromocode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Notifications_SendPromocode_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NotificationsServer).SendPromocode(ctx, req.(*PromocodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Notifications_ServiceDesc is the grpc.ServiceDesc for Notifications service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Notifications_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "notifications.Notifications",
	HandlerType: (*NotificationsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NotifyAboutNewTitleOnModeration",
			Handler:    _Notifications_NotifyAboutNewTitleOnModeration_Handler,
		},
		{
			MethodName: "NotifyAboutNewChapterOnModeration",
			Handler:    _Notifications_NotifyAboutNewChapterOnModeration_Handler,
		},
		{
			MethodName: "NotifyAboutReleaseOfNewChapterInTitle",
			Handler:    _Notifications_NotifyAboutReleaseOfNewChapterInTitle_Handler,
		},
		{
			MethodName: "SendPromocode",
			Handler:    _Notifications_SendPromocode_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "notifications.proto",
}
