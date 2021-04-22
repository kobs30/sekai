package keeper

import (
	"github.com/KiraCore/sekai/x/staking/types"
	customstakingtypes "github.com/KiraCore/sekai/x/staking/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// Activate a validator
func (k Keeper) Activate(ctx sdk.Context, valAddress sdk.ValAddress) error {
	validator, err := k.GetValidator(ctx, valAddress)
	if err != nil {
		return err
	}

	if validator.IsPaused() {
		return sdkerrors.Wrap(customstakingtypes.ErrValidatorPaused, "Can NOT activate paused validator, you must unpause")
	}

	k.setStatusToValidator(ctx, validator, customstakingtypes.Active)
	k.addReactivatingValidator(ctx, validator)
	k.RemoveRemovingValidator(ctx, validator)

	return nil
}

// ResetWholeValidatorRank reset whole validators' status, rank and streak
func (k Keeper) ResetWholeValidatorRank(ctx sdk.Context) {
	// TODO: is it correct to use this iterator @Jonathan?
	k.IterateValidators(ctx, func(index int64, validator *types.Validator) (stop bool) {
		validator.Status = customstakingtypes.Active
		validator.Rank = 0
		validator.Streak = 0
		k.AddValidator(ctx, *validator)
		return false
	})
}

// Inactivate inactivate the validator
func (k Keeper) Inactivate(ctx sdk.Context, valAddress sdk.ValAddress) error { // inactivate a validator
	validator, err := k.GetValidator(ctx, valAddress)
	if err != nil {
		return err
	}

	if validator.IsPaused() {
		return customstakingtypes.ErrValidatorPaused
	}

	networkProperties := k.govkeeper.GetNetworkProperties(ctx)
	validator.Status = customstakingtypes.Inactive
	validator.Rank = validator.Rank * int64(100-networkProperties.InactiveRankDecreasePercent) / 100

	k.AddValidator(ctx, validator)
	k.addRemovingValidator(ctx, validator)
	k.RemoveReactivatingValidator(ctx, validator)

	return nil
}

// HandleValidatorSignature manage rank and streak by block miss / sign result
func (k Keeper) HandleValidatorSignature(ctx sdk.Context, valAddress sdk.ValAddress, missed bool) error {
	validator, err := k.GetValidator(ctx, valAddress)
	if err != nil {
		return err
	}
	networkProperties := k.govkeeper.GetNetworkProperties(ctx)
	if missed {
		// set validator streak by 0 and decrease rank by X
		validator.Streak = 0
		validator.Rank -= int64(networkProperties.MischanceRankDecreaseAmount)
		if validator.Rank < 0 {
			validator.Rank = 0
		}
	} else {
		// increase streak and reset rank if streak is higher than rank
		validator.Streak++
		if validator.Streak > validator.Rank {
			validator.Rank = validator.Streak
		}
	}
	k.AddValidator(ctx, validator)
	return nil
}

// Pause a validator
func (k Keeper) Pause(ctx sdk.Context, valAddress sdk.ValAddress) error {
	validator, err := k.GetValidator(ctx, valAddress)
	if err != nil {
		return err
	}

	if validator.IsInactivated() {
		return customstakingtypes.ErrValidatorInactive
	}

	k.setStatusToValidator(ctx, validator, customstakingtypes.Paused)
	k.addRemovingValidator(ctx, validator)
	k.RemoveReactivatingValidator(ctx, validator)

	return nil
}

// Unpause unpause the validator
func (k Keeper) Unpause(ctx sdk.Context, valAddress sdk.ValAddress) error { // inactivate a validator
	validator, err := k.GetValidator(ctx, valAddress)
	if err != nil {
		return err
	}

	if validator.IsInactivated() {
		return customstakingtypes.ErrValidatorInactive
	}

	k.setStatusToValidator(ctx, validator, customstakingtypes.Active)
	k.addReactivatingValidator(ctx, validator)
	k.RemoveRemovingValidator(ctx, validator)

	return nil
}

// Jail a validator
func (k Keeper) Jail(ctx sdk.Context, valAddress sdk.ValAddress) error {
	validator, err := k.GetValidator(ctx, valAddress)
	if err != nil {
		return err
	}

	k.setStatusToValidator(ctx, validator, customstakingtypes.Jailed)
	k.addRemovingValidator(ctx, validator)
	k.setJailValidatorInfo(ctx, validator)
	k.RemoveReactivatingValidator(ctx, validator)

	return nil
}

// GetValidatorJailInfo returns information about a jailed validor, found is false
// if there is no validator, so a validator that is not jailed should return false.
func (k Keeper) GetValidatorJailInfo(ctx sdk.Context, valAddress sdk.ValAddress) (customstakingtypes.ValidatorJailInfo, bool) {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(GetValidatorJailInfoKey(valAddress))
	if bz == nil {
		return customstakingtypes.ValidatorJailInfo{}, false
	}

	var info customstakingtypes.ValidatorJailInfo
	k.cdc.MustUnmarshalBinaryBare(bz, &info)

	return info, true
}

// Unjail a validator
func (k Keeper) Unjail(ctx sdk.Context, valAddress sdk.ValAddress) error {
	validator, err := k.GetValidator(ctx, valAddress)
	if err != nil {
		return err
	}

	k.setStatusToValidator(ctx, validator, customstakingtypes.Active)
	k.addReactivatingValidator(ctx, validator)
	k.removeJailValidatorInfo(ctx, validator)
	k.RemoveRemovingValidator(ctx, validator)

	return nil
}

func (k Keeper) setStatusToValidator(ctx sdk.Context, validator customstakingtypes.Validator, status customstakingtypes.ValidatorStatus) {
	validator.Status = status
	k.AddValidator(ctx, validator)
}

func (k Keeper) setJailValidatorInfo(ctx sdk.Context, validator customstakingtypes.Validator) {
	jailInfo := customstakingtypes.ValidatorJailInfo{
		Time: ctx.BlockTime(),
	}

	bz := k.cdc.MustMarshalBinaryBare(jailInfo)

	store := ctx.KVStore(k.storeKey)
	store.Set(GetValidatorJailInfoKey(validator.ValKey), bz)
}

func (k Keeper) removeJailValidatorInfo(ctx sdk.Context, validator customstakingtypes.Validator) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(GetValidatorJailInfoKey(validator.ValKey))
}

