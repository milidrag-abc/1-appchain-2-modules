package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateAsset = "create_asset"
	TypeMsgUpdateAsset = "update_asset"
	TypeMsgDeleteAsset = "delete_asset"
)

var _ sdk.Msg = &MsgCreateAsset{}

func NewMsgCreateAsset(creator string, assetId string, version uint64, state string, vcAvailable bool, assetType string, owner string, address string, assetSize uint64, properties string) *MsgCreateAsset {
	return &MsgCreateAsset{
		Creator:     creator,
		AssetId:     assetId,
		Version:     version,
		State:       state,
		VcAvailable: vcAvailable,
		AssetType:   assetType,
		Owner:       owner,
		Address:     address,
		AssetSize:   assetSize,
		Properties:  properties,
	}
}

func (msg *MsgCreateAsset) Route() string {
	return RouterKey
}

func (msg *MsgCreateAsset) Type() string {
	return TypeMsgCreateAsset
}

func (msg *MsgCreateAsset) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateAsset) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateAsset) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateAsset{}

func NewMsgUpdateAsset(creator string, id uint64, assetId string, version uint64, state string, vcAvailable bool, assetType string, owner string, address string, assetSize uint64, properties string) *MsgUpdateAsset {
	return &MsgUpdateAsset{
		Id:          id,
		Creator:     creator,
		AssetId:     assetId,
		Version:     version,
		State:       state,
		VcAvailable: vcAvailable,
		AssetType:   assetType,
		Owner:       owner,
		Address:     address,
		AssetSize:   assetSize,
		Properties:  properties,
	}
}

func (msg *MsgUpdateAsset) Route() string {
	return RouterKey
}

func (msg *MsgUpdateAsset) Type() string {
	return TypeMsgUpdateAsset
}

func (msg *MsgUpdateAsset) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateAsset) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateAsset) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteAsset{}

func NewMsgDeleteAsset(creator string, id uint64) *MsgDeleteAsset {
	return &MsgDeleteAsset{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteAsset) Route() string {
	return RouterKey
}

func (msg *MsgDeleteAsset) Type() string {
	return TypeMsgDeleteAsset
}

func (msg *MsgDeleteAsset) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteAsset) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteAsset) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
