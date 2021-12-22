// Code generated by protoc-gen-go. DO NOT EDIT.
// source: kira/gov/network_properties.proto

package gov

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/golang/protobuf/proto"
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

type NetworkProperty int32

const (
	NetworkProperty_MIN_TX_FEE                     NetworkProperty = 0
	NetworkProperty_MAX_TX_FEE                     NetworkProperty = 1
	NetworkProperty_VOTE_QUORUM                    NetworkProperty = 2
	NetworkProperty_MINIMUM_PROPOSAL_END_TIME      NetworkProperty = 3
	NetworkProperty_PROPOSAL_ENACTMENT_TIME        NetworkProperty = 4
	NetworkProperty_MIN_PROPOSAL_END_BLOCKS        NetworkProperty = 5
	NetworkProperty_MIN_PROPOSAL_ENACTMENT_BLOCKS  NetworkProperty = 6
	NetworkProperty_ENABLE_FOREIGN_FEE_PAYMENTS    NetworkProperty = 7
	NetworkProperty_MISCHANCE_RANK_DECREASE_AMOUNT NetworkProperty = 8
	NetworkProperty_MAX_MISCHANCE                  NetworkProperty = 9
	NetworkProperty_MISCHANCE_CONFIDENCE           NetworkProperty = 10
	NetworkProperty_INACTIVE_RANK_DECREASE_PERCENT NetworkProperty = 11
	NetworkProperty_POOR_NETWORK_MAX_BANK_SEND     NetworkProperty = 12
	NetworkProperty_MIN_VALIDATORS                 NetworkProperty = 13
	NetworkProperty_UNJAIL_MAX_TIME                NetworkProperty = 14
	NetworkProperty_ENABLE_TOKEN_WHITELIST         NetworkProperty = 15
	NetworkProperty_ENABLE_TOKEN_BLACKLIST         NetworkProperty = 16
	NetworkProperty_MIN_IDENTITY_APPROVAL_TIP      NetworkProperty = 17
	NetworkProperty_UNIQUE_IDENTITY_KEYS           NetworkProperty = 18
)

var NetworkProperty_name = map[int32]string{
	0:  "MIN_TX_FEE",
	1:  "MAX_TX_FEE",
	2:  "VOTE_QUORUM",
	3:  "MINIMUM_PROPOSAL_END_TIME",
	4:  "PROPOSAL_ENACTMENT_TIME",
	5:  "MIN_PROPOSAL_END_BLOCKS",
	6:  "MIN_PROPOSAL_ENACTMENT_BLOCKS",
	7:  "ENABLE_FOREIGN_FEE_PAYMENTS",
	8:  "MISCHANCE_RANK_DECREASE_AMOUNT",
	9:  "MAX_MISCHANCE",
	10: "MISCHANCE_CONFIDENCE",
	11: "INACTIVE_RANK_DECREASE_PERCENT",
	12: "POOR_NETWORK_MAX_BANK_SEND",
	13: "MIN_VALIDATORS",
	14: "UNJAIL_MAX_TIME",
	15: "ENABLE_TOKEN_WHITELIST",
	16: "ENABLE_TOKEN_BLACKLIST",
	17: "MIN_IDENTITY_APPROVAL_TIP",
	18: "UNIQUE_IDENTITY_KEYS",
}

var NetworkProperty_value = map[string]int32{
	"MIN_TX_FEE":                     0,
	"MAX_TX_FEE":                     1,
	"VOTE_QUORUM":                    2,
	"MINIMUM_PROPOSAL_END_TIME":      3,
	"PROPOSAL_ENACTMENT_TIME":        4,
	"MIN_PROPOSAL_END_BLOCKS":        5,
	"MIN_PROPOSAL_ENACTMENT_BLOCKS":  6,
	"ENABLE_FOREIGN_FEE_PAYMENTS":    7,
	"MISCHANCE_RANK_DECREASE_AMOUNT": 8,
	"MAX_MISCHANCE":                  9,
	"MISCHANCE_CONFIDENCE":           10,
	"INACTIVE_RANK_DECREASE_PERCENT": 11,
	"POOR_NETWORK_MAX_BANK_SEND":     12,
	"MIN_VALIDATORS":                 13,
	"UNJAIL_MAX_TIME":                14,
	"ENABLE_TOKEN_WHITELIST":         15,
	"ENABLE_TOKEN_BLACKLIST":         16,
	"MIN_IDENTITY_APPROVAL_TIP":      17,
	"UNIQUE_IDENTITY_KEYS":           18,
}

