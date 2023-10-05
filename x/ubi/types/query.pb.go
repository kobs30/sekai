// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kira/ubi/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/KiraCore/sekai/x/gov/types"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type QueryUBIRecordsRequest struct {
}

func (m *QueryUBIRecordsRequest) Reset()         { *m = QueryUBIRecordsRequest{} }
func (m *QueryUBIRecordsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryUBIRecordsRequest) ProtoMessage()    {}
func (*QueryUBIRecordsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce4683e7ea12df2a, []int{0}
}
func (m *QueryUBIRecordsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryUBIRecordsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryUBIRecordsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryUBIRecordsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryUBIRecordsRequest.Merge(m, src)
}
func (m *QueryUBIRecordsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryUBIRecordsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryUBIRecordsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryUBIRecordsRequest proto.InternalMessageInfo

type QueryUBIRecordsResponse struct {
	Records []UBIRecord `protobuf:"bytes,1,rep,name=records,proto3" json:"records"`
}

func (m *QueryUBIRecordsResponse) Reset()         { *m = QueryUBIRecordsResponse{} }
func (m *QueryUBIRecordsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryUBIRecordsResponse) ProtoMessage()    {}
func (*QueryUBIRecordsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce4683e7ea12df2a, []int{1}
}
func (m *QueryUBIRecordsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryUBIRecordsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryUBIRecordsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryUBIRecordsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryUBIRecordsResponse.Merge(m, src)
}
func (m *QueryUBIRecordsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryUBIRecordsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryUBIRecordsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryUBIRecordsResponse proto.InternalMessageInfo

func (m *QueryUBIRecordsResponse) GetRecords() []UBIRecord {
	if m != nil {
		return m.Records
	}
	return nil
}

type QueryUBIRecordByNameRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (m *QueryUBIRecordByNameRequest) Reset()         { *m = QueryUBIRecordByNameRequest{} }
func (m *QueryUBIRecordByNameRequest) String() string { return proto.CompactTextString(m) }
func (*QueryUBIRecordByNameRequest) ProtoMessage()    {}
func (*QueryUBIRecordByNameRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce4683e7ea12df2a, []int{2}
}
func (m *QueryUBIRecordByNameRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryUBIRecordByNameRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryUBIRecordByNameRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryUBIRecordByNameRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryUBIRecordByNameRequest.Merge(m, src)
}
func (m *QueryUBIRecordByNameRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryUBIRecordByNameRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryUBIRecordByNameRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryUBIRecordByNameRequest proto.InternalMessageInfo

func (m *QueryUBIRecordByNameRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type QueryUBIRecordByNameResponse struct {
	Record *UBIRecord `protobuf:"bytes,1,opt,name=record,proto3" json:"record,omitempty"`
}

func (m *QueryUBIRecordByNameResponse) Reset()         { *m = QueryUBIRecordByNameResponse{} }
func (m *QueryUBIRecordByNameResponse) String() string { return proto.CompactTextString(m) }
func (*QueryUBIRecordByNameResponse) ProtoMessage()    {}
func (*QueryUBIRecordByNameResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ce4683e7ea12df2a, []int{3}
}
func (m *QueryUBIRecordByNameResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryUBIRecordByNameResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryUBIRecordByNameResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryUBIRecordByNameResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryUBIRecordByNameResponse.Merge(m, src)
}
func (m *QueryUBIRecordByNameResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryUBIRecordByNameResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryUBIRecordByNameResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryUBIRecordByNameResponse proto.InternalMessageInfo

func (m *QueryUBIRecordByNameResponse) GetRecord() *UBIRecord {
	if m != nil {
		return m.Record
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryUBIRecordsRequest)(nil), "kira.ubi.QueryUBIRecordsRequest")
	proto.RegisterType((*QueryUBIRecordsResponse)(nil), "kira.ubi.QueryUBIRecordsResponse")
	proto.RegisterType((*QueryUBIRecordByNameRequest)(nil), "kira.ubi.QueryUBIRecordByNameRequest")
	proto.RegisterType((*QueryUBIRecordByNameResponse)(nil), "kira.ubi.QueryUBIRecordByNameResponse")
}

func init() { proto.RegisterFile("kira/ubi/query.proto", fileDescriptor_ce4683e7ea12df2a) }

var fileDescriptor_ce4683e7ea12df2a = []byte{
	// 407 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x41, 0x8b, 0xd3, 0x40,
	0x18, 0x4d, 0x6a, 0xad, 0x3a, 0x1e, 0x84, 0xb1, 0xda, 0x92, 0xb6, 0xb1, 0x46, 0x2a, 0x45, 0x21,
	0x43, 0xdb, 0x1f, 0x20, 0xc4, 0x93, 0x14, 0x0a, 0x06, 0xbc, 0x78, 0x91, 0x49, 0x1d, 0xe2, 0xd0,
	0x26, 0x5f, 0x3a, 0x93, 0x14, 0x8b, 0x78, 0xf1, 0xa8, 0x17, 0xc1, 0x3f, 0xd5, 0x63, 0xc1, 0x8b,
	0x27, 0x59, 0xda, 0xfd, 0x21, 0x4b, 0x26, 0xd3, 0xee, 0x76, 0x37, 0xbb, 0x7b, 0xfb, 0x78, 0xdf,
	0x7b, 0x6f, 0xde, 0xf7, 0x12, 0x54, 0x9f, 0x71, 0x41, 0x49, 0x16, 0x70, 0xb2, 0xc8, 0x98, 0x58,
	0xb9, 0x89, 0x80, 0x14, 0xf0, 0xfd, 0x1c, 0x75, 0xb3, 0x80, 0x5b, 0xaf, 0xa6, 0x20, 0x23, 0x90,
	0x24, 0xa0, 0x92, 0x15, 0x14, 0xb2, 0x1c, 0x04, 0x2c, 0xa5, 0x03, 0x92, 0xd0, 0x90, 0xc7, 0x34,
	0xe5, 0x10, 0x17, 0x2a, 0xab, 0x1e, 0x42, 0x08, 0x6a, 0x24, 0xf9, 0xa4, 0xd1, 0x76, 0x08, 0x10,
	0xce, 0x19, 0xa1, 0x09, 0x27, 0x34, 0x8e, 0x21, 0x55, 0x12, 0xa9, 0xb7, 0xf8, 0xf0, 0x7e, 0x16,
	0x70, 0x8d, 0x35, 0x14, 0x16, 0xc2, 0x92, 0x24, 0x02, 0x12, 0x90, 0x74, 0x5e, 0x2c, 0x9c, 0x26,
	0x7a, 0xfa, 0x3e, 0x8f, 0xf0, 0xc1, 0x7b, 0xe7, 0xb3, 0x29, 0x88, 0xcf, 0xd2, 0x67, 0x8b, 0x8c,
	0xc9, 0xd4, 0x99, 0xa0, 0xc6, 0x95, 0x8d, 0x4c, 0x20, 0x96, 0x0c, 0x8f, 0xd0, 0x3d, 0x51, 0x40,
	0x4d, 0xb3, 0x7b, 0xa7, 0xff, 0x70, 0xf8, 0xd8, 0xdd, 0x5f, 0xe7, 0x1e, 0xe8, 0x5e, 0x75, 0xfd,
	0xff, 0x99, 0xe1, 0xef, 0x99, 0xce, 0x00, 0xb5, 0x8e, 0xfd, 0xbc, 0xd5, 0x84, 0x46, 0x4c, 0x3f,
	0x87, 0x31, 0xaa, 0xc6, 0x34, 0x62, 0x4d, 0xb3, 0x6b, 0xf6, 0x1f, 0xf8, 0x6a, 0x76, 0xc6, 0xa8,
	0x5d, 0x2e, 0xd1, 0x39, 0x5e, 0xa3, 0x5a, 0xe1, 0xae, 0x54, 0xe5, 0x31, 0x7c, 0x4d, 0x19, 0xfe,
	0xaa, 0xa0, 0xbb, 0xca, 0x0d, 0x67, 0xe8, 0xd1, 0xa5, 0xcb, 0x70, 0xf7, 0x5c, 0x59, 0x5e, 0x87,
	0xf5, 0xfc, 0x06, 0x46, 0x11, 0xc7, 0xe9, 0xfc, 0xf8, 0x7b, 0xfa, 0xa7, 0xd2, 0xc0, 0x4f, 0xc8,
	0xc5, 0x2f, 0xf0, 0x49, 0x17, 0x80, 0x7f, 0x9a, 0xa8, 0x5e, 0x76, 0x0e, 0xee, 0x5d, 0x67, 0x7d,
	0xd4, 0x90, 0xf5, 0xf2, 0x36, 0x9a, 0x8e, 0xf1, 0x42, 0xc5, 0xe8, 0xe0, 0x56, 0x59, 0x0c, 0xf2,
	0x2d, 0x6f, 0xf6, 0xbb, 0xf7, 0x66, 0xbd, 0xb5, 0xcd, 0xcd, 0xd6, 0x36, 0x4f, 0xb6, 0xb6, 0xf9,
	0x7b, 0x67, 0x1b, 0x9b, 0x9d, 0x6d, 0xfc, 0xdb, 0xd9, 0xc6, 0xc7, 0x5e, 0xc8, 0xd3, 0x2f, 0x59,
	0xe0, 0x4e, 0x21, 0x22, 0x63, 0x2e, 0xe8, 0x5b, 0x10, 0x8c, 0x48, 0x36, 0xa3, 0x9c, 0x7c, 0x55,
	0x66, 0xe9, 0x2a, 0x61, 0x32, 0xa8, 0xa9, 0xff, 0x67, 0x74, 0x16, 0x00, 0x00, 0xff, 0xff, 0x90,
	0x01, 0xfc, 0xac, 0xee, 0x02, 0x00, 0x00,
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
	// QueryUBIRecords - query names of all UBIRecords
	QueryUBIRecords(ctx context.Context, in *QueryUBIRecordsRequest, opts ...grpc.CallOption) (*QueryUBIRecordsResponse, error)
	// QueryUBIRecordByName - query specific UBIRecord by name
	QueryUBIRecordByName(ctx context.Context, in *QueryUBIRecordByNameRequest, opts ...grpc.CallOption) (*QueryUBIRecordByNameResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) QueryUBIRecords(ctx context.Context, in *QueryUBIRecordsRequest, opts ...grpc.CallOption) (*QueryUBIRecordsResponse, error) {
	out := new(QueryUBIRecordsResponse)
	err := c.cc.Invoke(ctx, "/kira.ubi.Query/QueryUBIRecords", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) QueryUBIRecordByName(ctx context.Context, in *QueryUBIRecordByNameRequest, opts ...grpc.CallOption) (*QueryUBIRecordByNameResponse, error) {
	out := new(QueryUBIRecordByNameResponse)
	err := c.cc.Invoke(ctx, "/kira.ubi.Query/QueryUBIRecordByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// QueryUBIRecords - query names of all UBIRecords
	QueryUBIRecords(context.Context, *QueryUBIRecordsRequest) (*QueryUBIRecordsResponse, error)
	// QueryUBIRecordByName - query specific UBIRecord by name
	QueryUBIRecordByName(context.Context, *QueryUBIRecordByNameRequest) (*QueryUBIRecordByNameResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) QueryUBIRecords(ctx context.Context, req *QueryUBIRecordsRequest) (*QueryUBIRecordsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryUBIRecords not implemented")
}
func (*UnimplementedQueryServer) QueryUBIRecordByName(ctx context.Context, req *QueryUBIRecordByNameRequest) (*QueryUBIRecordByNameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method QueryUBIRecordByName not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_QueryUBIRecords_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryUBIRecordsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryUBIRecords(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kira.ubi.Query/QueryUBIRecords",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryUBIRecords(ctx, req.(*QueryUBIRecordsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_QueryUBIRecordByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryUBIRecordByNameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).QueryUBIRecordByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/kira.ubi.Query/QueryUBIRecordByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).QueryUBIRecordByName(ctx, req.(*QueryUBIRecordByNameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "kira.ubi.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "QueryUBIRecords",
			Handler:    _Query_QueryUBIRecords_Handler,
		},
		{
			MethodName: "QueryUBIRecordByName",
			Handler:    _Query_QueryUBIRecordByName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "kira/ubi/query.proto",
}

func (m *QueryUBIRecordsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryUBIRecordsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryUBIRecordsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryUBIRecordsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryUBIRecordsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryUBIRecordsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Records) > 0 {
		for iNdEx := len(m.Records) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Records[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryUBIRecordByNameRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryUBIRecordByNameRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryUBIRecordByNameRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryUBIRecordByNameResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryUBIRecordByNameResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryUBIRecordByNameResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Record != nil {
		{
			size, err := m.Record.MarshalToSizedBuffer(dAtA[:i])
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
func (m *QueryUBIRecordsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryUBIRecordsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Records) > 0 {
		for _, e := range m.Records {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *QueryUBIRecordByNameRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryUBIRecordByNameResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Record != nil {
		l = m.Record.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryUBIRecordsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryUBIRecordsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryUBIRecordsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
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
func (m *QueryUBIRecordsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryUBIRecordsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryUBIRecordsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Records", wireType)
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
			m.Records = append(m.Records, UBIRecord{})
			if err := m.Records[len(m.Records)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
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
func (m *QueryUBIRecordByNameRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryUBIRecordByNameRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryUBIRecordByNameRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
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
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
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
func (m *QueryUBIRecordByNameResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryUBIRecordByNameResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryUBIRecordByNameResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Record", wireType)
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
			if m.Record == nil {
				m.Record = &UBIRecord{}
			}
			if err := m.Record.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
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
