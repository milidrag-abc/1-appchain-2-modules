package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"remap/x/adverts/types"
)

func (k msgServer) CreateAdvert(goCtx context.Context, msg *types.MsgCreateAdvert) (*types.MsgCreateAdvertResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)





	var advert = types.Advert{
		Creator:         msg.Creator,
		AdvertId:        msg.AdvertId,
		State:           msg.State,
		AssetId:         msg.AssetId,
		Price:           msg.Price,
		MinRentalPeriod: msg.MinRentalPeriod,
		MaxRentalPeriod: msg.MaxRentalPeriod,
		Conditions:      msg.Conditions,
		ExpirationDate:  msg.ExpirationDate,
	}

	id := k.AppendAdvert(
		ctx,
		advert,
	)

	return &types.MsgCreateAdvertResponse{
		Id: id,
	}, nil
}

func (k msgServer) UpdateAdvert(goCtx context.Context, msg *types.MsgUpdateAdvert) (*types.MsgUpdateAdvertResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var advert = types.Advert{
		Creator:         msg.Creator,
		Id:              msg.Id,
		AdvertId:        msg.AdvertId,
		State:           msg.State,
		AssetId:         msg.AssetId,
		Price:           msg.Price,
		MinRentalPeriod: msg.MinRentalPeriod,
		MaxRentalPeriod: msg.MaxRentalPeriod,
		Conditions:      msg.Conditions,
		ExpirationDate:  msg.ExpirationDate,
	}

	// Checks that the element exists
	val, found := k.GetAdvert(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.SetAdvert(ctx, advert)

	return &types.MsgUpdateAdvertResponse{}, nil
}

func (k msgServer) DeleteAdvert(goCtx context.Context, msg *types.MsgDeleteAdvert) (*types.MsgDeleteAdvertResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetAdvert(ctx, msg.Id)
	if !found {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveAdvert(ctx, msg.Id)

	return &types.MsgDeleteAdvertResponse{}, nil
}
