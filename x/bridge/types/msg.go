package types

import (
	"github.com/KiraCore/sekai/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func NewMsgChangeCosmosEtherium(from sdk.AccAddress, to string, in_amount sdk.Coins, out_amount int64) *MsgChangeCosmosEtherium {
	return &MsgChangeCosmosEtherium{from, to, in_amount, out_amount}
}

func (m *MsgChangeCosmosEtherium) Route() string {
	return ModuleName
}

func (m *MsgChangeCosmosEtherium) Type() string {
	return types.MsgTypeChangeCosmosEtherium
}

func (m *MsgChangeCosmosEtherium) ValidateBasic() error {
	return nil
}

func (m *MsgChangeCosmosEtherium) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgChangeCosmosEtherium) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{
		m.From,
	}
}

func NewMsgChangeEtheriumCosmos(addr sdk.AccAddress, from string, to sdk.AccAddress, in_amount uint64, out_amount sdk.Coin) *MsgChangeEtheriumCosmos {
	return &MsgChangeEtheriumCosmos{addr, from, to, in_amount, out_amount}
}

func (m *MsgChangeEtheriumCosmos) Route() string {
	return ModuleName
}

func (m *MsgChangeEtheriumCosmos) Type() string {
	return types.MsgTypeChangeCosmosEtherium
}

func (m *MsgChangeEtheriumCosmos) ValidateBasic() error {
	return nil
}

func (m *MsgChangeEtheriumCosmos) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(m)
	return sdk.MustSortJSON(bz)
}

func (m *MsgChangeEtheriumCosmos) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{
		m.Addr,
	}
}
