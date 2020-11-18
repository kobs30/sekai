// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: permission.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type PermValue int32

const (
	// PERMISSION_ZERO is a no-op permission.
	PermZero PermValue = 0
	// PERMISSION_SET_PERMISSIONS defines the permission that allows to Set Permissions to other actors.
	PermSetPermissions PermValue = 1
	// PERMISSION_CLAIM_VALIDATOR defines the permission that allows to Claim a validator Seat.
	PermClaimValidator PermValue = 2
	// PERMISSION_CLAIM_COUNCILOR defines the permission that allows to Claim a Councilor Seat.
	PermClaimCouncilor PermValue = 3
	// PERMISSION_CREATE_SET_PERMISSIONS_PROPOSAL defines the permission needed to create proposals for setting permissions.
	PermCreateSetPermissionsProposal PermValue = 4
	// PERMISSION_VOTE_SET_PERMISSIONS_PROPOSAL defines the permission that an actor must have in order to vote a
	// Proposal to set permissions.
	PermVoteSetPermissionProposal PermValue = 5
	//  PERMISSION_UPSERT_TOKEN_ALIAS
	PermUpsertTokenAlias PermValue = 6
	// PERMISSION_CHANGE_TX_FEE
	PermChangeTxFee PermValue = 7
	// PERMISSION_UPSERT_TOKEN_RATE
	PermUpsertTokenRate PermValue = 8
	// PERMISSION_UPSERT_ROLE makes possible to add, modify and assign roles.
	PermUpsertRole PermValue = 9
	// PERMISSION_UPSERT_DATA_REGISTRY_PROPOSAL makes possible to create a proposal to change the Data Registry.
	PermUpsertDataRegistryProposal PermValue = 10
	// PERMISSION_VOTE_UPSERT_DATA_REGISTRY_PROPOSAL makes possible to create a proposal to change the Data Registry.
	PermVoteUpsertDataRegistryProposal PermValue = 11
	// PERMISSION_CREATE_SET_NETWORK_PROPERTY_PROPOSAL defines the permission needed to create proposals for setting network property.
	PermCreateSetNetworkPropertyProposal PermValue = 12
	// PERMISSION_VOTE_SET_NETWORK_PROPERTY_PROPOSAL defines the permission that an actor must have in order to vote a
	// Proposal to set network property.
	PermVoteSetNetworkPropertyProposal PermValue = 13
	// PERMISSION_CREATE_UPSERT_TOKEN_ALIAS_PROPOSAL defines the permission needed to create proposals for upsert token proposal.
	PermCreateUpsertTokenAliasProposal PermValue = 14
)

var PermValue_name = map[int32]string{
	0:  "PERMISSION_ZERO",
	1:  "PERMISSION_SET_PERMISSIONS",
	2:  "PERMISSION_CLAIM_VALIDATOR",
	3:  "PERMISSION_CLAIM_COUNCILOR",
	4:  "PERMISSION_CREATE_SET_PERMISSIONS_PROPOSAL",
	5:  "PERMISSION_VOTE_SET_PERMISSIONS_PROPOSAL",
	6:  "PERMISSION_UPSERT_TOKEN_ALIAS",
	7:  "PERMISSION_CHANGE_TX_FEE",
	8:  "PERMISSION_UPSERT_TOKEN_RATE",
	9:  "PERMISSION_UPSERT_ROLE",
	10: "PERMISSION_UPSERT_DATA_REGISTRY_PROPOSAL",
	11: "PERMISSION_VOTE_UPSERT_DATA_REGISTRY_PROPOSAL",
	12: "PERMISSION_CREATE_SET_NETWORK_PROPERTY_PROPOSAL",
	13: "PERMISSION_VOTE_SET_NETWORK_PROPERTY_PROPOSAL",
	14: "PERMISSION_CREATE_UPSERT_TOKEN_ALIAS_PROPOSAL",
}

var PermValue_value = map[string]int32{
	"PERMISSION_ZERO":                                 0,
	"PERMISSION_SET_PERMISSIONS":                      1,
	"PERMISSION_CLAIM_VALIDATOR":                      2,
	"PERMISSION_CLAIM_COUNCILOR":                      3,
	"PERMISSION_CREATE_SET_PERMISSIONS_PROPOSAL":      4,
	"PERMISSION_VOTE_SET_PERMISSIONS_PROPOSAL":        5,
	"PERMISSION_UPSERT_TOKEN_ALIAS":                   6,
	"PERMISSION_CHANGE_TX_FEE":                        7,
	"PERMISSION_UPSERT_TOKEN_RATE":                    8,
	"PERMISSION_UPSERT_ROLE":                          9,
	"PERMISSION_UPSERT_DATA_REGISTRY_PROPOSAL":        10,
	"PERMISSION_VOTE_UPSERT_DATA_REGISTRY_PROPOSAL":   11,
	"PERMISSION_CREATE_SET_NETWORK_PROPERTY_PROPOSAL": 12,
	"PERMISSION_VOTE_SET_NETWORK_PROPERTY_PROPOSAL":   13,
	"PERMISSION_CREATE_UPSERT_TOKEN_ALIAS_PROPOSAL":   14,
}

