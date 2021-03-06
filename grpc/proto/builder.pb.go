// Code generated by protoc-gen-go. DO NOT EDIT.
// source: builder.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	builder.proto

It has these top-level messages:
	BuildParams
	VerboseParams
	Response
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type BuildParams struct {
	HardwareFolders         string `protobuf:"bytes,1,opt,name=hardwareFolders" json:"hardwareFolders,omitempty"`
	ToolsFolders            string `protobuf:"bytes,2,opt,name=toolsFolders" json:"toolsFolders,omitempty"`
	BuiltInLibrariesFolders string `protobuf:"bytes,3,opt,name=builtInLibrariesFolders" json:"builtInLibrariesFolders,omitempty"`
	OtherLibrariesFolders   string `protobuf:"bytes,4,opt,name=otherLibrariesFolders" json:"otherLibrariesFolders,omitempty"`
	SketchLocation          string `protobuf:"bytes,5,opt,name=sketchLocation" json:"sketchLocation,omitempty"`
	FQBN                    string `protobuf:"bytes,6,opt,name=fQBN" json:"fQBN,omitempty"`
	ArduinoAPIVersion       string `protobuf:"bytes,7,opt,name=arduinoAPIVersion" json:"arduinoAPIVersion,omitempty"`
	CustomBuildProperties   string `protobuf:"bytes,8,opt,name=customBuildProperties" json:"customBuildProperties,omitempty"`
	BuildCachePath          string `protobuf:"bytes,9,opt,name=buildCachePath" json:"buildCachePath,omitempty"`
	BuildPath               string `protobuf:"bytes,10,opt,name=buildPath" json:"buildPath,omitempty"`
	WarningsLevel           string `protobuf:"bytes,11,opt,name=warningsLevel" json:"warningsLevel,omitempty"`
	CodeCompleteAt          string `protobuf:"bytes,12,opt,name=codeCompleteAt" json:"codeCompleteAt,omitempty"`
	Verbose                 bool   `protobuf:"varint,13,opt,name=verbose" json:"verbose,omitempty"`
}

func (m *BuildParams) Reset()                    { *m = BuildParams{} }
func (m *BuildParams) String() string            { return proto1.CompactTextString(m) }
func (*BuildParams) ProtoMessage()               {}
func (*BuildParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BuildParams) GetHardwareFolders() string {
	if m != nil {
		return m.HardwareFolders
	}
	return ""
}

func (m *BuildParams) GetToolsFolders() string {
	if m != nil {
		return m.ToolsFolders
	}
	return ""
}

func (m *BuildParams) GetBuiltInLibrariesFolders() string {
	if m != nil {
		return m.BuiltInLibrariesFolders
	}
	return ""
}

func (m *BuildParams) GetOtherLibrariesFolders() string {
	if m != nil {
		return m.OtherLibrariesFolders
	}
	return ""
}

func (m *BuildParams) GetSketchLocation() string {
	if m != nil {
		return m.SketchLocation
	}
	return ""
}

func (m *BuildParams) GetFQBN() string {
	if m != nil {
		return m.FQBN
	}
	return ""
}

func (m *BuildParams) GetArduinoAPIVersion() string {
	if m != nil {
		return m.ArduinoAPIVersion
	}
	return ""
}

func (m *BuildParams) GetCustomBuildProperties() string {
	if m != nil {
		return m.CustomBuildProperties
	}
	return ""
}

func (m *BuildParams) GetBuildCachePath() string {
	if m != nil {
		return m.BuildCachePath
	}
	return ""
}

func (m *BuildParams) GetBuildPath() string {
	if m != nil {
		return m.BuildPath
	}
	return ""
}

func (m *BuildParams) GetWarningsLevel() string {
	if m != nil {
		return m.WarningsLevel
	}
	return ""
}

func (m *BuildParams) GetCodeCompleteAt() string {
	if m != nil {
		return m.CodeCompleteAt
	}
	return ""
}

func (m *BuildParams) GetVerbose() bool {
	if m != nil {
		return m.Verbose
	}
	return false
}

type VerboseParams struct {
	Verbose bool `protobuf:"varint,1,opt,name=verbose" json:"verbose,omitempty"`
}

func (m *VerboseParams) Reset()                    { *m = VerboseParams{} }
func (m *VerboseParams) String() string            { return proto1.CompactTextString(m) }
func (*VerboseParams) ProtoMessage()               {}
func (*VerboseParams) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *VerboseParams) GetVerbose() bool {
	if m != nil {
		return m.Verbose
	}
	return false
}

