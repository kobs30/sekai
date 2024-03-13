package keeper

import (
	"github.com/KiraCore/sekai/x/bridge/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) SetChangeCosmosEtheriumRecord(ctx sdk.Context, record types.ChangeComsosEtheriumRecord) {
	store := ctx.KVStore(k.storeKey)
	key := append([]byte(types.PrefixKeyBridgeCosmosEtheriumRecord), record.From...)

	store.Set(key, k.cdc.MustMarshal(&record))
}

func (k Keeper) SetChangeEtheriumCosmosRecord(ctx sdk.Context, record types.ChangeEtheriumComsosRecord) {
	store := ctx.KVStore(k.storeKey)
	key := append([]byte(types.PrefixKeyBridgeEtheriumCosmosRecord), record.To...)

	store.Set(key, k.cdc.MustMarshal(&record))
}

func (k Keeper) GetChangeCosmosEtheriumRecord(ctx sdk.Context, address sdk.AccAddress) *types.ChangeComsosEtheriumRecord {
	prefixStore := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PrefixKeyBridgeCosmosEtheriumRecord))
	bz := prefixStore.Get(address)

	if bz == nil {
		return nil
	}

	info := new(types.ChangeComsosEtheriumRecord)
	k.cdc.MustUnmarshal(bz, info)

	return info
}

func (k Keeper) GetChangeEtheriumCosmosRecord(ctx sdk.Context, address sdk.AccAddress) *types.ChangeEtheriumComsosRecord {
	prefixStore := prefix.NewStore(ctx.KVStore(k.storeKey), []byte(types.PrefixKeyBridgeCosmosEtheriumRecord))
	bz := prefixStore.Get(address)

	if bz == nil {
		return nil
	}

	info := new(types.ChangeEtheriumComsosRecord)
	k.cdc.MustUnmarshal(bz, info)

	return info
}