func (x PermValue) String() string {
	return proto.EnumName(PermValue_name, int32(x))
}

func (PermValue) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_c837ef01cbda0ad8, []int{0}
}

func init() {
	proto.RegisterEnum("kira.gov.PermValue", PermValue_name, PermValue_value)
}

func init() { proto.RegisterFile("permission.proto", fileDescriptor_c837ef01cbda0ad8) }

var fileDescriptor_c837ef01cbda0ad8 = []byte{
	// 586 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x94, 0xdf, 0x4a, 0xdc, 0x40,
	0x18, 0xc5, 0x77, 0x5b, 0x6b, 0xd7, 0xa9, 0xd5, 0x10, 0xc5, 0xca, 0x50, 0xc3, 0x28, 0xb6, 0x88,
	0xd0, 0x0d, 0xa5, 0x50, 0x28, 0xbd, 0x28, 0xd3, 0x75, 0xb4, 0xc1, 0x75, 0x27, 0x4c, 0x46, 0x5b,
	0x85, 0x12, 0x46, 0x3b, 0xc4, 0xb0, 0xeb, 0xce, 0x32, 0x19, 0xad, 0xbe, 0x41, 0xc9, 0x55, 0x5f,
	0x20, 0x50, 0xe8, 0xcb, 0xf4, 0xd2, 0xcb, 0x5e, 0x16, 0x7d, 0x91, 0xb2, 0xab, 0x6e, 0x76, 0xa3,
	0x5d, 0x7a, 0x97, 0x3f, 0xdf, 0xf9, 0xcd, 0xf9, 0x4e, 0x0e, 0x01, 0x56, 0x47, 0xea, 0xa3, 0x38,
	0x49, 0x62, 0xd5, 0xae, 0x76, 0xb4, 0x32, 0xca, 0xae, 0x34, 0x63, 0x2d, 0xaa, 0x91, 0x3a, 0x81,
	0xb3, 0x91, 0x8a, 0x54, 0xef, 0xa1, 0xdb, 0xbd, 0xba, 0x7a, 0xbf, 0xfa, 0xa3, 0x02, 0x26, 0x7c,
	0xa9, 0x8f, 0x76, 0x44, 0xeb, 0x58, 0xda, 0x8b, 0x60, 0xda, 0x27, 0x6c, 0xcb, 0x0b, 0x02, 0x8f,
	0x36, 0xc2, 0x3d, 0xc2, 0xa8, 0x55, 0x82, 0x93, 0x69, 0x86, 0x2a, 0xdd, 0x99, 0x3d, 0xa9, 0x95,
	0xfd, 0x1a, 0xc0, 0x81, 0x91, 0x80, 0xf0, 0x30, 0xbf, 0x0d, 0xac, 0x32, 0x9c, 0x4b, 0x33, 0x64,
	0x77, 0xa7, 0x03, 0x69, 0xfc, 0xbe, 0x9b, 0xa4, 0xa0, 0xab, 0xd5, 0xb1, 0xb7, 0x15, 0xee, 0xe0,
	0xba, 0xb7, 0x86, 0x39, 0x65, 0xd6, 0xbd, 0x5c, 0x57, 0x6b, 0x89, 0xb8, 0x6b, 0x27, 0xfe, 0x22,
	0x8c, 0xd2, 0x77, 0xea, 0x6a, 0x74, 0xbb, 0x51, 0xf3, 0xea, 0x94, 0x59, 0xf7, 0x0b, 0xba, 0x9a,
	0x3a, 0x6e, 0x1f, 0xc4, 0x2d, 0xa5, 0x6d, 0x0e, 0x56, 0x07, 0x75, 0x8c, 0x60, 0x4e, 0x8a, 0x76,
	0x43, 0x9f, 0x51, 0x9f, 0x06, 0xb8, 0x6e, 0x8d, 0xc1, 0xe5, 0x34, 0x43, 0xa8, 0xc7, 0xd1, 0x52,
	0x18, 0x39, 0xec, 0xde, 0xd7, 0xaa, 0xa3, 0x12, 0xd1, 0xb2, 0x29, 0x58, 0x19, 0xa0, 0xee, 0xd0,
	0x51, 0xcc, 0x07, 0x70, 0x31, 0xcd, 0xd0, 0x42, 0x2f, 0x5d, 0x55, 0x20, 0xf6, 0x81, 0x6f, 0xc1,
	0xc2, 0x00, 0x70, 0xdb, 0x0f, 0x08, 0xe3, 0x21, 0xa7, 0x9b, 0xa4, 0x11, 0xe2, 0xba, 0x87, 0x03,
	0x6b, 0x1c, 0xce, 0xa7, 0x19, 0x9a, 0xed, 0x4a, 0xb7, 0x3b, 0x89, 0xd4, 0x86, 0xab, 0xa6, 0x6c,
	0xe3, 0x56, 0x2c, 0x12, 0xfb, 0x25, 0x98, 0x1f, 0xdc, 0xf1, 0x03, 0x6e, 0x6c, 0x90, 0x90, 0x7f,
	0x0a, 0xd7, 0x09, 0xb1, 0x1e, 0xc2, 0x99, 0x34, 0x43, 0xd3, 0xbd, 0x8d, 0x0e, 0x45, 0x3b, 0x92,
	0xfc, 0x74, 0x5d, 0x4a, 0xfb, 0x0d, 0x78, 0xfa, 0xaf, 0xf3, 0x18, 0xe6, 0xc4, 0xaa, 0xc0, 0x27,
	0x69, 0x86, 0x66, 0x0a, 0xc7, 0x31, 0x61, 0xa4, 0x5d, 0x05, 0x73, 0xb7, 0xa5, 0x8c, 0xd6, 0x89,
	0x35, 0x01, 0xed, 0x34, 0x43, 0x53, 0xb9, 0x88, 0xa9, 0x96, 0xb4, 0xfd, 0xa1, 0xac, 0xae, 0xe7,
	0xd7, 0x30, 0xc7, 0x21, 0x23, 0x1b, 0x5e, 0xc0, 0xd9, 0x6e, 0x9e, 0x15, 0x80, 0x4b, 0x69, 0x86,
	0x9c, 0x9c, 0xb0, 0x26, 0x8c, 0x60, 0x32, 0x8a, 0x13, 0xa3, 0xcf, 0xfa, 0x61, 0xed, 0x82, 0x17,
	0xc5, 0xf4, 0x47, 0x63, 0x1f, 0xc1, 0xe7, 0x69, 0x86, 0x96, 0x6e, 0x3e, 0xc1, 0x08, 0xf4, 0x67,
	0xe0, 0xde, 0x5d, 0x97, 0x06, 0xe1, 0x1f, 0x29, 0xdb, 0xec, 0x31, 0x09, 0xe3, 0x03, 0xf0, 0x49,
	0xb8, 0x92, 0x66, 0x68, 0x79, 0xa8, 0x33, 0x0d, 0x69, 0xbe, 0x2a, 0xdd, 0xec, 0x62, 0xa5, 0x36,
	0x23, 0x9d, 0x8f, 0x86, 0x3f, 0x1e, 0x76, 0xfe, 0xdf, 0xe8, 0x6b, 0xe7, 0xb7, 0x8b, 0x94, 0xa3,
	0xa7, 0x72, 0xf4, 0x95, 0xef, 0x62, 0xaf, 0x6e, 0xd0, 0x70, 0xec, 0xdb, 0x4f, 0xa7, 0xf4, 0xfe,
	0xdd, 0xaf, 0x0b, 0xa7, 0x7c, 0x7e, 0xe1, 0x94, 0xff, 0x5c, 0x38, 0xe5, 0xef, 0x97, 0x4e, 0xe9,
	0xfc, 0xd2, 0x29, 0xfd, 0xbe, 0x74, 0x4a, 0x7b, 0xcf, 0xa2, 0xd8, 0x1c, 0x1e, 0xef, 0x57, 0x0f,
	0xd4, 0x91, 0xbb, 0x19, 0x6b, 0x51, 0x53, 0x5a, 0xba, 0x89, 0x6c, 0x8a, 0xd8, 0x3d, 0x75, 0x23,
	0x75, 0xe2, 0x9a, 0xb3, 0x8e, 0x4c, 0xf6, 0xc7, 0x7b, 0xbf, 0x9a, 0x57, 0x7f, 0x03, 0x00, 0x00,
	0xff, 0xff, 0x9a, 0x09, 0xcf, 0x04, 0x9e, 0x04, 0x00, 0x00,
}
