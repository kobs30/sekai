package types

import (
	stakingtypes "github.com/KiraCore/sekai/x/staking/types"
	tokenstypes "github.com/KiraCore/sekai/x/tokens/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// StakingKeeper expected staking keeper
type StakingKeeper interface {
	GetValidator(sdk.Context, sdk.ValAddress) (stakingtypes.Validator, error)
}

type BankKeeper interface {
	GetAllBalances(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins
	MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	SendCoinsFromModuleToModule(ctx sdk.Context, senderModule, recipientModule string, amt sdk.Coins) error
}

// TokensKeeper defines expected interface needed to get token rate
type TokensKeeper interface {
	GetTokenRate(ctx sdk.Context, denom string) *tokenstypes.TokenRate
}