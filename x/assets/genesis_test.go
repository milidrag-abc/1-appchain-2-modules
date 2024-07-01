package assets_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "remap/testutil/keeper"
	"remap/testutil/nullify"
	"remap/x/assets"
	"remap/x/assets/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		AssetList: []types.Asset{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		AssetCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.AssetsKeeper(t)
	assets.InitGenesis(ctx, *k, genesisState)
	got := assets.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.AssetList, got.AssetList)
	require.Equal(t, genesisState.AssetCount, got.AssetCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