// GetRemovingValidatorSet returns the keys of the validators that needs to be removed from
// the set.
func (k Keeper) GetRemovingValidatorSet(ctx sdk.Context) [][]byte {
	store := ctx.KVStore(k.storeKey)

	iter := sdk.KVStorePrefixIterator(store, RemovingValidatorQueue)
	defer iter.Close()

	var validatorKeys [][]byte
	for ; iter.Valid(); iter.Next() {
		validatorKeys = append(validatorKeys, iter.Value())
	}

	return validatorKeys
}

// GetReactivatingValidatorSet returns the keys of the validators need to be reactivated, this
// is used in the Enblock function to reactivate those validators who have been unpaused
// or activated.
func (k Keeper) GetReactivatingValidatorSet(ctx sdk.Context) [][]byte {
	store := ctx.KVStore(k.storeKey)

	iter := sdk.KVStorePrefixIterator(store, ReactivatingValidatorQueue)
	defer iter.Close()

	var validatorKeys [][]byte
	for ; iter.Valid(); iter.Next() {
		validatorKeys = append(validatorKeys, iter.Value())
	}

	return validatorKeys
}

func (k Keeper) addRemovingValidator(ctx sdk.Context, validator types.Validator) {
	store := ctx.KVStore(k.storeKey)
	store.Set(GetRemovingValidatorKey(validator.ValKey), validator.ValKey)
}

func (k Keeper) RemoveRemovingValidator(ctx sdk.Context, validator types.Validator) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(GetRemovingValidatorKey(validator.ValKey))
}

func (k Keeper) addReactivatingValidator(ctx sdk.Context, validator types.Validator) {
	store := ctx.KVStore(k.storeKey)
	store.Set(GetReactivatingValidatorKey(validator.ValKey), validator.ValKey)
}

func (k Keeper) RemoveReactivatingValidator(ctx sdk.Context, validator types.Validator) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(GetReactivatingValidatorKey(validator.ValKey))
}