func (x NetworkProperty) String() string {
	return proto.EnumName(NetworkProperty_name, int32(x))
}

func (NetworkProperty) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_98011a6048e5dde3, []int{0}
}

type MsgSetNetworkProperties struct {
	NetworkProperties    *NetworkProperties `protobuf:"bytes,1,opt,name=network_properties,json=networkProperties,proto3" json:"network_properties,omitempty"`
	Proposer             []byte             `protobuf:"bytes,2,opt,name=proposer,proto3" json:"proposer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *MsgSetNetworkProperties) Reset()         { *m = MsgSetNetworkProperties{} }
func (m *MsgSetNetworkProperties) String() string { return proto.CompactTextString(m) }
func (*MsgSetNetworkProperties) ProtoMessage()    {}
func (*MsgSetNetworkProperties) Descriptor() ([]byte, []int) {
	return fileDescriptor_98011a6048e5dde3, []int{0}
}

func (m *MsgSetNetworkProperties) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MsgSetNetworkProperties.Unmarshal(m, b)
}
func (m *MsgSetNetworkProperties) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MsgSetNetworkProperties.Marshal(b, m, deterministic)
}
func (m *MsgSetNetworkProperties) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgSetNetworkProperties.Merge(m, src)
}
func (m *MsgSetNetworkProperties) XXX_Size() int {
	return xxx_messageInfo_MsgSetNetworkProperties.Size(m)
}
func (m *MsgSetNetworkProperties) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgSetNetworkProperties.DiscardUnknown(m)
}

var xxx_messageInfo_MsgSetNetworkProperties proto.InternalMessageInfo

func (m *MsgSetNetworkProperties) GetNetworkProperties() *NetworkProperties {
	if m != nil {
		return m.NetworkProperties
	}
	return nil
}

func (m *MsgSetNetworkProperties) GetProposer() []byte {
	if m != nil {
		return m.Proposer
	}
	return nil
}

type NetworkPropertyValue struct {
	Value                uint64   `protobuf:"varint,1,opt,name=value,proto3" json:"value,omitempty"`
	StrValue             string   `protobuf:"bytes,2,opt,name=str_value,json=strValue,proto3" json:"str_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NetworkPropertyValue) Reset()         { *m = NetworkPropertyValue{} }
func (m *NetworkPropertyValue) String() string { return proto.CompactTextString(m) }
func (*NetworkPropertyValue) ProtoMessage()    {}
func (*NetworkPropertyValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_98011a6048e5dde3, []int{1}
}

func (m *NetworkPropertyValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkPropertyValue.Unmarshal(m, b)
}
func (m *NetworkPropertyValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkPropertyValue.Marshal(b, m, deterministic)
}
func (m *NetworkPropertyValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkPropertyValue.Merge(m, src)
}
func (m *NetworkPropertyValue) XXX_Size() int {
	return xxx_messageInfo_NetworkPropertyValue.Size(m)
}
func (m *NetworkPropertyValue) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkPropertyValue.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkPropertyValue proto.InternalMessageInfo

func (m *NetworkPropertyValue) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *NetworkPropertyValue) GetStrValue() string {
	if m != nil {
		return m.StrValue
	}
	return ""
}

