package keeper

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"github.com/KiraCore/sekai/x/custody/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"time"
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

func (s msgServer) CreateCustody(goCtx context.Context, msg *types.MsgCreteCustodyRecord) (*types.MsgCreteCustodyRecordResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	record := types.CustodyRecord{
		Address:         msg.Address,
		CustodySettings: &msg.CustodySettings,
	}

	s.keeper.SetCustodyRecord(ctx, record)

	return &types.MsgCreteCustodyRecordResponse{}, nil
}

func (s msgServer) AddToCustodians(goCtx context.Context, msg *types.MsgAddToCustodyCustodians) (*types.MsgAddToCustodyCustodiansResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	record := types.CustodyCustodiansRecord{
		Address:           msg.Address,
		CustodyCustodians: s.keeper.GetCustodyCustodiansByAddress(ctx, msg.Address),
	}

	if record.CustodyCustodians == nil {
		record.CustodyCustodians = new(types.CustodyCustodianList)
		record.CustodyCustodians.Addresses = map[string]bool{}
	}

	for _, address := range msg.AddAddress {
		record.CustodyCustodians.Addresses[address.String()] = true
	}

	keyRecord := types.CustodyKeyRecord{
		Address: msg.Address,
		Key:     msg.NewKey,
	}

	s.keeper.SetCustodyRecordKey(ctx, keyRecord)
	s.keeper.AddToCustodyCustodians(ctx, record)

	return &types.MsgAddToCustodyCustodiansResponse{}, nil
}

func (s msgServer) RemoveFromCustodians(goCtx context.Context, msg *types.MsgRemoveFromCustodyCustodians) (*types.MsgRemoveFromCustodyCustodiansResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	record := types.CustodyCustodiansRecord{
		Address:           msg.Address,
		CustodyCustodians: s.keeper.GetCustodyCustodiansByAddress(ctx, msg.Address),
	}

	if record.CustodyCustodians == nil {
		return nil, errors.Wrap(types.ErrNoWhiteLists, "Can not remove from the empty whitelist")
	}

	if !record.CustodyCustodians.Addresses[msg.RemoveAddress.String()] {
		return nil, errors.Wrap(types.ErrNoWhiteListsElement, "Can not remove missing element from the whitelist")
	}

	record.CustodyCustodians.Addresses[msg.RemoveAddress.String()] = false
	s.keeper.AddToCustodyCustodians(ctx, record)

	return &types.MsgRemoveFromCustodyCustodiansResponse{}, nil
}

func (s msgServer) AddToCustodyPool(goCtx context.Context, msg *types.MsgAddToCustodyPool) (*types.MsgAddToCustodyPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	record := types.CustodyPool{
		Address:      msg.Address,
		Transactions: s.keeper.GetCustodyPoolByAddress(ctx, msg.Address),
	}

	TxBoby := msg.Transaction
	hash := sha256.Sum256(TxBoby.Value)
	hashString := hex.EncodeToString(hash[:])

	transactionRecord := types.TransactionRecord{
		Transaction: msg.Transaction,
		Votes:       0,
	}

	if record.Transactions == nil {
		record.Transactions = new(types.TransactionPool)
		record.Transactions.Record = map[string]*types.TransactionRecord{}
	}

	if record.Transactions.Record[hashString].Transaction != nil {
		record.Transactions.Record[hashString].Votes += 1
	} else {
		record.Transactions.Record[hashString] = &transactionRecord
	}

	s.keeper.AddToCustodyPool(ctx, record)

	return &types.MsgAddToCustodyPoolResponse{}, nil
}

func (s msgServer) RemoveFromCustodyPool(goCtx context.Context, msg *types.MsgRemoveFromCustodyPool) (*types.MsgRemoveFromCustodyPoolResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	record := types.CustodyPool{
		Address:      msg.Address,
		Transactions: s.keeper.GetCustodyPoolByAddress(ctx, msg.Address),
	}

	body := msg.Transaction
	hash := sha256.Sum256(body.Value)
	hashString := hex.EncodeToString(hash[:])

	record.Transactions.Record[hashString] = nil

	s.keeper.AddToCustodyPool(ctx, record)

	return &types.MsgRemoveFromCustodyPoolResponse{}, nil
}

func (s msgServer) DropCustodians(goCtx context.Context, msg *types.MsgDropCustodyCustodians) (*types.MsgDropCustodyCustodiansResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	s.keeper.DropCustodyCustodiansByAddress(ctx, msg.Address)

	return &types.MsgDropCustodyCustodiansResponse{}, nil
}

