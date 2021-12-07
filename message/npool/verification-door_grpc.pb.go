// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package npool

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// VerificationDoorClient is the client API for VerificationDoor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type VerificationDoorClient interface {
	// Method Version
	Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VersionResponse, error)
	// get google authentication qr code url
	GetQRcodeURL(ctx context.Context, in *GetQRcodeURLRequest, opts ...grpc.CallOption) (*GetQRcodeURLResponse, error)
	// verify user google authentication(user's input code)
	VerifyGoogleAuth(ctx context.Context, in *VerifyGoogleAuthRequest, opts ...grpc.CallOption) (*VerifyGoogleAuthResponse, error)
	// delete user google authentication record
	DeleteUserGoogleAuth(ctx context.Context, in *DeleteUserGoogleAuthRequest, opts ...grpc.CallOption) (*DeleteUserGoogleAuthResponse, error)
	// send email to user
	SendEmail(ctx context.Context, in *SendEmailRequest, opts ...grpc.CallOption) (*SendEmailResponse, error)
	// send sms to user(todo......)
	SendSms(ctx context.Context, in *SendSmsRequest, opts ...grpc.CallOption) (*SendSmsResponse, error)
	// verify code user input. (can verify email code and sms code, verify sms code is todo......)
	VerifyCode(ctx context.Context, in *VerifyCodeRequest, opts ...grpc.CallOption) (*VerifyCodeResponse, error)
	// verify google recaptcha.
	VerifyGoogleRecaptcha(ctx context.Context, in *VerifyGoogleRecaptchaRequest, opts ...grpc.CallOption) (*VerifyGoogleRecaptchaResponse, error)
	// get captcher imgine url
	GetCaptcherImg(ctx context.Context, in *GetCaptcherImgRequest, opts ...grpc.CallOption) (*GetCaptcherImgResponse, error)
	// verify captcher input
	VerifyCaptcher(ctx context.Context, in *VerifyCaptcherRequest, opts ...grpc.CallOption) (*VerifyCaptcherResponse, error)
}

type verificationDoorClient struct {
	cc grpc.ClientConnInterface
}

func NewVerificationDoorClient(cc grpc.ClientConnInterface) VerificationDoorClient {
	return &verificationDoorClient{cc}
}

