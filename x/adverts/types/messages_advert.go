package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateAdvert = "create_advert"
	TypeMsgUpdateAdvert = "update_advert"
	TypeMsgDeleteAdvert = "delete_advert"
)

var _ sdk.Msg = &MsgCreateAdvert{}

func NewMsgCreateAdvert(creator string, advertId string, state string, assetId string, price uint64, minRentalPeriod uint64, maxRentalPeriod uint64, conditions string, expirationDate uint64) *MsgCreateAdvert {
	return &MsgCreateAdvert{
		Creator:         creator,
		AdvertId:        advertId,
		State:           state,
		AssetId:         assetId,
		Price:           price,
		MinRentalPeriod: minRentalPeriod,
		MaxRentalPeriod: maxRentalPeriod,
		Conditions:      conditions,
		ExpirationDate:  expirationDate,
	}
}

func (msg *MsgCreateAdvert) Route() string {
	return RouterKey
}

func (msg *MsgCreateAdvert) Type() string {
	return TypeMsgCreateAdvert
}

func (msg *MsgCreateAdvert) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateAdvert) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateAdvert) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateAdvert{}

func NewMsgUpdateAdvert(creator string, id uint64, advertId string, state string, assetId string, price uint64, minRentalPeriod uint64, maxRentalPeriod uint64, conditions string, expirationDate uint64) *MsgUpdateAdvert {
	return &MsgUpdateAdvert{
		Id:              id,
		Creator:         creator,
		AdvertId:        advertId,
		State:           state,
		AssetId:         assetId,
		Price:           price,
		MinRentalPeriod: minRentalPeriod,
		MaxRentalPeriod: maxRentalPeriod,
		Conditions:      conditions,
		ExpirationDate:  expirationDate,
	}
}

func (msg *MsgUpdateAdvert) Route() string {
	return RouterKey
}

func (msg *MsgUpdateAdvert) Type() string {
	return TypeMsgUpdateAdvert
}

func (msg *MsgUpdateAdvert) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateAdvert) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateAdvert) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteAdvert{}

func NewMsgDeleteAdvert(creator string, id uint64) *MsgDeleteAdvert {
	return &MsgDeleteAdvert{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteAdvert) Route() string {
	return RouterKey
}

func (msg *MsgDeleteAdvert) Type() string {
	return TypeMsgDeleteAdvert
}

func (msg *MsgDeleteAdvert) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteAdvert) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteAdvert) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
