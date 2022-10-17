package types

import (
	"github.com/KiraCore/sekai/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func NewMsgCreateCustody(addr sdk.AccAddress, custodySettings CustodySettings, oldKey string, newKey string) *MsgCreteCustodyRecord {
	return &MsgCreteCustodyRecord{addr, custodySettings, oldKey, newKey}
}

func (m *MsgCreteCustodyRecord) Route() string {
	return ModuleName
}

func (m *MsgCreteCustodyRecord) Type() string {
	return types.MsgTypeCreateCustody
}

func (m *MsgCreteCustodyRecord) ValidateBasic() error {
	return nil
}

func (m *MsgCreteCustodyRecord) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgCreteCustodyRecord) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{
		m.Address,
	}
}

func NewMsgApproveCustodyTransaction(from sdk.AccAddress, to sdk.AccAddress, hash string) *MsgApproveCustodyTransaction {
	return &MsgApproveCustodyTransaction{from, to, hash}
}

func (m *MsgApproveCustodyTransaction) Route() string {
	return ModuleName
}

func (m *MsgApproveCustodyTransaction) Type() string {
	return types.MsgTypeAddToCustodyCustodians
}

func (m *MsgApproveCustodyTransaction) ValidateBasic() error {
	return nil
}

func (m *MsgApproveCustodyTransaction) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgApproveCustodyTransaction) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{
		m.FromAddress,
	}
}

func NewMsgDeclineCustodyTransaction(from sdk.AccAddress, to sdk.AccAddress, hash string) *MsgDeclineCustodyTransaction {
	return &MsgDeclineCustodyTransaction{from, to, hash}
}

func (m *MsgDeclineCustodyTransaction) Route() string {
	return ModuleName
}

func (m *MsgDeclineCustodyTransaction) Type() string {
	return types.MsgTypeAddToCustodyCustodians
}

func (m *MsgDeclineCustodyTransaction) ValidateBasic() error {
	return nil
}

func (m *MsgDeclineCustodyTransaction) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgDeclineCustodyTransaction) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{
		m.FromAddress,
	}
}

func NewMsgPasswordConfirmTransaction(from sdk.AccAddress, sender sdk.AccAddress, hash string, password string) *MsgPasswordConfirmTransaction {
	return &MsgPasswordConfirmTransaction{from, sender, hash, password}
}

func (m *MsgPasswordConfirmTransaction) Route() string {
	return ModuleName
}

func (m *MsgPasswordConfirmTransaction) Type() string {
	return types.MsgPasswordConfirmTransaction
}

func (m *MsgPasswordConfirmTransaction) ValidateBasic() error {
	return nil
}

func (m *MsgPasswordConfirmTransaction) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgPasswordConfirmTransaction) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{
		m.FromAddress,
	}
}

func NewMsgAddToCustodyCustodians(addr sdk.AccAddress, newAddr []sdk.AccAddress, oldKey string, newKey string) *MsgAddToCustodyCustodians {
	return &MsgAddToCustodyCustodians{addr, newAddr, oldKey, newKey}
}

func (m *MsgAddToCustodyCustodians) Route() string {
	return ModuleName
}

func (m *MsgAddToCustodyCustodians) Type() string {
	return types.MsgTypeAddToCustodyCustodians
}

func (m *MsgAddToCustodyCustodians) ValidateBasic() error {
	return nil
}

func (m *MsgAddToCustodyCustodians) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgAddToCustodyCustodians) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{
		m.Address,
	}
}

func NewMsgRemoveFromCustodyCustodians(addr sdk.AccAddress, newAddr sdk.AccAddress, oldKey string, newKey string) *MsgRemoveFromCustodyCustodians {
	return &MsgRemoveFromCustodyCustodians{addr, newAddr, oldKey, newKey}
}

func (m *MsgRemoveFromCustodyCustodians) Route() string {
	return ModuleName
}

func (m *MsgRemoveFromCustodyCustodians) Type() string {
	return types.MsgTypeRemoveFromCustodyCustodians
}

func (m *MsgRemoveFromCustodyCustodians) ValidateBasic() error {
	return nil
}

func (m *MsgRemoveFromCustodyCustodians) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgRemoveFromCustodyCustodians) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{
		m.Address,
	}
}

func NewMsgDropCustodyCustodians(addr sdk.AccAddress, oldKey string, newKey string) *MsgDropCustodyCustodians {
	return &MsgDropCustodyCustodians{addr, oldKey, newKey}
}

func (m *MsgDropCustodyCustodians) Route() string {
	return ModuleName
}

func (m *MsgDropCustodyCustodians) Type() string {
	return types.MsgTypeDropCustodyCustodians
}

func (m *MsgDropCustodyCustodians) ValidateBasic() error {
	return nil
}

func (m *MsgDropCustodyCustodians) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgDropCustodyCustodians) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{
		m.Address,
	}
}