type NetworkProperties struct {
	MinTxFee                    uint64   `protobuf:"varint,1,opt,name=min_tx_fee,json=minTxFee,proto3" json:"min_tx_fee,omitempty"`
	MaxTxFee                    uint64   `protobuf:"varint,2,opt,name=max_tx_fee,json=maxTxFee,proto3" json:"max_tx_fee,omitempty"`
	VoteQuorum                  uint64   `protobuf:"varint,3,opt,name=vote_quorum,json=voteQuorum,proto3" json:"vote_quorum,omitempty"`
	MinimumProposalEndTime      uint64   `protobuf:"varint,4,opt,name=minimum_proposal_end_time,json=minimumProposalEndTime,proto3" json:"minimum_proposal_end_time,omitempty"`
	ProposalEnactmentTime       uint64   `protobuf:"varint,5,opt,name=proposal_enactment_time,json=proposalEnactmentTime,proto3" json:"proposal_enactment_time,omitempty"`
	MinProposalEndBlocks        uint64   `protobuf:"varint,6,opt,name=min_proposal_end_blocks,json=minProposalEndBlocks,proto3" json:"min_proposal_end_blocks,omitempty"`
	MinProposalEnactmentBlocks  uint64   `protobuf:"varint,7,opt,name=min_proposal_enactment_blocks,json=minProposalEnactmentBlocks,proto3" json:"min_proposal_enactment_blocks,omitempty"`
	EnableForeignFeePayments    bool     `protobuf:"varint,8,opt,name=enable_foreign_fee_payments,json=enableForeignFeePayments,proto3" json:"enable_foreign_fee_payments,omitempty"`
	MischanceRankDecreaseAmount uint64   `protobuf:"varint,9,opt,name=mischance_rank_decrease_amount,json=mischanceRankDecreaseAmount,proto3" json:"mischance_rank_decrease_amount,omitempty"`
	MaxMischance                uint64   `protobuf:"varint,10,opt,name=max_mischance,json=maxMischance,proto3" json:"max_mischance,omitempty"`
	MischanceConfidence         uint64   `protobuf:"varint,11,opt,name=mischance_confidence,json=mischanceConfidence,proto3" json:"mischance_confidence,omitempty"`
	InactiveRankDecreasePercent uint64   `protobuf:"varint,12,opt,name=inactive_rank_decrease_percent,json=inactiveRankDecreasePercent,proto3" json:"inactive_rank_decrease_percent,omitempty"`
	MinValidators               uint64   `protobuf:"varint,13,opt,name=min_validators,json=minValidators,proto3" json:"min_validators,omitempty"`
	PoorNetworkMaxBankSend      uint64   `protobuf:"varint,14,opt,name=poor_network_max_bank_send,json=poorNetworkMaxBankSend,proto3" json:"poor_network_max_bank_send,omitempty"`
	UnjailMaxTime               uint64   `protobuf:"varint,15,opt,name=unjail_max_time,json=unjailMaxTime,proto3" json:"unjail_max_time,omitempty"`
	EnableTokenWhitelist        bool     `protobuf:"varint,16,opt,name=enable_token_whitelist,json=enableTokenWhitelist,proto3" json:"enable_token_whitelist,omitempty"`
	EnableTokenBlacklist        bool     `protobuf:"varint,17,opt,name=enable_token_blacklist,json=enableTokenBlacklist,proto3" json:"enable_token_blacklist,omitempty"`
	MinIdentityApprovalTip      uint64   `protobuf:"varint,18,opt,name=min_identity_approval_tip,json=minIdentityApprovalTip,proto3" json:"min_identity_approval_tip,omitempty"`
	UniqueIdentityKeys          string   `protobuf:"bytes,19,opt,name=unique_identity_keys,json=uniqueIdentityKeys,proto3" json:"unique_identity_keys,omitempty"`
	XXX_NoUnkeyedLiteral        struct{} `json:"-"`
	XXX_unrecognized            []byte   `json:"-"`
	XXX_sizecache               int32    `json:"-"`
}

func (m *NetworkProperties) Reset()         { *m = NetworkProperties{} }
func (m *NetworkProperties) String() string { return proto.CompactTextString(m) }
func (*NetworkProperties) ProtoMessage()    {}
func (*NetworkProperties) Descriptor() ([]byte, []int) {
	return fileDescriptor_98011a6048e5dde3, []int{2}
}

