package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "remap/testutil/keeper"
	"remap/testutil/nullify"
	"remap/x/adverts/keeper"
	"remap/x/adverts/types"
)

func createNAdvert(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Advert {
	items := make([]types.Advert, n)
	for i := range items {
		items[i].Id = keeper.AppendAdvert(ctx, items[i])
	}
	return items
}

func TestAdvertGet(t *testing.T) {
	keeper, ctx := keepertest.AdvertsKeeper(t)
	items := createNAdvert(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetAdvert(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestAdvertRemove(t *testing.T) {
	keeper, ctx := keepertest.AdvertsKeeper(t)
	items := createNAdvert(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAdvert(ctx, item.Id)
		_, found := keeper.GetAdvert(ctx, item.Id)
		require.False(t, found)
	}
}

func TestAdvertGetAll(t *testing.T) {
	keeper, ctx := keepertest.AdvertsKeeper(t)
	items := createNAdvert(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllAdvert(ctx)),
	)
}

func TestAdvertCount(t *testing.T) {
	keeper, ctx := keepertest.AdvertsKeeper(t)
	items := createNAdvert(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetAdvertCount(ctx))
}