func (c *verificationDoorClient) Version(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*VersionResponse, error) {
	out := new(VersionResponse)
	err := c.cc.Invoke(ctx, "/verification.door.v1.VerificationDoor/Version", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *verificationDoorClient) GetQRcodeURL(ctx context.Context, in *GetQRcodeURLRequest, opts ...grpc.CallOption) (*GetQRcodeURLResponse, error) {
	out := new(GetQRcodeURLResponse)
	err := c.cc.Invoke(ctx, "/verification.door.v1.VerificationDoor/GetQRcodeURL", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *verificationDoorClient) VerifyGoogleAuth(ctx context.Context, in *VerifyGoogleAuthRequest, opts ...grpc.CallOption) (*VerifyGoogleAuthResponse, error) {
	out := new(VerifyGoogleAuthResponse)
	err := c.cc.Invoke(ctx, "/verification.door.v1.VerificationDoor/VerifyGoogleAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *verificationDoorClient) DeleteUserGoogleAuth(ctx context.Context, in *DeleteUserGoogleAuthRequest, opts ...grpc.CallOption) (*DeleteUserGoogleAuthResponse, error) {
	out := new(DeleteUserGoogleAuthResponse)
	err := c.cc.Invoke(ctx, "/verification.door.v1.VerificationDoor/DeleteUserGoogleAuth", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *verificationDoorClient) SendEmail(ctx context.Context, in *SendEmailRequest, opts ...grpc.CallOption) (*SendEmailResponse, error) {
	out := new(SendEmailResponse)
	err := c.cc.Invoke(ctx, "/verification.door.v1.VerificationDoor/SendEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *verificationDoorClient) SendSms(ctx context.Context, in *SendSmsRequest, opts ...grpc.CallOption) (*SendSmsResponse, error) {
	out := new(SendSmsResponse)
	err := c.cc.Invoke(ctx, "/verification.door.v1.VerificationDoor/SendSms", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *verificationDoorClient) VerifyCode(ctx context.Context, in *VerifyCodeRequest, opts ...grpc.CallOption) (*VerifyCodeResponse, error) {
	out := new(VerifyCodeResponse)
	err := c.cc.Invoke(ctx, "/verification.door.v1.VerificationDoor/VerifyCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *verificationDoorClient) VerifyGoogleRecaptcha(ctx context.Context, in *VerifyGoogleRecaptchaRequest, opts ...grpc.CallOption) (*VerifyGoogleRecaptchaResponse, error) {
	out := new(VerifyGoogleRecaptchaResponse)
	err := c.cc.Invoke(ctx, "/verification.door.v1.VerificationDoor/VerifyGoogleRecaptcha", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *verificationDoorClient) GetCaptcherImg(ctx context.Context, in *GetCaptcherImgRequest, opts ...grpc.CallOption) (*GetCaptcherImgResponse, error) {
	out := new(GetCaptcherImgResponse)
	err := c.cc.Invoke(ctx, "/verification.door.v1.VerificationDoor/GetCaptcherImg", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *verificationDoorClient) VerifyCaptcher(ctx context.Context, in *VerifyCaptcherRequest, opts ...grpc.CallOption) (*VerifyCaptcherResponse, error) {
	out := new(VerifyCaptcherResponse)
	err := c.cc.Invoke(ctx, "/verification.door.v1.VerificationDoor/VerifyCaptcher", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VerificationDoorServer is the server API for VerificationDoor service.
// All implementations must embed UnimplementedVerificationDoorServer
// for forward compatibility
type VerificationDoorServer interface {
	// Method Version
	Version(context.Context, *emptypb.Empty) (*VersionResponse, error)
	// get google authentication qr code url
	GetQRcodeURL(context.Context, *GetQRcodeURLRequest) (*GetQRcodeURLResponse, error)
	// verify user google authentication(user's input code)
	VerifyGoogleAuth(context.Context, *VerifyGoogleAuthRequest) (*VerifyGoogleAuthResponse, error)
	// delete user google authentication record
	DeleteUserGoogleAuth(context.Context, *DeleteUserGoogleAuthRequest) (*DeleteUserGoogleAuthResponse, error)
	// send email to user
	SendEmail(context.Context, *SendEmailRequest) (*SendEmailResponse, error)
	// send sms to user(todo......)
	SendSms(context.Context, *SendSmsRequest) (*SendSmsResponse, error)
	// verify code user input. (can verify email code and sms code, verify sms code is todo......)
	VerifyCode(context.Context, *VerifyCodeRequest) (*VerifyCodeResponse, error)
	// verify google recaptcha.
	VerifyGoogleRecaptcha(context.Context, *VerifyGoogleRecaptchaRequest) (*VerifyGoogleRecaptchaResponse, error)
	// get captcher imgine url
	GetCaptcherImg(context.Context, *GetCaptcherImgRequest) (*GetCaptcherImgResponse, error)
	// verify captcher input
	VerifyCaptcher(context.Context, *VerifyCaptcherRequest) (*VerifyCaptcherResponse, error)
	mustEmbedUnimplementedVerificationDoorServer()
}

// UnimplementedVerificationDoorServer must be embedded to have forward compatible implementations.
type UnimplementedVerificationDoorServer struct {
}

func (UnimplementedVerificationDoorServer) Version(context.Context, *emptypb.Empty) (*VersionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedVerificationDoorServer) GetQRcodeURL(context.Context, *GetQRcodeURLRequest) (*GetQRcodeURLResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetQRcodeURL not implemented")
}
func (UnimplementedVerificationDoorServer) VerifyGoogleAuth(context.Context, *VerifyGoogleAuthRequest) (*VerifyGoogleAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyGoogleAuth not implemented")
}
func (UnimplementedVerificationDoorServer) DeleteUserGoogleAuth(context.Context, *DeleteUserGoogleAuthRequest) (*DeleteUserGoogleAuthResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteUserGoogleAuth not implemented")
}
func (UnimplementedVerificationDoorServer) SendEmail(context.Context, *SendEmailRequest) (*SendEmailResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendEmail not implemented")
}
func (UnimplementedVerificationDoorServer) SendSms(context.Context, *SendSmsRequest) (*SendSmsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendSms not implemented")
}
func (UnimplementedVerificationDoorServer) VerifyCode(context.Context, *VerifyCodeRequest) (*VerifyCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyCode not implemented")
}
func (UnimplementedVerificationDoorServer) VerifyGoogleRecaptcha(context.Context, *VerifyGoogleRecaptchaRequest) (*VerifyGoogleRecaptchaResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyGoogleRecaptcha not implemented")
}
func (UnimplementedVerificationDoorServer) GetCaptcherImg(context.Context, *GetCaptcherImgRequest) (*GetCaptcherImgResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCaptcherImg not implemented")
}
func (UnimplementedVerificationDoorServer) VerifyCaptcher(context.Context, *VerifyCaptcherRequest) (*VerifyCaptcherResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyCaptcher not implemented")
}
func (UnimplementedVerificationDoorServer) mustEmbedUnimplementedVerificationDoorServer() {}

// UnsafeVerificationDoorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to VerificationDoorServer will
// result in compilation errors.
type UnsafeVerificationDoorServer interface {
	mustEmbedUnimplementedVerificationDoorServer()
}

func RegisterVerificationDoorServer(s grpc.ServiceRegistrar, srv VerificationDoorServer) {
	s.RegisterService(&VerificationDoor_ServiceDesc, srv)
}

func _VerificationDoor_Version_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerificationDoorServer).Version(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/verification.door.v1.VerificationDoor/Version",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerificationDoorServer).Version(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _VerificationDoor_GetQRcodeURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetQRcodeURLRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerificationDoorServer).GetQRcodeURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/verification.door.v1.VerificationDoor/GetQRcodeURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerificationDoorServer).GetQRcodeURL(ctx, req.(*GetQRcodeURLRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VerificationDoor_VerifyGoogleAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyGoogleAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerificationDoorServer).VerifyGoogleAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/verification.door.v1.VerificationDoor/VerifyGoogleAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerificationDoorServer).VerifyGoogleAuth(ctx, req.(*VerifyGoogleAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VerificationDoor_DeleteUserGoogleAuth_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteUserGoogleAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerificationDoorServer).DeleteUserGoogleAuth(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/verification.door.v1.VerificationDoor/DeleteUserGoogleAuth",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerificationDoorServer).DeleteUserGoogleAuth(ctx, req.(*DeleteUserGoogleAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VerificationDoor_SendEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendEmailRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerificationDoorServer).SendEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/verification.door.v1.VerificationDoor/SendEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerificationDoorServer).SendEmail(ctx, req.(*SendEmailRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VerificationDoor_SendSms_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SendSmsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerificationDoorServer).SendSms(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/verification.door.v1.VerificationDoor/SendSms",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerificationDoorServer).SendSms(ctx, req.(*SendSmsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VerificationDoor_VerifyCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyCodeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerificationDoorServer).VerifyCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/verification.door.v1.VerificationDoor/VerifyCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerificationDoorServer).VerifyCode(ctx, req.(*VerifyCodeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VerificationDoor_VerifyGoogleRecaptcha_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyGoogleRecaptchaRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerificationDoorServer).VerifyGoogleRecaptcha(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/verification.door.v1.VerificationDoor/VerifyGoogleRecaptcha",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerificationDoorServer).VerifyGoogleRecaptcha(ctx, req.(*VerifyGoogleRecaptchaRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VerificationDoor_GetCaptcherImg_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetCaptcherImgRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerificationDoorServer).GetCaptcherImg(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/verification.door.v1.VerificationDoor/GetCaptcherImg",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerificationDoorServer).GetCaptcherImg(ctx, req.(*GetCaptcherImgRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VerificationDoor_VerifyCaptcher_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyCaptcherRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VerificationDoorServer).VerifyCaptcher(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/verification.door.v1.VerificationDoor/VerifyCaptcher",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VerificationDoorServer).VerifyCaptcher(ctx, req.(*VerifyCaptcherRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// VerificationDoor_ServiceDesc is the grpc.ServiceDesc for VerificationDoor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var VerificationDoor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "verification.door.v1.VerificationDoor",
	HandlerType: (*VerificationDoorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Version",
			Handler:    _VerificationDoor_Version_Handler,
		},
		{
			MethodName: "GetQRcodeURL",
			Handler:    _VerificationDoor_GetQRcodeURL_Handler,
		},
		{
			MethodName: "VerifyGoogleAuth",
			Handler:    _VerificationDoor_VerifyGoogleAuth_Handler,
		},
		{
			MethodName: "DeleteUserGoogleAuth",
			Handler:    _VerificationDoor_DeleteUserGoogleAuth_Handler,
		},
		{
			MethodName: "SendEmail",
			Handler:    _VerificationDoor_SendEmail_Handler,
		},
		{
			MethodName: "SendSms",
			Handler:    _VerificationDoor_SendSms_Handler,
		},
		{
			MethodName: "VerifyCode",
			Handler:    _VerificationDoor_VerifyCode_Handler,
		},
		{
			MethodName: "VerifyGoogleRecaptcha",
			Handler:    _VerificationDoor_VerifyGoogleRecaptcha_Handler,
		},
		{
			MethodName: "GetCaptcherImg",
			Handler:    _VerificationDoor_GetCaptcherImg_Handler,
		},
		{
			MethodName: "VerifyCaptcher",
			Handler:    _VerificationDoor_VerifyCaptcher_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "npool/verification-door.proto",
}