func NewMsgAddToCustodyWhiteList(addr sdk.AccAddress, newAddr []sdk.AccAddress, oldKey string, newKey string) *MsgAddToCustodyWhiteList {
	return &MsgAddToCustodyWhiteList{addr, newAddr, oldKey, newKey}
}

func (m *MsgAddToCustodyWhiteList) Route() string {
	return ModuleName
}

func (m *MsgAddToCustodyWhiteList) Type() string {
	return types.MsgTypeAddToCustodyWhiteList
}

func (m *MsgAddToCustodyWhiteList) ValidateBasic() error {
	return nil
}

func (m *MsgAddToCustodyWhiteList) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgAddToCustodyWhiteList) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{
		m.Address,
	}
}

func NewMsgRemoveFromCustodyWhiteList(addr sdk.AccAddress, newAddr sdk.AccAddress, oldKey string, newKey string) *MsgRemoveFromCustodyWhiteList {
	return &MsgRemoveFromCustodyWhiteList{addr, newAddr, oldKey, newKey}
}

func (m *MsgRemoveFromCustodyWhiteList) Route() string {
	return ModuleName
}

func (m *MsgRemoveFromCustodyWhiteList) Type() string {
	return types.MsgTypeRemoveFromCustodyWhiteList
}

func (m *MsgRemoveFromCustodyWhiteList) ValidateBasic() error {
	return nil
}

func (m *MsgRemoveFromCustodyWhiteList) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgRemoveFromCustodyWhiteList) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{
		m.Address,
	}
}

func NewMsgDropCustodyWhiteList(addr sdk.AccAddress, oldKey string, newKey string) *MsgDropCustodyWhiteList {
	return &MsgDropCustodyWhiteList{addr, oldKey, newKey}
}

func (m *MsgDropCustodyWhiteList) Route() string {
	return ModuleName
}

func (m *MsgDropCustodyWhiteList) Type() string {
	return types.MsgTypeDropCustodyWhiteList
}

func (m *MsgDropCustodyWhiteList) ValidateBasic() error {
	return nil
}

func (m *MsgDropCustodyWhiteList) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgDropCustodyWhiteList) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{
		m.Address,
	}
}

func NewMsgAddToCustodyLimits(addr sdk.AccAddress, denom string, amount uint64, limit string, oldKey string, newKey string) *MsgAddToCustodyLimits {
	return &MsgAddToCustodyLimits{addr, denom, amount, limit, oldKey, newKey}
}

func (m *MsgAddToCustodyLimits) Route() string {
	return ModuleName
}

func (m *MsgAddToCustodyLimits) Type() string {
	return types.MsgTypeAddToCustodyWhiteList
}

func (m *MsgAddToCustodyLimits) ValidateBasic() error {
	return nil
}

func (m *MsgAddToCustodyLimits) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgAddToCustodyLimits) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{
		m.Address,
	}
}

func NewMsgRemoveFromCustodyLimits(addr sdk.AccAddress, denom string, oldKey string, newKey string) *MsgRemoveFromCustodyLimits {
	return &MsgRemoveFromCustodyLimits{addr, denom, oldKey, newKey}
}

func (m *MsgRemoveFromCustodyLimits) Route() string {
	return ModuleName
}

func (m *MsgRemoveFromCustodyLimits) Type() string {
	return types.MsgTypeRemoveFromCustodyWhiteList
}

func (m *MsgRemoveFromCustodyLimits) ValidateBasic() error {
	return nil
}

func (m *MsgRemoveFromCustodyLimits) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgRemoveFromCustodyLimits) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{
		m.Address,
	}
}

func NewMsgDropCustodyLimits(addr sdk.AccAddress, oldKey string, newKey string) *MsgDropCustodyLimits {
	return &MsgDropCustodyLimits{addr, oldKey, newKey}
}

func (m *MsgDropCustodyLimits) Route() string {
	return ModuleName
}

func (m *MsgDropCustodyLimits) Type() string {
	return types.MsgTypeDropCustodyWhiteList
}

func (m *MsgDropCustodyLimits) ValidateBasic() error {
	return nil
}

func (m *MsgDropCustodyLimits) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgDropCustodyLimits) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{
		m.Address,
	}
}

func NewMsgSend(fromAddr, toAddr sdk.AccAddress, amount sdk.Coins, password string) *MsgSend {
	return &MsgSend{FromAddress: fromAddr.String(), ToAddress: toAddr.String(), Amount: amount, Password: password}
}

func (m MsgSend) Route() string {
	return ModuleName
}

func (m MsgSend) Type() string {
	return types.MsgTypeSend
}

func (m MsgSend) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(m.FromAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid sender address (%s)", err)
	}

	_, err = sdk.AccAddressFromBech32(m.ToAddress)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "Invalid recipient address (%s)", err)
	}

	if !m.Amount.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, m.Amount.String())
	}

	if !m.Amount.IsAllPositive() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, m.Amount.String())
	}

	return nil
}

func (m MsgSend) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}

func (m MsgSend) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(m.FromAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}