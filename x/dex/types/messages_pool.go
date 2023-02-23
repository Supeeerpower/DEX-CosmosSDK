package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreatePool = "create_pool"
	TypeMsgUpdatePool = "update_pool"
	TypeMsgDeletePool = "delete_pool"
)

var _ sdk.Msg = &MsgCreatePool{}

func NewMsgCreatePool(creator string, amount string, height int32) *MsgCreatePool {
	return &MsgCreatePool{
		Creator: creator,
		Amount:  amount,
		Height:  height,
	}
}

func (msg *MsgCreatePool) Route() string {
	return RouterKey
}

func (msg *MsgCreatePool) Type() string {
	return TypeMsgCreatePool
}

func (msg *MsgCreatePool) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreatePool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreatePool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdatePool{}

func NewMsgUpdatePool(creator string, id uint64, amount string, height int32) *MsgUpdatePool {
	return &MsgUpdatePool{
		Id:      id,
		Creator: creator,
		Amount:  amount,
		Height:  height,
	}
}

func (msg *MsgUpdatePool) Route() string {
	return RouterKey
}

func (msg *MsgUpdatePool) Type() string {
	return TypeMsgUpdatePool
}

func (msg *MsgUpdatePool) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdatePool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdatePool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeletePool{}

func NewMsgDeletePool(creator string, id uint64) *MsgDeletePool {
	return &MsgDeletePool{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeletePool) Route() string {
	return RouterKey
}

func (msg *MsgDeletePool) Type() string {
	return TypeMsgDeletePool
}

func (msg *MsgDeletePool) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeletePool) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeletePool) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
