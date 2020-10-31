// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: query.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type TokenAliasRequest struct {
	Symbol string `protobuf:"bytes,1,opt,name=symbol,proto3" json:"symbol,omitempty"`
}

func (m *TokenAliasRequest) Reset()         { *m = TokenAliasRequest{} }
func (m *TokenAliasRequest) String() string { return proto.CompactTextString(m) }
func (*TokenAliasRequest) ProtoMessage()    {}
func (*TokenAliasRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{0}
}
func (m *TokenAliasRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TokenAliasRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TokenAliasRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TokenAliasRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenAliasRequest.Merge(m, src)
}
func (m *TokenAliasRequest) XXX_Size() int {
	return m.Size()
}
func (m *TokenAliasRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenAliasRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TokenAliasRequest proto.InternalMessageInfo

func (m *TokenAliasRequest) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

type TokenAliasResponse struct {
	Data *TokenAlias `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *TokenAliasResponse) Reset()         { *m = TokenAliasResponse{} }
func (m *TokenAliasResponse) String() string { return proto.CompactTextString(m) }
func (*TokenAliasResponse) ProtoMessage()    {}
func (*TokenAliasResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{1}
}
func (m *TokenAliasResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TokenAliasResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TokenAliasResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TokenAliasResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenAliasResponse.Merge(m, src)
}
func (m *TokenAliasResponse) XXX_Size() int {
	return m.Size()
}
func (m *TokenAliasResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenAliasResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TokenAliasResponse proto.InternalMessageInfo

func (m *TokenAliasResponse) GetData() *TokenAlias {
	if m != nil {
		return m.Data
	}
	return nil
}

type AllTokenAliasesRequest struct {
}

func (m *AllTokenAliasesRequest) Reset()         { *m = AllTokenAliasesRequest{} }
func (m *AllTokenAliasesRequest) String() string { return proto.CompactTextString(m) }
func (*AllTokenAliasesRequest) ProtoMessage()    {}
func (*AllTokenAliasesRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{2}
}
func (m *AllTokenAliasesRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AllTokenAliasesRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AllTokenAliasesRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AllTokenAliasesRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllTokenAliasesRequest.Merge(m, src)
}
func (m *AllTokenAliasesRequest) XXX_Size() int {
	return m.Size()
}
func (m *AllTokenAliasesRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AllTokenAliasesRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AllTokenAliasesRequest proto.InternalMessageInfo

type AllTokenAliasesResponse struct {
	Data []*TokenAlias `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
}

func (m *AllTokenAliasesResponse) Reset()         { *m = AllTokenAliasesResponse{} }
func (m *AllTokenAliasesResponse) String() string { return proto.CompactTextString(m) }
func (*AllTokenAliasesResponse) ProtoMessage()    {}
func (*AllTokenAliasesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{3}
}
func (m *AllTokenAliasesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AllTokenAliasesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AllTokenAliasesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AllTokenAliasesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllTokenAliasesResponse.Merge(m, src)
}
func (m *AllTokenAliasesResponse) XXX_Size() int {
	return m.Size()
}
func (m *AllTokenAliasesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AllTokenAliasesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AllTokenAliasesResponse proto.InternalMessageInfo

func (m *AllTokenAliasesResponse) GetData() []*TokenAlias {
	if m != nil {
		return m.Data
	}
	return nil
}

type TokenAliasesByDenomRequest struct {
	Denoms []string `protobuf:"bytes,1,rep,name=denoms,proto3" json:"denoms,omitempty"`
}

func (m *TokenAliasesByDenomRequest) Reset()         { *m = TokenAliasesByDenomRequest{} }
func (m *TokenAliasesByDenomRequest) String() string { return proto.CompactTextString(m) }
func (*TokenAliasesByDenomRequest) ProtoMessage()    {}
func (*TokenAliasesByDenomRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{4}
}
func (m *TokenAliasesByDenomRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TokenAliasesByDenomRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TokenAliasesByDenomRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TokenAliasesByDenomRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenAliasesByDenomRequest.Merge(m, src)
}
func (m *TokenAliasesByDenomRequest) XXX_Size() int {
	return m.Size()
}
func (m *TokenAliasesByDenomRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenAliasesByDenomRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TokenAliasesByDenomRequest proto.InternalMessageInfo

func (m *TokenAliasesByDenomRequest) GetDenoms() []string {
	if m != nil {
		return m.Denoms
	}
	return nil
}

type TokenAliasesByDenomResponse struct {
	Data map[string]*TokenAlias `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *TokenAliasesByDenomResponse) Reset()         { *m = TokenAliasesByDenomResponse{} }
func (m *TokenAliasesByDenomResponse) String() string { return proto.CompactTextString(m) }
func (*TokenAliasesByDenomResponse) ProtoMessage()    {}
func (*TokenAliasesByDenomResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5c6ac9b241082464, []int{5}
}
func (m *TokenAliasesByDenomResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TokenAliasesByDenomResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TokenAliasesByDenomResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TokenAliasesByDenomResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenAliasesByDenomResponse.Merge(m, src)
}
func (m *TokenAliasesByDenomResponse) XXX_Size() int {
	return m.Size()
}
func (m *TokenAliasesByDenomResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenAliasesByDenomResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TokenAliasesByDenomResponse proto.InternalMessageInfo

func (m *TokenAliasesByDenomResponse) GetData() map[string]*TokenAlias {
	if m != nil {
		return m.Data
	}
	return nil
}

func init() {
	proto.RegisterType((*TokenAliasRequest)(nil), "kira.tokens.TokenAliasRequest")
	proto.RegisterType((*TokenAliasResponse)(nil), "kira.tokens.TokenAliasResponse")
	proto.RegisterType((*AllTokenAliasesRequest)(nil), "kira.tokens.AllTokenAliasesRequest")
	proto.RegisterType((*AllTokenAliasesResponse)(nil), "kira.tokens.AllTokenAliasesResponse")
	proto.RegisterType((*TokenAliasesByDenomRequest)(nil), "kira.tokens.TokenAliasesByDenomRequest")
	proto.RegisterType((*TokenAliasesByDenomResponse)(nil), "kira.tokens.TokenAliasesByDenomResponse")
	proto.RegisterMapType((map[string]*TokenAlias)(nil), "kira.tokens.TokenAliasesByDenomResponse.DataEntry")
}

func init() { proto.RegisterFile("query.proto", fileDescriptor_5c6ac9b241082464) }

var fileDescriptor_5c6ac9b241082464 = []byte{
	// 386 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4d, 0x4f, 0xea, 0x50,
	0x10, 0x6d, 0xe1, 0x41, 0xc2, 0x34, 0x2f, 0x79, 0xef, 0x2e, 0x80, 0xf4, 0x25, 0x7d, 0xa6, 0x9a,
	0x48, 0x42, 0x6c, 0x13, 0x74, 0x61, 0xdc, 0x81, 0x08, 0x0b, 0x37, 0xda, 0xb8, 0x72, 0x77, 0x91,
	0x89, 0x36, 0xfd, 0xb8, 0xd0, 0x7b, 0x6b, 0xec, 0xbf, 0xf0, 0x8f, 0xb8, 0xf6, 0x2f, 0xb8, 0x64,
	0xe9, 0xd2, 0xc0, 0x1f, 0x31, 0xfd, 0x10, 0x8a, 0x7c, 0xb9, 0xeb, 0xcc, 0x9c, 0x99, 0x73, 0x7a,
	0x4e, 0x2e, 0x28, 0xe3, 0x10, 0x83, 0xc8, 0x18, 0x05, 0x4c, 0x30, 0xa2, 0x38, 0x76, 0x40, 0x0d,
	0xc1, 0x1c, 0xf4, 0xb9, 0xaa, 0x50, 0xd7, 0xa6, 0x3c, 0x9d, 0xe8, 0x4d, 0xf8, 0x7b, 0x13, 0xb7,
	0xdb, 0x71, 0xcf, 0xc2, 0x71, 0x88, 0x5c, 0x90, 0x2a, 0x94, 0x79, 0xe4, 0x0d, 0x98, 0x5b, 0x97,
	0xf7, 0xe4, 0x46, 0xc5, 0xca, 0x2a, 0xbd, 0x0d, 0x24, 0x0f, 0xe6, 0x23, 0xe6, 0x73, 0x24, 0x4d,
	0xf8, 0x35, 0xa4, 0x82, 0x26, 0x58, 0xa5, 0x55, 0x33, 0x72, 0x5c, 0x46, 0x0e, 0x9e, 0x80, 0xf4,
	0x3a, 0x54, 0xdb, 0xae, 0xbb, 0x68, 0xe3, 0x17, 0xa9, 0xde, 0x83, 0xda, 0xca, 0x64, 0x85, 0xa1,
	0xb8, 0x9b, 0xe1, 0x04, 0xd4, 0xfc, 0x91, 0x4e, 0xd4, 0x45, 0x9f, 0x79, 0xb9, 0x5f, 0x1b, 0xc6,
	0x35, 0x4f, 0x8e, 0x55, 0xac, 0xac, 0xd2, 0x5f, 0x65, 0xf8, 0xb7, 0x76, 0x2d, 0x93, 0xd0, 0x5b,
	0x92, 0xd0, 0xda, 0x20, 0x61, 0x65, 0xcf, 0xe8, 0x52, 0x41, 0x2f, 0x7c, 0x11, 0x44, 0xa9, 0x3a,
	0xf5, 0x0a, 0x2a, 0xf3, 0x16, 0xf9, 0x03, 0x45, 0x07, 0xa3, 0xcc, 0xe4, 0xf8, 0x93, 0x1c, 0x41,
	0xe9, 0x91, 0xba, 0x21, 0xd6, 0x0b, 0xdb, 0xcd, 0x4c, 0x51, 0x67, 0x85, 0x53, 0xb9, 0xf5, 0x52,
	0x80, 0xd2, 0x75, 0x9c, 0x35, 0xb1, 0xe0, 0x77, 0x1f, 0xc5, 0x02, 0x45, 0xb4, 0x4d, 0xeb, 0xa9,
	0x19, 0xea, 0xff, 0x8d, 0xf3, 0x54, 0xbd, 0x2e, 0x11, 0x0a, 0xa4, 0x8f, 0xe2, 0x5b, 0x30, 0x64,
	0x7f, 0x69, 0x71, 0x7d, 0xa0, 0xea, 0xc1, 0x76, 0xd0, 0x9c, 0xc2, 0x83, 0xea, 0x92, 0xec, 0xb9,
	0x89, 0xe4, 0x70, 0xb7, 0xcd, 0x29, 0x55, 0xe3, 0xa7, 0x79, 0xe8, 0x52, 0xa7, 0xf3, 0x36, 0xd5,
	0xe4, 0xc9, 0x54, 0x93, 0x3f, 0xa6, 0x9a, 0xfc, 0x3c, 0xd3, 0xa4, 0xc9, 0x4c, 0x93, 0xde, 0x67,
	0x9a, 0x74, 0xdb, 0xb8, 0xb7, 0xc5, 0x43, 0x38, 0x30, 0xee, 0x98, 0x67, 0x5e, 0xda, 0x01, 0x3d,
	0x67, 0x01, 0x9a, 0x1c, 0x1d, 0x6a, 0x9b, 0x4f, 0x66, 0x7a, 0xdb, 0x14, 0xd1, 0x08, 0xf9, 0xa0,
	0x9c, 0x3c, 0x9e, 0xe3, 0xcf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x21, 0x12, 0x7a, 0x04, 0x65, 0x03,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Returns the token alias
	GetTokenAlias(ctx context.Context, in *TokenAliasRequest, opts ...grpc.CallOption) (*TokenAliasResponse, error)
	GetAllTokenAliases(ctx context.Context, in *AllTokenAliasesRequest, opts ...grpc.CallOption) (*AllTokenAliasesResponse, error)
	GetTokenAliasesByDenom(ctx context.Context, in *TokenAliasesByDenomRequest, opts ...grpc.CallOption) (*TokenAliasesByDenomResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) GetTokenAlias(ctx context.Context, in *TokenAliasRequest, opts ...grpc.CallOption) (*TokenAliasResponse, error) {
	out := new(TokenAliasResponse)
	err := c.cc.Invoke(ctx, "/kira.tokens.Query/GetTokenAlias", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetAllTokenAliases(ctx context.Context, in *AllTokenAliasesRequest, opts ...grpc.CallOption) (*AllTokenAliasesResponse, error) {
	out := new(AllTokenAliasesResponse)
	err := c.cc.Invoke(ctx, "/kira.tokens.Query/GetAllTokenAliases", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetTokenAliasesByDenom(ctx context.Context, in *TokenAliasesByDenomRequest, opts ...grpc.CallOption) (*TokenAliasesByDenomResponse, error) {
	out := new(TokenAliasesByDenomResponse)
	err := c.cc.Invoke(ctx, "/kira.tokens.Query/GetTokenAliasesByDenom", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Returns the token alias
	GetTokenAlias(context.Context, *TokenAliasRequest) (*TokenAliasResponse, error)
	GetAllTokenAliases(context.Context, *AllTokenAliasesRequest) (*AllTokenAliasesResponse, error)
	GetTokenAliasesByDenom(context.Context, *TokenAliasesByDenomRequest) (*TokenAliasesByDenomResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) GetTokenAlias(ctx context.Context, req *TokenAliasRequest) (*TokenAliasResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTokenAlias not implemented")
}
func (*UnimplementedQueryServer) GetAllTokenAliases(ctx context.Context, req *AllTokenAliasesRequest) (*AllTokenAliasesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllTokenAliases not implemented")
}
func (*UnimplementedQueryServer) GetTokenAliasesByDenom(ctx context.Context, req *TokenAliasesByDenomRequest) (*TokenAliasesByDenomResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTokenAliasesByDenom not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_GetTokenAlias_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenAliasRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetTokenAlias(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kira.tokens.Query/GetTokenAlias",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetTokenAlias(ctx, req.(*TokenAliasRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetAllTokenAliases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AllTokenAliasesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetAllTokenAliases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kira.tokens.Query/GetAllTokenAliases",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetAllTokenAliases(ctx, req.(*AllTokenAliasesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetTokenAliasesByDenom_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TokenAliasesByDenomRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetTokenAliasesByDenom(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kira.tokens.Query/GetTokenAliasesByDenom",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetTokenAliasesByDenom(ctx, req.(*TokenAliasesByDenomRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kira.tokens.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTokenAlias",
			Handler:    _Query_GetTokenAlias_Handler,
		},
		{
			MethodName: "GetAllTokenAliases",
			Handler:    _Query_GetAllTokenAliases_Handler,
		},
		{
			MethodName: "GetTokenAliasesByDenom",
			Handler:    _Query_GetTokenAliasesByDenom_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "query.proto",
}

func (m *TokenAliasRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenAliasRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TokenAliasRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *TokenAliasResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenAliasResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TokenAliasResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Data != nil {
		{
			size, err := m.Data.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AllTokenAliasesRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AllTokenAliasesRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AllTokenAliasesRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *AllTokenAliasesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AllTokenAliasesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AllTokenAliasesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		for iNdEx := len(m.Data) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Data[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *TokenAliasesByDenomRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenAliasesByDenomRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TokenAliasesByDenomRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Denoms) > 0 {
		for iNdEx := len(m.Denoms) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Denoms[iNdEx])
			copy(dAtA[i:], m.Denoms[iNdEx])
			i = encodeVarintQuery(dAtA, i, uint64(len(m.Denoms[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *TokenAliasesByDenomResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenAliasesByDenomResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TokenAliasesByDenomResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Data) > 0 {
		for k := range m.Data {
			v := m.Data[k]
			baseI := i
			if v != nil {
				{
					size, err := v.MarshalToSizedBuffer(dAtA[:i])
					if err != nil {
						return 0, err
					}
					i -= size
					i = encodeVarintQuery(dAtA, i, uint64(size))
				}
				i--
				dAtA[i] = 0x12
			}
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintQuery(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintQuery(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TokenAliasRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *TokenAliasResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *AllTokenAliasesRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *AllTokenAliasesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Data) > 0 {
		for _, e := range m.Data {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *TokenAliasesByDenomRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Denoms) > 0 {
		for _, s := range m.Denoms {
			l = len(s)
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *TokenAliasesByDenomResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Data) > 0 {
		for k, v := range m.Data {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovQuery(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovQuery(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovQuery(uint64(mapEntrySize))
		}
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TokenAliasRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TokenAliasRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenAliasRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TokenAliasResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TokenAliasResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenAliasResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &TokenAlias{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AllTokenAliasesRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AllTokenAliasesRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AllTokenAliasesRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *AllTokenAliasesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: AllTokenAliasesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AllTokenAliasesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = append(m.Data, &TokenAlias{})
			if err := m.Data[len(m.Data)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TokenAliasesByDenomRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TokenAliasesByDenomRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenAliasesByDenomRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denoms", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denoms = append(m.Denoms, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TokenAliasesByDenomResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TokenAliasesByDenomResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenAliasesByDenomResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = make(map[string]*TokenAlias)
			}
			var mapkey string
			var mapvalue *TokenAlias
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowQuery
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					wire |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				fieldNum := int32(wire >> 3)
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowQuery
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthQuery
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthQuery
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowQuery
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthQuery
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthQuery
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &TokenAlias{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipQuery(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthQuery
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Data[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)