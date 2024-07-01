package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "remap/testutil/keeper"
	"remap/testutil/nullify"
	"remap/x/assets/keeper"
	"remap/x/assets/types"
)

func createNAsset(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Asset {
	items := make([]types.Asset, n)
	for i := range items {
		items[i].Id = keeper.AppendAsset(ctx, items[i])
	}
	return items
}

func TestAssetGet(t *testing.T) {
	keeper, ctx := keepertest.AssetsKeeper(t)
	items := createNAsset(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetAsset(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestAssetRemove(t *testing.T) {
	keeper, ctx := keepertest.AssetsKeeper(t)
	items := createNAsset(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveAsset(ctx, item.Id)
		_, found := keeper.GetAsset(ctx, item.Id)
		require.False(t, found)
	}
}

func TestAssetGetAll(t *testing.T) {
	keeper, ctx := keepertest.AssetsKeeper(t)
	items := createNAsset(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllAsset(ctx)),
	)
}

func TestAssetCount(t *testing.T) {
	keeper, ctx := keepertest.AssetsKeeper(t)
	items := createNAsset(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetAssetCount(ctx))
}
