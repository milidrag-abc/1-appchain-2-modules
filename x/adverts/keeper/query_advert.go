package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"remap/x/adverts/types"
)

func (k Keeper) AdvertAll(goCtx context.Context, req *types.QueryAllAdvertRequest) (*types.QueryAllAdvertResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var adverts []types.Advert
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	advertStore := prefix.NewStore(store, types.KeyPrefix(types.AdvertKey))

	pageRes, err := query.Paginate(advertStore, req.Pagination, func(key []byte, value []byte) error {
		var advert types.Advert
		if err := k.cdc.Unmarshal(value, &advert); err != nil {
			return err
		}

		adverts = append(adverts, advert)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllAdvertResponse{Advert: adverts, Pagination: pageRes}, nil
}

func (k Keeper) Advert(goCtx context.Context, req *types.QueryGetAdvertRequest) (*types.QueryGetAdvertResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	advert, found := k.GetAdvert(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetAdvertResponse{Advert: advert}, nil
}
