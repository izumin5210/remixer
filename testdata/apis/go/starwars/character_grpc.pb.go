// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package starwars_pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion7

// CharacterServiceClient is the client API for CharacterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CharacterServiceClient interface {
	GetCharacter(ctx context.Context, in *GetCharacterRequest, opts ...grpc.CallOption) (*Character, error)
	GetDroid(ctx context.Context, in *GetDroidRequest, opts ...grpc.CallOption) (*Droid, error)
	GetHuman(ctx context.Context, in *GetHumanRequest, opts ...grpc.CallOption) (*Human, error)
}

type characterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewCharacterServiceClient(cc grpc.ClientConnInterface) CharacterServiceClient {
	return &characterServiceClient{cc}
}

func (c *characterServiceClient) GetCharacter(ctx context.Context, in *GetCharacterRequest, opts ...grpc.CallOption) (*Character, error) {
	out := new(Character)
	err := c.cc.Invoke(ctx, "/testapi.starwars.CharacterService/GetCharacter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *characterServiceClient) GetDroid(ctx context.Context, in *GetDroidRequest, opts ...grpc.CallOption) (*Droid, error) {
	out := new(Droid)
	err := c.cc.Invoke(ctx, "/testapi.starwars.CharacterService/GetDroid", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *characterServiceClient) GetHuman(ctx context.Context, in *GetHumanRequest, opts ...grpc.CallOption) (*Human, error) {
	out := new(Human)
	err := c.cc.Invoke(ctx, "/testapi.starwars.CharacterService/GetHuman", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CharacterServiceServer is the server API for CharacterService service.
// All implementations must embed UnimplementedCharacterServiceServer
// for forward compatibility
type CharacterServiceServer interface {
	GetCharacter(context.Context, *GetCharacterRequest) (*Character, error)
	GetDroid(context.Context, *GetDroidRequest) (*Droid, error)
	GetHuman(context.Context, *GetHumanRequest) (*Human, error)
	mustEmbedUnimplementedCharacterServiceServer()
}

// UnimplementedCharacterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedCharacterServiceServer struct {
}

func (UnimplementedCharacterServiceServer) GetCharacter(context.Context, *GetCharacterRequest) (*Character, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCharacter not implemented")
}
func (UnimplementedCharacterServiceServer) GetDroid(context.Context, *GetDroidRequest) (*Droid, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDroid not implemented")
}
func (UnimplementedCharacterServiceServer) GetHuman(context.Context, *GetHumanRequest) (*Human, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHuman not implemented")
}
func (UnimplementedCharacterServiceServer) mustEmbedUnimplementedCharacterServiceServer() {}

// UnsafeCharacterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CharacterServiceServer will
// result in compilation errors.
type UnsafeCharacterServiceServer interface {
	mustEmbedUnimplementedCharacterServiceServer()
}

func RegisterCharacterServiceServer(s grpc.ServiceRegistrar, srv CharacterServiceServer) {
	s.RegisterService(&_CharacterService_serviceDesc, srv)
}

func _CharacterService_GetCharacter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCharacterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharacterServiceServer).GetCharacter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testapi.starwars.CharacterService/GetCharacter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharacterServiceServer).GetCharacter(ctx, req.(*GetCharacterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CharacterService_GetDroid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDroidRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharacterServiceServer).GetDroid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testapi.starwars.CharacterService/GetDroid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharacterServiceServer).GetDroid(ctx, req.(*GetDroidRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CharacterService_GetHuman_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetHumanRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CharacterServiceServer).GetHuman(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/testapi.starwars.CharacterService/GetHuman",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CharacterServiceServer).GetHuman(ctx, req.(*GetHumanRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CharacterService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "testapi.starwars.CharacterService",
	HandlerType: (*CharacterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetCharacter",
			Handler:    _CharacterService_GetCharacter_Handler,
		},
		{
			MethodName: "GetDroid",
			Handler:    _CharacterService_GetDroid_Handler,
		},
		{
			MethodName: "GetHuman",
			Handler:    _CharacterService_GetHuman_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "starwars/character.proto",
}
