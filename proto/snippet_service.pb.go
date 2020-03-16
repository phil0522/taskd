// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/snippet_service.proto

package phil0522_taskd

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ShellSnippetRequest struct {
	CurrentDirectory string `protobuf:"bytes,1,opt,name=current_directory,json=currentDirectory,proto3" json:"current_directory,omitempty"`
	// The current input.
	CmdPrefix string `protobuf:"bytes,2,opt,name=cmd_prefix,json=cmdPrefix,proto3" json:"cmd_prefix,omitempty"`
	// Category of the snippet, this is group name part of snippet name.
	Category             string   `protobuf:"bytes,3,opt,name=category,proto3" json:"category,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShellSnippetRequest) Reset()         { *m = ShellSnippetRequest{} }
func (m *ShellSnippetRequest) String() string { return proto.CompactTextString(m) }
func (*ShellSnippetRequest) ProtoMessage()    {}
func (*ShellSnippetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd244997a2724ff9, []int{0}
}

func (m *ShellSnippetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShellSnippetRequest.Unmarshal(m, b)
}
func (m *ShellSnippetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShellSnippetRequest.Marshal(b, m, deterministic)
}
func (m *ShellSnippetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShellSnippetRequest.Merge(m, src)
}
func (m *ShellSnippetRequest) XXX_Size() int {
	return xxx_messageInfo_ShellSnippetRequest.Size(m)
}
func (m *ShellSnippetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ShellSnippetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ShellSnippetRequest proto.InternalMessageInfo

func (m *ShellSnippetRequest) GetCurrentDirectory() string {
	if m != nil {
		return m.CurrentDirectory
	}
	return ""
}

func (m *ShellSnippetRequest) GetCmdPrefix() string {
	if m != nil {
		return m.CmdPrefix
	}
	return ""
}

func (m *ShellSnippetRequest) GetCategory() string {
	if m != nil {
		return m.Category
	}
	return ""
}

type ShellSnippetResponse struct {
	ShellSnippets        []*ShellSnippet `protobuf:"bytes,1,rep,name=shell_snippets,json=shellSnippets,proto3" json:"shell_snippets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ShellSnippetResponse) Reset()         { *m = ShellSnippetResponse{} }
func (m *ShellSnippetResponse) String() string { return proto.CompactTextString(m) }
func (*ShellSnippetResponse) ProtoMessage()    {}
func (*ShellSnippetResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd244997a2724ff9, []int{1}
}

func (m *ShellSnippetResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShellSnippetResponse.Unmarshal(m, b)
}
func (m *ShellSnippetResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShellSnippetResponse.Marshal(b, m, deterministic)
}
func (m *ShellSnippetResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShellSnippetResponse.Merge(m, src)
}
func (m *ShellSnippetResponse) XXX_Size() int {
	return xxx_messageInfo_ShellSnippetResponse.Size(m)
}
func (m *ShellSnippetResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ShellSnippetResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ShellSnippetResponse proto.InternalMessageInfo

func (m *ShellSnippetResponse) GetShellSnippets() []*ShellSnippet {
	if m != nil {
		return m.ShellSnippets
	}
	return nil
}

type ShellSnippet struct {
	SnippetName          string   `protobuf:"bytes,1,opt,name=snippet_name,json=snippetName,proto3" json:"snippet_name,omitempty"`
	SnippetDescription   string   `protobuf:"bytes,2,opt,name=snippet_description,json=snippetDescription,proto3" json:"snippet_description,omitempty"`
	SnippetCommand       string   `protobuf:"bytes,3,opt,name=snippet_command,json=snippetCommand,proto3" json:"snippet_command,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ShellSnippet) Reset()         { *m = ShellSnippet{} }
func (m *ShellSnippet) String() string { return proto.CompactTextString(m) }
func (*ShellSnippet) ProtoMessage()    {}
func (*ShellSnippet) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd244997a2724ff9, []int{2}
}

func (m *ShellSnippet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ShellSnippet.Unmarshal(m, b)
}
func (m *ShellSnippet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ShellSnippet.Marshal(b, m, deterministic)
}
func (m *ShellSnippet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ShellSnippet.Merge(m, src)
}
func (m *ShellSnippet) XXX_Size() int {
	return xxx_messageInfo_ShellSnippet.Size(m)
}
func (m *ShellSnippet) XXX_DiscardUnknown() {
	xxx_messageInfo_ShellSnippet.DiscardUnknown(m)
}

var xxx_messageInfo_ShellSnippet proto.InternalMessageInfo

func (m *ShellSnippet) GetSnippetName() string {
	if m != nil {
		return m.SnippetName
	}
	return ""
}

func (m *ShellSnippet) GetSnippetDescription() string {
	if m != nil {
		return m.SnippetDescription
	}
	return ""
}

func (m *ShellSnippet) GetSnippetCommand() string {
	if m != nil {
		return m.SnippetCommand
	}
	return ""
}

type QuitServerRequest struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QuitServerRequest) Reset()         { *m = QuitServerRequest{} }
func (m *QuitServerRequest) String() string { return proto.CompactTextString(m) }
func (*QuitServerRequest) ProtoMessage()    {}
func (*QuitServerRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd244997a2724ff9, []int{3}
}

func (m *QuitServerRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuitServerRequest.Unmarshal(m, b)
}
func (m *QuitServerRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuitServerRequest.Marshal(b, m, deterministic)
}
func (m *QuitServerRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuitServerRequest.Merge(m, src)
}
func (m *QuitServerRequest) XXX_Size() int {
	return xxx_messageInfo_QuitServerRequest.Size(m)
}
func (m *QuitServerRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QuitServerRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QuitServerRequest proto.InternalMessageInfo

type QuitServerResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QuitServerResponse) Reset()         { *m = QuitServerResponse{} }
func (m *QuitServerResponse) String() string { return proto.CompactTextString(m) }
func (*QuitServerResponse) ProtoMessage()    {}
func (*QuitServerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_dd244997a2724ff9, []int{4}
}

func (m *QuitServerResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QuitServerResponse.Unmarshal(m, b)
}
func (m *QuitServerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QuitServerResponse.Marshal(b, m, deterministic)
}
func (m *QuitServerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QuitServerResponse.Merge(m, src)
}
func (m *QuitServerResponse) XXX_Size() int {
	return xxx_messageInfo_QuitServerResponse.Size(m)
}
func (m *QuitServerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QuitServerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QuitServerResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ShellSnippetRequest)(nil), "phil0522.taskd.ShellSnippetRequest")
	proto.RegisterType((*ShellSnippetResponse)(nil), "phil0522.taskd.ShellSnippetResponse")
	proto.RegisterType((*ShellSnippet)(nil), "phil0522.taskd.ShellSnippet")
	proto.RegisterType((*QuitServerRequest)(nil), "phil0522.taskd.QuitServerRequest")
	proto.RegisterType((*QuitServerResponse)(nil), "phil0522.taskd.QuitServerResponse")
}

func init() {
	proto.RegisterFile("proto/snippet_service.proto", fileDescriptor_dd244997a2724ff9)
}

var fileDescriptor_dd244997a2724ff9 = []byte{
	// 340 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0x4e, 0xc2, 0x40,
	0x10, 0xc6, 0xad, 0x24, 0x46, 0x06, 0xac, 0x32, 0x70, 0x20, 0xa8, 0x09, 0x54, 0x13, 0x49, 0x4c,
	0x8a, 0xa9, 0xf1, 0x09, 0xe0, 0x6c, 0x94, 0xc6, 0x93, 0x87, 0x66, 0xdd, 0x8e, 0xb2, 0x91, 0xb6,
	0xeb, 0xee, 0x42, 0xf4, 0xe0, 0x13, 0xf8, 0x66, 0x3e, 0x95, 0xa1, 0xdd, 0x0a, 0xd5, 0x84, 0xe3,
	0xfe, 0xbe, 0x6f, 0xfe, 0x7c, 0xd3, 0xc2, 0xb1, 0x54, 0x99, 0xc9, 0x46, 0x3a, 0x15, 0x52, 0x92,
	0x89, 0x34, 0xa9, 0xa5, 0xe0, 0xe4, 0xe7, 0x14, 0x5d, 0x39, 0x13, 0xf3, 0xab, 0x9b, 0x20, 0xf0,
	0x0d, 0xd3, 0xaf, 0xb1, 0xf7, 0x09, 0xed, 0x70, 0x46, 0xf3, 0x79, 0x58, 0xb8, 0xa7, 0xf4, 0xb6,
	0x20, 0x6d, 0xf0, 0x12, 0x5a, 0x7c, 0xa1, 0x14, 0xa5, 0x26, 0x8a, 0x85, 0x22, 0x6e, 0x32, 0xf5,
	0xd1, 0x75, 0xfa, 0xce, 0xb0, 0x3e, 0x3d, 0xb2, 0xc2, 0xa4, 0xe4, 0x78, 0x0a, 0xc0, 0x93, 0x38,
	0x92, 0x8a, 0x9e, 0xc5, 0x7b, 0x77, 0x37, 0x77, 0xd5, 0x79, 0x12, 0xdf, 0xe5, 0x00, 0x7b, 0xb0,
	0xcf, 0x99, 0xa1, 0x97, 0x55, 0x8b, 0x5a, 0x2e, 0xfe, 0xbe, 0xbd, 0x47, 0xe8, 0x54, 0xc7, 0x6b,
	0x99, 0xa5, 0x9a, 0x70, 0x0c, 0xae, 0x5e, 0xf1, 0xc8, 0xa6, 0xd0, 0x5d, 0xa7, 0x5f, 0x1b, 0x36,
	0x82, 0x13, 0xbf, 0xba, 0xbf, 0x5f, 0xa9, 0x3e, 0xd0, 0x1b, 0x2f, 0xed, 0x7d, 0x39, 0xd0, 0xdc,
	0xd4, 0x71, 0x00, 0xcd, 0xf2, 0x2a, 0x29, 0x4b, 0xc8, 0x06, 0x6a, 0x58, 0x76, 0xcb, 0x12, 0xc2,
	0x11, 0xb4, 0x4b, 0x4b, 0x4c, 0x9a, 0x2b, 0x21, 0x8d, 0xc8, 0x52, 0x1b, 0x0a, 0xad, 0x34, 0x59,
	0x2b, 0x78, 0x01, 0x87, 0x65, 0x01, 0xcf, 0x92, 0x84, 0xa5, 0xb1, 0x0d, 0xe9, 0x5a, 0x3c, 0x2e,
	0xa8, 0xd7, 0x86, 0xd6, 0xfd, 0x42, 0x98, 0x90, 0xd4, 0x92, 0x94, 0xbd, 0xb3, 0xd7, 0x01, 0xdc,
	0x84, 0x45, 0xfa, 0xe0, 0xdb, 0x01, 0xd7, 0xee, 0x1c, 0x16, 0x5f, 0x0f, 0x19, 0x60, 0x48, 0x4c,
	0xf1, 0x59, 0x25, 0xd0, 0xd9, 0xd6, 0x73, 0x14, 0x33, 0x7a, 0xe7, 0xdb, 0x4d, 0xc5, 0x4c, 0x6f,
	0x07, 0x1f, 0x00, 0xd6, 0xbb, 0xe0, 0xe0, 0x6f, 0xd5, 0xbf, 0xe5, 0x7b, 0xde, 0x36, 0x4b, 0xd9,
	0xf6, 0x69, 0x2f, 0xff, 0xf1, 0xae, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xe3, 0xf7, 0x13, 0x25,
	0x97, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SnippetServiceClient is the client API for SnippetService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SnippetServiceClient interface {
	SearchShellSnippet(ctx context.Context, in *ShellSnippetRequest, opts ...grpc.CallOption) (*ShellSnippetResponse, error)
	QuitServer(ctx context.Context, in *QuitServerRequest, opts ...grpc.CallOption) (*QuitServerResponse, error)
}

type snippetServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSnippetServiceClient(cc grpc.ClientConnInterface) SnippetServiceClient {
	return &snippetServiceClient{cc}
}

func (c *snippetServiceClient) SearchShellSnippet(ctx context.Context, in *ShellSnippetRequest, opts ...grpc.CallOption) (*ShellSnippetResponse, error) {
	out := new(ShellSnippetResponse)
	err := c.cc.Invoke(ctx, "/phil0522.taskd.SnippetService/SearchShellSnippet", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *snippetServiceClient) QuitServer(ctx context.Context, in *QuitServerRequest, opts ...grpc.CallOption) (*QuitServerResponse, error) {
	out := new(QuitServerResponse)
	err := c.cc.Invoke(ctx, "/phil0522.taskd.SnippetService/QuitServer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SnippetServiceServer is the server API for SnippetService service.
type SnippetServiceServer interface {
	SearchShellSnippet(context.Context, *ShellSnippetRequest) (*ShellSnippetResponse, error)
	QuitServer(context.Context, *QuitServerRequest) (*QuitServerResponse, error)
}

// UnimplementedSnippetServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSnippetServiceServer struct {
}

func (*UnimplementedSnippetServiceServer) SearchShellSnippet(ctx context.Context, req *ShellSnippetRequest) (*ShellSnippetResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchShellSnippet not implemented")
}
func (*UnimplementedSnippetServiceServer) QuitServer(ctx context.Context, req *QuitServerRequest) (*QuitServerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QuitServer not implemented")
}

func RegisterSnippetServiceServer(s *grpc.Server, srv SnippetServiceServer) {
	s.RegisterService(&_SnippetService_serviceDesc, srv)
}

func _SnippetService_SearchShellSnippet_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ShellSnippetRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnippetServiceServer).SearchShellSnippet(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/phil0522.taskd.SnippetService/SearchShellSnippet",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnippetServiceServer).SearchShellSnippet(ctx, req.(*ShellSnippetRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _SnippetService_QuitServer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QuitServerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SnippetServiceServer).QuitServer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/phil0522.taskd.SnippetService/QuitServer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SnippetServiceServer).QuitServer(ctx, req.(*QuitServerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SnippetService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "phil0522.taskd.SnippetService",
	HandlerType: (*SnippetServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SearchShellSnippet",
			Handler:    _SnippetService_SearchShellSnippet_Handler,
		},
		{
			MethodName: "QuitServer",
			Handler:    _SnippetService_QuitServer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/snippet_service.proto",
}