type Response struct {
	Line string `protobuf:"bytes,1,opt,name=line" json:"line,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto1.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Response) GetLine() string {
	if m != nil {
		return m.Line
	}
	return ""
}

func init() {
	proto1.RegisterType((*BuildParams)(nil), "proto.BuildParams")
	proto1.RegisterType((*VerboseParams)(nil), "proto.VerboseParams")
	proto1.RegisterType((*Response)(nil), "proto.Response")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Builder service

type BuilderClient interface {
	// A server-to-client streaming RPC.
	//
	// Obtains the Features available within the given Rectangle.  Results are
	// streamed rather than returned at once (e.g. in a response message with a
	// repeated field), as the rectangle may cover a large area and contain a
	// huge number of features.
	Build(ctx context.Context, in *BuildParams, opts ...grpc.CallOption) (Builder_BuildClient, error)
	Autocomplete(ctx context.Context, in *BuildParams, opts ...grpc.CallOption) (*Response, error)
	DropCache(ctx context.Context, in *VerboseParams, opts ...grpc.CallOption) (*Response, error)
}

type builderClient struct {
	cc *grpc.ClientConn
}

func NewBuilderClient(cc *grpc.ClientConn) BuilderClient {
	return &builderClient{cc}
}

func (c *builderClient) Build(ctx context.Context, in *BuildParams, opts ...grpc.CallOption) (Builder_BuildClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_Builder_serviceDesc.Streams[0], c.cc, "/proto.Builder/Build", opts...)
	if err != nil {
		return nil, err
	}
	x := &builderBuildClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Builder_BuildClient interface {
	Recv() (*Response, error)
	grpc.ClientStream
}

type builderBuildClient struct {
	grpc.ClientStream
}

func (x *builderBuildClient) Recv() (*Response, error) {
	m := new(Response)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *builderClient) Autocomplete(ctx context.Context, in *BuildParams, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/proto.Builder/Autocomplete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *builderClient) DropCache(ctx context.Context, in *VerboseParams, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/proto.Builder/DropCache", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Builder service

type BuilderServer interface {
	// A server-to-client streaming RPC.
	//
	// Obtains the Features available within the given Rectangle.  Results are
	// streamed rather than returned at once (e.g. in a response message with a
	// repeated field), as the rectangle may cover a large area and contain a
	// huge number of features.
	Build(*BuildParams, Builder_BuildServer) error
	Autocomplete(context.Context, *BuildParams) (*Response, error)
	DropCache(context.Context, *VerboseParams) (*Response, error)
}

func RegisterBuilderServer(s *grpc.Server, srv BuilderServer) {
	s.RegisterService(&_Builder_serviceDesc, srv)
}

func _Builder_Build_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BuildParams)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(BuilderServer).Build(m, &builderBuildServer{stream})
}

type Builder_BuildServer interface {
	Send(*Response) error
	grpc.ServerStream
}

type builderBuildServer struct {
	grpc.ServerStream
}

func (x *builderBuildServer) Send(m *Response) error {
	return x.ServerStream.SendMsg(m)
}

func _Builder_Autocomplete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BuildParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuilderServer).Autocomplete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Builder/Autocomplete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuilderServer).Autocomplete(ctx, req.(*BuildParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Builder_DropCache_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerboseParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BuilderServer).DropCache(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Builder/DropCache",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BuilderServer).DropCache(ctx, req.(*VerboseParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _Builder_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Builder",
	HandlerType: (*BuilderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Autocomplete",
			Handler:    _Builder_Autocomplete_Handler,
		},
		{
			MethodName: "DropCache",
			Handler:    _Builder_DropCache_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Build",
			Handler:       _Builder_Build_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "builder.proto",
}

func init() { proto1.RegisterFile("builder.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 420 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0x6f, 0x6b, 0x13, 0x41,
	0x10, 0xc6, 0x3d, 0x4d, 0x9a, 0x64, 0x9a, 0x58, 0x1c, 0x14, 0x17, 0x11, 0x29, 0xa1, 0x48, 0x04,
	0x09, 0x45, 0x2b, 0xf8, 0x36, 0x57, 0x11, 0x0a, 0x41, 0xce, 0xbc, 0xe8, 0xfb, 0xbd, 0xcd, 0xe8,
	0x2d, 0x5e, 0x6e, 0x8e, 0xdd, 0x4d, 0xfb, 0x59, 0xfc, 0x06, 0x7e, 0x4c, 0xd9, 0x3f, 0xd1, 0x5e,
	0x13, 0xc1, 0x57, 0x99, 0x3c, 0xcf, 0x6f, 0x6e, 0x9f, 0xdb, 0x99, 0x83, 0x49, 0xb9, 0xd5, 0xf5,
	0x9a, 0xcc, 0xbc, 0x35, 0xec, 0x18, 0xfb, 0xe1, 0x67, 0xfa, 0xb3, 0x07, 0xc7, 0xb9, 0x37, 0x0a,
	0x69, 0xe4, 0xc6, 0xe2, 0x0c, 0x4e, 0x2a, 0x69, 0xd6, 0xb7, 0xd2, 0xd0, 0x67, 0xf6, 0xb8, 0x15,
	0xd9, 0x69, 0x36, 0x1b, 0xad, 0xee, 0xcb, 0x38, 0x85, 0xb1, 0x63, 0xae, 0xed, 0x0e, 0x7b, 0x18,
	0xb0, 0x8e, 0x86, 0x1f, 0xe1, 0xb9, 0x3f, 0xd5, 0x5d, 0x35, 0x4b, 0x5d, 0x1a, 0x69, 0x34, 0xfd,
	0xc1, 0x1f, 0x05, 0xfc, 0x5f, 0x36, 0x5e, 0xc0, 0x33, 0x76, 0x15, 0x99, 0xbd, 0xbe, 0x5e, 0xe8,
	0x3b, 0x6c, 0xe2, 0x6b, 0x78, 0x6c, 0x7f, 0x90, 0x53, 0xd5, 0x92, 0x95, 0x74, 0x9a, 0x1b, 0xd1,
	0x0f, 0xf8, 0x3d, 0x15, 0x11, 0x7a, 0xdf, 0xbe, 0xe6, 0x5f, 0xc4, 0x51, 0x70, 0x43, 0x8d, 0x6f,
	0xe1, 0x89, 0x34, 0xeb, 0xad, 0x6e, 0x78, 0x51, 0x5c, 0x5d, 0x93, 0xb1, 0xbe, 0x7d, 0x10, 0x80,
	0x7d, 0xc3, 0xe7, 0x53, 0x5b, 0xeb, 0x78, 0x13, 0x2f, 0xcf, 0x70, 0x4b, 0xc6, 0x69, 0xb2, 0x62,
	0x18, 0xf3, 0x1d, 0x34, 0x7d, 0xbe, 0x30, 0x85, 0x4b, 0xa9, 0x2a, 0x2a, 0xa4, 0xab, 0xc4, 0x28,
	0xe6, 0xeb, 0xaa, 0xf8, 0x12, 0x46, 0x65, 0x1c, 0x8a, 0xab, 0x04, 0x04, 0xe4, 0xaf, 0x80, 0x67,
	0x30, 0xb9, 0x95, 0xa6, 0xd1, 0xcd, 0x77, 0xbb, 0xa4, 0x1b, 0xaa, 0xc5, 0x71, 0x20, 0xba, 0xa2,
	0x3f, 0x4b, 0xf1, 0x9a, 0x2e, 0x79, 0xd3, 0xd6, 0xe4, 0x68, 0xe1, 0xc4, 0x38, 0x9e, 0xd5, 0x55,
	0x51, 0xc0, 0xe0, 0x86, 0x4c, 0xc9, 0x96, 0xc4, 0xe4, 0x34, 0x9b, 0x0d, 0x57, 0xbb, 0xbf, 0xd3,
	0x37, 0x30, 0xb9, 0x8e, 0x65, 0x5a, 0x8e, 0x3b, 0x68, 0xd6, 0x45, 0x5f, 0xc1, 0x70, 0x45, 0xb6,
	0xe5, 0xc6, 0x92, 0xbf, 0xdc, 0x5a, 0x37, 0x94, 0xf6, 0x26, 0xd4, 0xef, 0x7e, 0x65, 0x30, 0xc8,
	0xe3, 0xfe, 0xe1, 0x39, 0xf4, 0x43, 0x89, 0x18, 0x57, 0x71, 0x7e, 0x67, 0xff, 0x5e, 0x9c, 0x24,
	0x6d, 0xf7, 0xb4, 0xe9, 0x83, 0xf3, 0x0c, 0x3f, 0xc0, 0x78, 0xb1, 0x75, 0xac, 0x52, 0xe8, 0xff,
	0x6c, 0xc4, 0x0b, 0x18, 0x7d, 0x32, 0xdc, 0x86, 0x6b, 0xc5, 0xa7, 0xc9, 0xef, 0xbc, 0xd1, 0x81,
	0xae, 0xfc, 0x0c, 0x50, 0xa9, 0x79, 0x9a, 0xf8, 0x3c, 0x7d, 0x34, 0xf9, 0x38, 0xa5, 0x2f, 0x3c,
	0x5e, 0x64, 0xe5, 0x51, 0xe8, 0x7b, 0xff, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xb2, 0x39, 0x66, 0x51,
	0x56, 0x03, 0x00, 0x00,
}
