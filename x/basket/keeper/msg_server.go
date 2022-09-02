package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/KiraCore/sekai/x/basket/types"
)

type msgServer struct {
	keeper Keeper
	cgk    types.CustomGovKeeper
}

// NewMsgServerImpl returns an implementation of the bank MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper, cgk types.CustomGovKeeper) types.MsgServer {
	return &msgServer{
		keeper: keeper,
		cgk:    cgk,
	}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) DisableBasketDeposits(
	goCtx context.Context,
	msg *types.MsgDisableBasketDeposits,
) (*types.MsgDisableBasketDepositsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_ = ctx
	return &types.MsgDisableBasketDepositsResponse{}, nil
}

func (k msgServer) DisableBasketWithdraws(
	goCtx context.Context,
	msg *types.MsgDisableBasketDeposits,
) (*types.MsgDisableBasketDepositsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_ = ctx
	return &types.MsgDisableBasketDepositsResponse{}, nil
}

func (k msgServer) DisableBasketSwaps(
	goCtx context.Context,
	msg *types.MsgDisableBasketSwaps,
) (*types.MsgDisableBasketSwapsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_ = ctx
	return &types.MsgDisableBasketSwapsResponse{}, nil
}

func (k msgServer) BasketTokenMint(
	goCtx context.Context,
	msg *types.MsgBasketTokenMint,
) (*types.MsgBasketTokenMintResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_ = ctx
	return &types.MsgBasketTokenMintResponse{}, nil
}

func (k msgServer) BasketTokenBurn(
	goCtx context.Context,
	msg *types.MsgBasketTokenBurn,
) (*types.MsgBasketTokenBurnResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_ = ctx
	return &types.MsgBasketTokenBurnResponse{}, nil
}

func (k msgServer) BasketTokenSwap(
	goCtx context.Context,
	msg *types.MsgBasketTokenSwap,
) (*types.MsgBasketTokenSwapResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_ = ctx
	return &types.MsgBasketTokenSwapResponse{}, nil
}

func (k msgServer) BasketClaimRewards(
	goCtx context.Context,
	msg *types.MsgBasketClaimRewards,
) (*types.MsgBasketClaimRewardsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	_ = ctx
	return &types.MsgBasketClaimRewardsResponse{}, nil
}