func (m *NetworkProperties) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NetworkProperties.Unmarshal(m, b)
}
func (m *NetworkProperties) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NetworkProperties.Marshal(b, m, deterministic)
}
func (m *NetworkProperties) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NetworkProperties.Merge(m, src)
}
func (m *NetworkProperties) XXX_Size() int {
	return xxx_messageInfo_NetworkProperties.Size(m)
}
func (m *NetworkProperties) XXX_DiscardUnknown() {
	xxx_messageInfo_NetworkProperties.DiscardUnknown(m)
}

var xxx_messageInfo_NetworkProperties proto.InternalMessageInfo

func (m *NetworkProperties) GetMinTxFee() uint64 {
	if m != nil {
		return m.MinTxFee
	}
	return 0
}

func (m *NetworkProperties) GetMaxTxFee() uint64 {
	if m != nil {
		return m.MaxTxFee
	}
	return 0
}

func (m *NetworkProperties) GetVoteQuorum() uint64 {
	if m != nil {
		return m.VoteQuorum
	}
	return 0
}

func (m *NetworkProperties) GetMinimumProposalEndTime() uint64 {
	if m != nil {
		return m.MinimumProposalEndTime
	}
	return 0
}

func (m *NetworkProperties) GetProposalEnactmentTime() uint64 {
	if m != nil {
		return m.ProposalEnactmentTime
	}
	return 0
}

func (m *NetworkProperties) GetMinProposalEndBlocks() uint64 {
	if m != nil {
		return m.MinProposalEndBlocks
	}
	return 0
}

func (m *NetworkProperties) GetMinProposalEnactmentBlocks() uint64 {
	if m != nil {
		return m.MinProposalEnactmentBlocks
	}
	return 0
}

func (m *NetworkProperties) GetEnableForeignFeePayments() bool {
	if m != nil {
		return m.EnableForeignFeePayments
	}
	return false
}

func (m *NetworkProperties) GetMischanceRankDecreaseAmount() uint64 {
	if m != nil {
		return m.MischanceRankDecreaseAmount
	}
	return 0
}

func (m *NetworkProperties) GetMaxMischance() uint64 {
	if m != nil {
		return m.MaxMischance
	}
	return 0
}

func (m *NetworkProperties) GetMischanceConfidence() uint64 {
	if m != nil {
		return m.MischanceConfidence
	}
	return 0
}

func (m *NetworkProperties) GetInactiveRankDecreasePercent() uint64 {
	if m != nil {
		return m.InactiveRankDecreasePercent
	}
	return 0
}

func (m *NetworkProperties) GetMinValidators() uint64 {
	if m != nil {
		return m.MinValidators
	}
	return 0
}

func (m *NetworkProperties) GetPoorNetworkMaxBankSend() uint64 {
	if m != nil {
		return m.PoorNetworkMaxBankSend
	}
	return 0
}

func (m *NetworkProperties) GetUnjailMaxTime() uint64 {
	if m != nil {
		return m.UnjailMaxTime
	}
	return 0
}

func (m *NetworkProperties) GetEnableTokenWhitelist() bool {
	if m != nil {
		return m.EnableTokenWhitelist
	}
	return false
}

func (m *NetworkProperties) GetEnableTokenBlacklist() bool {
	if m != nil {
		return m.EnableTokenBlacklist
	}
	return false
}

func (m *NetworkProperties) GetMinIdentityApprovalTip() uint64 {
	if m != nil {
		return m.MinIdentityApprovalTip
	}
	return 0
}

func (m *NetworkProperties) GetUniqueIdentityKeys() string {
	if m != nil {
		return m.UniqueIdentityKeys
	}
	return ""
}

func init() {
	proto.RegisterEnum("kira.gov.NetworkProperty", NetworkProperty_name, NetworkProperty_value)
	proto.RegisterType((*MsgSetNetworkProperties)(nil), "kira.gov.MsgSetNetworkProperties")
	proto.RegisterType((*NetworkPropertyValue)(nil), "kira.gov.NetworkPropertyValue")
	proto.RegisterType((*NetworkProperties)(nil), "kira.gov.NetworkProperties")
}

func init() {
	proto.RegisterFile("kira/gov/network_properties.proto", fileDescriptor_98011a6048e5dde3)
}

