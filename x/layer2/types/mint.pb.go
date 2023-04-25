// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: kira/layer2/mint.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

type TokenInfo struct {
	TokenType   string                                 `protobuf:"bytes,1,opt,name=token_type,json=tokenType,proto3" json:"token_type,omitempty"`
	Denom       string                                 `protobuf:"bytes,2,opt,name=denom,proto3" json:"denom,omitempty"`
	Name        string                                 `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Symbol      string                                 `protobuf:"bytes,4,opt,name=symbol,proto3" json:"symbol,omitempty"`
	Icon        string                                 `protobuf:"bytes,5,opt,name=icon,proto3" json:"icon,omitempty"`
	Description string                                 `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	Website     string                                 `protobuf:"bytes,7,opt,name=website,proto3" json:"website,omitempty"`
	Social      string                                 `protobuf:"bytes,8,opt,name=social,proto3" json:"social,omitempty"`
	Decimals    uint64                                 `protobuf:"varint,9,opt,name=decimals,proto3" json:"decimals,omitempty"`
	Cap         github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,10,opt,name=cap,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"cap"`
	Supply      github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,11,opt,name=supply,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"supply"`
	Holders     uint64                                 `protobuf:"varint,12,opt,name=holders,proto3" json:"holders,omitempty"`
	Fee         github_com_cosmos_cosmos_sdk_types.Int `protobuf:"bytes,13,opt,name=fee,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Int" json:"fee"`
	Owner       string                                 `protobuf:"bytes,14,opt,name=owner,proto3" json:"owner,omitempty"`
	Metadata    string                                 `protobuf:"bytes,15,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Hash        string                                 `protobuf:"bytes,16,opt,name=hash,proto3" json:"hash,omitempty"`
}

func (m *TokenInfo) Reset()         { *m = TokenInfo{} }
func (m *TokenInfo) String() string { return proto.CompactTextString(m) }
func (*TokenInfo) ProtoMessage()    {}
func (*TokenInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_066d1ba7ae573fa3, []int{0}
}
func (m *TokenInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TokenInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TokenInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TokenInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TokenInfo.Merge(m, src)
}
func (m *TokenInfo) XXX_Size() int {
	return m.Size()
}
func (m *TokenInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_TokenInfo.DiscardUnknown(m)
}

var xxx_messageInfo_TokenInfo proto.InternalMessageInfo

func (m *TokenInfo) GetTokenType() string {
	if m != nil {
		return m.TokenType
	}
	return ""
}

func (m *TokenInfo) GetDenom() string {
	if m != nil {
		return m.Denom
	}
	return ""
}

func (m *TokenInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TokenInfo) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *TokenInfo) GetIcon() string {
	if m != nil {
		return m.Icon
	}
	return ""
}

func (m *TokenInfo) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *TokenInfo) GetWebsite() string {
	if m != nil {
		return m.Website
	}
	return ""
}

func (m *TokenInfo) GetSocial() string {
	if m != nil {
		return m.Social
	}
	return ""
}

func (m *TokenInfo) GetDecimals() uint64 {
	if m != nil {
		return m.Decimals
	}
	return 0
}

func (m *TokenInfo) GetHolders() uint64 {
	if m != nil {
		return m.Holders
	}
	return 0
}

func (m *TokenInfo) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *TokenInfo) GetMetadata() string {
	if m != nil {
		return m.Metadata
	}
	return ""
}

func (m *TokenInfo) GetHash() string {
	if m != nil {
		return m.Hash
	}
	return ""
}

func init() {
	proto.RegisterType((*TokenInfo)(nil), "kira.layer2.TokenInfo")
}

func init() { proto.RegisterFile("kira/layer2/mint.proto", fileDescriptor_066d1ba7ae573fa3) }

var fileDescriptor_066d1ba7ae573fa3 = []byte{
	// 429 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcf, 0x6a, 0xdb, 0x40,
	0x10, 0xc6, 0xad, 0xc6, 0x71, 0xe2, 0x75, 0xff, 0xb1, 0x84, 0x30, 0x18, 0x2a, 0x9b, 0x1e, 0x8a,
	0x2f, 0xf5, 0x42, 0xfb, 0x02, 0x25, 0x85, 0x40, 0xe8, 0xcd, 0xe4, 0xd4, 0x4b, 0x59, 0x49, 0x63,
	0x79, 0xb1, 0x76, 0x47, 0x68, 0xd7, 0xa4, 0x7a, 0x8b, 0x3e, 0x40, 0x1f, 0xa0, 0x8f, 0x92, 0x63,
	0x8e, 0xa5, 0x87, 0x50, 0xec, 0x17, 0x09, 0x3b, 0xb2, 0x43, 0xce, 0x39, 0x69, 0x7e, 0xdf, 0xcc,
	0x68, 0xbf, 0xdd, 0x19, 0x71, 0xbe, 0x36, 0x8d, 0x56, 0x95, 0x6e, 0xb1, 0xf9, 0xa4, 0xac, 0x71,
	0x61, 0x5e, 0x37, 0x14, 0x48, 0x8e, 0xa2, 0x3e, 0xef, 0xf4, 0xf1, 0x59, 0x49, 0x25, 0xb1, 0xae,
	0x62, 0xd4, 0x95, 0x8c, 0x27, 0x25, 0x51, 0x59, 0xa1, 0x62, 0xca, 0x36, 0x4b, 0x15, 0x8c, 0x45,
	0x1f, 0xb4, 0xad, 0xbb, 0x82, 0xf7, 0xbf, 0xfb, 0x62, 0x78, 0x4d, 0x6b, 0x74, 0x57, 0x6e, 0x49,
	0xf2, 0x9d, 0x10, 0x21, 0xc2, 0x8f, 0xd0, 0xd6, 0x08, 0xc9, 0x34, 0x99, 0x0d, 0x17, 0x43, 0x56,
	0xae, 0xdb, 0x1a, 0xe5, 0x99, 0x38, 0x2e, 0xd0, 0x91, 0x85, 0x17, 0x9c, 0xe9, 0x40, 0x4a, 0xd1,
	0x77, 0xda, 0x22, 0x1c, 0xb1, 0xc8, 0xb1, 0x3c, 0x17, 0x03, 0xdf, 0xda, 0x8c, 0x2a, 0xe8, 0xb3,
	0xba, 0xa7, 0x58, 0x6b, 0x72, 0x72, 0x70, 0xdc, 0xd5, 0xc6, 0x58, 0x4e, 0xc5, 0xa8, 0x40, 0x9f,
	0x37, 0xa6, 0x0e, 0x86, 0x1c, 0x0c, 0x38, 0xf5, 0x54, 0x92, 0x20, 0x4e, 0x6e, 0x30, 0xf3, 0x26,
	0x20, 0x9c, 0x70, 0xf6, 0x80, 0x7c, 0x0e, 0xe5, 0x46, 0x57, 0x70, 0xba, 0x3f, 0x87, 0x49, 0x8e,
	0xc5, 0x69, 0x81, 0xb9, 0xb1, 0xba, 0xf2, 0x30, 0x9c, 0x26, 0xb3, 0xfe, 0xe2, 0x91, 0xe5, 0x17,
	0x71, 0x94, 0xeb, 0x1a, 0x44, 0x6c, 0xb8, 0x98, 0xdf, 0xde, 0x4f, 0x7a, 0xff, 0xee, 0x27, 0x1f,
	0x4a, 0x13, 0x56, 0x9b, 0x6c, 0x9e, 0x93, 0x55, 0x39, 0x79, 0x4b, 0x7e, 0xff, 0xf9, 0xe8, 0x8b,
	0xb5, 0x8a, 0xcf, 0xe1, 0xe7, 0x57, 0x2e, 0x2c, 0x62, 0xab, 0xbc, 0x14, 0x03, 0xbf, 0xa9, 0xeb,
	0xaa, 0x85, 0xd1, 0xb3, 0x7e, 0xb2, 0xef, 0x8e, 0xf7, 0x5a, 0x51, 0x55, 0x60, 0xe3, 0xe1, 0x25,
	0x9b, 0x3c, 0x60, 0xf4, 0xb8, 0x44, 0x84, 0x57, 0xcf, 0xf3, 0xb8, 0x44, 0x9e, 0x15, 0xdd, 0x38,
	0x6c, 0xe0, 0x75, 0x37, 0x2b, 0x86, 0xf8, 0x2e, 0x16, 0x83, 0x2e, 0x74, 0xd0, 0xf0, 0x86, 0x13,
	0x8f, 0x1c, 0x67, 0xb3, 0xd2, 0x7e, 0x05, 0x6f, 0xbb, 0xd9, 0xc4, 0xf8, 0xe2, 0xf2, 0xcf, 0x36,
	0x4d, 0x6e, 0xb7, 0x69, 0x72, 0xb7, 0x4d, 0x93, 0xff, 0xdb, 0x34, 0xf9, 0xb5, 0x4b, 0x7b, 0x77,
	0xbb, 0xb4, 0xf7, 0x77, 0x97, 0xf6, 0xbe, 0xcf, 0x9e, 0x18, 0xfa, 0x66, 0x1a, 0xfd, 0x95, 0x1a,
	0x54, 0x1e, 0xd7, 0xda, 0xa8, 0x9f, 0x87, 0x7d, 0x65, 0x5b, 0xd9, 0x80, 0xb7, 0xed, 0xf3, 0x43,
	0x00, 0x00, 0x00, 0xff, 0xff, 0x04, 0x35, 0x96, 0x68, 0xcb, 0x02, 0x00, 0x00,
}

func (this *TokenInfo) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TokenInfo)
	if !ok {
		that2, ok := that.(TokenInfo)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.TokenType != that1.TokenType {
		return false
	}
	if this.Denom != that1.Denom {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Symbol != that1.Symbol {
		return false
	}
	if this.Icon != that1.Icon {
		return false
	}
	if this.Description != that1.Description {
		return false
	}
	if this.Website != that1.Website {
		return false
	}
	if this.Social != that1.Social {
		return false
	}
	if this.Decimals != that1.Decimals {
		return false
	}
	if !this.Cap.Equal(that1.Cap) {
		return false
	}
	if !this.Supply.Equal(that1.Supply) {
		return false
	}
	if this.Holders != that1.Holders {
		return false
	}
	if !this.Fee.Equal(that1.Fee) {
		return false
	}
	if this.Owner != that1.Owner {
		return false
	}
	if this.Metadata != that1.Metadata {
		return false
	}
	if this.Hash != that1.Hash {
		return false
	}
	return true
}
func (m *TokenInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TokenInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *TokenInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Hash) > 0 {
		i -= len(m.Hash)
		copy(dAtA[i:], m.Hash)
		i = encodeVarintMint(dAtA, i, uint64(len(m.Hash)))
		i--
		dAtA[i] = 0x1
		i--
		dAtA[i] = 0x82
	}
	if len(m.Metadata) > 0 {
		i -= len(m.Metadata)
		copy(dAtA[i:], m.Metadata)
		i = encodeVarintMint(dAtA, i, uint64(len(m.Metadata)))
		i--
		dAtA[i] = 0x7a
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintMint(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x72
	}
	{
		size := m.Fee.Size()
		i -= size
		if _, err := m.Fee.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x6a
	if m.Holders != 0 {
		i = encodeVarintMint(dAtA, i, uint64(m.Holders))
		i--
		dAtA[i] = 0x60
	}
	{
		size := m.Supply.Size()
		i -= size
		if _, err := m.Supply.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x5a
	{
		size := m.Cap.Size()
		i -= size
		if _, err := m.Cap.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x52
	if m.Decimals != 0 {
		i = encodeVarintMint(dAtA, i, uint64(m.Decimals))
		i--
		dAtA[i] = 0x48
	}
	if len(m.Social) > 0 {
		i -= len(m.Social)
		copy(dAtA[i:], m.Social)
		i = encodeVarintMint(dAtA, i, uint64(len(m.Social)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.Website) > 0 {
		i -= len(m.Website)
		copy(dAtA[i:], m.Website)
		i = encodeVarintMint(dAtA, i, uint64(len(m.Website)))
		i--
		dAtA[i] = 0x3a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintMint(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x32
	}
	if len(m.Icon) > 0 {
		i -= len(m.Icon)
		copy(dAtA[i:], m.Icon)
		i = encodeVarintMint(dAtA, i, uint64(len(m.Icon)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Symbol) > 0 {
		i -= len(m.Symbol)
		copy(dAtA[i:], m.Symbol)
		i = encodeVarintMint(dAtA, i, uint64(len(m.Symbol)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintMint(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Denom) > 0 {
		i -= len(m.Denom)
		copy(dAtA[i:], m.Denom)
		i = encodeVarintMint(dAtA, i, uint64(len(m.Denom)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.TokenType) > 0 {
		i -= len(m.TokenType)
		copy(dAtA[i:], m.TokenType)
		i = encodeVarintMint(dAtA, i, uint64(len(m.TokenType)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMint(dAtA []byte, offset int, v uint64) int {
	offset -= sovMint(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *TokenInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TokenType)
	if l > 0 {
		n += 1 + l + sovMint(uint64(l))
	}
	l = len(m.Denom)
	if l > 0 {
		n += 1 + l + sovMint(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovMint(uint64(l))
	}
	l = len(m.Symbol)
	if l > 0 {
		n += 1 + l + sovMint(uint64(l))
	}
	l = len(m.Icon)
	if l > 0 {
		n += 1 + l + sovMint(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovMint(uint64(l))
	}
	l = len(m.Website)
	if l > 0 {
		n += 1 + l + sovMint(uint64(l))
	}
	l = len(m.Social)
	if l > 0 {
		n += 1 + l + sovMint(uint64(l))
	}
	if m.Decimals != 0 {
		n += 1 + sovMint(uint64(m.Decimals))
	}
	l = m.Cap.Size()
	n += 1 + l + sovMint(uint64(l))
	l = m.Supply.Size()
	n += 1 + l + sovMint(uint64(l))
	if m.Holders != 0 {
		n += 1 + sovMint(uint64(m.Holders))
	}
	l = m.Fee.Size()
	n += 1 + l + sovMint(uint64(l))
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovMint(uint64(l))
	}
	l = len(m.Metadata)
	if l > 0 {
		n += 1 + l + sovMint(uint64(l))
	}
	l = len(m.Hash)
	if l > 0 {
		n += 2 + l + sovMint(uint64(l))
	}
	return n
}

func sovMint(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMint(x uint64) (n int) {
	return sovMint(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TokenInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMint
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
			return fmt.Errorf("proto: TokenInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TokenInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Denom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Symbol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Symbol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Icon", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Icon = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Website", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Website = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Social", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Social = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Decimals", wireType)
			}
			m.Decimals = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Decimals |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Cap", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Cap.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 11:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Supply", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Supply.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Holders", wireType)
			}
			m.Holders = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Holders |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Fee.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Metadata", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Metadata = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 16:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Hash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMint
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
func skipMint(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMint
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
					return 0, ErrIntOverflowMint
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
					return 0, ErrIntOverflowMint
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
				return 0, ErrInvalidLengthMint
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMint
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMint
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMint        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMint          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMint = fmt.Errorf("proto: unexpected end of group")
)
