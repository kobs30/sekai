package keeper

import (
	"github.com/KiraCore/sekai/x/custody/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"
)

// Keeper is for managing token module
type Keeper struct {
	cdc      codec.BinaryCodec
	storeKey sdk.StoreKey
}

// NewKeeper returns instance of a keeper
func NewKeeper(storeKey sdk.StoreKey, cdc codec.BinaryCodec) Keeper {
	return Keeper{
		cdc:      cdc,
		storeKey: storeKey,
	}
}

// BondDenom returns the denom that is basically used for fee payment
func (k Keeper) BondDenom(ctx sdk.Context) string {
	return "ukex"
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}