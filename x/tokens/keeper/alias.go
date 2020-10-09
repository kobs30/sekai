package keeper

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/KiraCore/sekai/x/tokens/types"
)

// GetTokenAlias returns a token alias
func (k Keeper) GetTokenAlias(ctx sdk.Context, symbol string) *types.TokenAlias {
	prefixStore := prefix.NewStore(ctx.KVStore(k.storeKey), PrefixKeyTokenAlias)
	bz := prefixStore.Get([]byte(symbol))
	if bz == nil {
		return nil
	}

	alias := new(types.TokenAlias)
	k.cdc.MustUnmarshalBinaryBare(bz, alias)

	return alias
}

// ListTokenAlias returns all list of token alias
func (k Keeper) ListTokenAlias(ctx sdk.Context) []types.TokenAlias {
	var tokenAliases []types.TokenAlias

	// get iterator for token aliases
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(PrefixKeyTokenAlias))

	for ; iterator.Valid(); iterator.Next() {
		symbol := string(iterator.Key())
		tokenAlias := k.GetTokenAlias(ctx, symbol)
		if tokenAlias != nil {
			tokenAliases = append(tokenAliases, *tokenAlias)
		}
	}
	return tokenAliases
}

// UpsertTokenAlias upsert a token alias to the registry
func (k Keeper) UpsertTokenAlias(ctx sdk.Context, alias types.TokenAlias) error {
	store := ctx.KVStore(k.storeKey)
	// we use symbol of TokenAlias as an ID inside KVStore storage
	tokenAliasStoreID := append([]byte(PrefixKeyTokenAlias), []byte(alias.Symbol)...)

	for _, denom := range alias.Denoms {
		denomTokenStoreID := append([]byte(PrefixKeyDenomToken), []byte(denom)...)

		if store.Has(denomTokenStoreID) {
			symbol := string(store.Get(denomTokenStoreID))
			if symbol != alias.Symbol {
				return fmt.Errorf("%s denom is already registered for %s token alias", denom, symbol)
			}
		}
		store.Set(denomTokenStoreID, []byte(alias.Symbol))
	}
	store.Set(tokenAliasStoreID, k.cdc.MustMarshalBinaryBare(&alias))
	return nil
}

// DeleteTokenAlias delete token alias by symbol
func (k Keeper) DeleteTokenAlias(ctx sdk.Context, symbol string) error {
	store := ctx.KVStore(k.storeKey)
	// we use symbol of TokenAlias as an ID inside KVStore storage
	tokenAliasStoreID := append([]byte(PrefixKeyTokenAlias), []byte(symbol)...)

	if !store.Has(tokenAliasStoreID) {
		return fmt.Errorf("no alias is available for %s symbol", symbol)
	}

	var alias types.TokenAlias
	bz := store.Get(tokenAliasStoreID)
	k.cdc.MustUnmarshalBinaryBare(bz, &alias)

	for _, denom := range alias.Denoms {
		denomTokenStoreID := append([]byte(PrefixKeyDenomToken), []byte(denom)...)
		store.Delete(denomTokenStoreID)
	}

	store.Delete(tokenAliasStoreID)
	return nil
}
