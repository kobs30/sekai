// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kira/gov/allowed_messages.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// defines allowed messages by network status, we only use this for poor network where
// number of validators is less than min_validators network property
type AllowedMessages struct {
	Messages []string `protobuf:"bytes,1,rep,name=messages,proto3" json:"messages,omitempty"`
}

func (m *AllowedMessages) Reset()         { *m = AllowedMessages{} }
func (m *AllowedMessages) String() string { return proto.CompactTextString(m) }
func (*AllowedMessages) ProtoMessage()    {}
func (*AllowedMessages) Descriptor() ([]byte, []int) {
	return fileDescriptor_c40df6bae73197f9, []int{0}
}
func (m *AllowedMessages) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AllowedMessages) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AllowedMessages.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AllowedMessages) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AllowedMessages.Merge(m, src)
}
func (m *AllowedMessages) XXX_Size() int {
	return m.Size()
}
func (m *AllowedMessages) XXX_DiscardUnknown() {
	xxx_messageInfo_AllowedMessages.DiscardUnknown(m)
}

var xxx_messageInfo_AllowedMessages proto.InternalMessageInfo

func (m *AllowedMessages) GetMessages() []string {
	if m != nil {
		return m.Messages
	}
	return nil
}

func init() {
	proto.RegisterType((*AllowedMessages)(nil), "kira.gov.AllowedMessages")
}

func init() { proto.RegisterFile("kira/gov/allowed_messages.proto", fileDescriptor_c40df6bae73197f9) }

var fileDescriptor_c40df6bae73197f9 = []byte{
	// 171 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xcf, 0xce, 0x2c, 0x4a,
	0xd4, 0x4f, 0xcf, 0x2f, 0xd3, 0x4f, 0xcc, 0xc9, 0xc9, 0x2f, 0x4f, 0x4d, 0x89, 0xcf, 0x4d, 0x2d,
	0x2e, 0x4e, 0x4c, 0x4f, 0x2d, 0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x00, 0x29, 0xd0,
	0x4b, 0xcf, 0x2f, 0x93, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x0b, 0xea, 0x83, 0x58, 0x10, 0x79,
	0x25, 0x5d, 0x2e, 0x7e, 0x47, 0x88, 0x4e, 0x5f, 0xa8, 0x46, 0x21, 0x29, 0x2e, 0x0e, 0x98, 0x21,
	0x12, 0x8c, 0x0a, 0xcc, 0x1a, 0x9c, 0x41, 0x70, 0xbe, 0x93, 0xfd, 0x89, 0x47, 0x72, 0x8c, 0x17,
	0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c,
	0x37, 0x1e, 0xcb, 0x31, 0x44, 0xa9, 0xa6, 0x67, 0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7,
	0xea, 0x7b, 0x67, 0x16, 0x25, 0x3a, 0xe7, 0x17, 0xa5, 0xea, 0x17, 0xa7, 0x66, 0x27, 0x66, 0xea,
	0x57, 0x80, 0x1d, 0x58, 0x52, 0x59, 0x90, 0x5a, 0x9c, 0xc4, 0x06, 0xb6, 0xd6, 0x18, 0x10, 0x00,
	0x00, 0xff, 0xff, 0x5a, 0x7e, 0x33, 0x2a, 0xb9, 0x00, 0x00, 0x00,
}

func (m *AllowedMessages) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AllowedMessages) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AllowedMessages) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Messages) > 0 {
		for iNdEx := len(m.Messages) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Messages[iNdEx])
			copy(dAtA[i:], m.Messages[iNdEx])
			i = encodeVarintAllowedMessages(dAtA, i, uint64(len(m.Messages[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintAllowedMessages(dAtA []byte, offset int, v uint64) int {
	offset -= sovAllowedMessages(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AllowedMessages) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Messages) > 0 {
		for _, s := range m.Messages {
			l = len(s)
			n += 1 + l + sovAllowedMessages(uint64(l))
		}
	}
	return n
}

func sovAllowedMessages(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAllowedMessages(x uint64) (n int) {
	return sovAllowedMessages(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AllowedMessages) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAllowedMessages
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
			return fmt.Errorf("proto: AllowedMessages: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AllowedMessages: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Messages", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAllowedMessages
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
				return ErrInvalidLengthAllowedMessages
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAllowedMessages
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Messages = append(m.Messages, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipAllowedMessages(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAllowedMessages
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
func skipAllowedMessages(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAllowedMessages
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
					return 0, ErrIntOverflowAllowedMessages
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
					return 0, ErrIntOverflowAllowedMessages
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
				return 0, ErrInvalidLengthAllowedMessages
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAllowedMessages
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAllowedMessages
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAllowedMessages        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAllowedMessages          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAllowedMessages = fmt.Errorf("proto: unexpected end of group")
)
