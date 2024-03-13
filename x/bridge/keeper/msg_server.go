package keeper

import (
	"context"
	"github.com/KiraCore/sekai/x/bridge/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgServer struct {
	keeper Keeper
}

// NewMsgServerImpl returns an implementation of the bank MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{
		keeper: keeper,
	}
}

var _ types.MsgServer = msgServer{}

func (s msgServer) ChangeCosmosEtherium(goCtx context.Context, msg *types.MsgChangeCosmosEtherium) (*types.MsgChangeCosmosEtheriumResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	record := types.ChangeComsosEtheriumRecord{
		From:      msg.From,
		To:        msg.To,
		InAmount:  msg.InAmount,
		OutAmount: msg.OutAmount,
	}

	s.keeper.SetChangeCosmosEtheriumRecord(ctx, record)

	return &types.MsgChangeCosmosEtheriumResponse{}, nil
}

func (s msgServer) ChangeEtheriumCosmos(goCtx context.Context, msg *types.MsgChangeEtheriumCosmos) (*types.MsgChangeEtheriumCosmosResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	record := types.ChangeEtheriumComsosRecord{
		From:      msg.From,
		To:        msg.To,
		InAmount:  msg.InAmount,
		OutAmount: msg.OutAmount,
	}

	s.keeper.SetChangeEtheriumCosmosRecord(ctx, record)

	return &types.MsgChangeEtheriumCosmosResponse{}, nil
}
