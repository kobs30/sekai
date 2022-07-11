package types

import (
	kiratypes "github.com/KiraCore/sekai/types"
)

// DefaultGenesis returns the default CustomGo genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		NextRoleId: 3,
		Roles: []Role{
			{
				Id:          uint32(RoleSudo),
				Sid:         "sudo",
				Description: "Sudo role",
			},
			{
				Id:          uint32(RoleValidator),
				Sid:         "validator",
				Description: "Validator role",
			},
		},
		RolePermissions: map[uint64]*Permissions{
			RoleSudo: NewPermissions([]PermValue{
				PermSetPermissions,
				PermClaimValidator,
				PermClaimCouncilor,
				PermUpsertTokenAlias,
				// PermChangeTxFee, // do not give this permission to sudo account - test does not pass
				PermUpsertTokenRate,
				PermUpsertRole,
				PermCreateSetNetworkPropertyProposal,
				PermVoteSetNetworkPropertyProposal,
				PermCreateUpsertDataRegistryProposal,
				PermVoteUpsertDataRegistryProposal,
				PermCreateUpsertTokenAliasProposal,
				PermVoteUpsertTokenAliasProposal,
				PermCreateUpsertTokenRateProposal,
				PermVoteUpsertTokenRateProposal,
				PermCreateUnjailValidatorProposal,
				PermVoteUnjailValidatorProposal,
				PermCreateRoleProposal,
				PermVoteCreateRoleProposal,
				PermCreateSetProposalDurationProposal,
				PermVoteSetProposalDurationProposal,
				PermCreateTokensWhiteBlackChangeProposal,
				PermVoteTokensWhiteBlackChangeProposal,
				PermCreateSetPoorNetworkMessagesProposal,
				PermVoteSetPoorNetworkMessagesProposal,
				PermWhitelistAccountPermissionProposal,
				PermVoteWhitelistAccountPermissionProposal,
				PermCreateResetWholeValidatorRankProposal,
				PermVoteResetWholeValidatorRankProposal,
				PermCreateSoftwareUpgradeProposal,
				PermVoteSoftwareUpgradeProposal,
				PermSetClaimValidatorPermission,
				PermBlacklistAccountPermissionProposal,
				PermVoteBlacklistAccountPermissionProposal,
				PermRemoveWhitelistedAccountPermissionProposal,
				PermVoteRemoveWhitelistedAccountPermissionProposal,
				PermRemoveBlacklistedAccountPermissionProposal,
				PermVoteRemoveBlacklistedAccountPermissionProposal,
				PermWhitelistRolePermissionProposal,
				PermVoteWhitelistRolePermissionProposal,
				PermBlacklistRolePermissionProposal,
				PermVoteBlacklistRolePermissionProposal,
				PermRemoveWhitelistedRolePermissionProposal,
				PermVoteRemoveWhitelistedRolePermissionProposal,
				PermRemoveBlacklistedRolePermissionProposal,
				PermVoteRemoveBlacklistedRolePermissionProposal,
				PermAssignRoleToAccountProposal,
				PermVoteAssignRoleToAccountProposal,
				PermUnassignRoleFromAccountProposal,
				PermVoteUnassignRoleFromAccountProposal,
				PermRemoveRoleProposal,
				PermVoteRemoveRoleProposal,
				PermCreateUpsertUBIProposal,
				PermVoteUpsertUBIProposal,
				PermCreateRemoveUBIProposal,
				PermVoteRemoveUBIProposal,
			}, nil),
			uint64(RoleValidator): NewPermissions([]PermValue{PermClaimValidator}, nil),
		},
		StartingProposalId: 1,
		NetworkProperties: &NetworkProperties{
			MinTxFee:                    100,
			MaxTxFee:                    1000000,
			VoteQuorum:                  33,
			MinimumProposalEndTime:      300, // 300 seconds / 5 mins
			ProposalEnactmentTime:       300, // 300 seconds / 5 mins
			MinProposalEndBlocks:        2,
			MinProposalEnactmentBlocks:  1,
			EnableForeignFeePayments:    true,
			MischanceRankDecreaseAmount: 10,
			MischanceConfidence:         10,
			MaxMischance:                110,
			InactiveRankDecreasePercent: 50,      // 50%
			PoorNetworkMaxBankSend:      1000000, // 1M ukex
			MinValidators:               1,
			UnjailMaxTime:               600, // 600  seconds / 10 mins
			EnableTokenWhitelist:        false,
			EnableTokenBlacklist:        true,
			MinIdentityApprovalTip:      200,
			UniqueIdentityKeys:          "moniker,username",
			UbiHardcap:                  6000_000,
			ValidatorsFeeShare:          50,
			InflationRate:               18,       // 18%
			InflationPeriod:             31557600, // 1 year
			UnstakingPeriod:             2629800,  // 1 month
			MaxDelegators:               100,
			MinDelegationPushout:        10,
		},
		ExecutionFees: []*ExecutionFee{
			{
				TransactionType:   kiratypes.MsgTypeClaimValidator,
				ExecutionFee:      100,
				FailureFee:        1,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				TransactionType:   kiratypes.MsgTypeClaimCouncilor,
				ExecutionFee:      100,
				FailureFee:        1,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				TransactionType:   "claim-proposal-type-x",
				ExecutionFee:      100,
				FailureFee:        1,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				TransactionType:   "vote-proposal-type-x",
				ExecutionFee:      100,
				FailureFee:        1,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				TransactionType:   "submit-proposal-type-x",
				ExecutionFee:      10,
				FailureFee:        1,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				TransactionType:   "veto-proposal-type-x",
				ExecutionFee:      100,
				FailureFee:        1,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				TransactionType:   kiratypes.MsgTypeUpsertTokenAlias,
				ExecutionFee:      100,
				FailureFee:        1,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				TransactionType:   kiratypes.MsgTypeActivate,
				ExecutionFee:      100,
				FailureFee:        1000,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				TransactionType:   kiratypes.MsgTypePause,
				ExecutionFee:      100,
				FailureFee:        100,
				Timeout:           10,
				DefaultParameters: 0,
			},
			{
				TransactionType:   kiratypes.MsgTypeUnpause,
				ExecutionFee:      100,
				FailureFee:        100,
				Timeout:           10,
				DefaultParameters: 0,
			},
		},
		PoorNetworkMessages: &AllowedMessages{
			Messages: []string{
				kiratypes.MsgTypeSubmitProposal,
				kiratypes.MsgTypeSetNetworkProperties,
				kiratypes.MsgTypeVoteProposal,
				kiratypes.MsgTypeClaimCouncilor,
				kiratypes.MsgTypeWhitelistPermissions,
				kiratypes.MsgTypeBlacklistPermissions,
				kiratypes.MsgTypeCreateRole,
				kiratypes.MsgTypeAssignRole,
				kiratypes.MsgTypeRemoveRole,
				kiratypes.MsgTypeWhitelistRolePermission,
				kiratypes.MsgTypeBlacklistRolePermission,
				kiratypes.MsgTypeRemoveWhitelistRolePermission,
				kiratypes.MsgTypeRemoveBlacklistRolePermission,
				kiratypes.MsgTypeClaimValidator,
				kiratypes.MsgTypeActivate,
				kiratypes.MsgTypePause,
				kiratypes.MsgTypeUnpause,
				kiratypes.MsgTypeRegisterIdentityRecords,
				kiratypes.MsgTypeEditIdentityRecord,
				kiratypes.MsgTypeRequestIdentityRecordsVerify,
				kiratypes.MsgTypeHandleIdentityRecordsVerifyRequest,
				kiratypes.MsgTypeCancelIdentityRecordsVerifyRequest,
			},
		},
		LastIdentityRecordId:        0,
		LastIdRecordVerifyRequestId: 0,
	}
}