var fileDescriptor_98011a6048e5dde3 = []byte{
	// 1183 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x56, 0xcd, 0x6e, 0xdb, 0x46,
	0x17, 0x8d, 0x12, 0x27, 0x96, 0xc7, 0x7f, 0x34, 0xa3, 0xd8, 0x0c, 0x9d, 0x2f, 0xe6, 0x97, 0x22,
	0x41, 0x50, 0x20, 0x56, 0xd3, 0x34, 0x05, 0x1a, 0xa0, 0x0b, 0x4a, 0x1e, 0x37, 0x8c, 0xc4, 0x9f,
	0x50, 0x94, 0x92, 0x74, 0x33, 0x18, 0x4b, 0x13, 0x79, 0x2a, 0x71, 0x46, 0x21, 0x29, 0x45, 0x7e,
	0x83, 0x42, 0xef, 0x20, 0xa0, 0x40, 0x5f, 0xa1, 0xcf, 0xd3, 0x45, 0x17, 0x45, 0x9f, 0xa1, 0xab,
	0x62, 0x86, 0xa4, 0xac, 0x9f, 0xd8, 0x2b, 0x51, 0xf7, 0x9c, 0x73, 0xef, 0x9d, 0x3b, 0xf7, 0x10,
	0x04, 0xff, 0xef, 0xd1, 0x08, 0x97, 0xbb, 0x7c, 0x54, 0x66, 0x24, 0xf9, 0xcc, 0xa3, 0x1e, 0x1a,
	0x44, 0x7c, 0x40, 0xa2, 0x84, 0x92, 0xf8, 0x78, 0x10, 0xf1, 0x84, 0xab, 0x45, 0x41, 0x39, 0xee,
	0xf2, 0x91, 0x5e, 0xea, 0xf2, 0x2e, 0x97, 0xc1, 0xb2, 0x78, 0x4a, 0xf1, 0x47, 0x7f, 0x14, 0xc0,
	0x81, 0x1d, 0x77, 0x1b, 0x24, 0x71, 0xd2, 0x14, 0xde, 0x2c, 0x83, 0xfa, 0x06, 0xa8, 0xab, 0x79,
	0xb5, 0x82, 0x51, 0x78, 0xba, 0xf9, 0xed, 0xe1, 0x71, 0x9e, 0xf8, 0x78, 0x45, 0xe8, 0xef, 0xb1,
	0x95, 0x5c, 0x36, 0x28, 0x8a, 0x1c, 0x3c, 0x26, 0x91, 0x76, 0xd3, 0x28, 0x3c, 0xdd, 0xaa, 0x3c,
	0xff, 0xf7, 0xcf, 0xa3, 0x67, 0x5d, 0x9a, 0x9c, 0x0f, 0xcf, 0x8e, 0xdb, 0x3c, 0x2c, 0xb7, 0x79,
	0x1c, 0xf2, 0x38, 0xfb, 0x79, 0x16, 0x77, 0x7a, 0xe5, 0xe4, 0x62, 0x40, 0xe2, 0x63, 0xb3, 0xdd,
	0x36, 0x3b, 0x9d, 0x88, 0xc4, 0xb1, 0x3f, 0x4b, 0xf1, 0xc8, 0x05, 0xa5, 0xc5, 0xb2, 0x17, 0x2d,
	0xdc, 0x1f, 0x12, 0xb5, 0x04, 0x6e, 0x8f, 0xc4, 0x83, 0xec, 0x72, 0xcd, 0x4f, 0xff, 0xa8, 0x87,
	0x60, 0x23, 0x4e, 0x22, 0x94, 0x22, 0xa2, 0xfa, 0x86, 0x5f, 0x8c, 0x93, 0x48, 0x4a, 0x5e, 0xad,
	0xfd, 0xf3, 0xdb, 0x51, 0xe1, 0xd1, 0x5f, 0xeb, 0x60, 0x6f, 0x75, 0x02, 0x0f, 0x00, 0x08, 0x29,
	0x43, 0xc9, 0x18, 0x7d, 0x24, 0x79, 0xce, 0x62, 0x48, 0x59, 0x30, 0x3e, 0x25, 0x44, 0xa2, 0x78,
	0x9c, 0xa3, 0x37, 0x33, 0x14, 0x8f, 0x53, 0xf4, 0x08, 0x6c, 0x8e, 0x78, 0x42, 0xd0, 0xa7, 0x21,
	0x8f, 0x86, 0xa1, 0x76, 0x4b, 0xc2, 0x40, 0x84, 0xde, 0xca, 0x88, 0xfa, 0x03, 0xb8, 0x1f, 0x52,
	0x46, 0xc3, 0x61, 0x88, 0xd2, 0x73, 0xe1, 0x3e, 0x22, 0xac, 0x83, 0x12, 0x1a, 0x12, 0x6d, 0x4d,
	0xd2, 0xf7, 0x33, 0x82, 0x97, 0xe1, 0x90, 0x75, 0x02, 0x1a, 0x12, 0xf5, 0x7b, 0x70, 0x30, 0x27,
	0xc1, 0xed, 0x24, 0x24, 0x2c, 0x49, 0x85, 0xb7, 0xa5, 0xf0, 0xde, 0x60, 0xa6, 0xc8, 0x50, 0xa9,
	0x7b, 0x09, 0x0e, 0xc4, 0x79, 0x16, 0xca, 0x9d, 0xf5, 0x79, 0xbb, 0x17, 0x6b, 0x77, 0xa4, 0xae,
	0x14, 0x52, 0x36, 0x57, 0xac, 0x22, 0x31, 0xd5, 0x04, 0xff, 0x5b, 0x92, 0xe5, 0x25, 0x33, 0xf1,
	0xba, 0x14, 0xeb, 0x0b, 0xe2, 0x8c, 0x92, 0xa5, 0xf8, 0x11, 0x1c, 0x12, 0x86, 0xcf, 0xfa, 0x04,
	0x7d, 0xe4, 0x11, 0xa1, 0x5d, 0x26, 0x66, 0x86, 0x06, 0xf8, 0x42, 0x70, 0x62, 0xad, 0x68, 0x14,
	0x9e, 0x16, 0x7d, 0x2d, 0xa5, 0x9c, 0xa6, 0x8c, 0x53, 0x42, 0xbc, 0x0c, 0x57, 0xab, 0xe0, 0x61,
	0x48, 0xe3, 0xf6, 0x39, 0x66, 0x6d, 0x82, 0x22, 0xcc, 0x7a, 0xa8, 0x43, 0xda, 0x11, 0xc1, 0x31,
	0x41, 0x38, 0xe4, 0x43, 0x96, 0x68, 0x1b, 0xb2, 0x85, 0xc3, 0x19, 0xcb, 0xc7, 0xac, 0x77, 0x92,
	0x71, 0x4c, 0x49, 0x51, 0xbf, 0x02, 0xdb, 0xe2, 0xbe, 0x66, 0x14, 0x0d, 0x48, 0xcd, 0x56, 0x88,
	0xc7, 0x76, 0x1e, 0x53, 0x9f, 0x83, 0xd2, 0x65, 0xa5, 0x36, 0x67, 0x1f, 0x69, 0x87, 0x08, 0xee,
	0xa6, 0xe4, 0xde, 0x9d, 0x61, 0xd5, 0x19, 0x24, 0x9a, 0xa3, 0xe2, 0xb8, 0x74, 0xb4, 0xdc, 0xdb,
	0x80, 0x44, 0x6d, 0xc2, 0x12, 0x6d, 0x2b, 0x6d, 0x2e, 0x67, 0xcd, 0xf7, 0xe6, 0xa5, 0x14, 0xf5,
	0x31, 0xd8, 0x11, 0x33, 0x1e, 0xe1, 0x3e, 0xed, 0xe0, 0x84, 0x47, 0xb1, 0xb6, 0x2d, 0x45, 0xdb,
	0x21, 0x65, 0xad, 0x59, 0x50, 0x7d, 0x05, 0xf4, 0x01, 0xe7, 0x11, 0xca, 0x8d, 0x29, 0x0e, 0x74,
	0x26, 0x6a, 0xc6, 0x84, 0x75, 0xb4, 0x9d, 0x74, 0x6b, 0x04, 0x23, 0x5b, 0x66, 0x1b, 0x8f, 0x2b,
	0x98, 0xf5, 0x1a, 0x84, 0x75, 0xd4, 0x27, 0x60, 0x77, 0xc8, 0x7e, 0xc1, 0xb4, 0x2f, 0x55, 0x72,
	0x5b, 0x76, 0xd3, 0x1a, 0x69, 0xd8, 0xc6, 0x63, 0xb9, 0x25, 0xdf, 0x81, 0xfd, 0xec, 0xae, 0x12,
	0xde, 0x23, 0x0c, 0x7d, 0x3e, 0xa7, 0x09, 0xe9, 0xd3, 0x38, 0xd1, 0x14, 0x79, 0x4d, 0xa5, 0x14,
	0x0d, 0x04, 0xf8, 0x2e, 0xc7, 0x56, 0x54, 0x67, 0x7d, 0xdc, 0xee, 0x49, 0xd5, 0xde, 0x8a, 0xaa,
	0x92, 0x63, 0x99, 0x09, 0x90, 0x98, 0x64, 0x42, 0x93, 0x0b, 0x84, 0x07, 0x83, 0x88, 0x8f, 0x70,
	0x1f, 0x25, 0x74, 0xa0, 0xa9, 0x33, 0x13, 0x58, 0x19, 0x6e, 0x66, 0x70, 0x40, 0x07, 0xea, 0x37,
	0xa0, 0x34, 0x64, 0xf4, 0xd3, 0x90, 0x5c, 0xaa, 0x7b, 0xe4, 0x22, 0xd6, 0xee, 0x4a, 0x83, 0xab,
	0x29, 0x96, 0x0b, 0x6b, 0xe4, 0x22, 0xfe, 0xfa, 0xef, 0x75, 0xb0, 0xbb, 0xf4, 0xda, 0x10, 0x26,
	0xb6, 0x2d, 0x07, 0x05, 0xef, 0xd1, 0x29, 0x84, 0xca, 0x0d, 0x7d, 0x6b, 0x32, 0x35, 0x8a, 0xf6,
	0x9c, 0xc5, 0x6d, 0xf3, 0x7d, 0x8e, 0x16, 0x32, 0x74, 0xce, 0xe2, 0x2d, 0x37, 0x80, 0xe8, 0x6d,
	0xd3, 0xf5, 0x9b, 0xb6, 0x72, 0x53, 0xdf, 0x99, 0x4c, 0x0d, 0xd0, 0x5a, 0xb0, 0xb8, 0x6d, 0x39,
	0x96, 0xdd, 0xb4, 0x91, 0xe7, 0xbb, 0x9e, 0xdb, 0x30, 0xeb, 0x08, 0x3a, 0x27, 0x28, 0xb0, 0x6c,
	0xa8, 0xdc, 0xd2, 0xf5, 0xc9, 0xd4, 0xd8, 0xb7, 0xaf, 0xb4, 0xf8, 0x9c, 0xc4, 0xac, 0x06, 0x36,
	0x74, 0x82, 0x54, 0xb8, 0xa6, 0xdf, 0x9f, 0x4c, 0x8d, 0x7b, 0xde, 0x55, 0x16, 0x17, 0xe7, 0x59,
	0x28, 0x57, 0xa9, 0xbb, 0xd5, 0x5a, 0x43, 0xb9, 0xad, 0x6b, 0x93, 0xa9, 0x51, 0xb2, 0xaf, 0xb0,
	0xf8, 0x92, 0x2c, 0x2f, 0x99, 0x89, 0xef, 0xe8, 0x0f, 0x27, 0x53, 0x43, 0xb7, 0xaf, 0xb5, 0x38,
	0x74, 0xcc, 0x4a, 0x1d, 0xa2, 0x53, 0xd7, 0x87, 0xd6, 0x4f, 0x8e, 0x98, 0x19, 0xf2, 0xcc, 0x0f,
	0x22, 0x4d, 0x43, 0x59, 0xd7, 0x1f, 0x4c, 0xa6, 0x86, 0x06, 0xaf, 0xb1, 0xb8, 0x6d, 0x35, 0xaa,
	0xaf, 0x4d, 0xa7, 0x0a, 0x91, 0x6f, 0x3a, 0x35, 0x74, 0x02, 0xab, 0x3e, 0x34, 0x1b, 0x10, 0x99,
	0xb6, 0xdb, 0x74, 0x02, 0xa5, 0xa8, 0x1f, 0x4d, 0xa6, 0xc6, 0xa1, 0x7d, 0xbd, 0xc5, 0xc5, 0x7d,
	0xcd, 0x12, 0x29, 0x1b, 0xba, 0x32, 0x99, 0x1a, 0x5b, 0xf6, 0x92, 0xc5, 0x2f, 0x2b, 0x55, 0x5d,
	0xe7, 0xd4, 0x3a, 0x81, 0x82, 0x0b, 0xf4, 0x83, 0xc9, 0xd4, 0xb8, 0x6b, 0x7f, 0xd9, 0xe2, 0x96,
	0x98, 0x88, 0xd5, 0x5a, 0xee, 0xcd, 0x83, 0x7e, 0x15, 0x3a, 0x81, 0xb2, 0x99, 0x36, 0x67, 0x5d,
	0x63, 0xf1, 0x57, 0x40, 0xf7, 0x5c, 0xd7, 0x47, 0x0e, 0x0c, 0xde, 0xb9, 0x7e, 0x0d, 0x89, 0x4e,
	0x2b, 0x22, 0x59, 0x03, 0x3a, 0x27, 0xca, 0x56, 0xba, 0x0e, 0xde, 0x97, 0xbd, 0xfb, 0x18, 0xec,
	0x88, 0xfb, 0x69, 0x99, 0x75, 0xeb, 0xc4, 0x0c, 0x5c, 0xbf, 0xa1, 0x6c, 0xeb, 0x7b, 0x93, 0xa9,
	0xb1, 0x6d, 0x2f, 0xbc, 0x1e, 0x9e, 0x80, 0xdd, 0xa6, 0xf3, 0xc6, 0xb4, 0xea, 0x32, 0xb9, 0xdc,
	0x96, 0x9d, 0x94, 0xd7, 0x5c, 0xb6, 0x78, 0x76, 0x57, 0x81, 0x5b, 0x83, 0x0e, 0x7a, 0xf7, 0xda,
	0x0a, 0x60, 0xdd, 0x6a, 0x04, 0xca, 0x6e, 0xba, 0x24, 0xf0, 0x0a, 0x8b, 0x2f, 0xa8, 0x2a, 0x75,
	0xb3, 0x5a, 0x93, 0x2a, 0x65, 0x45, 0xb5, 0x60, 0x71, 0xd1, 0xba, 0x18, 0x72, 0x60, 0x05, 0x1f,
	0x90, 0xe9, 0x79, 0xbe, 0xdb, 0x32, 0xeb, 0x28, 0xb0, 0x3c, 0x65, 0x6f, 0x66, 0x82, 0x2b, 0x2c,
	0xde, 0x74, 0xac, 0xb7, 0x4d, 0x78, 0xa9, 0xae, 0xc1, 0x0f, 0x0d, 0x45, 0xd5, 0xf7, 0x27, 0x53,
	0x43, 0x6d, 0xae, 0x58, 0x5c, 0x5f, 0xfb, 0xf5, 0xf7, 0x87, 0x37, 0x2a, 0x2f, 0x7f, 0x7e, 0x31,
	0xf7, 0x65, 0x51, 0xa3, 0x11, 0xae, 0xf2, 0x88, 0x94, 0x63, 0xd2, 0xc3, 0xb4, 0x6c, 0x39, 0x01,
	0xf4, 0xdf, 0x97, 0xe5, 0xf7, 0xcf, 0xb3, 0x2e, 0x61, 0xe5, 0xfc, 0x2b, 0xea, 0xec, 0x8e, 0x8c,
	0xbd, 0xf8, 0x2f, 0x00, 0x00, 0xff, 0xff, 0x43, 0x49, 0xec, 0x05, 0x58, 0x09, 0x00, 0x00,
}
