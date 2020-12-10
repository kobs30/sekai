// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
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

type MsgUpsertTokenAliasResponse struct {
}

func (m *MsgUpsertTokenAliasResponse) Reset()         { *m = MsgUpsertTokenAliasResponse{} }
func (m *MsgUpsertTokenAliasResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpsertTokenAliasResponse) ProtoMessage()    {}
func (*MsgUpsertTokenAliasResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fd2153dc07d3b5c, []int{0}
}
func (m *MsgUpsertTokenAliasResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpsertTokenAliasResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpsertTokenAliasResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpsertTokenAliasResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpsertTokenAliasResponse.Merge(m, src)
}
func (m *MsgUpsertTokenAliasResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpsertTokenAliasResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpsertTokenAliasResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpsertTokenAliasResponse proto.InternalMessageInfo

type MsgUpsertTokenRateResponse struct {
}

func (m *MsgUpsertTokenRateResponse) Reset()         { *m = MsgUpsertTokenRateResponse{} }
func (m *MsgUpsertTokenRateResponse) String() string { return proto.CompactTextString(m) }
func (*MsgUpsertTokenRateResponse) ProtoMessage()    {}
func (*MsgUpsertTokenRateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fd2153dc07d3b5c, []int{1}
}
func (m *MsgUpsertTokenRateResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgUpsertTokenRateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgUpsertTokenRateResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgUpsertTokenRateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgUpsertTokenRateResponse.Merge(m, src)
}
func (m *MsgUpsertTokenRateResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgUpsertTokenRateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgUpsertTokenRateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgUpsertTokenRateResponse proto.InternalMessageInfo

type MsgProposalUpsertTokenAliasResponse struct {
	ProposalID uint64 `protobuf:"varint,1,opt,name=proposalID,proto3" json:"proposalID,omitempty"`
}

func (m *MsgProposalUpsertTokenAliasResponse) Reset()         { *m = MsgProposalUpsertTokenAliasResponse{} }
func (m *MsgProposalUpsertTokenAliasResponse) String() string { return proto.CompactTextString(m) }
func (*MsgProposalUpsertTokenAliasResponse) ProtoMessage()    {}
func (*MsgProposalUpsertTokenAliasResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fd2153dc07d3b5c, []int{2}
}
func (m *MsgProposalUpsertTokenAliasResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgProposalUpsertTokenAliasResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgProposalUpsertTokenAliasResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgProposalUpsertTokenAliasResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgProposalUpsertTokenAliasResponse.Merge(m, src)
}
func (m *MsgProposalUpsertTokenAliasResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgProposalUpsertTokenAliasResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgProposalUpsertTokenAliasResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgProposalUpsertTokenAliasResponse proto.InternalMessageInfo

func (m *MsgProposalUpsertTokenAliasResponse) GetProposalID() uint64 {
	if m != nil {
		return m.ProposalID
	}
	return 0
}

type MsgProposalUpsertTokenRatesResponse struct {
	ProposalID uint64 `protobuf:"varint,1,opt,name=proposalID,proto3" json:"proposalID,omitempty"`
}

func (m *MsgProposalUpsertTokenRatesResponse) Reset()         { *m = MsgProposalUpsertTokenRatesResponse{} }
func (m *MsgProposalUpsertTokenRatesResponse) String() string { return proto.CompactTextString(m) }
func (*MsgProposalUpsertTokenRatesResponse) ProtoMessage()    {}
func (*MsgProposalUpsertTokenRatesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0fd2153dc07d3b5c, []int{3}
}
func (m *MsgProposalUpsertTokenRatesResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgProposalUpsertTokenRatesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgProposalUpsertTokenRatesResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgProposalUpsertTokenRatesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgProposalUpsertTokenRatesResponse.Merge(m, src)
}
func (m *MsgProposalUpsertTokenRatesResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgProposalUpsertTokenRatesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgProposalUpsertTokenRatesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgProposalUpsertTokenRatesResponse proto.InternalMessageInfo

func (m *MsgProposalUpsertTokenRatesResponse) GetProposalID() uint64 {
	if m != nil {
		return m.ProposalID
	}
	return 0
}

func init() {
	proto.RegisterType((*MsgUpsertTokenAliasResponse)(nil), "kira.tokens.MsgUpsertTokenAliasResponse")
	proto.RegisterType((*MsgUpsertTokenRateResponse)(nil), "kira.tokens.MsgUpsertTokenRateResponse")
	proto.RegisterType((*MsgProposalUpsertTokenAliasResponse)(nil), "kira.tokens.MsgProposalUpsertTokenAliasResponse")
	proto.RegisterType((*MsgProposalUpsertTokenRatesResponse)(nil), "kira.tokens.MsgProposalUpsertTokenRatesResponse")
}

func init() { proto.RegisterFile("tx.proto", fileDescriptor_0fd2153dc07d3b5c) }

var fileDescriptor_0fd2153dc07d3b5c = []byte{
	// 310 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x28, 0xa9, 0xd0, 0x2b,
	0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0xce, 0xce, 0x2c, 0x4a, 0xd4, 0x2b, 0xc9, 0xcf, 0x4e, 0xcd,
	0x2b, 0x96, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x8b, 0xeb, 0x83, 0x58, 0x10, 0x25, 0x52, 0xdc,
	0x89, 0x39, 0x99, 0x89, 0xc5, 0x50, 0x0e, 0x5f, 0x41, 0x51, 0x7e, 0x41, 0x7e, 0x71, 0x62, 0x0e,
	0x94, 0xcf, 0x55, 0x94, 0x58, 0x92, 0x0a, 0x61, 0x2b, 0xc9, 0x72, 0x49, 0xfb, 0x16, 0xa7, 0x87,
	0x16, 0x14, 0xa7, 0x16, 0x95, 0x84, 0x80, 0x4c, 0x74, 0x04, 0x69, 0x0c, 0x4a, 0x2d, 0x2e, 0xc8,
	0xcf, 0x2b, 0x4e, 0x55, 0x92, 0xe1, 0x92, 0x42, 0x95, 0x0e, 0x4a, 0x2c, 0x49, 0x85, 0xcb, 0xba,
	0x72, 0x29, 0xfb, 0x16, 0xa7, 0x07, 0x40, 0x4d, 0xc7, 0x65, 0x88, 0x90, 0x1c, 0x17, 0x17, 0xcc,
	0x05, 0x9e, 0x2e, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0x2c, 0x41, 0x48, 0x22, 0xb8, 0x8d, 0x01, 0x59,
	0x46, 0xb4, 0x31, 0x46, 0x9b, 0x98, 0xb9, 0x98, 0x7d, 0x8b, 0xd3, 0x85, 0xe2, 0xb8, 0x04, 0xd0,
	0x9d, 0x22, 0xa4, 0xa0, 0x87, 0x14, 0x66, 0x7a, 0x58, 0x7c, 0x2c, 0xa5, 0x41, 0x48, 0x05, 0xdc,
	0x1d, 0xd1, 0x5c, 0xfc, 0x68, 0x6e, 0x14, 0x92, 0xc7, 0xa3, 0x19, 0xa4, 0x40, 0x4a, 0x9d, 0x80,
	0x02, 0xb8, 0xe1, 0x65, 0x5c, 0x12, 0xb8, 0xc2, 0x53, 0x08, 0xc3, 0x89, 0xb8, 0x54, 0x4a, 0x19,
	0x10, 0xab, 0x92, 0x80, 0xbd, 0xe0, 0x08, 0x20, 0xca, 0x5e, 0xb0, 0x4a, 0xa2, 0xec, 0x45, 0x89,
	0x54, 0x27, 0xa7, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71,
	0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0xd2, 0x48, 0xcf,
	0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0xf7, 0xce, 0x2c, 0x4a, 0x74, 0xce, 0x2f,
	0x4a, 0xd5, 0x2f, 0x4e, 0xcd, 0x4e, 0xcc, 0xd4, 0xaf, 0xd0, 0x87, 0xd8, 0xa0, 0x5f, 0x52, 0x59,
	0x90, 0x5a, 0x9c, 0xc4, 0x06, 0x4e, 0xca, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x02, 0x76,
	0x16, 0x5c, 0x22, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	// UpsertTokenAlias defines a method to upsert token alias
	UpsertTokenAlias(ctx context.Context, in *MsgUpsertTokenAlias, opts ...grpc.CallOption) (*MsgUpsertTokenAliasResponse, error)
	// UpsertTokenRate defines a method to upsert token rate
	UpsertTokenRate(ctx context.Context, in *MsgUpsertTokenRate, opts ...grpc.CallOption) (*MsgUpsertTokenRateResponse, error)
	// ProposalUpsertTokenAlias defines a method to create a proposal to upsert token alias
	ProposalUpsertTokenAlias(ctx context.Context, in *MsgProposalUpsertTokenAlias, opts ...grpc.CallOption) (*MsgProposalUpsertTokenAliasResponse, error)
	// ProposalUpsertTokenRates defines a method to create a proposal to upsert token rates
	ProposalUpsertTokenRates(ctx context.Context, in *MsgProposalUpsertTokenRates, opts ...grpc.CallOption) (*MsgProposalUpsertTokenRatesResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) UpsertTokenAlias(ctx context.Context, in *MsgUpsertTokenAlias, opts ...grpc.CallOption) (*MsgUpsertTokenAliasResponse, error) {
	out := new(MsgUpsertTokenAliasResponse)
	err := c.cc.Invoke(ctx, "/kira.tokens.Msg/UpsertTokenAlias", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) UpsertTokenRate(ctx context.Context, in *MsgUpsertTokenRate, opts ...grpc.CallOption) (*MsgUpsertTokenRateResponse, error) {
	out := new(MsgUpsertTokenRateResponse)
	err := c.cc.Invoke(ctx, "/kira.tokens.Msg/UpsertTokenRate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ProposalUpsertTokenAlias(ctx context.Context, in *MsgProposalUpsertTokenAlias, opts ...grpc.CallOption) (*MsgProposalUpsertTokenAliasResponse, error) {
	out := new(MsgProposalUpsertTokenAliasResponse)
	err := c.cc.Invoke(ctx, "/kira.tokens.Msg/ProposalUpsertTokenAlias", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ProposalUpsertTokenRates(ctx context.Context, in *MsgProposalUpsertTokenRates, opts ...grpc.CallOption) (*MsgProposalUpsertTokenRatesResponse, error) {
	out := new(MsgProposalUpsertTokenRatesResponse)
	err := c.cc.Invoke(ctx, "/kira.tokens.Msg/ProposalUpsertTokenRates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// UpsertTokenAlias defines a method to upsert token alias
	UpsertTokenAlias(context.Context, *MsgUpsertTokenAlias) (*MsgUpsertTokenAliasResponse, error)
	// UpsertTokenRate defines a method to upsert token rate
	UpsertTokenRate(context.Context, *MsgUpsertTokenRate) (*MsgUpsertTokenRateResponse, error)
	// ProposalUpsertTokenAlias defines a method to create a proposal to upsert token alias
	ProposalUpsertTokenAlias(context.Context, *MsgProposalUpsertTokenAlias) (*MsgProposalUpsertTokenAliasResponse, error)
	// ProposalUpsertTokenRates defines a method to create a proposal to upsert token rates
	ProposalUpsertTokenRates(context.Context, *MsgProposalUpsertTokenRates) (*MsgProposalUpsertTokenRatesResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) UpsertTokenAlias(ctx context.Context, req *MsgUpsertTokenAlias) (*MsgUpsertTokenAliasResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertTokenAlias not implemented")
}
func (*UnimplementedMsgServer) UpsertTokenRate(ctx context.Context, req *MsgUpsertTokenRate) (*MsgUpsertTokenRateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpsertTokenRate not implemented")
}
func (*UnimplementedMsgServer) ProposalUpsertTokenAlias(ctx context.Context, req *MsgProposalUpsertTokenAlias) (*MsgProposalUpsertTokenAliasResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProposalUpsertTokenAlias not implemented")
}
func (*UnimplementedMsgServer) ProposalUpsertTokenRates(ctx context.Context, req *MsgProposalUpsertTokenRates) (*MsgProposalUpsertTokenRatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProposalUpsertTokenRates not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_UpsertTokenAlias_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpsertTokenAlias)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpsertTokenAlias(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kira.tokens.Msg/UpsertTokenAlias",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpsertTokenAlias(ctx, req.(*MsgUpsertTokenAlias))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_UpsertTokenRate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgUpsertTokenRate)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).UpsertTokenRate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kira.tokens.Msg/UpsertTokenRate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).UpsertTokenRate(ctx, req.(*MsgUpsertTokenRate))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ProposalUpsertTokenAlias_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgProposalUpsertTokenAlias)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ProposalUpsertTokenAlias(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kira.tokens.Msg/ProposalUpsertTokenAlias",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ProposalUpsertTokenAlias(ctx, req.(*MsgProposalUpsertTokenAlias))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ProposalUpsertTokenRates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgProposalUpsertTokenRates)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ProposalUpsertTokenRates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kira.tokens.Msg/ProposalUpsertTokenRates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ProposalUpsertTokenRates(ctx, req.(*MsgProposalUpsertTokenRates))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kira.tokens.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpsertTokenAlias",
			Handler:    _Msg_UpsertTokenAlias_Handler,
		},
		{
			MethodName: "UpsertTokenRate",
			Handler:    _Msg_UpsertTokenRate_Handler,
		},
		{
			MethodName: "ProposalUpsertTokenAlias",
			Handler:    _Msg_ProposalUpsertTokenAlias_Handler,
		},
		{
			MethodName: "ProposalUpsertTokenRates",
			Handler:    _Msg_ProposalUpsertTokenRates_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tx.proto",
}

func (m *MsgUpsertTokenAliasResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpsertTokenAliasResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpsertTokenAliasResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgUpsertTokenRateResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgUpsertTokenRateResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgUpsertTokenRateResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgProposalUpsertTokenAliasResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgProposalUpsertTokenAliasResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgProposalUpsertTokenAliasResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ProposalID != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.ProposalID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *MsgProposalUpsertTokenRatesResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgProposalUpsertTokenRatesResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgProposalUpsertTokenRatesResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.ProposalID != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.ProposalID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgUpsertTokenAliasResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgUpsertTokenRateResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgProposalUpsertTokenAliasResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProposalID != 0 {
		n += 1 + sovTx(uint64(m.ProposalID))
	}
	return n
}

func (m *MsgProposalUpsertTokenRatesResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ProposalID != 0 {
		n += 1 + sovTx(uint64(m.ProposalID))
	}
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgUpsertTokenAliasResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUpsertTokenAliasResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpsertTokenAliasResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgUpsertTokenRateResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgUpsertTokenRateResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgUpsertTokenRateResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgProposalUpsertTokenAliasResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgProposalUpsertTokenAliasResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgProposalUpsertTokenAliasResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalID", wireType)
			}
			m.ProposalID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProposalID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgProposalUpsertTokenRatesResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgProposalUpsertTokenRatesResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgProposalUpsertTokenRatesResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProposalID", wireType)
			}
			m.ProposalID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ProposalID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)