package keeper

import (
	"context"
	"github.com/KiraCore/sekai/x/bridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Querier struct {
	keeper Keeper
}

func NewQuerier(keeper Keeper) types.QueryServer {
	return &Querier{keeper: keeper}
}

var _ types.QueryServer = Querier{}

func (q Querier) ChangeCosmosEtheriumByAddress(c context.Context, request *types.ChangeCosmosEtheriumByAddressRequest) (*types.ChangeCosmosEtheriumByAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	record := q.keeper.GetChangeCosmosEtheriumRecord(ctx, request.Addr)

	return &types.ChangeCosmosEtheriumByAddressResponse{
		From:      record.From,
		To:        record.To,
		InAmount:  record.InAmount,
		OutAmount: record.OutAmount,
	}, nil
}

func (q Querier) ChangeEtheriumCosmosByAddress(c context.Context, request *types.ChangeEtheriumCosmosByAddressRequest) (*types.ChangeEtheriumCosmosByAddressResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	record := q.keeper.GetChangeEtheriumCosmosRecord(ctx, request.Addr)

	return &types.ChangeEtheriumCosmosByAddressResponse{
		From:      record.From,
		To:        record.To,
		InAmount:  record.InAmount,
		OutAmount: record.OutAmount,
	}, nil
}