func (s msgServer) AddToWhiteList(goCtx context.Context, msg *types.MsgAddToCustodyWhiteList) (*types.MsgAddToCustodyWhiteListResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	record := types.CustodyWhiteListRecord{
		Address:          msg.Address,
		CustodyWhiteList: s.keeper.GetCustodyWhiteListByAddress(ctx, msg.Address),
	}

	if record.CustodyWhiteList == nil {
		record.CustodyWhiteList = new(types.CustodyWhiteList)
		record.CustodyWhiteList.Addresses = map[string]bool{}
	}

	for _, address := range msg.AddAddress {
		record.CustodyWhiteList.Addresses[address.String()] = true
	}

	s.keeper.AddToCustodyWhiteList(ctx, record)

	return &types.MsgAddToCustodyWhiteListResponse{}, nil
}

func (s msgServer) RemoveFromWhiteList(goCtx context.Context, msg *types.MsgRemoveFromCustodyWhiteList) (*types.MsgRemoveFromCustodyWhiteListResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	record := types.CustodyWhiteListRecord{
		Address:          msg.Address,
		CustodyWhiteList: s.keeper.GetCustodyWhiteListByAddress(ctx, msg.Address),
	}

	if record.CustodyWhiteList == nil {
		return nil, errors.Wrap(types.ErrNoWhiteLists, "Can not remove from the empty whitelist")
	}

	if !record.CustodyWhiteList.Addresses[msg.RemoveAddress.String()] {
		return nil, errors.Wrap(types.ErrNoWhiteListsElement, "Can not remove missing element from the whitelist")
	}

	record.CustodyWhiteList.Addresses[msg.RemoveAddress.String()] = false
	s.keeper.AddToCustodyWhiteList(ctx, record)

	return &types.MsgRemoveFromCustodyWhiteListResponse{}, nil
}

func (s msgServer) DropWhiteList(goCtx context.Context, msg *types.MsgDropCustodyWhiteList) (*types.MsgDropCustodyWhiteListResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	s.keeper.DropCustodyWhiteListByAddress(ctx, msg.Address)

	return &types.MsgDropCustodyWhiteListResponse{}, nil
}

func (s msgServer) AddToLimits(goCtx context.Context, msg *types.MsgAddToCustodyLimits) (*types.MsgAddToCustodyLimitsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	record := types.CustodyLimitRecord{
		Address:       msg.Address,
		CustodyLimits: s.keeper.GetCustodyLimitsByAddress(ctx, msg.Address),
	}

	if record.CustodyLimits == nil {
		record.CustodyLimits = new(types.CustodyLimits)
		record.CustodyLimits.Limits = map[string]*types.CustodyLimit{}
	}

	custodyLimit := types.CustodyLimit{
		Amount: msg.Amount,
		Limit:  msg.Limit,
	}

	record.CustodyLimits.Limits[msg.Denom] = &custodyLimit
	s.keeper.AddToCustodyLimits(ctx, record)

	return &types.MsgAddToCustodyLimitsResponse{}, nil
}

func (s msgServer) RemoveFromLimits(goCtx context.Context, msg *types.MsgRemoveFromCustodyLimits) (*types.MsgRemoveFromCustodyLimitsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	record := types.CustodyLimitRecord{
		Address:       msg.Address,
		CustodyLimits: s.keeper.GetCustodyLimitsByAddress(ctx, msg.Address),
	}

	if record.CustodyLimits == nil {
		return nil, errors.Wrap(types.ErrNoWhiteLists, "Can not remove from the empty limits")
	}

	if record.CustodyLimits.Limits[msg.Denom] == nil {
		return nil, errors.Wrap(types.ErrNoWhiteListsElement, "Can not remove missing element from the limits")
	}

	record.CustodyLimits.Limits[msg.Denom] = new(types.CustodyLimit)
	s.keeper.AddToCustodyLimits(ctx, record)

	return &types.MsgRemoveFromCustodyLimitsResponse{}, nil
}

func (s msgServer) DropLimits(goCtx context.Context, msg *types.MsgDropCustodyLimits) (*types.MsgDropCustodyLimitsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	s.keeper.DropCustodyLimitsByAddress(ctx, msg.Address)

	return &types.MsgDropCustodyLimitsResponse{}, nil
}

func (s msgServer) AddToLimitsStatus(goCtx context.Context, msg *types.MsgAddToCustodyLimitsStatus) (*types.MsgAddToCustodyLimitsStatusResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	record := types.CustodyLimitStatusRecord{
		Address:         msg.Address,
		CustodyStatuses: s.keeper.GetCustodyLimitsStatusByAddress(ctx, msg.Address),
	}

	if record.CustodyStatuses == nil {
		record.CustodyStatuses = new(types.CustodyStatuses)
		record.CustodyStatuses.Statuses = map[string]*types.CustodyStatus{}
	}

	custodyStatus := types.CustodyStatus{
		Amount: msg.Amount,
		Time:   time.Now().Unix(),
	}

	record.CustodyStatuses.Statuses[msg.Denom] = &custodyStatus
	s.keeper.AddToCustodyLimitsStatus(ctx, record)

	return &types.MsgAddToCustodyLimitsStatusResponse{}, nil
}